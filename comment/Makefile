PB = protoc

.PHONY: protoc
protoc:
	$(PB) --proto_path=proto --proto_path=../third_party --go_out=. --go-grpc_out=. --validate_out="lang=go:." proto/comment.proto
	protoc-go-inject-tag -input=proto/pb/comment.pb.go