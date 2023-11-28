# Phony Targets
.PHONY: all build test install-tools lint run run-debug invoke

# Meta Targets
all: install-tools build test lint

# Build Targets
build: go-mod-tidy go-build

# Build and run the appliction
build-and-run:
	@echo ">>>>> Starting app"
	@go mod tidy -go=1.21 && go build -o critic && ./critic

go-mod-tidy:
	@echo "🧹 Running go mod tidy"
	@go mod tidy -go=1.21

go-build:
	@echo "🔨 Building Go binaries"
	@go build -ldflags="-s -w" ./...

# Test Targets
test: test-basic

test-basic:
	@echo "🧪 Running tests"
	@go test -cover ./...

test-verbose:
	@echo "📝 Running tests with verbose output"
	@go test -v -cover ./...

test-race:
	CGO_ENABLED=1 go test -race -cover ./...

# Tooling
install-tools:
	@echo "🛠️ Installing tools"
	@go install mvdan.cc/gofumpt@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Linting
lint: lint-golangci lint-fumpt lint-markdown

lint-fumpt:
	@echo "🧹 Running gofumpt linter"
	@gofumpt -l -w .

lint-golangci:
	@echo "🐳 Running golangci linters"
	@golangci-lint run

lint-go:
	@echo "🐳 Running Go linters in Docker"
	@docker run -t --rm -v $$(pwd):/app -w /app golangci/golangci-lint:v1.54.2 golangci-lint run -v \
		-E bodyclose \
		-E exportloopref \
		-E forcetypeassert \
		-E goconst \
		-E gocritic \
		-E misspell \
		-E noctx \
		-E nolintlint \
		-E prealloc \
		-E predeclared \
		-E reassign \
		-E sqlclosecheck \
		-E stylecheck \
		-E varnamelen \
		-E wastedassign \
		-E staticcheck

lint-markdown:
	@echo "📚 Running Markdown linters with npm"
	@if [ -z $$(which markdownlint) ]; then npm install -g markdownlint-cli; fi
	@markdownlint $$(find ./. -name '*.md')
