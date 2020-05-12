# swe
Задача для SWE (Golang)

В проекте используется grpc (https://grpc.io). Для база даных используется BadgerDB is 
an embeddable, persistent and fast key-value (KV) database (https://github.com/dgraph-io/badger)



Первый метод API -- GetNumber, он возвращает число. Сначала это число -- ноль. 

Второй метод API -- IncrementNumber, он инкрементирует это число так, что при 
следующем вызове GetNumber вернет на единицу больше. 

Третий метод SetSettings -- метод настроек. Настройки нужно сделать следующие:
Размер инкремента (чтобы увеличивалось не на один, а на два, три или тысячу). 
Размер инкремента должен быть положительным.
Размер верхней границы. Если я поставлю ее 1000, то при доинкременчивании 
до 1000 число сбрасывается на ноль само.

Число и настройки должны переживать перезапуск приложения: если я доинкрементил до 10, 
то при следующем запуске сервера GetNumber должно сразу возвращать 10, т.е. должно быть 
персистентное хранилище. В качестве хранилища должна быть любая OS БД, способная пережить 
выключение питания сервера, но не просто файл, так как файл не позволит наращивать 
функционал приложения. 

```
RPC API

Сначала нужно установить настройки и получить id инкремента, и можно редактировать настройки
SetSettings(ctx context.Context, in *Request) (*Response, error)

Возвращает число по id инкремента
GetNumber(ctx context.Context, in *Request) (*Response, error)

Увеличивает число по id инкремента на размер инкремента.
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

	//SetSettings - настройка инкремента, size это размер инкремента, 
	//max это размер верхней границы,
	//возвращает id инкремента, и по этой id делаеться инкремент 
	//(IncrementNumber) и возвращает число (GetNumber)
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
