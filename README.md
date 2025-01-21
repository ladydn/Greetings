# Greetings

### Inicializar el modulo
go mod init github.com/ladydn/Greetings

### editar el modulo
go mod edit -replace github.com/ladydn/Greetings=../greetings

### Agregar los modulos faltantes
go mod tidy