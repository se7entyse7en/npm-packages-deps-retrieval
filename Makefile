.PHONY: proto-gen
proto-gen:
	cd proto && protoc -I=. --go_out=plugins=grpc:../internal/api --js_out=import_style=commonjs:../app/src --grpc-web_out=import_style=commonjs,mode=grpcwebtext:../app/src api.proto
	sed -i '' '/GENERATED CODE -- DO NOT EDIT!/a\
/* eslint-disable */' app/src/api_grpc_web_pb.js
	sed -i '' '/GENERATED CODE -- DO NOT EDIT!/a\
/* eslint-disable */' app/src/api_pb.js

.PHONY: docker-build-envoy
docker-build-envoy:
	docker build -t se7entyse7en/npm-pdr-envoy -f docker/envoy/Dockerfile .

.PHONY: docker-build-cmd
docker-build-cmd:
	docker build -t se7entyse7en/npm-pdr-cmd -f docker/cmd/Dockerfile .

.PHONY: docker-build-app
docker-build-app:
	docker build -t se7entyse7en/npm-pdr-app -f docker/app/Dockerfile .

.PHONY: docker-build
docker-build: docker-build-envoy docker-build-cmd docker-build-app

.PHONY: start
start:
	docker-compose up -d mongodb rabbitmq envoy
	@sleep 10
	docker-compose up -d --scale=worker=8 api worker dispatcher
	@sleep 5
	docker-compose up -d app

.PHONY: stop
stop:
	docker-compose down -v
