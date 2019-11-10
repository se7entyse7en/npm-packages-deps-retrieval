.PHONY: proto-gen
proto-gen:
	cd proto && protoc -I=. --go_out=plugins=grpc:../internal/api --js_out=import_style=commonjs:../app/src --grpc-web_out=import_style=commonjs,mode=grpcwebtext:../app/src api.proto
	sed -i '' '/GENERATED CODE -- DO NOT EDIT!/a\
/* eslint-disable */' app/src/api_grpc_web_pb.js
	sed -i '' '/GENERATED CODE -- DO NOT EDIT!/a\
/* eslint-disable */' app/src/api_pb.js
