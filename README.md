# economic_cdn


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
go run origin_server/main.go 8088
go run server_level2/main.go 8084 8088
go run server_level1/main.go 8081 8084
go run user/main.go 8081 ./user/books/bacteria_wiki.txt


```

# run testing for hyperbolic cache
```
go test ./hyperbolic_cache
```