# Knofu

### Пример составления текста ошибки:

```go
import "fmt"
import "errors"

...


var errMessage string = "Some error occured"
var errorStatus int = 400

errorWithCode := fmt.Sprintf("%s||%s", errorStatus, errMessage)
var err error = errors.New(errorWithCode)
```

### Пример ответа с ошибкой:

```json5
{
  "message": "email: Email is not in the right format",
  "path": "/api/user/register",
  "status": 400,
  "timestamp": "24-09-22 22:19:02 +03"
}
```
