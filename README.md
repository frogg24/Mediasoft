# REST API для управления группами людей

REST API сервис для управления группами людей и участниками этих групп.

Проект позволяет создавать группы, задавать иерархию групп, добавлять людей в группы, менять данные людей и получать списки участников как в конкретной группе, так и с учётом дочерних групп.

## Функциональность

### Группы

- создание группы;
- получение группы по ID;
- получение списка групп;
- обновление группы;
- удаление группы;
- поддержка дочерних групп.

### Люди

- создание человека;
- получение человека по ID;
- получение списка людей;
- обновление данных человека;
- изменение группы человека;
- удаление человека.

# Полезные команды

**Запустить проект**
```
docker compose up --build
```

После запуска API будет доступно по адресу: ```http://localhost:8080```
PostgreSQL будет доступен на порту: ```localhost:12323```

**Остановить проект**
```
docker compose down
```

**Остановить проект и удалить данные БД**
```
docker compose down -v
```

# Функции

**Создание группы**
```
curl -X POST http://localhost:8080/groups \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Backend",
    "parentgroup": null
  }'
```

**Создание дочерней группы**
```
curl -X POST http://localhost:8080/groups \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Go Developers",
    "parentgroup": 1
  }'
```

**Получить список групп**
```
curl http://localhost:8080/groups

```
**Получить группу по ID**
```
curl http://localhost:8080/groups/1
```

**Обновить группу**
```
curl -X PUT http://localhost:8080/groups/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Development Department",
    "parentgroup": null
  }'
```

**Удалить группу**
```
curl -X DELETE http://localhost:8080/groups/1
```

**Создать человека**
```
curl -X POST http://localhost:8080/persons \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Ivan",
    "lastname": "Ivanov",
    "birthdate": "2000-01-01T00:00:00Z",
    "groupid": 1
  }'
```

**Получить список людей**
```
curl http://localhost:8080/persons
```

**Получить человека по ID**
```
curl http://localhost:8080/persons/1
```

**Обновить человека**
```
curl -X PUT http://localhost:8080/persons/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Petr",
    "lastname": "Petrov",
    "birthdate": "1999-05-10T00:00:00Z",
    "groupid": 2
  }'
```

**Удалить человека**
```
curl -X DELETE http://localhost:8080/persons/1
```

**Получить людей только из конкретной группы**
```
curl http://localhost:8080/groupspersons/1
```

**Получить людей из группы и всех дочерних групп**
```
curl http://localhost:8080/allgroupspersons/1
```

**Получить количество людей только в конкретной группе**
```
curl http://localhost:8080/countgroupspersons/1
```

**Получить количество людей в группе и дочерних группах**
```
curl http://localhost:8080/countallgroupspersons/1
```