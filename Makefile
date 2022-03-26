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
	--descriptor_set_out services/party/descriptor.pb \
	--include_imports \
	packages/grpc/party/*.proto

protos-common:
	protoc \
	--proto_path packages/grpc \
	--go_out packages/grpc/ --go_opt paths=source_relative \
	--go-grpc_out packages/grpc/ --go-grpc_opt paths=source_relative  \
	--include_imports \
	packages/grpc/common/*.proto

protos-user:
	protoc \
	--proto_path packages/grpc \
	--go_out packages/grpc/ --go_opt paths=source_relative \
	--go-grpc_out packages/grpc/ --go-grpc_opt paths=source_relative  \
	--descriptor_set_out services/user/descriptor.pb \
	--include_imports \
	packages/grpc/user/*.proto
		
protos-story:
	protoc \
	--proto_path packages/grpc \
	--go_out packages/grpc/ --go_opt paths=source_relative \
	--go-grpc_out packages/grpc/ --go-grpc_opt paths=source_relative  \
	--descriptor_set_out services/story/descriptor.pb \
	--include_imports \
	packages/grpc/story/*.proto

protos-comment:
	protoc \
	--proto_path packages/grpc \
	--go_out packages/grpc/ --go_opt paths=source_relative \
	--go-grpc_out packages/grpc/ --go-grpc_opt paths=source_relative  \
	--descriptor_set_out services/comment/descriptor.pb \
	--include_imports \
	packages/grpc/comment/*.proto


# You can also view the base64 in the terminal through this command:
# cat comment_descriptor.pb | base64
encode-descriptors:
	cd packages/grpc/utils/encode_descriptors; ./encode_descriptors

# minikube ssh