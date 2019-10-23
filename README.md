# GormBenchmark
Analysis tool for Go GORM

**Start docker**
```docker-compose up -d```

**Start profiler:**  
```go tool pprof -seconds 60 http://0.0.0.0:8080/debug/pprof/profile```

**Run benchmark:**  
```ab -k -c 8 -n 100000 "http://0.0.0.0:8080/query?n=50"```  
n = number of rows

*At **Profiler** terminal: use command ```web``` to see the graph.
```
$ go tool pprof -seconds 60 http://0.0.0.0:8080/debug/pprof/profile
Fetching profile over HTTP from http://0.0.0.0:8080/debug/pprof/profile?seconds=60
Please wait... (1m0s)
Saved profile in /Users/lftv20190401/pprof/pprof.hello-world.samples.cpu.003.pb.gz
File: hello-world
Type: cpu
Time: Oct 23, 2019 at 2:33pm (+07)
Duration: 1mins, Total samples = 32.72s (54.53%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) web
```
