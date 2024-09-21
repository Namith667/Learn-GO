package main

import (
	"context"
	"fmt"
	"log"
	"time"

	//"golang.org/x/text/cases"
)

type Response struct {
	value int
	err error
}

func main(){
	start:=time.Now()
	ctx := context.Background()
	userID:= 10
	val,err:=fetchAPIData(ctx, userID)

	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("value: ",val)
	fmt.Println("Took ",time.Since(start))
}
func fetchAPIData(ctx context.Context,userID int)(int ,error){
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond * 100)
	defer cancel()

	respch:= make(chan Response)

	go func(){
		val,err:=ThirdPartyAPI()
		respch <- Response{
			value: val,
			err: err,
		}
		if err != nil{
			fmt.Errorf("Failed to fetch")
		}
	}()
	for{
		select{
			case <- ctx.Done():
				return 0, fmt.Errorf("fetching from 3rd party api took too long")
			case resp := <- respch:
				return resp.value , resp.err		
		} 
	}
}

func ThirdPartyAPI()(int ,error){
	time.Sleep(time.Millisecond * 500)
	return 200,nil
}