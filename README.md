# Template сервиса авторизации

## Описание

Это шаблон простого сервиса авторизации.

Стек:
 - PostgreSQL
 - gRPC

Конфигурация указывается в `.env` файле. Пример минимальной конфигурации находится в файле `example.env`

В `makefile` в корне проекта есть полезные команды для разработки и эксплуатации.

## Запуск проекта

Установка зависимостей
```sh
go mod tidy
```
Запуск проекта
```sh
make run
```

## API

### auth.AuthService/CreateUser

#### Request
```json
{
    "user": {
        "username": "user", 
        "password": "password1"
    }
}
```

#### Response
```json
{
  "user": {
    "id": "7",
    "username": "user",
    "password": "password1"
  }
}
```

### auth.AuthService/CreateToken

#### Request
```json
{
    "user": {
        "username": "user", 
        "password": "password1"
    }
}
```

#### Response
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTUyNzAyNjEsImlhdCI6MTc1NTI2OTY2MSwic3ViIjoiNyJ9.SQprAUGFN8ECa_cztHBfSJfJF44PG2tBiaFNaEdHy5U"
}
```

### auth.AuthService/RecreateToken

#### Request
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTUyNzAyNjEsImlhdCI6MTc1NTI2OTY2MSwic3ViIjoiNyJ9.SQprAUGFN8ECa_cztHBfSJfJF44PG2tBiaFNaEdHy5U"
}
```

#### Response
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTUyNzAzMjcsImlhdCI6MTc1NTI2OTcyNywic3ViIjoiNSJ9.8iyEuBh9yLYJtC7pyV8hXTmOjBt74b-QbI0do-OU7jo"
}
```