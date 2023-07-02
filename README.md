# Inventory - Kreasitex Jakarta
**Robert**

### Installation :
##### 1. Go mod vendor
```sh
go mod vendor
go mod tidy
```

### Library :
##### 1. Gin Gonic
```sh
go get github.com/gin-gonic/gin
```

##### 2. postgres
```sh
go get github.com/lib/pq
go get gorm.io/driver/postgres
```
##### 3. sqlc
```sh
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
```
##### 4. middlewares
```sh
go get github.com/go-playground/validator/v10
go get github.com/golang-jwt/jwt/v4
go get golang.org/x/crypto
```

### Running :
```sh
go run main.go
```

### Postman :
[![Run in Postman](https://run.pstmn.io/button.svg)]()