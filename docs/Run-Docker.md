## Comandos para rodar o docker


```sh

# Dentro do diretorio do projeto

docker build -t api-gin .

# depois

docker run -p 8080:8080 api-gin

```

> Acesse `http://localhost:8080`

- Se tudo estiver certo, a API deve reponser normalmente


