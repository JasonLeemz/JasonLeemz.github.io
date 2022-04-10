package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	ctx := context.Background()
	//ctx, cancelFnc := context.WithCancel(ctx) // 用于结束正在运行的任务
	//cancelFnc()

	wg := &sync.WaitGroup{}

	ch := make(chan string, 3)

	weather := make([]WeatherResponse, 0)
	for i := 0; i < 3; i++ {
		go func() {
			city := <-ch
			w, err := GetWeather(ctx, city)
			wg.Done()
			if err == nil {
				weather = append(weather, *w)
			}
		}()
	}

	ProductCity(ctx, wg, ch)
	//time.Sleep(5 * time.Second)
	wg.Wait()

	fmt.Println(weather)

	//time.Sleep(10 * time.Second)

}

func ProductCity(ctx context.Context, wg *sync.WaitGroup, ch chan string) {
	ch <- "beijing"
	ch <- "shanghai"
	ch <- "changchun"

	wg.Add(3)
	//ch <- "meihekou"
	//ch <- "tieling"
	//ch <- "shijiazhuang"
	//ch <- "xianggang"
	//ch <- "xinjiapo"

	//return ch
}
