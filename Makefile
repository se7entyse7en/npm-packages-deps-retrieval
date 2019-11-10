.PHONY: proto-gen
proto-gen:
	cd proto && protoc -I=. --go_out=plugins=grpc:../internal/api api.proto
