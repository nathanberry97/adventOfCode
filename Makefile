.DEFAULT_GOAL := explain

.PHONY: explain
explain:
	@echo "Advent of Code"
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage: \033[36m\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: init
init: ## Initialise the project: install node dependencies
	@npm ci

.PHONY: test
test: ## Test a specific project: make test PROJECT=<project-name> DAY=<day>
	@if [ -z "$(PROJECT)" ]; then \
		echo "Error: PROJECT variable not set."; \
		exit 1; \
    fi
	@if [ -z "$(DAY)" ]; then \
		echo "Error: DAY variable not set."; \
		exit 1; \
    fi
	@echo "Testing $(PROJECT) for day $(DAY)"
	@DAY=$(DAY) npx nx test $(PROJECT) -- --reporter=verbose --watch

.PHONY: build
build: ## Build a specific project: make build PROJECT=<project-name>
	@if [ -z "$(PROJECT)" ]; then \
		echo "Error: PROJECT variable not set."; \
		exit 1; \
    fi
	@npx nx build $(PROJECT) --silent > /dev/null 2>&1

.PHONY: run
run: build ## Run a specific project: make run PROJECT=<project-name> DAY=<day>
	@if [ -z "$(PROJECT)" ]; then \
		echo "Error: PROJECT variable not set."; \
		exit 1; \
    fi
	@if [ -z "$(DAY)" ]; then \
		echo "Error: DAY variable not set."; \
		exit 1; \
    fi
	@node build/AoC/$$(echo $(PROJECT) | sed 's/aoc-//g')/main.cjs --day=$(DAY)

.PHONY: generate-new-year
generate-new-year: ## Generate a new year: make new-year YEAR=<year>
	@if [ -z "$(YEAR)" ]; then \
		echo "Error: YEAR variable not set."; \
		exit 1; \
	fi
	@echo "Generating year $(YEAR)"
	@cd tools/generators/new-year && npx tsc --skipLibCheck --esModuleInterop --module commonjs --target ES2020 index.ts
	@npx nx g ./tools/generators:new-year --year=$(YEAR); \
	EXIT_CODE=$$?; \
	rm -f tools/generators/new-year/index.js; \
	exit $$EXIT_CODE

.PHONY: list-aoc
list-aoc: ## List AoC projects in this repo
	@npx nx show projects --type=app

.PHONY: clean
clean: ## Remove build output and Nx cache
	@rm -rf build/
	@npx nx reset
