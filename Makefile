pb-gen:
	mkdir -p pkg/pb_gopuno
	protoc -I=proto --go_out=pkg/pb_gopuno --go-grpc_out=pkg/pb_gopuno --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative proto/gopuno.proto
