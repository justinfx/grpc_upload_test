# gRPC File Upload Test

This is a simple test comparing the difference between uploading file data either by a unary
request (all in one go), or via a stream (broken into chunks).

Start server

```
go run main.go
```

Run client

```
go run ./client/main.go <path to some file>
```

Regenerate the proto file, if changed (needs `protoc` compiler and grpc plugins installed):

```
go generate
```

You can adjust the `chunkSize` value in `client/main.go` to test the impact of breaking a large
file into smaller parts during a stream request (default is 5MB chunks).

## Results

Average of 10 runs

| Size | Server Unary | Server Stream | Client Unary | Client Stream |
|---:|---|---|---|---|
| 500B | 35μs | 665μs | 775μs | 2.3ms |
| 10K | 44μs | 1.4ms | 1ms | 3ms |
| 30M | 9ms | 54ms | 90ms | 60ms |
| 300M | 63ms | 621ms | 753ms | 624ms |
