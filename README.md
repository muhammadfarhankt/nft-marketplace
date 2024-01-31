<h1>NFT Marketplace</h1>
NFT Bidding Platform. Users can mint, buy & sell NFTs.

## Tech Stack

- Go Programming Language
- Clean Architecture (Hexagonal) & Microservices (DDD - Domain Driven Design)
- gRPC & Apache Kafka
- Fiber Framework
- PostgreSQL & MongoDB 
- JWT
- Kubernetes

  
<h2>Required</h2>
<ul>
    <li><a href="https://go.dev/">Golang</a></li>
    <li><a href="https://www.docker.com/">Docker</a></li>
    <li><a href="https://www.postman.com/">Postman</a></li>
    <li><a href="https://code.visualstudio.com/">IDE (Vscode)</a></li>
    <li><a href="https://cloud.google.com/sdk/docs/install">GCP CLI</a></li>
</ul>

<h2>Start PostgreSQLon Docker 🐋</h2>

```bash
docker run --name nft_marketplace_test -e POSTGRES_USER=user -e POSTGRES_PASSWORD=123456 -p 4444:5432 -d postgres:alpine
```

<h2>Execute a container and CREATE a new database</h2>

```bash
docker exec -it nft_marketplace_test bash 
psql -U user
CREATE DATABASE nft_marketplace_test;
\l
```

<h2>Migrate command</h2>

```bash
# Migrate up
migrate -database 'postgres://user:123456@localhost:4444/nft_marketplace_test?sslmode=disable' -source <path> -verbose up 

# Migrate down
migrate -database 'postgres://user:123456@localhost:4444/nft_marketplace_test?sslmode=disable' -source <path> -verbose down 

```
<h2>.env Example</h2>

```bash
APP_HOST=127.0.0.1
APP_PORT=3000
APP_NAME=nft-marketplace
APP_VERSION=v0.1.0
APP_BODY_LIMIT=10490000 //10 MB
APP_API_KEY=pwnYdkPTacwhH2O1
APP_ADMIN_KEY=uKgDUvbpIJ44dvHx
APP_READ_TIMEOUT=60
APP_WRITE_TIMEOUT=60
APP_FILE_LIMIT=2097000 //2 MB
APP_GCP_BUCKET=nft-marketplace-dev-bucket

JWT_API_KEY=JwtApiKeycwhH2O1
JWT_ADMIN_KEY=JwtAdminKeyHxfdeG
JWT_SECRET_KEY=JwtSecretKey1KrA0
JWT_ACCESS_EXPIRES=86400 //1 Day
JWT_REFRESH_EXPIRES=604800 //7 Days

DB_HOST=127.0.0.1
DB_PORT=4444
DB_PROTOCOL=tcp
DB_USERNAME=user
DB_PASSWORD=123456
DB_DATABASE=nft_marketplace_test
DB_SSL_MODE=disable
DB_MAX_CONNECTIONS=25
```

### Run Project
```bash
go run main.go
```
