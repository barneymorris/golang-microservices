# golang-microservices

Just practicing to create own microservices using Golang. Nothing interesting.

There are 4 microservices
1) Frontend
2) Authentication service
3) Broker service (send request to other microservices)
4) Logger microservices

It users Postgres for authenticating and Mongo for logging. Make sure you edit all source files just to get microservices works well.


# How to run

1) Go to project folder
2) run command "docker-compose up -d --build front-end broker-service logger-service authentication-service postgres mongo"
3) Go to http://localhost:80 ant check that is work correct
