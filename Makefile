.PHONY: start down unit-test integration-test

start:
	docker-compose up

down:
	docker-compose down

unit-test:
	docker-compose -f docker-compose.tests.yml up unit

integration-test:
	docker-compose -f docker-compose.tests.yml up integration
