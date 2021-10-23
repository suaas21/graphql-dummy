build:
	./build.sh

run: build
	go run ./cmd/book-author/main.go ./cmd/book-author/graphql.go serve --config config.yaml

serve:
	docker-compose down
	docker-compose up -d
