# The Essentials to get Go Micro Working

Much of this is from https://micro.mu/docs/writing-a-go-service.html.

Define a .proto file in which you describe the API to the services that will be needed. Note the package name as that will be used later.

Use protoc to compile your .proto file into go files:

`protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. [your_file].proto`

You can search the generated .go file(s) for "API" to find the function signatures you will need to implement.

## Server

Create a main.go file to implement the server. Be sure to import go-micro:

`import "github.com/micro/go-micro`

Then, initialize a new service:

`service := micro.NewService()`

You can name the service by passing in options:

```Golang
service := micro.NewService(micro.Name("my-service"))
```

Note that the name used above DOES NOT need to match the package name used in the original .proto file. But whatever name you choose here, the client will need to use the same name when searching for the service.

Next, implement the handler function(s). Again, look in the .go generated files (by protoc) to find the interface function descriptions. Then, in the main.go file, do something like:

```Golang
import proto "github.com/micro/examples/service/proto"

type Greeter struct {}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
  rsp.Greeting = "Hello" + req.Name
  return nil
}
```

Register the handler:

```Golang
service := micro.NewService(micro.Name("greeter"))
proto.RegisterGreeterHandler(service.Server(), new(Greeter))
```

Then run the service:

```Golang
if err := service.Run(); err != nil {
  log.Fatal(err)
}
```

## Client

Do something like this to consume the service:

```Golang
// create the greeter client using the service name and client
greeter := proto.NewGreeterService("greeter", service.Client())

// request the Hello method on the Greeter handler
rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{
	Name: "John",
})
if err != nil {
	fmt.Println(err)
	return
}

fmt.Println(rsp.Greeter)
```

Note that the name of the service used in the NewGreeterService call must match the name used by the server code when it names and registers its service.
