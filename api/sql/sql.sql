CREATE DATABASE IF NOT EXISTS devbook;
use devbook;

drop table if EXISTS publicacoes;
drop table if EXISTS seguidores;
drop table if exists usuarios;

create table usuarios(
    id int auto_increment primary key,
    nome varchar(100) not null,
    nick varchar(50) not null unique,
    email varchar(100) not null unique,
    senha varchar(100) not null,
    criadoEm datetime default current_timestamp())
    engine=InnoDB default charset=utf8;


create table seguidores (
    usuario_id int not null,
    Foreign Key (usuario_id) REFERENCES usuarios(id)
    on delete cascade,

    seguidor_id int not null,
    Foreign Key (seguidor_id) REFERENCES usuarios(id)
    on delete cascade,

    PRIMARY KEY (usuario_id, seguidor_id)
    ) engine=InnoDB default charset=utf8;

create table publicacoes(
    id int auto_increment primary key,
    titulo VARCHAR(100) not null,
    conteudo VARCHAR(1000) not null,
    autor_id int not null,
    Foreign Key (autor_id) REFERENCES usuarios(id) 
    on delete cascade,

    curtidas int DEFAULT 0,
    criadoEm datetime default current_timestamp()
    ) engine=InnoDB default charset=utf8;


