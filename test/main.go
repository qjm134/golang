package main

import "fmt"

type f interface {
	do()
}

type f1 struct{}

func (f1) do() {}

type f2 struct{}

func (f2) do() {}

func main() {
	var a interface{}
	a = 1
	b, ok := a.(string)
	fmt.Println(b, ok)

	/*
		a := make([]int, 2)
		fmt.Println(a)
	*/
	/*
				var b byte
				b = 'a'
				fmt.Printf("%c", b)

				exit := make(chan interface{})
				var data chan int = make(chan int, 2)
				//var m map[int]int

				go func() {
					data<-1
					close(data)

					fmt.Println(len(data))


					close(exit)
				}()

				//<-data
				//m[1] = 1
				 <-exit

				m := struct{}{}
				fmt.Println(unsafe.Sizeof(m))


				wg := sync.WaitGroup{}
				wg.Add(2)
				data := make(chan int)
				for i:=0; i<2; i++ {
					go func(ind int) {
						defer wg.Done()
						d := <-data
						fmt.Println(ind, d)
					}(i)
				}

				close(data)
				//data <- 4
				wg.Wait()


			start := make(chan bool)
			start <- true
			go func() {
				select {
				case <-start:
					time.Sleep(time.Second * 4)
				case <-time.After(time.Second):
					fmt.Println("")
				}
			}()



		s := make([]int, 0, 2)
		s = append(s, 0)
		fmt.Println(s)
		func(s []int) {
			s[0] = 1
			fmt.Println(s)

		}(s)
		fmt.Println(s)


	*/
	/*
		a := []int{1}
		b := []int{1}

		c := map[int]int{1: 1}
		d := map[int]int{1: 1}

		type ee struct {
			a int
			b []int
		}

		e1 := ee{1, []int{1}}
		e2 := ee{1, []int{1}}

		ff1 := f1{}
		ff2 := f2{}

		if ff1 == ff2 {
			fmt.Print("equal")
		}


	*/
}
