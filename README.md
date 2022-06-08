## golang
https://go.dev/tour/welcome/1

### step01 - basic

`go run step01/main.go`
`go build -o app ./step01`

### step02 - call some go packages

`go run step02/main.go`
`go build -o app ./step01`

### step03 - call external packages

`go mod init cloudfield.cz/skol`
`go get github.com/sirupsen/logrus`
`go run step03/main.go`

### step04 - custom module

`go run step04/main.go`

### step05 - structs

`go build -o app ./step05`
`app`

### step06 - threads

`go run step06/main.go`

### step07 - exception handling

`go run step07/main.go`


### step08 - webserver

`go get github.com/gorilla/mux`
`go run step08/main.go`

`curl localhost:9000`
`curl localhost:9000/cats`
`curl localhost:9000/cats/1`

### step09 - webserver + gorm
https://rshipp.com/go-web-api/

`go get github.com/jinzhu/gorm`
`go get github.com/mattn/go-sqlite3`
`go run step09/main.go`

`curl -d 'name=Name+1&description=Desc+1&url=http' -X POST localhost:9000/stars`
`curl -d 'name=Name+2&description=Desc+2&url=http' -X POST localhost:9000/stars`
`curl localhost:9000/stars`
