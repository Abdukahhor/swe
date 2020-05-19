# swe
Sample project in clean architecture (Golang)

In project is used grpc (https://grpc.io) and for database BadgerDB is 
an embeddable, persistent and fast key-value (KV) database (https://github.com/dgraph-io/badger)



The first API method is GetNumber, it returns a number. At first this number is zero.

The second API method is IncrementNumber; it increments this number so that when
the next call to GetNumber will return one more.

The third SetSettings method is the settings method. Settings you need to do the following:
The size of the increment (so that it does not increase by one, but by two, three or a thousand).
The increment size must be positive.
The max number of increment. If set it to 1000, then when incrementing
up to 1000 the number is reset to zero by itself.

The number and settings must survive restarting the application: if have incremented to 10,
then the next time the server starts, GetNumber should immediately return 10, i.e. should be
persistent storage. The storage should be any OS database that can survive
turning off the power of the server, but not just a file, since the file will not allow to grow
application functionality.

```
RPC API

First need to set the settings and returns the increment id, and can edit the settings
SetSettings(ctx context.Context, in *Request) (*Response, error)

Returns a number by increment id
GetNumber(ctx context.Context, in *Request) (*Response, error)

Increments the number by increment id by the size of the increment.
IncrementNumber(ctx context.Context, in *Request) (*Response, error) 

```


Example of grpc client 

```Go

package main

import (
	"context"
	"fmt"

	pb "github.com/abdukahhor/swe/handlers/rpcpb"
	"google.golang.org/grpc"
)

func main() {
	c, err := grpc.Dial(":8087", grpc.WithInsecure())
	defer c.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	client := pb.NewIncrementClient(c)
	ctx := context.Background()

	//SetSettings - sets settings of increment, size of increment, 
	//max number of increment,
	//returns id increment,
	rsp, err := client.SetSettings(ctx, &pb.Request{Size: 2, Max: 50})
	if err != nil {
		fmt.Println("SetSettings", err)
		return
	}
	if rsp.Code != 1 {
		fmt.Println("SetSettings", rsp.Message)
		return
	}
	fmt.Println("SetSettings, id of increment: ", rsp.Id)

	rsp, _ = client.GetNumber(ctx, &pb.Request{Id: rsp.Id})
	if rsp.Code != 1 {
		fmt.Println("GetNumber, error message: ", rsp.Message)
		return
	}
	fmt.Println("GetNumber:", rsp.Number)

	rsp, _ = client.IncrementNumber(ctx, &pb.Request{Id: rsp.Id})
	if rsp.Code != 1 {
		fmt.Println("IncrementNumber, error message", rsp.Message)
		return
	}
	if rsp.Code == 1 {
		fmt.Println("IncrementNumber, successfully increased")
	}

	rsp, _ = client.GetNumber(ctx, &pb.Request{Id: rsp.Id})
	if rsp.Code != 1 {
		fmt.Println("GetNumber, error message: ", rsp.Message)
		return
	}
	fmt.Println("GetNumber:", rsp.Number)

	rsp, _ = client.IncrementNumber(ctx, &pb.Request{Id: rsp.Id})
	if rsp.Code != 1 {
		fmt.Println("IncrementNumber, error message", rsp.Message)
		return
	}
	if rsp.Code == 1 {
		fmt.Println("IncrementNumber, successfully increased")
	}

	rsp, _ = client.GetNumber(ctx, &pb.Request{Id: rsp.Id})
	if rsp.Code != 1 {
		fmt.Println("GetNumber, error message: ", rsp.Message)
		return
	}
	fmt.Println("GetNumber:", rsp.Number)

	rsp, _ = client.GetNumber(ctx, &pb.Request{})
	if rsp.Code != 1 {
		fmt.Println("GetNumber, error message: ", rsp.Message)
		return
	}
	fmt.Println("GetNumber:", rsp.Number)
}


```
