.PHONY: test tools gen-db gen-api
tools:
	@go install github.com/mfridman/tparse@latest
	@go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	@go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

test:
	@echo "Running tests..."
	@go test ./... -json -cover -count 1 | tparse -notests  --progress --pass

gen-db:
	@echo "Running migrations..."
	@go run . database migrate
	@echo "Generating queries..."
	@go tool sqlc generate --file ./etc/db/sqlc.yaml

gen-api:
	@echo "Generating API..."
	@go tool oapi-codegen -config ./etc/api/server.yaml ./etc/api/spec.jsonc

gen-ui:
	@echo "Generating UI..."
	@cd ui && pnpm generate