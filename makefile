.PHONY: help up down restart stop

up:
	docker-compose -f docker/docker-compose.yml up -d

down:
	docker-compose -f docker/docker-compose.yml down

restart:
	docker-compose -f docker/docker-compose.yml restart

stop:
	docker-compose -f docker/docker-compose.yml stop
