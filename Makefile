.PHONY: docker-build
docker-build:
	@docker build -t encryption-service -f build/Dockerfile.encryption . \
		&& docker build -t gateway-service -f build/Dockerfile.gateway .

.PHONY: compose-up
compose-up:
	docker-compose -f docker-compose.yml up -d encryption-service  gateway-service && docker-compose logs -f


.PHONY: generate-encryptor-api
generate-encryptor-api:
	protoc -I api/proto --go_out=pkg --go-grpc_out=pkg api/proto/encryptor.proto

.PHONY: start
start: docker-build compose-up

.PHONY: stop
stop:
	docker-compose down --volumes --remove-orphans

