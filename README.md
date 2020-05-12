# swe
Задача для SWE (Golang)



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
