
Compile the example.

```
make build
```

Run the Zipkin backend.

```
docker run -d -p 9411:9411 openzipkin/zipkin
```

Run the program

```
./opentracing-go-example
```

Browse to http://localhost:9411


