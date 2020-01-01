GOVERSION=$(shell go version)
THIS_GOOS=$(word 1,$(subst /, ,$(lastword $(GOVERSION))))
THIS_GOARCH=$(word 2,$(subst /, ,$(lastword $(GOVERSION))))
GOOS?=$(THIS_GOOS)
GOARCH?=$(THIS_GOARCH)
DIR_BUILD=build
DIR_SCRIPT=./misc/ops
DIR_CACHE=$(DIR_BUILD)/.cache
APP=$(DIR_BUILD)/$(GOOS)_$(GOARCH)/timeline

default: timeline

.PHONY: \
	timeline \
	clean \

$(DIR_CACHE)/mod:
	mkdir -p $(DIR_CACHE)
	ln -sfn $(GOPATH)/pkg/mod $@

timeline: $(DIR_CACHE)/mod
	go build -v -o $(APP) ./cmd/*.go

clean:
	-rm -rf $(DIR_BUILD)