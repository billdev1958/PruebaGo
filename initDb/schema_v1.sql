-- Tabla de dependencias
CREATE TABLE IF NOT EXISTS users(
	id       UUID,
	username VARCHAR(70) NOT NULL,
	phone    VARCHAR(10) NOT NULL,
	email    VARCHAR(100) NOT NULL,
    password VARCHAR(150) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);