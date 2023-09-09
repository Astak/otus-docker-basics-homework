# OTUS | Основы работы с Docker | Домашнее задание

## Как запустить
```
docker build -t otus-docker-basics-homework .
docker run --rm -d -p 80:8000 --name web-service-gin otus-docker-basics-homework
```

## Как проверить
```
curl http://localhost/health
```