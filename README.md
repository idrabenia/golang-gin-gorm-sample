# golang-gin-gorm-sample

Golang CRUD application.

Used Libraries and Frameworks:
- [Gin](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/index.html)
- [Wire](https://github.com/google/wire)

To run on Windows OS need to install [GCC](https://jmeubank.github.io/tdm-gcc/download/)

To run application you need:
```shell
go get github.com/google/wire/cmd/wire
go mod download

cd ./src
go generate wire.go
go run .
```
