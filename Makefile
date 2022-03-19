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

proto-comtypes:
	protoc --go_out=. --go_opt=paths=source_relative packages/comtypes/*.proto

protos-party:
	protoc \
	--proto_path packages/grpc \
	--go_out packages/grpc/ --go_opt paths=source_relative \
	--go-grpc_out packages/grpc/ --go-grpc_opt paths=source_relative  \
	--descriptor_set_out packages/grpc/party/descriptor \
	packages/grpc/party/*.proto

protos-common:
	protoc \
	--proto_path packages/grpc \
	--go_out packages/grpc/ --go_opt paths=source_relative \
	--go-grpc_out packages/grpc/ --go-grpc_opt paths=source_relative  \
	--descriptor_set_out packages/grpc/common/descriptor \
	packages/grpc/common/*.proto

protos-user:
	protoc \
	--proto_path packages/grpc \
	--go_out packages/grpc/ --go_opt paths=source_relative \
	--go-grpc_out packages/grpc/ --go-grpc_opt paths=source_relative  \
	--descriptor_set_out packages/grpc/user/descriptor \
	packages/grpc/user/*.proto
		
protos-story:
	protoc \
	--proto_path packages/grpc \
	--go_out packages/grpc/ --go_opt paths=source_relative \
	--go-grpc_out packages/grpc/ --go-grpc_opt paths=source_relative  \
	--descriptor_set_out packages/grpc/story/descriptor \
	packages/grpc/story/*.proto

protos-comment:
	protoc \
	--proto_path packages/grpc \
	--go_out packages/grpc/ --go_opt paths=source_relative \
	--go-grpc_out packages/grpc/ --go-grpc_opt paths=source_relative  \
	--descriptor_set_out packages/grpc/comment/descriptor \
	packages/grpc/comment/*.proto
