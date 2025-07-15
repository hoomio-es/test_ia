-- Crear usuario administrador por defecto
INSERT INTO users (username, email, password_hash, is_admin)
VALUES ('admin', 'admin@example.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', true);

-- Insertar permisos para el usuario admin a todas las propiedades
INSERT INTO user_properties (user_id, property_id, can_read, can_write)
SELECT 
    (SELECT id FROM users WHERE username = 'admin'),
    p.id,
    true,
    true
FROM properties p;
