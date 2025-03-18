.PHONY: all hot-reload

all: run

hot-reload:
	nodemon --exec go run main.go --signal SIGTERM

run:
	go run main.go
