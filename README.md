# dmirou/tribonaccirest

Application ```dmirou/tribonaccirest``` implements Restful Api for calculating tribonacci numbers using an [matrix exponentiation](https://www.geeksforgeeks.org/matrix-exponentiation/).

## Run with docker

1. To build new image use
```
docker build --rm -f <path/to>/tribonaccirest/Dockerfile -t tribonaccirest:latest <path/to>/tribonaccirest
```

2. To run created image
```
docker run --rm -d -p 8080:8080 --name tribonaccirest tribonaccirest:latest
```

3. Now you can send request to ```<protocol>://<your-domain>:8080/tribonacci/{n}```, where {n} is positive integer and receive the response.

For example, if your domain is "localhost", you run command
```
curl http://localhost:8080/tribonacci/5
```  
the response will be 
```
{"code":200,"desc":"OK","data":{"n":5,"tribonacci":2}}
```  

## Application restrictions

* The source number "n" must be a positive integer.
* Max calculation time of one tribonacci number is 4 seconds (4000 ms). If the calculation is more than 4 seconds you receive the result with code 460 and empty data. You can change max calculation time by changing the constant ```maxTribonacciCalcTimeInMs``` in ```app.go```
