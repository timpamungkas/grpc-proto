GO_MODULE := github.com/timpamungkas/course-grpc-proto

.PHONY: clean
clean:
ifeq ($(OS), Windows_NT)
	if exist "protogen" rd /s /q protogen
	mkdir protogen\go
else
	rm -fR ./protogen \
	mkdir -p ./protogen/go
endif

.PHONY: protoc-go
protoc-go:
	protoc --go_opt=module=${GO_MODULE} --go_out=. \
	--go-grpc_opt=module=${GO_MODULE} --go-grpc_out=. \
	./hello/*.proto

.PHONY: build
build: clean protoc-go