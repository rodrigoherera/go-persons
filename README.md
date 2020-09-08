# Person API

## REST Endpoints

## GET

WIP

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

```cmd
/v1/login/:id
```

Ejemplo de respuesta:

```golang
{
    "Name": "token",
    "Value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjgiLCJleHAiOjE1OTk1NzM5ODJ9.yDGxeuuRCc2SfnDRCYfD68pFFH1ndSN7zDH04PdeLcM",
    "Expires": "2020-09-08 14:06:22.3582725 +0000 UTC m=+980.435677301"
}
```

-------------------------------------------------------

```cmd
/v1/person
```

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

Retorna el ID del nuevo POST.

Ejemplo de respuesta:

```golang
8
```

-------------------------------------------------------

## DELETE

```cmd
/v1/person/:id
```

Ejemplo de respuesta:

```golang
OK
```

-------------------------------------------------------

## PUT

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

Ejemplo de respuesta:

```golang
OK
```
