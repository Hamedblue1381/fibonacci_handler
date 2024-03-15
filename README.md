<h1>Fibonacci Next And Previous Number</h1> 
</br>

Running the Server
To run the server, use the following command:
```bash
make server
```

Endpoints
* /next: Get the next number.
* /prev: Get the previous number.

Request Requirements
* Set Content-Type to either application/json or application/msgpack.
* Set Accept to either application/json or application/msgpack.
* The request must contain a number field.

Example cURL request:
```bash
curl -X POST http://localhost:8080/next -H "Content-Type: application/json" -H "Accept: application/json" -d '{"number": 5}'
```

Testing
To run tests, use the following command:
```bash
make test
```
