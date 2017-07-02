build-parser:
	docker-compose build parser

stop-parser:
	docker-compose stop parser

remove-parser:
	docker-compose rm --all -f parser

up-parser:
	docker-compose up parser
