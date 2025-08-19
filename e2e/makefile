make create-user:
	grpcurl -plaintext -d '{"user":{"username": "user", "password": "password1"}}' \
	localhost:50011 auth.AuthService/CreateUser

make create-token:
	grpcurl -plaintext -d '{"user":{"username": "user", "password": "password1"}}' \
	localhost:50011 auth.AuthService/CreateToken

make recreate-token:
	grpcurl -plaintext -d '{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTU2Mjc1OTEsImlhdCI6MTc1NTYyNjk5MSwic3ViIjoiMjEifQ.Y8BJ9GdY5uqqExohG5QzHpb7KGuBMnJpYR_55VtDo38"}' \
	localhost:50011 auth.AuthService/RecreateToken