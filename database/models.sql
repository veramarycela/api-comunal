CREATE TABLE IF NOT EXISTS residente (
    cedula VARCHAR NOT NULL,
    nombre VARCHAR NOT NULL,
    apellido VARCHAR NOT NULL,
    fecha_nac DATE NOT NULL,
    direccion VARCHAR NOT NULL, 
    telefono VARCHAR NOT NULL, 
    tipo VARCHAR NOT NULL,   
    CONSTRAINT pk_residente PRIMARY KEY(cedula)
);

CREATE TABLE IF NOT EXISTS documento (
    codigo serial NOT NULL,
    tipo VARCHAR NOT NULL,
    dirigido VARCHAR,
    contenido VARCHAR,
    autorizado VARCHAR, 
    CONSTRAINT pk_documento PRIMARY KEY(codigo)
);

CREATE TABLE IF NOT EXISTS solicitud (
    codigo serial NOT NULL,
    cedula_residente VARCHAR NOT NULL,
    codigo_documento serial NOT NULL,
    fecha DATE,
    CONSTRAINT pk_solicitud PRIMARY KEY(codigo),
    CONSTRAINT fk_residente FOREIGN KEY(cedula_residente) REFERENCES residente(cedula),
    CONSTRAINT fk_documento FOREIGN KEY(codigo_documento) REFERENCES documento(codigo)
);

CREATE TABLE IF NOT EXISTS usuario (
    id serial NOT NULL,
    username VARCHAR NOT NULL UNIQUE,
    password varchar NOT NULL,
    privilegios VARCHAR NOT NULL,
    rol VARCHAR NOT NULL,
    cedula_residente VARCHAR NOT NULL,
    CONSTRAINT pk_usuario PRIMARY KEY(id),
    CONSTRAINT fk_residente FOREIGN KEY(cedula_residente) REFERENCES residente(cedula) 
);
