GOPATH:=$(shell go env GOPATH)
GOROOT:=$(shell go env GOROOT)
API_PROTO_FILES=$(shell find protoc -name *.proto)


.PHONY: init
# init
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@v0.6.1
	go install github.com/envoyproxy/protoc-gen-validate@latest


.PHONY: install
# install self package
install:
	go install


.PHONY: api
# make api
api:
	protoc --foo_out=./out --go_out=./out $(API_PROTO_FILES)


# ent
ent:
	go run entgo.io/ent/cmd/ent generate ${ENT_DIR}/schema --template ${ENT_DIR}/template --feature sql/modifier --feature sql/upsert

.PHONY: all
# generate all
all:
	make api;

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
