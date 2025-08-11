package main

import (
	"fmt"

	"github.com/Ridaore/ridaore-go/ridaore"
)

func main() {
	opt := ridaore.Options{
		Port: 7000,
		Host: "localhost",
	}

	client := ridaore.New(&opt);
	if err := client.Dial(); err != nil {
		fmt.Println(err.Error())
		return
	}
	
	if err := client.Set("nim:1", "2204735"); err != nil {
		fmt.Println(err.Error())
		return;
	}
	fmt.Println("Setted");

	res, err := client.Get("nim:1");
	if err != nil {
		fmt.Println(err.Error());
		return
	}

	fmt.Println(res);
}