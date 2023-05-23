# Go-ToDo

## Description

This project is just a small PoC for me to learn how to develop an API using GoLang and Gin Framework

## API Endpoints

### Global Endpoints

- Request

```http
GET /health
```

- Response

```json
{
    "status": "green"
}
```

### Tasks Endpoints

#### Tasks List

- Request

```http
GET /tasks
```

- Response

```json
{
    "tasks": [
        {
            "id": 0,
            "desc": "..."
        },
        {
            "id": 1,
            "desc": "..."
        }
    ]
}
```

#### Create Task

- Request

```http
GET /tasks/create
{
    "desc": "..."
}
```

- Response

```json
{
    "acknowmedge": true
}
```
