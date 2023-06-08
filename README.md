# Golang gRPC

## About gRPC
gRPC was created by google. The basic idea is that it makes you server the http to protocol.
It uses this with along another technology called protobuf to make extremely quick and extremely performant RPCs.

#
## Install gRPC
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

export PATH="$PATH:$(go env GOPATH)/bin"
```
#

## About this project
In this tutorial i'm going to be making a very simplestic application, we'll essentially create a server that will allow us to add and mutiply numbers together, then we'll have a client which will serve the server to api.

#
## Proto file.
This will tell Go, on how to encode and decode verious piece of data. This protocal buffers are interesting because they are language agnostic, so you can use them with any language. 

```proto
message Request {
    int64 a = 1;
    int64 b = 2;
}
```
This is the request message, it has two fields, a and b, both of which are int64s. The numbers 1 and 2 are the tags, they are used to identify your fields in the message binary format, and should be unique. Tags from 1 to 15 take one byte to encode, including the identifying number and the field type. Tags from 16 to 2047 take two bytes. So you should reserve the tags 1 to 15 for very frequently occurring message elements. Remember to leave some room for frequently occurring elements that might be added in the future.

```proto
service AddService {
    rpc Add(Request) returns (Response);
}
```
This is the service definition. It defines the RPC methods that can be called remotely with their parameters and return types. To define a method, you specify:

#

## Server
The server struct in main.go file will implement the interface that was generated in the server_grpc.pb.go file.
i.e. AddServiceServer interface.<br /><br />
Our server type need to implement both of the function in the interface, Add and Multiply.
<br /> <br />
Run the main.go file in client and server folder

request method: <b> GET <br/> </b>
<b>END POINT <br />
add: http://localhost:8080/add/1/2 <br />
multiply: http://localhost:8080/multiply/1/2 <br />
</b>


