# website_status_monitor

This tool is a command line program that monitors websites listed in some txt file by checking their HTTP status code.

To use it, pass a .txt file path that contains the list of sites in the initialization command or when the program requests it.

Example of use

```Shell
$ go run main.go sites.txt

Iniciando...
1- Iniciar monitoramento
2- Sair do programa
1

Monitorando... https://www.google.com.br
Site: https://www.google.com.br está online. Status Code: 200


Monitorando... https://www.facebook.com
Site: https://www.facebook.com está online. Status Code: 200


Monitorando... https://twitter.com
Site: https://twitter.com está com problemas. Status Code: 302


Monitorando... https://www.instagram.com
Site: https://www.instagram.com está online. Status Code: 200


1- Iniciar monitoramento
2- Sair do programa
|
```
