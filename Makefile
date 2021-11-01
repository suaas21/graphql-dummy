build:
	./build.sh

run: build
	go run ./cmd/graphql-dummy/main.go ./cmd/graphql-dummy/graphql.go serve --config config.yaml

serve:
	docker-compose down
	docker-compose up -d
