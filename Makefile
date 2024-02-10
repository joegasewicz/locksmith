NAME=locksmith:v1.0.0

build:
	docker build --tag bandnoticeboard/$(NAME) .

run:
	docker run --add-host=host.docker.internal:host-gateway --env-file ./.env-local bandnoticeboard/$(NAME)

push:
	docker push bandnoticeboard/$(NAME)

go_run_api:
	go run api.go

docker-compose:
	docker compose -f docker-compose.local.yaml up