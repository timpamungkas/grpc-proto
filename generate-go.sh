GO_MODULE := github.com/timpamungkas/course-grpc-proto

# init tools
sudo apt-get install -y protobuf-compiler golang-goprotobuf-dev
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# cleanup
rm -fR ./protogen
mkdir -p ./protogen/go

# build
protoc --go_opt=module=$GO_MODULE --go_out=. \
	--go-grpc_opt=module=$GO_MODULE --go-grpc_out=. \
	./proto/hello/*.proto ./proto/payment/*.proto ./proto/transaction/*.proto