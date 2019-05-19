PRACTICE_BUILD_DIR = $(PRACTICE_DIR)/build
PRACTICE_C_DIR = $(PRACTICE_DIR)/c

.PHONY: practice-all
practice-all: practice-format practice-build practice-test

.PHONY: practice-build
practice-build: $(patsubst %.c, %.o, $(wildcard $(PRACTICE_C_DIR)/*.c))

.PHONY: practice-clean
practice-clean:
	rm -f $(PRACTICE_C_DIR)/*.o

.PHONY: practice-format
practice-format:
	@cd $(TOP_DIR)
	@find -name '*.go' | xargs -n 1 go fmt

.PHONY: practice-test
practice-test: practice-test-go
	@find $(PRACTICE_DIR) -name test.sh | sort | while read f; do \
	  echo $$f; \
	  $$f; \
	  echo; \
	done

.PHONY: practice-test-go
practice-test-go:
	@cd $(TOP_DIR)
	@find ./practice -name cc | sort | xargs go test
