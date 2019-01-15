// package main

// import "fmt"

// func main(){
// 	// str := "hello,world\n"
// 	// a := 1
// 	// fmt.Print(str)
// 	// if a == 1{
// 	// 	fmt.Print("strt is exist")
// 	var marks int
// 	var grade string
// 	switch marks{
// 		case 90: grade = "A"
// 		case 80: grade = "B"
// 		default:grade = "C"
// 	}
// 	fmt.Println("Please input your grade:")
// 	fmt.Scanln(&marks)
// 	fmt.Println(marks)
// 	}

package main

import "fmt"
const MAX int= 3
func main(){
   /* 定义局部变量 */
   var grade string 
   var marks int 
   var array = []int{1,1,3,6,5,6,9,8,7,4,5}
   var ptr [MAX]*int;
	for i:=0;i<MAX;i++{
		ptr[i] = &array[i]
		fmt.Println(*ptr[i])
	}
   fmt.Println("请输入你的成绩:")
   fmt.Scanln(&marks)
   switch {
	  case marks>=90: grade = "A"
      case marks>=80 && marks<90: grade = "B"
      case marks<80 && marks>=70 : grade = "C"
      default: grade = "D"  
   }

   switch {
      case grade == "A" :
         fmt.Printf("优秀!\n" )
      case grade == "B", grade == "C" :
         fmt.Printf("良好\n" )      
      case grade == "D" :
         fmt.Printf("及格\n" )      
      case grade == "F":
         fmt.Printf("不及格\n" )
      default:
         fmt.Printf("差\n" )
   }
   fmt.Printf("你的等级是 %s\n", grade)
   fmt.Println(&marks)
   imp()
}


//结构体
type Books struct{
	title string
	author string
	subject string
	book_id int
}


func imp(){
	var book1 Books
	book1.title = "go"
	book1.author = "Vincent"
	book1.subject = "maths"
	book1.book_id = 123
	printbooks(book1)
	printbook(&book1)
}

func printbooks(book Books){
	fmt.Printf("book title is :%s\n",book.title)
	fmt.Printf("book author is :%s\n",book.author)
	fmt.Printf("book subject is :%s\n",book.subject)
	fmt.Printf("book book_id is :%d\n",book.book_id)
}
func printbook(book *Books){
	fmt.Printf("book title is :%s\n",book.title)
}