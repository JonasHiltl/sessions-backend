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