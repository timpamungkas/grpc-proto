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
protoc:
	protoc --go_opt=module=${GO_MODULE} --go_out=. ./hello/*.proto

.PHONY: build
build: clean protoc