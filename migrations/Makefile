migration_path = .

migration:
	migrate create -ext sql -dir $(migration_path) -seq $(NAME)

migration-up:
	migrate -database ${POSTGRESQL_URL} -path $(migration_path) up

migration-down:
	migrate -database ${POSTGRESQL_URL} -path $(migration_path) down

migration-force: # usage: make migration-force VERSION=value
	migrate -path $(migration_path) -database ${POSTGRESQL_URL} force $(VERSION)