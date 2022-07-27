# Курс OTUS Microservice architect
## ДЗ №2 - «Основы работы с Kubernetes (часть 2)»

### TL;DR
1. `make k8s-deploy`
2. `make newman`
3. `make k8s-remove` 

### Подробнее 
```
Сделать простейший RESTful CRUD по созданию, удалению, просмотру и обновлению пользователей.
Пример API - https://app.swaggerhub.com/apis/otus55/users/1.0.0
```
Приложение реализовано, [app/](app), апи полностью взят из примера выше

Доступны 5 стандартных REST-методов: GET, POST, PUT, PATCH, DELETE

Примеры работы с ними - в [коллекции](postman/collection.json)

```
Конфигурация приложения должна хранится в Configmaps.
```


```
Добавить базу данных для приложения.
```
Делал без внешних helm-репозиториев, в своем чарте: [statefulset](helm/chart/templates/postgres_statefulset.yaml) +
[service](helm/chart/templates/postgres_service.yaml)

```
Доступы к БД должны храниться в Secrets.
```
[app_secret.yaml](helm/chart/templates/app_secret.yaml)

```
Первоначальные миграции должны быть оформлены в качестве Job-ы, если это требуется.
```
Не потребовалось, сделаны в виде initContaner в [деплойменте приложения](helm/chart/templates/app_deployment.yaml)

```
Ingress-ы должны также вести на url arch.homework/ (как и в прошлом задании)
```
[Ведут](helm/chart/values.yaml#L14)

```
На выходе должны быть предоставлены:
1. ссылка на github
2. инструкция по запуску приложения.
3. команда установки БД из helm, вместе с файлом values.yaml.
4. команда применения первоначальных миграций
5. команда kubectl apply -f, которая запускает в правильном порядке манифесты кубернетеса
6. Postman коллекция, в которой будут представлены примеры запросов к сервису на создание, получение, изменение и удаление пользователя

+5 баллов за шаблонизацию приложения в helm чартах
```
2, 3, 4, 5 объединены в один вызов helm: `helm upgrade --install -n otus-msa-hw2 otus-msa-hw2 helm/chart`

Он же для удобства обернут в `make k8s-deploy`

Используется неймспейс `otus-msa-hw2`, создается автоматически по команде выше

Тесты прогоняются через `make newman`

`make k8s-remove` удаляет неймспейс 

