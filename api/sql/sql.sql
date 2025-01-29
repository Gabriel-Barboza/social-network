CREATE DATABASE IF NOT EXISTS devbook;
use devbook;

drop table if exists usuarios;


create table usuarios(
    id int auto_increment primary key,
    nome varchar(100) not null,
    nick varchar(50) not null unique,
    email varchar(100) not null unique,
    senha varchar(100) not null,
    criadoEm datetime default current_timestamp())
    engine=InnoDB default charset=utf8;


