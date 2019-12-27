package interface1

import (
    "fmt"
)

//定义接口，只要是实现了这个接口的方法都输入实现的
type Phone interface {
    call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
    fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
    fmt.Println("I am iPhone, I can call you!")
}

type Test struct {
}

func (test Test) call(){
	fmt.Println("I am test, I can call you!")
}

func main() {
    var phone Phone

    phone = new(NokiaPhone)
    phone.call()

    phone = new(IPhone)
    phone.call()

    phone = new(Test)
    phone.call()

}