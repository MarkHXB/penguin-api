APP = penguin-app
BIN_LINUX = bin/$(APP)
CMD_SRC = cmd/$(APP)/main.go
APP_START = bin/$(APP)

build:
	go build -o $(BIN_LINUX) $(CMD_SRC)

run: build
	$(APP_START) --addr 127.0.0.1 --port 8081

clean:
	rm -rf bin/