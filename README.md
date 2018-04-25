# dmirou/tribonaccirest

Application ```dmirou/tribonaccirest``` implements Restful Api for calculating tribonacci numbers.

## Run with docker

To build new image use
```
docker build --rm -f <path/to>/tribonaccirest/Dockerfile -t tribonaccirest:latest <path/to>/tribonaccirest
```

To run created image
```
docker run --rm -d -p 8080:8080 --name tribonaccirest tribonaccirest:latest
```

If you want to stop the running container, use the command 
```
docker stop tribonaccirest
```
