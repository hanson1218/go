package go_struct

import "fmt"

type Books struct {
   title string
   author string
   subject string
   book_id int
}

func UseStruct(){
	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})
	var Book1 Books        /* 声明 Book1 为 Books 类型 */

   	/* book 1 描述 */
   Book1.title = "Go 语言"
   Book1.author = "www.runoob.com"
   Book1.subject = "Go 语言教程"
   Book1.book_id = 6495407
   /* 打印 Book1 信息 */
   fmt.Printf( "Book 1 title : %s\n", Book1.title)
   fmt.Printf( "Book 1 author : %s\n", Book1.author)
   fmt.Printf( "Book 1 subject : %s\n", Book1.subject)
   fmt.Printf( "Book 1 book_id : %d\n", Book1.book_id)

    var Book2 Books        /* 声明 Book2 为 Books 类型 */
    /* book 2 描述 */
   Book2.title = "Python 教程"
   Book2.author = "www.runoob.com"
   Book2.subject = "Python 语言教程"
   Book2.book_id = 6495700
   printBook(Book2)
   //取Book2的指针
   printBookByPoint(&Book2)
}

/**结构体当参数传**/
func printBook( book Books ) {
   fmt.Printf( "Book title : %s\n", book.title)
   fmt.Printf( "Book author : %s\n", book.author)
   fmt.Printf( "Book subject : %s\n", book.subject)
   fmt.Printf( "Book book_id : %d\n", book.book_id)
}

/**通过结构体的指针传参**/
func printBookByPoint( book *Books ) {
   fmt.Printf( "Point Book title : %s\n", book.title)
   fmt.Printf( "Point Book author : %s\n", book.author)
   fmt.Printf( "Point Book subject : %s\n", book.subject)
   fmt.Printf( "Point Book book_id : %d\n", book.book_id)
}