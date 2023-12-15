# economic_cdn

We implement a Multi-Layered CDN with Hyperbolic Caching and a Health Monitoring System.

# run cdn code

run origin server
```
go run origin_server/main.go receive_port origin_stats_port server_name
```
run origin server stats
```
go run origin_server_stats/main.go receive_port
```
run server level 2
(cache_number decides the cache: 1 for lru, 2 for fifo, 3 for hyperbolic)
```
go run server_level2/main.go receive_port origin_server_port origin_stats_port server_name cache_number
```
run server level 1
```
go run server_level1/main.go receive_port level2_port origin_stats_port server_name cache_number
```
run user
```
go run user/main.go level1_port text_file_path
```

# example cdn run

This is an example run of a 2 layer CDN with hyperbolic caching
```
go run origin_server_stats/main.go 8087 
go run origin_server/main.go 8086 8087 origin
go run server_level2/main.go 8084 8086 8087 l2_server1 3  
go run server_level2/main.go 8085 8086 8087 l2_server2 3 
go run server_level1/main.go 8080 8084 8087 l1_server1 3 
go run server_level1/main.go 8081 8084 8087 l1_server2 3 
go run server_level1/main.go 8082 8085 8087 l1_server3 3 
go run server_level1/main.go 8083 8085 8087 l1_server4 3 
go run user/main.go 8080 ./user/books/bacteria_wiki.txt
go run user/main.go 8080 ./user/books/bumble_wiki.txt
go run user/main.go 8081 ./user/books/godzilla_wiki.txt
go run user/main.go 8081 ./user/books/mammal_wiki.txt
go run user/main.go 8082 ./user/books/moby_dick.txt
go run user/main.go 8082 ./user/books/moby_thick.txt
go run user/main.go 8083 ./user/books/moby_wiki.txt
go run user/main.go 8083 ./user/books/starwars_wiki.txt
```

This is an example run of a 1 layer CDN with hyperbolic caching
```
go run origin_server_stats/main.go 8087 
go run origin_server/main.go 8086 8087 origin
go run server_level2/main.go 8084 8086 8087 l2_server1 3  
go run server_level2/main.go 8085 8086 8087 l2_server2 3 
go run server_level1/main.go 8080 8086 8087 l1_server1 3 
go run server_level1/main.go 8081 8086 8087 l1_server2 3 
go run server_level1/main.go 8082 8086 8087 l1_server3 3 
go run server_level1/main.go 8083 8086 8087 l1_server4 3 
go run user/main.go 8080 ./user/books/bacteria_wiki.txt
go run user/main.go 8084 ./user/books/bumble_wiki.txt
go run user/main.go 8084 ./user/books/godzilla_wiki.txt
go run user/main.go 8081 ./user/books/mammal_wiki.txt
go run user/main.go 8082 ./user/books/moby_dick.txt
go run user/main.go 8085 ./user/books/moby_thick.txt
go run user/main.go 8085 ./user/books/moby_wiki.txt
go run user/main.go 8083 ./user/books/starwars_wiki.txt
```

Running with FIFO on 2 Layer CDN result in about 
- origin: 13 requests per second
- level 2: 7 requests per second
- level 1: 7 requests per second

Running with LRU on 2 Layer CDN result in about 
- origin: 12 requests per second
- level 2: 6.5 requests per second
- level 1: 7.5 requests per second

Running with Hyperbolic on 2 Layer CDN result in about 
- origin: 13 requests per second
- level 2: 7 requests per second
- level 1: 7 requests per second

Running with Hyperbolic on 1 Layer CDN results in about
- origin: 20 requests per second
- level 2: 4 requests per second
- level 1: 4 requests per second

# run cache performance evaluation
```
go run ./cache_performance/main.go file_path
```
Running for each example text input gives:
```
economic_cdn main % go run ./cache_performance/main.go ./user/books/bacteria_wiki.txt
LRU Hit Rate: 0.657841
Fifo Hit Rate: 0.624422
Hyperbolic Hit Rate: 0.711183
economic_cdn main % go run ./cache_performance/main.go ./user/books/bumble_wiki.txt  
LRU Hit Rate: 0.655849
Fifo Hit Rate: 0.626415
Hyperbolic Hit Rate: 0.684075
economic_cdn main % go run ./cache_performance/main.go ./user/books/godzilla_wiki.txt 
LRU Hit Rate: 0.623461
Fifo Hit Rate: 0.603841
Hyperbolic Hit Rate: 0.643081
economic_cdn main % go run ./cache_performance/main.go ./user/books/mammal_wiki.txt  
LRU Hit Rate: 0.647523
Fifo Hit Rate: 0.615518
Hyperbolic Hit Rate: 0.717130
economic_cdn main % go run ./cache_performance/main.go ./user/books/moby_dick.txt  
LRU Hit Rate: 0.664141
Fifo Hit Rate: 0.625730
Hyperbolic Hit Rate: 0.742928
economic_cdn main % go run ./cache_performance/main.go ./user/books/moby_thick.txt 
LRU Hit Rate: 0.644812
Fifo Hit Rate: 0.606336
Hyperbolic Hit Rate: 0.713223
economic_cdn main % go run ./cache_performance/main.go ./user/books/moby_wiki.txt 
LRU Hit Rate: 0.652166
Fifo Hit Rate: 0.621065
Hyperbolic Hit Rate: 0.707789
economic_cdn main % go run ./cache_performance/main.go ./user/books/starwars_wiki.txt 
LRU Hit Rate: 0.671937
Fifo Hit Rate: 0.644964
Hyperbolic Hit Rate: 0.705554
```

# unit testing

run cache testing
```
go test ./caches
```
