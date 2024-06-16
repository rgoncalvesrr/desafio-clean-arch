# Executando a Aplicação

Para executar a aplicação devemos subir os containers Docker; para isso devemos executar, no diretório raíz da aplicação, o comando ```docker-compose up -d```. Esse comando "subirá" o RabbitMQ e o banco de dados MySQL.

Após executar a instrução acima acesse o diretório ```cmd/ordersystem``` com o comando ```cd cmd/ordersystem``` (considero que você no diretório raíz do pojeto) e execute o comando ```go run main.go wire_gen.go```.

Se tudo correr como esperado será apresentado no terminal:
```
Starting web server on port :8000
Starting gRPC server on port 50051
Starting GraphQL server on port 8080
```

# Serviços e Portas

- Porta 8000: servidor Web - API REST
- Porta 50051: servidor gRPC
- Porta 8080: servidor GraphQL
- Porta 3306: MySQL
- Porta 15672: Interface RabbitMQ

# API REST

As requisições web podem ser feitas de várias formas. É possível o cliente [Postman](https://www.postman.com/), por exemplo, mas utilizei u plugin **REST Client** dentro do Visual Studio Código.

Há um arquivo na raíz do projeto chamado orders.http que contém as chamadas para criação e listagem das ordens.

# GraphQL

A aplicação traz um serviço que ouve na porta 8080. Para fazer requisições GraphQL execute o projeto e acesse a URL [GraphQL](http://localhost:8080) em qualquer navegado web.

A seguir um exemplo de entrada de consulta:

A entrada a seguir cria uma ordem:
```
mutation createOrder {
  createOrder(
    input: {
      id: "45",
      Price: 10,
      Tax: 60
    }
  ) {
    id,
    Price,
    Tax,
    FinalPrice
  }
}
```


A entrada a seguir lista os ordens no banco de dados:  
```
query queryOrders {
  listOrders{
    id,
    Price,
    Tax,
    FinalPrice
  } 
}
```

# gRPC

Para executa as instruções gRPC utilizamos o cliente EVANS. Este cliente está disponível no [github.com/ktr0731/evans](https://github.com/ktr0731/evans)

Siga as etapas de configuração descritas no github e depois execute o comando:

```
evans -r repl
```