# Build the container
build:
	docker-compose build msisdn

run:
	docker-compose run msisdn

up:
	docker-compose up --build msisdn

stop:
	docker-compose stop