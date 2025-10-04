# Переменные
APP_NAME=server
PORT?=8080
BIN_DIR=bin

# Цели
.PHONY: run build test

## run: запускает приложение в режиме разработки
run:
	@echo "Starting server on port $(PORT)..."
	set PORT=$(PORT) && go run ./cmd/$(APP_NAME)

## build: собирает бинарный файл
build:
	@echo "Building $(APP_NAME)..."
	@if not exist "$(BIN_DIR)" mkdir "$(BIN_DIR)"
	go build -o $(BIN_DIR)/$(APP_NAME).exe ./cmd/$(APP_NAME)
	@echo "Binary created: $(BIN_DIR)/$(APP_NAME).exe"

