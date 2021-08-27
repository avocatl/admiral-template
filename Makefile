test: 
	@docker-compose run --rm tester
.PHONY: test

build:
	@docker-compose build runner
.PHONY: build