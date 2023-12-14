# economic_cdn

reminder: make sure hyperbolic cache has all functions implemented (and evaluated in our presentation) from the cache interface

Description here

# run main code

run origin server
```
go run origin_server/main.go receive_port
```
run origin server stats
```
go run origin_server_stats/main.go receive_port
```
run server level 2
```
go run server_level2/main.go receive_port origin_server_port origin_stats_port
```
run server level 1
```
go run server_level1/main.go receive_port level2_port origin_stats_port
```
run user
```
go run user/main.go level1_port text_file_path
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

run cache testing
```
go test ./caches
```
