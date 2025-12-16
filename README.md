# Практическое занятие 12. Подключение Swagger/OpenAPI. Автоматическая генерация документации
# Свидовский М.А. ЭФМО-01-25

## Цель работы

В рамках работы была создана автоматически документация проекта pz11-notes-api при помощи Swagger/OpenAP. Был выбран подход code-first: документацция создается автоматически на основе комментариев "ручек" и последующего запуска Swagger'a (реализован ввод в консоль команды "make swagger"). 
Документация доступна при запуске проекта по "http://localhost:8080/docs/*" с возможностью теста API ("ручек"). 

## Notes API

<img width="1855" height="960" alt="image" src="https://github.com/user-attachments/assets/b5ce0f7d-4c4b-4b20-9ca6-3a34e73aca18" />

### Получение списка заметок

<img width="1282" height="747" alt="image" src="https://github.com/user-attachments/assets/4a23d5d2-5753-4480-9c12-e3080c384b49" />


### Создание заметки

<img width="1279" height="960" alt="image" src="https://github.com/user-attachments/assets/fb5fcda7-dad4-4574-af6f-31f71af5d63d" />

### Получение заметки по id 

<img width="1278" height="897" alt="image" src="https://github.com/user-attachments/assets/2bc8237a-50ef-4746-adaf-78eb6056457c" />

### Редактирование заметки по id 

<img width="1279" height="950" alt="image" src="https://github.com/user-attachments/assets/84da8c2a-cf00-416b-a616-225de70a851a" />

### Удаление заметки по id 

<img width="1273" height="614" alt="image" src="https://github.com/user-attachments/assets/11cc7568-e573-43a5-b9a2-08edce3dc6ae" />

## Запуск

Запуск осуществляется при помощи команд:

``` bash
swag init -g cmd/api/main.go -o docs // формирование документации, использовать при изменении кода

go run ./cmd/api // запуск проекта
```

Либо же при помощи:
``` bash

make swagger // замена первой команды из блока выше

make run // замена второй команды из блока выше
```

## Структура проекта 
<img width="385" height="686" alt="image" src="https://github.com/user-attachments/assets/37c77577-9dbf-4488-9a13-ae6a0864112a" />

## Доп. плюшки

### ReDoc

Был добавлен альтернативный роут с ReDoc ("http://localhost:8080/redoc")

<img width="1858" height="1008" alt="image" src="https://github.com/user-attachments/assets/045134cd-1bbd-407b-9ced-e60cea55e845" />


## Выводы

В ходе выполнения работы была создана документация (Swagger UI), добавлен альтернативный роут для документации (ReDoc) и автоматизированны команды для запуска (make run) и создания документации (make swagger) при помощи makefile.




