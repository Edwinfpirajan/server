-- Crear tabla Collaborators
CREATE TABLE collaborators (
    id serial PRIMARY KEY,
    document VARCHAR(25) UNIQUE,
    f_name VARCHAR(50) NOT NULL,
    l_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    bmail VARCHAR(100) NOT NULL,
    position VARCHAR(45) NOT NULL,
    state VARCHAR(10) NOT NULL,
    leader VARCHAR(50) NOT NULL,
    leader_document VARCHAR(20),
    process VARCHAR(50),
    subprocess VARCHAR(50),
    headquarters VARCHAR(30),
    created_at TIMESTAMP
);

-- Crear tabla Attendances
CREATE TABLE attendances (
    id serial PRIMARY KEY,
    arrival TIME,
    departure TIME,
    location VARCHAR(10),
    late BOOLEAN,
    early_arrival BOOLEAN,
    photo_arrival BYTEA,
    photo_departure BYTEA,
    created_at TIMESTAMP,
    fk_collaborator_id INT,
    CONSTRAINT fk_collaborator_id FOREIGN KEY (fk_collaborator_id) REFERENCES collaborators (id)
);

-- Crear tabla Schedules
CREATE TABLE schedules (
    id serial PRIMARY KEY,
    day VARCHAR(11),
    arrival_time VARCHAR,
    departure_time VARCHAR,
    fk_collaborator_id INTEGER,
    CONSTRAINT fk_collaborator_id FOREIGN KEY (fk_collaborator_id) REFERENCES collaborators(id)
);

-- Crear tabla TranslatedCollaborators
CREATE TABLE TranslatedCollaborators (
    id serial PRIMARY KEY,
    created_at TIMESTAMP,
    fk_collaborator_id INTEGER,
    CONSTRAINT fk_collaborator_id FOREIGN KEY (fk_collaborator_id) REFERENCES collaborators(id)
);

-- Crear tabla Users
CREATE TABLE users (
    id serial PRIMARY KEY,
    document VARCHAR(25),
    f_name VARCHAR(50) NOT NULL,
    l_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(12) NOT NULL,
    created_at TIMESTAMP,
    fk_role_id INTEGER REFERENCES roles(id) ON DELETE CASCADE
);

-- Crear tabla Roles
CREATE TABLE roles (
    id serial PRIMARY KEY,
    name VARCHAR(25)
);

-- Insertar datos en Users y Headquarters
INSERT INTO users (f_name, l_name, email, fk_role_id, password)
VALUES ('Edwin Fernando', 'Pirajan Arevalo', 'epirajan@smart.edu.co', 1, 'Smart2023++');

CREATE TABLE headquarters (
    id serial PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

INSERT INTO headquarters (name)
VALUES
    ('ADMINISTRATIVO'), ('ARKADIA'), ('BELLO'), ('CALASANZ'), ('CALIMA'), ('CEDRITOS'),
    ('CENTRO INTERNACIONAL'), ('CENTRO MAYOR'), ('CENTRO MEDELLIN'), ('CHAPINERO'), ('CHIA'),
    ('ENVIGADO'), ('FLORIDABLANCA'), ('FONTANAR'), ('HAYUELOS'), ('ITAGÜÍ'), ('MADELENA'),
    ('MODELIA'), ('MULTIPLAZA'), ('NUESTRO BOGOTÁ'), ('OLAYA'), ('PALATINO'), ('PIEDECUESTA'),
    ('PLAZA CENTRAL'), ('PLAZA DE LAS AMERICAS'), ('POBLADO'), ('RESTREPO'), ('SAN MARTÍN'),
    ('SANTAFÉ'), ('ONLINE'), ('SOACHA'), ('SUBA'), ('UNICENTRO DE OCCIDENTE'), ('VIRTUAL');
