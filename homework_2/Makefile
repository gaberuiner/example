build-all:
	cd checkout && GOOS=linux GOARCH=amd64 make build
	cd loms && GOOS=linux GOARCH=amd64 make build
	cd ProductService && GOOS=linux GOARCH=amd64 make build

run-all: build-all
	sudo docker compose up --force-recreate --build
	#docker-compose up --force-recreate --build

precommit:
	cd checkout && make precommit
	cd loms && make precommit
	cd ProductService && make precommit

docker-up:
	docker-compose build
	docker-compose up

migration:
	
	goose -dir ./checkout/migrations postgres "postgres://postgres:123@localhost:5433/checkout?sslmode=disable" up 
	goose -dir ./loms/migrations postgres "postgres://postgres_u:123@localhost:5434/loms?sslmode=disable" up