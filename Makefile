run-user:
	cd services/user; air

swagger-user:
	cd services/user; swag init --parseDependency --parseDepth 1

swagger-party:
	cd services/party; swag init --parseDependency --parseDepth 1

swagger-story:
	cd services/story; swag init --parseDependency --parseDepth 1

swagger-comment:
	cd services/comment; swag init --parseDependency --parseDepth 1

migrate:
	SCYLLA_HOSTS=localhost SCYLLA_KEYSPACE=sessions go run packages/scylla/migration/main.go

protos-types:
	protoc --go_out=. --go_opt=paths=source_relative packages/types/*.proto

protos-events:
	protoc  \
	--go_out . --go_opt paths=source_relative \
	--go-grpc_out . --go-grpc_opt paths=source_relative  \
	packages/events/*.proto

protos-relation:
	protoc  \
	--go_out . --go_opt paths=source_relative \
	--go-grpc_out . --go-grpc_opt paths=source_relative  \
	packages/grpc/relation/*.proto

protos-party:
	protoc \
	--go_out . --go_opt paths=source_relative \
	--go-grpc_out . --go-grpc_opt paths=source_relative  \
	packages/grpc/party/*.proto

protos-common:
	protoc \
	--go_out . --go_opt paths=source_relative \
	--go-grpc_out . --go-grpc_opt paths=source_relative  \
	packages/grpc/common/*.proto

protos-profile:
	protoc \
	--go_out . --go_opt paths=source_relative \
	--go-grpc_out . --go-grpc_opt paths=source_relative  \
	packages/grpc/profile/*.proto

protos-user:
	protoc \
	--go_out . --go_opt paths=source_relative \
	--go-grpc_out . --go-grpc_opt paths=source_relative  \
	packages/grpc/user/*.proto
		
protos-story:
	protoc \
	--go_out . --go_opt paths=source_relative \
	--go-grpc_out . --go-grpc_opt paths=source_relative  \
	packages/grpc/story/*.proto

protos-comment:
	protoc \
	--go_out . --go_opt paths=source_relative \
	--go-grpc_out . --go-grpc_opt paths=source_relative  \
	packages/grpc/comment/*.proto

protos-auth:
	protoc \
	--go_out . --go_opt paths=source_relative \
	--go-grpc_out . --go-grpc_opt paths=source_relative  \
	packages/grpc/auth/*.proto


# You can also view the base64 in the terminal through this command:
# cat comment_descriptor.pb | base64
encode-descriptors:
	cd packages/grpc/utils/encode_descriptors; ./encode_descriptors

kratos-up:
	docker-compose -f ./services/kratos/quickstart-docker.yml -f ./services/kratos/quickstart-standalone-docker.yml up -d --build

kratos-down:
	docker-compose -f ./services/kratos/quickstart-docker.yml -f ./services/kratos/quickstart-standalone-docker.yml down -v

# minikube ssh