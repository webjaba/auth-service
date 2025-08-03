make create-user:
	grpcurl -plaintext -d '{"user":{"login": "user1", "password": "password1"}}' \
	localhost:50001 auth.AuthService/CreateUser