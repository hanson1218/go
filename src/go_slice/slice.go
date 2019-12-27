package main

import (
	"fmt"
)


func main(){
	printSlice()
}

func printSlice(){
	//直接初始化切片，[]表示是切片类型，{1,2,3}初始化值依次是1,2,3.其cap=len=3
	arr :=[]int {1,2,3 } 
	fmt.Println(arr)
	//初始化切片s,是数组arr的引用
	slice := arr[:]
	fmt.Println(slice)
	//将arr中从下标startIndex到endIndex-1 下的元素创建为一个新的切片
	//左指针指向0，右指针指向2
	slice1 := arr[0:2]
	fmt.Println(slice1)
	//默认 endIndex 时将表示一直到arr的最后一个元素
	slice2 := arr[0:]
	fmt.Println(slice2)
	//默认 startIndex 时将表示从arr的第一个元素开始
	slice3 := arr[:3] 
	fmt.Println(slice3)
	printSliceMake()
}


func printSliceMake(){
	//通过make方法初始化slice
	arr1 :=make([]int,5,10) 
	fmt.Println("len=%d cap=%d slice=%v\n",len(arr1),cap(arr1),arr1)

	arr1[0]=2
	arr1[4]=8
	
	fmt.Println("len=%d cap=%d slice=%v\n",len(arr1),cap(arr1),arr1)
	arr1 = arr1[:cap(arr1)]
	//增加切片的长度，用容量赋值
	arr1[6]=8
	fmt.Println("len=%d cap=%d slice=%v\n",len(arr1),cap(arr1),arr1)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
    numbers1 := make([]int, len(arr1), (cap(arr1))*2)
    /* 拷贝 numbers 的内容到 numbers1 */
   copy(numbers1,arr1)
   fmt.Println("len=%d cap=%d slice=%v\n",len(numbers1),cap(numbers1),numbers1) 
}