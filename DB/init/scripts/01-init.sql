-- Conexión a base de datos
\c meli_nist


-- crear tabla user

CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    name VARCHAR(50) NOT NULL,
    password VARCHAR(200) NOT NULL,
    email VARCHAR(50) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP

);


-- crear tabla app

CREATE TABLE apps(
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(50) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);


-- crear tabla app_rules

CREATE TABLE app_rules(
    id SERIAL PRIMARY KEY,
    rule_id VARCHAR(20) NOT NULL,
    app_id INTEGER NOT NULL,
    deleted_at TIMESTAMP,
    updated_at TIMESTAMP,
    created_at TIMESTAMP,
    FOREIGN KEY(app_id) REFERENCES apps(id)
);


-- crear usuario inicial

INSERT INTO users(username,name,password,email,created_at) values ('admin','admin','$2a$10$HNYBxj1EV4yp381AOvzf8eVvG.0tX.vptVx1OdF5XVI/la.ICJSJq','admin@admin.com',NOW())