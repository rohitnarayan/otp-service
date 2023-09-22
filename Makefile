.PHONY: all

all: build fmt test

UNIT_TEST_PACKAGES=$(shell  go list ./... | grep -v "vendor")

APP_EXECUTABLE="out/otp-service"

build:
	mkdir -p out/
	go build -o $(APP_EXECUTABLE) ./*.go

fmt:
	go fmt ./...

db.setup: db.create db.migrate

db.create:
	$(APP_EXECUTABLE) create-db

db.migrate: build
	$(APP_EXECUTABLE) migrate

db.rollback: build
	$(APP_EXECUTABLE) rollback

db.drop:
	$(APP_EXECUTABLE) drop-db

db.reset: db.drop db.create db.migrate

test: db.reset
	@ENVIRONMENT=test go test $(UNIT_TEST_PACKAGES) -race -p=1