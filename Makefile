clean:
	docker compose rm -f
	docker volume rm -f northpole-shop_data-volume

run:
	docker compose up --build

test:
	go test ./...

docs:
	swag init -g cmd/api/main.go -o docs/openapi -ot yaml
	openapi-generator generate -i docs/openapi/swagger.yaml -g html -o ./docs/openapi

.PHONY: clean docs run
