build:
	export DOCKER_DEFAULT_PLATFORM=linux/amd64 && docker build --tag bandnoticeboard/nottoboard:identity-api-v0.1.1 .

run:
	docker run --add-host=host.docker.internal:host-gateway --env-file ./.env-local bandnoticeboard/nottoboard:identity-api-v0.1.1

push:
	docker push bandnoticeboard/nottoboard:identity-api-v0.1.1

go_run_api:
	go run api.go