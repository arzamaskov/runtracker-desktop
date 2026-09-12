GOOS := $(shell go env GOOS)

WAILS_FLAGS :=
ifeq ($(GOOS),linux)
WAILS_FLAGS := -tags webkit2_41
endif

.PHONY: setup dev build

setup:
	cd frontend && pnpm install

dev:
	wails dev $(WAILS_FLAGS)

build:
	wails build $(WAILS_FLAGS)
