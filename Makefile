# Default command

.PHONY: default
default:
	@make check

# Tests

.PHONY: tests
tests:
	@gotip test ./...

TEST_FUNC=^.*$$
ifdef t
TEST_FUNC=$(t)
endif
TEST_PKG=./...
ifdef p
TEST_PKG=./$(p)
endif

.PHONY: test
test:
	@gotip test -timeout 30s -run $(TEST_FUNC) $(TEST_PKG)

# Docs

.PHONY: docs
docs:
	@echo "\033[4mhttp://localhost:9995/pkg/github.com/drykit-go/testx/\033[0m"
	@godoc -http=localhost:9995
