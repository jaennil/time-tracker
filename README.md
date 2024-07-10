# Time Tracker API

### [technical specifications](technical_specifications.pdf)

## Build and run
1. create `.env` file in the project root directory by [.env.example](.env.example)
2. `go mod tidy`
3. `go run ./cmd/api/timetracker.go`

## Documentation
- Swagger:
    - http://localhost:8081/swagger/index.html

        or at another port specified in .env file

## Features
### Users
- create user:
    ```http
    POST /users
    {
        "passportNumber": "1234 567890"
    }
    ```
    output:
    ```json
    {
        "address": "г. Москва, ул. Ленина, д. 5, кв. 1",
        "id": 1,
        "name": "Иван",
        "passport_number": "567890",
        "passport_series": "1234",
        "patronymic": "Иванович",
        "surname": "Иванов"
    }
    ```
- delete user:
    ```http
    DELETE /users/:id
    ```
- update user:
    ```http
    PATCH /users/:id
    {
        "name": "Ivan",
        "surname": "Ivanov",
        "patronymic": "Ivanovich",
        "address": "г. East Lorine, ул. Jamey Extension, д. 93, кв. 54",
        "passport_series": "1234",
        "passport_number": "567890",
    }
    ```
- get users:
    ```http
    GET /users
    ```
    - pagination:
        - page
        - page_size

    - filtering:
        - id
        - name
        - surname
        - patronymic
        - address
        - passport_series
        - passport_number
### Tasks
- start task:
    ```http
    POST /tasks/start
    {
        "name": "do stuff",
        "user_id": 1
    }
    ```
- end task:
    ```http
    POST /tasks/end
    {
        "task_id": 1,
        "user_id": 1
    }
    ```
- get user activity:
    ```http
    GET /tasks/activity/:user_id
    ```
    parameters:
    - user_id
    - start_time (Ex. 2024-01-02T15:04:05Z)
    - end_time (Ex. 2025-01-02T15:04:05Z)

    output:
    ```json
    [
      {
        "hours": 1,
        "minutes": 37.830123801,
        "name": "homework"
      },
      {
        "hours": 0,
        "minutes": 12.321313131,
        "name": "cleaning"
      }
    ]
    ```