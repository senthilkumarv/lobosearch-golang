all: deps gen-proto build run

.PHONY: deps
deps:
	$(GOPATH)/bin/govendor sync

.PHONY: gen-proto
gen-proto: deps
	if [ -a $(GOPATH)/src/proto ] ; \
	then \
		rm -rf $(GOPATH)/src/proto ; \
	fi;
	mkdir $(GOPATH)/src/proto
	PATH=$(PATH):$(GOPATH)/bin protoc -I=protobufs --go_out $(GOPATH)/src/proto protobufs/*proto

.PHONY: build
build: gen-proto
	go build -o search *go

.PHONY: run
run: build
	POOL_SIZE=20 ./search

.PHONY: dev-run
dev-run: build
	DATABASE_HOST=127.0.0.1	DATABASE=catalog DATABASE_USER=djlobouser ./search
