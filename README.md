# Knofu


### Пример ответа с ошибкой (`ErrorsDeafultSchema`):

```json5
{
  "status": "error",
  "statusCode": 400,
  "path": "/api/user/register",
  "errors": {
    "lastName": "LastName field must not be blank",
    "password": "Password field must contain at least 8 symbols"
  },
  "timestamp": "24-09-23 10:49:14 +03"
}
```

### Содержимое `.env` файла:

```dotenv
# Go app
GO_PORT=8000
SECRET=example


# MySQL DB
DB_NAME=db_name

DB_USER=db_user
DB_PASSWORD=password

DB_HOST=1.2.3.4
DB_PORT=3306
```
