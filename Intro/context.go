//store and pass the info accross the different layer of application
//context variable and pass it through the main->router->handler->db function
//the ability to cancellation the job in between of the execution
//consume restful api..and if i am not done with it within 2mins i need to cancel that job
package main

import (
	"context"
	"fmt"
	"time"
)

/* func main() {
	ctx := context.Background()
	//sed some data in ctx
	//seedContext(ctx)
	ctx = seedContext(ctx)
	readCtx(ctx)

} */
func main() {
	ctx := context.Background() //empty context creation
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	go doSomething(ctx)
	select {
	case v := <-ctx.Done():
		fmt.Println("Timeline exceeded of 2 sec", v)
	}
	time.Sleep(time.Second * 3)
}

func doSomething(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timeout")
			err := ctx.Err()
			fmt.Println(err)
			return
		default:
			fmt.Println("Doing something bkwas")
		}
	}
}
func readCtx(ctx context.Context) {
	value := ctx.Value("one")
	fmt.Println(value)
}
func seedContext(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, "one", "111")
	return ctx
}
