NAME=fantasy_banking:v0.0.1

build:
	docker build --tag bandnoticeboard/$(NAME) .

run:
	docker run --add-host=host.docker.internal:host-gateway --env-file ./.env-local bandnoticeboard/$(NAME)

push:
	docker push bandnoticeboard/$(NAME)

docker-compose:
	docker compose -f docker-compose.dev.yaml up