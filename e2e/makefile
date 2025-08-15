make create-user:
	grpcurl -plaintext -d '{"user":{"username": "user", "password": "password1"}}' \
	localhost:50011 auth.AuthService/CreateUser

make create-token:
	grpcurl -plaintext -d '{"user":{"username": "user", "password": "password1"}}' \
	localhost:50011 auth.AuthService/CreateToken

make recreate-token:
	grpcurl -plaintext -d '{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTUyNjk5NDksImlhdCI6MTc1NTI2OTM0OSwic3ViIjoiNSJ9.6bFVBBTGOcIgL_BKnPUXQrbdgZPjROSQgAngLSvrjeg"}' \
	localhost:50011 auth.AuthService/RecreateToken