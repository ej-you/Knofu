package settings

import (
	"os"

	"github.com/joho/godotenv"
)


// загрузка переменных окружения
var _ error = godotenv.Load("./.env")


// распаковка переменных окружения по переменным
var Port string = os.Getenv("GO_PORT")


// формат логов
var LogFmt string = "[${time_rfc3339}] -- ${status} -- from ${remote_ip} to ${host} (${method} ${uri}) | ${bytes_in} ${bytes_out} | error: ${error} | -- User-Agent: ${user_agent}\n"
