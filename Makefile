
setup:
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/envoyproxy/protoc-gen-validate@v1.0.3-SNAPSHOT.3

executor:
	@make proto PKG=executor

.PHONY: proto
proto:
	@protoc \
		--proto_path=./proto/${PKG} \
		--proto_path=./proto/github.com \
		--go_out=. \
		--go-grpc_out=. \
		--validate_out="lang=go:pkg/${PKG}" --validate_opt="paths=source_relative" \
		proto/${PKG}/*.proto