build:
	@docker-compose build

run: build
	@docker-compose up --build

stop:
	@docker-compose down

clean: stop
	@docker-compose rm -f

