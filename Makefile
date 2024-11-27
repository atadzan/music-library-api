create-migration:
	migrate create -ext sql -dir db/migrations -tz Local $(name)
migrate:
	migrate -database 'postgres://pq_admin:strongPassword@localhost:6432/pq_db?sslmode=disable' -path db/migrations up
dependencies-up:
	docker-compose -f ./deployment/dependencies.yaml up -d
generate-docs:
	rm -rf docs && swag init -d ./cmd -o ./docs