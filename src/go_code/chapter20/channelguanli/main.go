package main

import (
	_ "context"
	_ "fmt"
	_ "math/rand"
	_ "sync"
	_ "time"
)

/*
并发核心在于协程管理
example 随即填充一个数组求平均数


//waitgroup用于线程同步
var num [100]int

	func main() {
		var group sync.WaitGroup
		group.Add(2)
		//Add，有多少个任务要完成，Done每个任务完成后都调用,Wait等待所有任务完成，在所有任务完成前是阻塞的
		go fill(num[:50], &group)
		go fill(num[50:], &group)

		group.Wait()
		fmt.Println(num)
		//如果不加group.Add与Wait。主线程执行到此可能还没有填充结束

		//创建channel来并发读取，读一个channel是会等待，所以等待数据读入完毕
		ch1 := make(chan int)
		go sum(num[:33], ch1)
		ch2 := make(chan int)
		go sum(num[33:66], ch2)
		ch3 := make(chan int)
		go sum(num[66:], ch3)

		sum1 := <-ch1
		sum2 := <-ch2
		sum3 := <-ch3
		avg := (sum1 + sum2 + sum3) / len(num)
		fmt.Println(avg)
	}

// 只读管道,channel就是在goroutine之间通讯的一个数据结构 <- chan <-

	func sum(s []int, result chan<- int) {
		var total int
		for _, n := range s {
			total += n
		}
		result <- total
	}

	func fill(s []int, group *sync.WaitGroup) {
		defer group.Done()
		for i := range s {
			s[i] = rand.Intn(101)
		}
	}

//使用 channel 关闭 goroutine,让goroutine去读一个channel如果读取到值就退出

func main() {
	ch := make(chan string, 6)
	go func() {
		for {
			v, ok := <-ch
			if !ok {
				fmt.Println("结束")
				return
			}
			fmt.Println(v)
		}
	}()

	ch <- "煎鱼还没进锅里..."
	ch <- "煎鱼进脑子里了！"
	close(ch)
	time.Sleep(time.Second)
}

if i, ok := <-ch1 ;ok{
		...
	}

select如果没有case语句可以执行（channel发生阻塞），有default语句，则执行default语句；
通道如果不主动close掉，读出通道全部数据即通道无数据后该协程就会阻塞
从已经关闭的通道接收数据或者正在接收数据时，将会接收到通道类型的零值，然后停止阻塞并返回。

func main() {
	val := make(chan interface{})
	i := 0
	go func() {
		for {
			select {
			case <-val:
				return
			default:
				i++
				fmt.Println(i)
			}
		}
	}()
	time.Sleep(time.Second)
	close(val)
}


//ctx控制groutine
func main() {
	ch := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				ch <- struct{}{}
				return
			default:
				fmt.Println("煎鱼还没到锅里...")
			}
			time.Sleep(500 * time.Millisecond)
		}
	}(ctx)

	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()

	<-ch
	fmt.Println("结束")
}
*/
