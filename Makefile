env:
	export $(xargs <.env)

build-product: # usage: make build-product VERSION=value
	docker build -t larien/product:$(VERSION) ./product

build-discount: # usage: make build-discount VERSION=value
	docker build -t larien/discount:$(VERSION) ./discount

build:
	docker-compose build

product:
	docker-compose run product

discount:
	docker-compose run discount

db:
	docker-compose run database

migrate:
	docker-compose run migrate

up:
	docker-compose up --remove-orphans

up-detach:
	docker-compose up -d --remove-orphans

down:
	docker-compose down

psql:
	docker exec -it database bash
	# su - postgres
	# psql -p 5432 -d product -U larien