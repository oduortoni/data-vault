# Root Makefile for both Rust (dv) and Go (go-dv)

# === Rust ===

RUST_DIR = dv
RUST_BINARY = $(RUST_DIR)/target/debug/dv

.PHONY: build-rust run-rust clean-rust

build-rust:
	cargo build --manifest-path=$(RUST_DIR)/Cargo.toml

run-rust:
	cargo run --manifest-path=$(RUST_DIR)/Cargo.toml

clean-rust:
	cargo clean --manifest-path=$(RUST_DIR)/Cargo.toml


# === Go ===

GO_DIR = go-dv
GO_MAIN = ./cmd/main.go
GO_BINARY = $(GO_DIR)/main

.PHONY: build-go run-go clean-go

GO_BINARY = main

build-go:
	@cd $(GO_DIR) && go build -o $(GO_BINARY) $(GO_MAIN)

run-go: build-go
	@./$(GO_DIR)/$(GO_BINARY)

clean-go:
	rm -f $(GO_DIR)/$(GO_BINARY)

# === Combined Convenience Targets ===

.PHONY: build run clean

build: build-rust build-go

run: run-rust run-go

clean: clean-rust clean-go
