# website_status_monitor

Command line program that monitors websites listed in some txt file by checking their HTTP status code.

To use it, pass a .txt file path that contains the list of sites in the initialization command or when the program requests it.

Example of use

```Shell
$ go run main.go sites.txt
Iniciando...
1- Iniciar monitoramento
2- Exibir logs
3- Sair do programa

1
Testando... https://www.google.com.br
Site: https://www.google.com.br está online. Status Code: 200

Testando... https://www.facebook.com
Site: https://www.facebook.com está online. Status Code: 200

Testando... https://twitter.com
Site: https://twitter.com está com problemas. Status Code: 302
Ocorreu um erro: Get "/": stopped after 10 redirects

Testando... https://www.instagram.com
Site: https://www.instagram.com está online. Status Code: 200

1- Iniciar monitoramento
2- Exibir logs
3- Sair do programa

2
Exibindo logs...
03/09/2023 14:20:09 - https://www.google.com.br - online
03/09/2023 14:20:09 - https://www.facebook.com - online
03/09/2023 14:20:12 - https://twitter.com - offline
03/09/2023 14:20:13 - https://www.instagram.com - online
03/09/2023 14:20:39 - https://www.google.com.br - online
03/09/2023 14:20:39 - https://www.facebook.com - online
03/09/2023 14:20:42 - https://twitter.com - offline
03/09/2023 14:20:42 - https://www.instagram.com - online
03/09/2023 14:22:22 - https://www.google.com.br - online
03/09/2023 14:22:23 - https://www.facebook.com - online
03/09/2023 14:22:26 - https://twitter.com - offline
03/09/2023 14:22:26 - https://www.instagram.com - online

1- Iniciar monitoramento
2- Exibir logs
3- Sair do programa

|
```
