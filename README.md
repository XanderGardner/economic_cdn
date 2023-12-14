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
go run user/main.go send_port
```

# run testing for hyperbolic cache
```
go test ./hyperbolic_cache
```