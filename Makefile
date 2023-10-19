.PHONY: bench
bench:
	go test --tags=bench -bench . -benchmem -count 6

.PHONY: test-integration
test-integration:
	@echo "Execute all integration test files"
	go test --tags=integration ./tests/integration/...

.PHONY: test
test: test-integration