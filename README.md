# Ingeniería de Software II - TP I

## Índice

1. [Introducción](#1-introducción)
2. [Desafíos Encontrados](#2-desafíos-encontrados)
3. [Pre-Requisitos](#3-pre-requisitos)
4. [Docker](#4-docker)
5. [Posibles Mejoras](#5-posibles-mejoras)
6. [Testing](#6-Testing)
7. [Comandos](#7-comandos)

## 1. Introducción

En este trabajo se buscan implementar los endpoints `GET`, `POST`, `GET/{id}` y `DELETE/{id}` siguiendo el contrato HTTP. Al llegar un nuevo request, la clase **Router** se encarga de recibirlo y derivar la consulta al método apropiado de la clase **Controller**.

Los mensajes de *snap* se guardan con un **UUID identificatorio** y un **timestamp**. Se busca usar el timestamp para organizar los snaps en orden cronológico inverso al recibir el comando `GET`.

Se tiene una base de datos donde se guardan los snaps. Los errores que pueden ocurrir en los diferentes endpoints se manejan siguiendo el formato de la **RFC 7807**.

## 2. Desafíos Encontrados

El uso de los endpoints de HTTP fue un desafío ya que nunca los había utilizado, desconocía tanto los formatos como las librerías (como lo es `gin`). Lo mismo me sucedió a la hora de testear.

Hacer uso de `cURL`, implementar la persistencia en memoria y el uso de Docker para desplegar la aplicación fueron otros desafíos que encontré ya que no tenía experiencia previa en estos temas.

Me hubiera gustado hacer uso de una **BDD propiamente dicha**, pero no contaba con el tiempo suficiente para implementarla. Muy probablemente hubiera sido un desafío adicional y muy interesante.

## 3. Pre-Requisitos

- **Lenguaje**:
    - `go version go1.24.1 linux/amd64`
- **Docker**:
    - `Version 28.0.1`
    - `API version 1.48`

## 4. Docker

#### Construir la imagen de docker:

    docker build -t classconnect-api .

#### Levantar el contenedor:

    docker run --env-file .env -p 8080:8080 classconnect-api

### 5. posibles-mejoras

- Agregar una base de datos (PostgreSQL o MySQL) para persistencia real
- Implementar un sistema de caché (Redis) para mejorar rendimiento

### 6. Testing

Se han implementado pruebas automatizadas usando [httptest](https://pkg.go.dev/net/http/httptest) y [Testify](https://pkg.go.dev/github.com/stretchr/testify/assert).

### 7. Comandos

> Los comandos `cURL` utilizados fueron los siguientes

### **GET**

Conseguir todos los cursos registrados en el sistema.

```sh
curl http://localhost:8080/courses
```

### **GET with ID**

Conseguir el curso correspondiente al {id}.

```sh
curl http://localhost:8080/courses/{id}
```

### **POST**

Crear y guardar un nuevo curso con el título y la descripción indicados.

```sh
curl -X POST http://localhost:8080/courses \
  -H "Content-Type: application/json" \
  -d '{"title": "Go Basics", "description": "Learn Golang"}'
```

### **DELETE**

Borrar el curso correspondiente al {id}.

```sh
curl -X DELETE http://localhost:8080/courses/{id}
```
