include Makefile.ledger
all: install

install: go.sum
		go install $(BUILD_FLAGS) ./cmd/nsd
		go install $(BUILD_FLAGS) ./cmd/nscli

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify

test:
	@go test  $(PACKAGES)
