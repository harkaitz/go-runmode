.POSIX:
.SUFFIXES:
.PHONY: all clean install check
all:
PROJECT=runmode
VERSION=1.0.0
PREFIX=/usr/local

all:
clean:
install:
check:
## -- BLOCK:go --
build/runmode$(EXE):
	mkdir -p build
	go build -o $@ $(GO_CONF) ./cmd/runmode
all: build/runmode$(EXE)
install: all
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp build/runmode$(EXE) $(DESTDIR)$(PREFIX)/bin
clean:
	rm -f build/runmode$(EXE)
## -- BLOCK:go --
