# Greetings

### Inicializar el modulo

~~~bash
go mod init github.com/ladydn/Greetings
~~~

### editar el modulo

~~~bash
go mod edit -replace github.com/ladydn/Greetings=../greetings
~~~

### Agregar los modulos faltantes

~~~bash
go mod tidy
~~~

## Compilar e instalar la aplicación

~~~bash
$ go build
~~~

~~~bash
$ ./hello
~~~

# Saludos en Go

Este paquete proporciona una forma simple de obtener saludos personalizados en Go.

## Instalación

Ejecuta el siguiente comando para instalar el paquete:

```bash
go get -u github.com/ladydn/greetings
```
## Uso
Aquí tienes un ejemplo de cómo utilizar el paquete en tu código:

```go
package main

import (
    "fmt"
    "github.com/ladydn/greetings"
)

func main() {
    message, err := greetings.Hello("Lady")

    if err != nil {
        fmt.Println("Ocurrió un error:", err)
        return
    }

    fmt.Println(message)
}

```
Este ejemplo importa el paquete github.com/ladydn/greetings y llama a la función Hello para obtener un saludo personalizado. Si ocurre un error, se imprime un mensaje de error.

