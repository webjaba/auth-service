make create-user:
	grpcurl -plaintext -d '{"user":{"username": "user", "password": "password1"}}' \
	localhost:50011 auth.AuthService/CreateUser