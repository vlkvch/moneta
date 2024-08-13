PROG = moneta
SRCDIR = ./cmd/cli
BINDIR = ./bin

all: $(PROG)

$(PROG):
	go build -o $(BINDIR)/$(PROG) $(SRCDIR)

clean:
	rm -f $(BINDIR)/$(PROG)
	rmdir $(BINDIR)

.PHONY: all clean
