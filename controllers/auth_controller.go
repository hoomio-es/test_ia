package controllers

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/sessions"
	"github.com/jgargallo/test_app/models"
	"github.com/jgargallo/test_app/middleware"
	"golang.org/x/crypto/bcrypt"
	"log"
)

var db *sql.DB

func LoginHandler(c *fiber.Ctx) error {
	if c.Method() == "GET" {
		return c.Render("auth/login.html", nil)
	}

	username := c.FormValue("username")
	password := c.FormValue("password")

	// Buscar usuario en la base de datos
	var user struct {
		ID       int
		Password string
		IsAdmin  bool
	}

	err := db.QueryRow("SELECT id, password_hash, is_admin FROM users WHERE username = $1", username).Scan(&user.ID, &user.Password, &user.IsAdmin)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusUnauthorized).SendString("Credenciales inválidas")
		}
		log.Printf("Error al buscar usuario: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error interno")
	}

	// Verificar contraseña
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Credenciales inválidas")
	}

	// Crear sesión
	session, err := middleware.store.Get(c.Request(), "user-session")
	if err != nil {
		log.Printf("Error al crear sesión: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error interno")
	}

	session.Values["authenticated"] = true
	session.Values["user_id"] = user.ID
	session.Values["username"] = username
	session.Values["is_admin"] = user.IsAdmin
	if err := session.Save(c.Request(), c.Response()); err != nil {
		log.Printf("Error al guardar sesión: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error interno")
	}

	return c.Redirect("/")
}

func RegisterHandler(c *fiber.Ctx) error {
	if c.Method() == "GET" {
		return c.Render("auth/register.html", nil)
	}

	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")
	passwordConfirm := c.FormValue("password_confirm")

	// Validar campos
	if password != passwordConfirm {
		return c.Status(fiber.StatusBadRequest).SendString("Las contraseñas no coinciden")
	}

	// Verificar si el usuario ya existe
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = $1 OR email = $2", username, email).Scan(&count)
	if err != nil {
		log.Printf("Error al verificar usuario: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error interno")
	}
	if count > 0 {
		return c.Status(fiber.StatusBadRequest).SendString("El nombre de usuario o correo ya están en uso")
	}

	// Hashear la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error al hashear contraseña: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error interno")
	}

	// Insertar usuario
	var userID int
	err = db.QueryRow("INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3) RETURNING id", 
		username, email, string(hashedPassword)).Scan(&userID)
	if err != nil {
		log.Printf("Error al insertar usuario: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error interno")
	}

	// Crear sesión
	session, err := middleware.store.Get(c.Request(), "user-session")
	if err != nil {
		log.Printf("Error al crear sesión: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error interno")
	}

	session.Values["authenticated"] = true
	session.Values["user_id"] = userID
	session.Values["username"] = username
	if err := session.Save(c.Request(), c.Response()); err != nil {
		log.Printf("Error al guardar sesión: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error interno")
	}

	return c.Redirect("/")
}

func LogoutHandler(c *fiber.Ctx) error {
	session, err := middleware.store.Get(c.Request(), "user-session")
	if err != nil {
		log.Printf("Error al obtener sesión: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error interno")
	}

	// Eliminar sesión
	session.Options.MaxAge = -1
	if err := session.Save(c.Request(), c.Response()); err != nil {
		log.Printf("Error al guardar sesión: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error interno")
	}

	return c.Redirect("/login")
}
