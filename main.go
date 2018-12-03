package main

import (
	"fmt"
)

//func main(){
//	c := make(chan int)
//	quit := make(chan int)
//	go func(){
//		for i := 0; i < 4; i++{
//			fmt.Println("in = ", <-c)
//		}
//		quit <- 1
//	}()
//	testMuti(c, quit)
//}
//func testMuti(c, quit chan int){
//	x, y := 0, 1
//	for {
//		select{
//		case c<-x:
//			fmt.Println("out = ", &c)
//			x, y = y, x+y
//		case <-quit:
//			fmt.Print("\nquit")
//			return
//		}
//	}
//}




//type Stu struct{
//	name string
//	age  int
//}

//func (s Stu) String() string{
//	return fmt.Sprintf("%s_%d", s.name, s.age)
//}

//func (s *Stu) String() string{
//	return fmt.Sprintf("%s__%d", s.name, s.age)
//}



func main(){
	fmt.Println("welcome to firstProject.")

}



