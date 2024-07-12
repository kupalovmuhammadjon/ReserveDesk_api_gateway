CURRENT_DIR := $(shell pwd)
DATABASE_URL="postgres://postgres:root@localhost:5432/reservedesk_auth_service?sslmode=disable"

runserver:
	@go run cmd/main.go

gen-proto:
	@./scripts/gen-proto.sh $(CURRENT_DIR)

tidy:
	@go mod tidy

mig-create:
	@if [ -z "$(name)" ]; then \
		read -p "Enter migration name: " name; \
	fi; \
	migrate create -ext sql -dir migrations -seq $$name

mig-up:
	@migrate -database "$(DATABASE_URL)" -path migrations up

mig-down:
	@migrate -database "$(DATABASE_URL)" -path migrations down

mig-force:
	@if [ -z "$(version)" ]; then \
		read -p "Enter migration version: " version; \
	fi; \
	migrate -database "$(DATABASE_URL)" -path migrations force $$version

permission:

	@chmod +x scripts/gen-proto.sh

swag-gen:
	$(HOME)/go/bin/swag init -g ./api/api.go -o api/docs

