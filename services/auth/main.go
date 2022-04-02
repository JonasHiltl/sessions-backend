package main

func main()  {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	mongo, err := repository.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer mongo.Client().Disconnect(ctx)

	dao := repository.NewDAO(mongo)

	tokenManager = service.NewTokenManager()

	addr := "0.0.0.0:8081"

	conn, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()

	aServer := handler.NewAuthServer(tokenManager)

}