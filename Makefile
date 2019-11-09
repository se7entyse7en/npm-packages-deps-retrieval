.PHONY: proto-gen
proto-gen:
	protoc -I=. --go_out=plugins=grpc:. internal/api/api.proto
