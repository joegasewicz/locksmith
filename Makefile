NAME=locksmith:v1.0.0

build:
	export DOCKER_DEFAULT_PLATFORM=linux/amd64 && docker build --tag bandnoticeboard/$(NAME) .

run:
	docker run --add-host=host.docker.internal:host-gateway --env-file ./.env-local bandnoticeboard/$(NAME)

push:
	docker push bandnoticeboard/$(NAME)

go_run_api:
	go run api.go