go-auth-api

Este projeto implementa um sistema simples de cadastro e login em Golang, utilizando o framework Gin e armazenando os dados de usuários em memória. A senha é armazenada de forma segura utilizando o algoritmo de hash bcrypt.

Requisitos

•
Go (versão 1.21 ou superior recomendada)

Estrutura do Projeto

O projeto consiste em um único arquivo main.go.

•
main.go: Contém toda a lógica da API, incluindo a definição da estrutura de usuário, armazenamento em memória com mutex para concorrência, e os handlers para os endpoints /register e /login.

Como Executar

1.
Clone o repositório (ou crie o arquivo main.go e go.mod):

2.
Instale as dependências:

3.
Compile e execute a aplicação:

Endpoints da API

1. Cadastro de Usuário

•
URL: /register

•
Método: POST

•
Body (JSON):

•
Respostas:

•
201 Created: Cadastro realizado com sucesso.

•
400 Bad Request: Dados de entrada inválidos.

•
409 Conflict: O nome de usuário já está em uso.



Exemplo de uso com curl:

Bash


curl -X POST http://localhost:8080/register -H "Content-Type: application/json" -d '{"username": "testeuser", "password": "testepassword"}'


2. Login de Usuário

•
URL: /login

•
Método: POST

•
Body (JSON):

•
Respostas:

•
200 OK: Login bem-sucedido.

•
400 Bad Request: Dados de entrada inválidos.

•
401 Unauthorized: Credenciais inválidas (usuário não existe ou senha incorreta).



Exemplo de uso com curl:

Bash


curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{"username": "testeuser", "password": "testepassword"}'


Detalhes da Implementação

•
Armazenamento: Os usuários são armazenados em um mapa Go (map[string]User) em memória. Os dados serão perdidos ao reiniciar a aplicação.

•
Concorrência: Um sync.RWMutex é utilizado para garantir que o acesso ao mapa de usuários seja seguro em ambientes concorrentes (múltiplas requisições simultâneas).

•
Segurança: A senha é hashada usando bcrypt antes de ser armazenada, garantindo que a senha original nunca seja guardada em texto puro.




