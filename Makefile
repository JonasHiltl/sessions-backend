migrate:
	SCYLLA_HOSTS=localhost SCYLLA_KEYSPACE=sessions go run packages/cqlx/migration/main.go

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

bindata-party:
	cd services/party/repository/migrations; go-bindata -pkg migrations .

# You can also view the base64 in the terminal through this command:
# cat comment_descriptor.pb | base64
encode-descriptors:
	cd packages/grpc/utils/encode_descriptors; ./encode_descriptors

# minikube ssh