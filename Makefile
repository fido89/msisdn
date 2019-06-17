# Build the container
build:
	docker-compose build --no-cache msisdn
	docker-compose run msisdn
	docker build -t msisdn .

run:
	docker run -i -t --rm -p=8080:8080 --name="msisdn" msisdn

up:
	docker-compose up --build msisdn

stop:
	docker stop msisdn