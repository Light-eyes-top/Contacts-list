# Параметры подключения к БД из конфига
DBUSER := $(shell yq '.postgres.user' etc/config.yml)
DBPASSWORD := $(shell yq '.postgres.password' etc/config.yml)
DBHOST := $(shell yq '.postgres.host' etc/config.yml)
DBPORT := $(shell yq '.postgres.port' etc/config.yml)
DBNAME := $(shell yq '.postgres.dbname' etc/config.yml)
SSLMODE := $(shell yq '.postgres.sslmode' etc/config.yml)

default: help

# Зависимости
dep:
	go mod tidy
	go mod vendor
.PHONY: dep

# FMT & GOIMPORT
fmt:
	go fmt ./... && goimports -w .
.PHONY: fmt

# Инициализация утилиты для миграций
mig-init:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/mikefarah/yq/v4@latest
.PHONY: mig-init

# Миграции вперед до конца
mig-up:
	migrate -source file://migrations -database 'postgres://$(DBUSER):$(DBPASSWORD)@$(DBHOST):$(DBPORT)/$(DBNAME)?sslmode=$(SSLMODE)' up
.PHONY: mig-up

# Миграции назад до конца
mig-down:
	migrate -source file://migrations -database 'postgres://$(DBUSER):$(DBPASSWORD)@$(DBHOST):$(DBPORT)/$(DBNAME)?sslmode=$(SSLMODE)' down
.PHONY: mig-down

# Генерация protobuf
pb-gen:
	protoc --go-grpc_out=internal/handlers/hendler_gRPC --go_out=internal/handlers/hendler_gRPC internal/handlers/hendler_gRPC/proto/contact.proto
.PHONY: pb-gen

# Help
h:
	@echo "Usage: make [target]"
	@echo "  target is:"
	@echo "    dep			- Исправление зависимостей"
	@echo "    fmt			- Форматирование кодовой базы"
	@echo "    mig-init		- Инициализация утилиты для миграций"
	@echo "    mig-down		- Миграция назад до конца"
	@echo "    mig-up		- Миграция вперёд до конца"
	@echo "    pb-gen       - Генерация protobuf"

.PHONY: h
help: h
.PHONY: help
