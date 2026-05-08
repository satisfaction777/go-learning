# Task Manager API

REST API для управления задачами. Написан на Go в рамках изучения backend-разработки.

## Стек

- **Go** - основной язык
- **Gin** - HTTP фреймворк
- **PostgreSQL** - база данных
- **GORM** - ORM для работы с базой
- **Docker** - запуск PostgreSQL

## Структура проекта

```
task-manager/
├── main.go          - точка входа, маршруты
├── models/
│   └── task.go      - модель задачи
├── handlers/
│   └── task.go      - обработчики запросов
├── db/
│   └── db.go        - подключение к базе данных
├── go.mod
└── go.sum
```

## Запуск

### 1. Запустить PostgreSQL через Docker

```bash
docker run --name go-postgres \
  -e POSTGRES_USER=admin \
  -e POSTGRES_PASSWORD=secret \
  -e POSTGRES_DB=goapp \
  -p 5432:5432 \
  -d postgres
```

### 2. Установить зависимости

```bash
go mod tidy
```

### 3. Запустить сервер

```bash
go run main.go
```

Сервер запустится на `http://localhost:8080`

## Эндпоинты

| Метод | URL | Описание |
|-------|-----|---------|
| GET | /tasks | Получить все задачи |
| GET | /tasks/:id | Получить задачу по ID |
| POST | /tasks | Создать задачу |
| PUT | /tasks/:id | Обновить задачу |
| DELETE | /tasks/:id | Удалить задачу |

## Примеры запросов

### Создать задачу
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title": "Купить продукты", "description": "Молоко, хлеб"}'
```

Ответ `201`:
```json
{
  "id": 1,
  "title": "Купить продукты",
  "description": "Молоко, хлеб",
  "done": false,
  "created_at": "2026-05-08T12:00:00Z"
}
```

### Получить все задачи
```bash
curl http://localhost:8080/tasks
```

Ответ `200`:
```json
[
  {
    "id": 1,
    "title": "Купить продукты",
    "description": "Молоко, хлеб",
    "done": false,
    "created_at": "2026-05-08T12:00:00Z"
  }
]
```

### Обновить задачу
```bash
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"title": "Купить продукты", "done": true}'
```

Ответ `200`:
```json
{
  "id": 1,
  "title": "Купить продукты",
  "description": "Молоко, хлеб",
  "done": true,
  "created_at": "2026-05-08T12:00:00Z"
}
```

### Удалить задачу
```bash
curl -X DELETE http://localhost:8080/tasks/1
```

Ответ `204` - нет тела.

## Модель задачи

```json
{
  "id": 1,
  "title": "Название задачи",
  "description": "Описание задачи",
  "done": false,
  "created_at": "2026-05-08T12:00:00Z"
}
```

| Поле | Тип | Описание |
|------|-----|---------|
| id | int | Уникальный идентификатор |
| title | string | Название задачи (обязательное) |
| description | string | Описание задачи |
| done | bool | Выполнена ли задача |
| created_at | datetime | Дата создания |
