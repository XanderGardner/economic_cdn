# economic_cdn

Description here

# run main code

run origin server
```
go run origin_server/main.go receive_port
```
run server level 2
```
go run server_level2/main.go receive_port send_port
```
run server level 1
```
go run server_level1/main.go receive_port send_port
```
run user
```
go run user/main.go send_port text_file_path
```

# example run

```
go run origin_server/main.go 8086
go run server_level2/main.go 8084 8086
go run server_level2/main.go 8085 8086
go run server_level1/main.go 8080 8084
go run server_level1/main.go 8081 8084
go run server_level1/main.go 8082 8085
go run server_level1/main.go 8083 8085
go run user/main.go 8080 ./user/books/bacteria_wiki.txt
go run user/main.go 8080 ./user/books/bumble_wiki.txt
go run user/main.go 8081 ./user/books/godzilla_wiki.txt
go run user/main.go 8081 ./user/books/mammal_wiki.txt
go run user/main.go 8082 ./user/books/moby_dick.txt
go run user/main.go 8082 ./user/books/moby_thick.txt
go run user/main.go 8083 ./user/books/moby_wiki.txt
go run user/main.go 8083 ./user/books/starwars_wiki.txt
```

# run testing

hyerbolic cache testing
```
go test ./hyperbolic_cache
```

queue based caches testing
```
go test ./queue_caches
```
