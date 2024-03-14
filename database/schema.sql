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
    rua VARCHAR(100) NOT NULL,
    cidade VARCHAR(100) NOT NULL,
    estado VARCHAR(2) NOT NULL,
    cep VARCHAR(8) NOT NULL,
    numero VARCHAR(10) NOT NULL,
    complemento VARCHAR(100)
);

-- Tabela children
CREATE TABLE IF NOT EXISTS childrens (
    id SERIAL,
    rg VARCHAR(20) PRIMARY KEY,
    responsavel VARCHAR(14),
    nome VARCHAR(100),
    escola VARCHAR(100),
    driver VARCHAR(100),
    endereco TEXT,
    FOREIGN KEY (responsavel) REFERENCES drivers(cpf)
);

-- Tabela schools
CREATE TABLE IF NOT EXISTS schools (
    id SERIAL,
    nome VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    cnpj VARCHAR(14) PRIMARY KEY,
    rua VARCHAR(100) NOT NULL,
    numero VARCHAR(10) NOT NULL,
    cep VARCHAR(8) NOT NULL,
    email VARCHAR(100) NOT NULL
);

-- Tabela school_drivers
CREATE TABLE IF NOT EXISTS schools_drivers (
    registro SERIAL PRIMARY KEY,
    school VARCHAR(14),
    driver VARCHAR(14),
    FOREIGN KEY (school) REFERENCES schools(cnpj),
    FOREIGN KEY (driver) REFERENCES drivers(cpf)
);

-- Tabela users_drivers
CREATE TABLE IF NOT EXISTS users_drivers (
    matricula SERIAL PRIMARY KEY,
    driver VARCHAR(14),
    child VARCHAR(20),
    FOREIGN KEY (driver) REFERENCES drivers(cpf),
    FOREIGN KEY (child) REFERENCES childrens(rg)
);

