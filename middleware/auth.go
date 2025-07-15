package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/sessions"
)

// AuthMiddleware verifica si el usuario está autenticado
func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Obtener la sesión
		session, err := sessions.NewCookieStore([]byte("your-secret-key")).Get(c.Request(), "user-session")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error de sesión")
		}

		// Verificar si el usuario está autenticado
		if !session.Values["authenticated"].(bool) {
			return c.Redirect("/login")
		}

		// Continuar con la siguiente middleware/ruta
		return c.Next()
	}
}
