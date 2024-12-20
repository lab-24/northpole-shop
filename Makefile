clean:
	docker compose rm -f
	docker volume rm -f northpole-shop_data-volume

run:
	docker compose up --build

docs:
	swag init -g cmd/api/main.go -o .swagger -ot yaml
	openapi-generator generate -i .swagger/swagger.yaml -g html -o ./docs/openapi

.PHONY: clean docs run
