package main

import (
	"context"
	"fmt"
	"time"
)

func Master(ctx context.Context, name string) {
	go Trainee(ctx, "T1")
	go Trainee(ctx, "T2")
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, " Master Job Done")
			return
		default:
			fmt.Println("Master ", name, " still working")
			time.Sleep(1 * time.Second)
		}
	}
}

func Trainee(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, " Trainee Job Done")
			return
		default:
			fmt.Println("Trainee ", name, " still working")
			time.Sleep(3 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go Master(ctx, "M1")
	time.Sleep(5 * time.Second)
	cancel()

}
