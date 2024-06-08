BUILD_FLAG=llvm17
GOCMD=go
GOBUILD=$(GOCMD) build -tags $(BUILD_FLAG)
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test -tags $(BUILD_FLAG)
BINARY_NAME=./dist/gvc
MAIN_FILE=gvc.go

antlr:
	java -jar ./tools/antlr-4.13.1-complete.jar -Dlanguage=Go -o internal/parser ./guajavita.g4 
# Build target
build: antlr
	$(GOBUILD) -o $(BINARY_NAME)  $(MAIN_FILE)

# Test target
test: antlr
	$(GOTEST) ./...
# Clean target
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
lint: antlr
	golangci-lint run
# Set phony targets
.PHONY: build clean

