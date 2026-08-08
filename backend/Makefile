.PHONY: db-up db-down db-reset migrate-up migrate-down run

db-up:
	docker-compose up -d

db-down:
	docker-compose down

db-reset:
	docker-compose down -v
	docker-compose up -d

migrate-up:
	Get-Content migrations/000001_create_users.up.sql | docker exec -i finance-db psql -U postgres -d finance_db
	Get-Content migrations/000002_create_transactions.up.sql | docker exec -i finance-db psql -U postgres -d finance_db

migrate-down:
	Get-Content migrations/000002_create_transactions.down.sql | docker exec -i finance-db psql -U postgres -d finance_db
	Get-Content migrations/000001_create_users.down.sql | docker exec -i finance-db psql -U postgres -d finance_db

run:
	go run cmd/main.go