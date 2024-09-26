info_log = "/logs/info-log.log"
error_log = "/logs/error-log.log"

source = ./main.go
dest = ./go_app


swag-update:
	@/home/danil/Documents/Golang/bin/swag init

swag-fmt:
	@/home/danil/Documents/Golang/bin/swag fmt


dev: swag-update swag-fmt
	go run $(source) dev

migrate:
	go run $(source) migrate

compile:
	go build -o $(dest) $(source)

prod:
	@echo "Running migrations..."
	/root/$(dest) migrate
	@echo "Running main app..."
	/root/$(dest) >> $(info_log) 2>> $(error_log)

