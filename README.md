# GormBenchmark
Analysis tool for Go GORM

**Start profiler:**  
```go tool pprof -seconds 30 http://0.0.0.0:8080/debug/pprof/profile```

**Run benchmark:**  
```ab -k -c 8 -n 100000 "http://0.0.0.0:8080/queries_rows?n=50"```
