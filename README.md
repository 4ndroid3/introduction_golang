# Creacion del modulo para correr GO.
Siempre se usa una URL, suele ser la url del proyecto en GitHub.
- go mod init example.com/greetings
Este comando va a crear el archivo go.mod
Dentro de go.mod es donde van a ir las dependencias de nuestro proyecto

# Packages
Todo programa de Go está hecho con 'packages'
Cualquier programa se corre desde el package principal main.

# Variables exportables
En Go, si una variable empieza con mayusculas es accesible para todo el programa, si la variable
empieza en minusculas solo se puede acceder a ella en el modulo donde se la declara.

