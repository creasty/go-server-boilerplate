.DEFAULT_GOAL := all

GO_BUILD_FLAGS := -v -ldflags="-s -w"
GO_TEST_FLAGS  := -v
PACKAGE_DIRS   := $(shell go list ./... 2> /dev/null | grep -v /vendor/)


#  App
#-----------------------------------------------
BIN       := bin/server-boilerplate
SRC_FILES := $(shell find . -name '*.go' -not -path './vendor/*')

$(BIN): $(SRC_FILES)
	@echo "Building $(BIN)"
	@go build $(GO_BUILD_FLAGS) -o $(BIN)


#  Commands
#-----------------------------------------------
all: $(BIN)

.PHONY: clean
clean:
	@rm -rf bin/*

.PHONY: lint
lint:
	@gofmt -e -d -s $(SRC_FILES) | awk '{ e = 1; print $0 } END { if (e) exit(1) }'
	@echo $(SRC_FILES) | xargs -n1 golint -set_exit_status
	@go vet $(PACKAGE_DIRS)

.PHONY: test
test: lint
	@go test $(GO_TEST_FLAGS) $(PACKAGE_DIRS)

.PHONY: run
run: $(BIN)
	@$(BIN)
