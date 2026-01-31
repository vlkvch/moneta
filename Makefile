PROGRAM := moneta
SOURCE_DIRECTORY := ./cmd/cli
BUILD_DIRECTORY := ./build

all: $(PROGRAM)

$(PROGRAM):
	@go build -o $(BUILD_DIRECTORY)/$(PROGRAM) $(SOURCE_DIRECTORY)

run: $(PROGRAM)
	@$(BUILD_DIRECTORY)/$(PROGRAM)

clean:
	@rm -f $(BUILD_DIRECTORY)/$(PROGRAM)
	@rmdir $(BUILD_DIRECTORY)

fmt:
	@gofmt -w .

.PHONY: all run clean fmt
