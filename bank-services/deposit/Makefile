APP_NAME=deposit
PORT=8081

.PHONY: build run docker-build docker-run clean

build:
	go build -o $(APP_NAME) .

run:
	go run main.go

docker-build:
	docker build -t $(APP_NAME):latest .

docker-run:
	docker run -p $(PORT):8081 $(APP_NAME):latest

clean:
	rm -rf $(APP_NAME)