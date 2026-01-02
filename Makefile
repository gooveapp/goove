.DEFAULT_GOAL := build

.PHONY: fmt vet build test clean
fmt:
	go fmt ./...
	@echo "Code formatted."

vet: fmt
	go vet ./...
	@echo "Code vetted."

tailwind: ## Build Tailwind CSS
	npx tailwindcss -i ./static/css/input.css -o ./static/css/output.css --minify

tailwind-watch: ## Watch and rebuild Tailwind CSS
	npx tailwindcss -i ./static/css/input.css -o ./static/css/output.css --watch

templ:
	templ generate
	@echo "Templates generated."

dev: templ ## Run with hot reload
	make tailwind-watch & air
	@echo "Starting development server..."

test: vet ## Run tests
	go test ./...
	@echo "All tests passed."

clean: ## Clean build artifacts
	rm -f goove
	rm -rf tmp/
	find . -name '*_templ.go' -delete
	@echo "Cleaned up build artifacts."

install-dev: ## Install development tools
	go install github.com/air-verse/air@latest
	go install github.com/a-h/templ/cmd/templ@latest
	@echo "Development tools installed."

docker-build: ## Build Docker image
	docker build -t goove:latest .
	@echo "Docker image built."

docker-run: ## Run Docker container
	docker run -d -p 8080:8080 -v goove-data:/data goove:latest
	@echo "Docker container started."

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

