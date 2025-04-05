.PHONY: all clean build solc abigen

# Project directories
CONTRACTS_DIR=./contracts/src
ARTIFACTS_DIR=./contracts/artifacts
BINDINGS_DIR=./backend/pkg/ethereum/contracts
GO_BUILD_DIR=./backend/cmd/server

# Solidity contracts
SOLIDITY_FILES=$(wildcard $(CONTRACTS_DIR)/*.sol)
ABI_FILES=$(patsubst $(CONTRACTS_DIR)/%.sol,$(ARTIFACTS_DIR)/%.abi,$(SOLIDITY_FILES))
BIN_FILES=$(patsubst $(CONTRACTS_DIR)/%.sol,$(ARTIFACTS_DIR)/%.bin,$(SOLIDITY_FILES))

# Go files to generate
GO_BINDINGS=$(patsubst $(CONTRACTS_DIR)/%.sol,$(BINDINGS_DIR)/%.go,$(SOLIDITY_FILES))

# Tools
SOLC=/usr/local/bin/solc
ABIGEN=/home/iv/Projects/go/bin/abigen

# Default target
all: build

# Build everything
build: solc abigen go-build

# Compile Solidity contracts
solc: $(ABI_FILES) $(BIN_FILES)

# Generate .abi and .bin files from Solidity contracts
$(ARTIFACTS_DIR)/%.abi $(ARTIFACTS_DIR)/%.bin: $(CONTRACTS_DIR)/%.sol
	@mkdir -p $(ARTIFACTS_DIR)
	$(SOLC) --abi --bin -o $(ARTIFACTS_DIR) $< --overwrite

# Generate Go bindings
abigen: $(GO_BINDINGS)

# Generate Go binding file for each contract
$(BINDINGS_DIR)/%.go: $(ARTIFACTS_DIR)/%.abi $(ARTIFACTS_DIR)/%.bin
	@mkdir -p $(BINDINGS_DIR)
	$(ABIGEN) --bin=$(ARTIFACTS_DIR)/$*.bin --abi=$(ARTIFACTS_DIR)/$*.abi --pkg=contracts --out=$@

# Build Go application
go-build:
	cd $(GO_BUILD_DIR) && go build -o ../../server

# Clean artifacts
clean:
	rm -rf $(ARTIFACTS_DIR)/*.abi $(ARTIFACTS_DIR)/*.bin $(BINDINGS_DIR)/*.go

# For development: just regenerate bindings and build Go app
rebuild: abigen go-build