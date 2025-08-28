# Basic go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFMT=$(GOCMD) fmt

test:
	@echo "🚀 Starting unit test..."
	$(GOTEST) ./...

clean:
	@echo "🚀 Cleaning test cache..."
	${GOCLEAN} -testcache

fmt:
	@echo "🚀 Start code formatting..."
	${GOFMT} ./...

race:
	@echo "🚀 Start race detector..."
	${GOTEST} ./... -race -v