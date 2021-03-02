.PHONY: run
run:build
	./grinforum

.PHONY: build
build: mod
	go build -o grinforum main.go

.PHONY: mod
mod:
	git add .
	go mod tidy

.PHONY: migrate
migrate: build
	./grinforum migrate

.PHONY: rollback
rollback: build
	./grinforum migrate rollbackLast