# OTUS | Основы работы с Docker | Домашнее задание

## Как запустить локально
```console
#Запустить все
docker-compose up -d

#Запустить Postgres и миграцию
docker-compose up -d migrate

#Запустить Postgres и сервис
docker-compose up -d api

#Запустить только Postgres
docker-compose up -d db

#Запустить Postgres и сервис с пересборкой контейнера
docker-compose up -d --build api

#Запустить сервис не в контейрере
go run . -configpath=config.yml
```
