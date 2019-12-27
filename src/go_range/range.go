package main

import "fmt"


func main(){
	testRange()
}

func testRange(){
	nums := []int{2, 3, 4}
	sum := 0
	//不需要使用该元素的序号，所以我们使用空白符"_"省略了
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }
}