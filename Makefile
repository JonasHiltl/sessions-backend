run-user:
	cd services/user; air

swagger-user:
	cd services/user; swag init -g services/user/main.go --parseDependency --parseDepth 1