# Person API

[![Build Status](https://travis-ci.org/rodrigoherera/go-persons.svg)](https://travis-ci.org/rodrigoherera/go-persons)
[![codecov.io](https://codecov.io/gh/rodrigoherera/go-persons/branch/master/graph/badge.svg)](https://codecov.io/gh/rodrigoherera/go-persons)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=rodrigoherera_go-persons&metric=alert_status)](https://sonarcloud.io/dashboard?id=rodrigoherera_go-persons)

## REST Endpoints

## GET

Endpoint:

```cmd
/
```

Ejemplo de respuesta:

```golang
{
    "name": "Person API",
    "version": "1.0"
}
```

-------------------------------------------------------

Endpoint:

```cmd
/v1/login
```

Se envía con un Basic Auth, previo creado el usuario con email y password.

Ejemplo de respuesta:

```golang
{
    "Email": "test@test.com",
    "Name": "Bearer token",
    "Value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJvZHJpZ28uaGVyZXJhQG1lcmNhZG9saWJyZS5jb20iLCJleHAiOjE2MDAwOTI0NDN9.bP3f-sRhF2Dse-fCKrUnxJW4kKfDfjli3cTGM6Qs0kI",
    "Expires": "2020-09-14 11:07:23.294109 -0300 -03 m=+4117.806553302"
}
```

-------------------------------------------------------

Endpoint:

```cmd
/v1/person
```

Ejemplo de respuesta:

```golang
[
    {
        "ID": 9,
        "name": "Test",
        "lastname": "Test",
        "age": 26,
        "dni": 1234567
    },
    {
        "ID": 10,
        "name": "Test 2",
        "lastname": "Test 3",
        "age": 44,
        "dni": 123456745
    }
]
```

-------------------------------------------------------

Endpoint:

```cmd
/v1/person/:id
```

Ejemplo de respuesta:

```golang
{
    "ID": 9,
    "name": "Test",
    "lastname": "Test",
    "age": 26,
    "dni": 1234567
}
```

-------------------------------------------------------

## POST

Endpoint :

```cmd
/v1/user
```

Se envía en el BODY del user a crear:

```golang
{
    "email": "test@gmail.com",
    "password": "Test3!sd_",
}
```

Retorna el ID y el email que se guardó del user.

```golang
{
    "id": "3",
    "email": "test@test.com"
}
```

-------------------------------------------------------

Endpoint:

```cmd
/v1/person
```

Se envía en el BODY la persona a crear:

```golang
{
    "name": "Test 2",
    "lastname": "Test 3",
    "age": 44,
    "dni": 123456745
}
```

Retorna la nueva persona creada.

Ejemplo de respuesta:

```golang
{
    "ID": 10,
    "name": "Test 2",
    "lastname": "Test 3",
    "age": 44,
    "dni": 123456745
}
```

-------------------------------------------------------

## DELETE

Endpoint:

```cmd
/v1/person/:id
```

Ejemplo de respuesta:

```golang
OK
```

-------------------------------------------------------

## PUT

Endpoint:

```cmd
/v1/person/:id
```

Se envía en el BODY los nuevos datos a updatear:

```golang
{
    "name": "Test 2",
    "lastname": "Test 3",
    "age": 44,
    "dni": 123456745
}
```

Ejemplo de respuesta:

```golang
OK
```
