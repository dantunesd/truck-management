.PHONY: start down unit-test integration-test

start:
	docker-compose up --force-recreate -d

down:
	docker-compose down

unit-test:
	docker-compose -f docker-compose.tests.yml up unit

integration-test: down start wait-start
	docker-compose -f docker-compose.tests.yml up integration

wait-start:
	@while ! docker logs app 2>&1 | grep -q "starting webserver"; do \
		sleep 1; \
		echo "waiting..."; \
	done; \