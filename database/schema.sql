-- Tabela responsibles
CREATE TABLE IF NOT EXISTS responsibles (
    id SERIAL,
    name VARCHAR(100) NOT NULL,
    cpf VARCHAR(20) PRIMARY KEY,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    street VARCHAR(100) NOT NULL,
    number VARCHAR(10) NOT NULL,
    complement VARCHAR(10),
    zip VARCHAR(8) NOT NULL
);

-- Tabela driver
CREATE TABLE IF NOT EXISTS drivers (
    id SERIAL,
    name VARCHAR(100) NOT NULL,
    cpf VARCHAR(14) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    cnh VARCHAR(20) PRIMARY KEY NOT NULL,
    qrcode VARCHAR(100) NOT NULL,
    street VARCHAR(100) NOT NULL,
    number VARCHAR(10) NOT NULL,
    complement VARCHAR(10),
    zip VARCHAR(8) NOT NULL
);

-- Tabela children
CREATE TABLE IF NOT EXISTS children (
    id SERIAL,
    rg VARCHAR(20) PRIMARY KEY,
    responsibles VARCHAR(14),
    shift VARCHAR(20) NOT NULL,
    name VARCHAR(100),
    school VARCHAR(100),
    driver VARCHAR(100),
    street VARCHAR(100) NOT NULL,
    number VARCHAR(10) NOT NULL,
    complement VARCHAR(10),
    zip VARCHAR(8) NOT NULL,
    FOREIGN KEY (responsibles) REFERENCES responsibles(cpf)
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
    email VARCHAR(100) NOT NULL,
    complement VARCHAR(10)
);

-- Tabela school_drivers (Relação entre Escola e Motorista)
CREATE TABLE IF NOT EXISTS schools_drivers (
    record SERIAL PRIMARY KEY,
    name_school VARCHAR(100) NOT NULL,
    school VARCHAR(14),
    email_school VARCHAR(100) NOT NULL,
    name_driver VARCHAR(100) NOT NULL,
    driver VARCHAR(14),
    email_driver VARCHAR(100) NOT NULL,
    FOREIGN KEY (school) REFERENCES schools(cnpj),
    FOREIGN KEY (driver) REFERENCES drivers(cnh)
);

-- Tabela responsibles_drivers
CREATE TABLE IF NOT EXISTS responsibles_drivers (
    record SERIAL PRIMARY KEY,
    name_driver VARCHAR(50),
    driver VARCHAR(14),
    name_child VARCHAR(50),
    child VARCHAR(20),
    FOREIGN KEY (driver) REFERENCES drivers(cnh),
    FOREIGN KEY (child) REFERENCES children(rg)
);

-- Tabela responsibles_schools
CREATE TABLE IF NOT EXISTS responsibles_schools (
    record SERIAL PRIMARY KEY,
    name_school VARCHAR(100) NOT NULL,
    school VARCHAR(14) NOT NULL,
    name_child VARCHAR(50) NOT NULL,
    child VARCHAR(10) NOT NULL,
    FOREIGN KEY (school) REFERENCES schools(cnpj),
    FOREIGN KEY (child) REFERENCES children(rg)
);

-- Tabela de Convites
CREATE TABLE IF NOT EXISTS invites (
    invite_id SERIAL PRIMARY KEY,
    requester VARCHAR(14), -- school
    school VARCHAR(100) NOT NULL, --name_school
    email_school VARCHAR(100) NOT NULL,
    guest VARCHAR(14), -- driver
    driver VARCHAR(100) NOT NULL, --name_driver
    email_driver VARCHAR(100) NOT NULL,
    status TEXT NOT NULL,
    FOREIGN KEY (requester) REFERENCES schools(cnpj),
    FOREIGN KEY (guest) REFERENCES drivers(cnh)
);

