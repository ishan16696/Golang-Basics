### Step 1

Install protobuf and grpc

```go
brew install protobuf
protoc --version

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### Step 2

- Define a `<fileName>.proto` file for protocol buffers.

1. Define Request/Response

   example:

    ```
    message HelloRequest {
        string name = 1;
    }

    message HelloResponse {
        string message = 1;
        uint32 number = 2;
    }
    ```

2. Defining the gRPC Methods

   ```
   service Hello {
       ...
       rpc SayHello (InputRequest) returns (OutputResponse) {}
       ...
   }
   ```

### Step 3

```
protoc --go_out=. --go-grpc_out=. <fileName>.proto
```

This command would generate the grpc files.

1. [`protoc-gen-go`](https://pkg.go.dev/github.com/golang/protobuf/protoc-gen-go), which generates Go code containing the messages defined in the Protobuf files.
2. [`protoc-gen-go-grpc`](https://pkg.go.dev/google.golang.org/grpc/cmd/protoc-gen-go-grpc), which generates Go bindings for services defined in the Protobuf files
