-- Tabela users
CREATE TABLE IF NOT EXISTS users (
    id SERIAL,
    name VARCHAR(100) NOT NULL,
    cpf VARCHAR(20) PRIMARY KEY,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL
);

-- Tabela driver
CREATE TABLE IF NOT EXISTS drivers (
    id SERIAL,
    cpf VARCHAR(14) PRIMARY KEY,
    cnh VARCHAR(20) NOT NULL,
    qrcode VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL
);

-- Tabela address
CREATE TABLE IF NOT EXISTS address (
    id SERIAL PRIMARY KEY,
    cpf VARCHAR(14) REFERENCES users(cpf),
    street VARCHAR(100) NOT NULL,
    city VARCHAR(100) NOT NULL,
    state VARCHAR(2) NOT NULL,
    zip VARCHAR(8) NOT NULL,
    number VARCHAR(10) NOT NULL,
    complement VARCHAR(100)
);

-- Tabela children
CREATE TABLE IF NOT EXISTS childrens (
    id SERIAL,
    rg VARCHAR(20) PRIMARY KEY,
    responsible VARCHAR(14),
    name VARCHAR(100),
    school VARCHAR(100),
    driver VARCHAR(100),
    address TEXT,
    FOREIGN KEY (responsible) REFERENCES users(cpf)
);

-- Tabela schools
CREATE TABLE IF NOT EXISTS schools (
    id SERIAL,
    name VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    cnpj VARCHAR(14) PRIMARY KEY,
    street VARCHAR(100) NOT NULL,
    number VARCHAR(10) NOT NULL,
    zip VARCHAR(8) NOT NULL,
    email VARCHAR(100) NOT NULL
);

-- Tabela school_drivers
CREATE TABLE IF NOT EXISTS schools_drivers (
    record SERIAL PRIMARY KEY,
    school VARCHAR(14),
    driver VARCHAR(14),
    FOREIGN KEY (school) REFERENCES schools(cnpj),
    FOREIGN KEY (driver) REFERENCES drivers(cpf)
);

-- Tabela users_drivers
CREATE TABLE IF NOT EXISTS users_drivers (
    registration SERIAL PRIMARY KEY,
    driver VARCHAR(14),
    child VARCHAR(20),
    FOREIGN KEY (driver) REFERENCES drivers(cpf),
    FOREIGN KEY (child) REFERENCES childrens(rg)
);

