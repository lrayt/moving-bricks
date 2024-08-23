pb:
	@cd "./dto" && protoc -I=./protobuf/ --go_out=. protobuf/*.proto
	@cd "./dto" && protoc -I=./protobuf --go-grpc_out=. protobuf/*.proto