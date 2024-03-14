-- Tabela users
CREATE TABLE IF NOT EXISTS user (
    id SERIAL,
    name VARCHAR(100) NOT NULL,
    cpf VARCHAR(20) PRIMARY KEY,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL
);

-- Tabela driver
CREATE TABLE IF NOT EXISTS driver (
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
CREATE TABLE IF NOT EXISTS children (
    id SERIAL PRIMARY KEY,
    rg VARCHAR(20) PRIMARY KEY,
    responsavel INT REFERENCES driver(id),
    nome VARCHAR(100),
    escola VARCHAR(100),
    driver VARCHAR(100),
    endereco INT REFERENCES address(id)
);

-- Tabela schools
CREATE TABLE IF NOT EXISTS school (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    cnpj VARCHAR(14) PRIMARY KEY,
    rua VARCHAR(100) NOT NULL,
    numero VARCHAR(10) NOT NULL,
    cep VARCHAR(8 NOT NULL,
    email VARCHAR(100) NOT NULL
);

-- Tabela school_drivers
CREATE TABLE IF NOT EXISTS school_driver (
    registro SERIAL PRIMARY KEY,
    school VARCHAR(14) REFERENCES schools(cnpj),
    driver VARCHAR(14) REFERENCES driver(cpf)
);

-- Tabela users_drivers
CREATE TABLE IF NOT EXISTS users_drivers (
    matricula SERIAL PRIMARY KEY,
    driver VARCHAR(14) REFERENCES driver(cpf),
    child VARCHAR(20) REFERENCES children(rg)
);

