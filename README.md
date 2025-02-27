### Steps to run this sample:
1) Run a [Temporal service](https://github.com/temporalio/samples-go/tree/main/#how-to-use).
2) Run the following command to start the worker
```bash
go run helloworld/worker/main.go
```
3) Run the following command to start the example
```bash
go run helloworld/starter/main.go
```
### Or run the binary:
1) Run a [Temporal service](https://github.com/temporalio/samples-go/tree/main/#how-to-use).
2) Compile the worker
```bash
go build -o worker worker/main.go
```

3) Compile the starter
```bash
go build -o starter starter/main.go
```

4) Run the worker
```bash
worker/worker
```

5) Run the starter
```bash
starter/starter
```
