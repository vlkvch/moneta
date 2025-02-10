PROG := moneta
SRCDIR := ./cmd/cli
BINDIR := ./bin

all: $(PROG)

$(PROG):
	@go build -o $(BINDIR)/$(PROG) $(SRCDIR)

run: $(PROG)
	@$(BINDIR)/$(PROG)

clean:
	@rm --force $(BINDIR)/$(PROG)
	@rmdir $(BINDIR)

fmt:
	@gofmt -w .

.PHONY: all run clean fmt
