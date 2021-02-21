package main 
import "fmt"

func update (a *int){
*a = *a + 10
}

func main(){

 var x int 
 var  y *int

 fmt.Println("The value of x  = ", x)
 fmt.Println("Memory address of x = ", &x )

 fmt.Println("The value of y  = ", y)
 fmt.Println("Memory address of y = ", &y )

 x = 10 //assignment 
 y = &x //referencing

 fmt.Println("The value of x  = ", x)  //accessing
 fmt.Println("The value of y  = ", y)
 fmt.Println("y dereferancing value is ", *y)
 
 update( &x)
fmt.Println(x)

}