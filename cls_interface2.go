package main
import "fmt"

func myfunc(args interface{}){

	fmt.Println("myFunc is called")
	fmt.Println(args)

	val, ok := args.(string)
	if !ok {
		fmt.Println("args is not str")
	}else{
		fmt.Println("args is str, val = ",val)
		fmt.Printf("val type is %T\n",val )
	}
}

type Book struct{
	name string
}

func main(){
	book :=Book{"gw"}
	myfunc(book)

	myfunc(111)
	myfunc("ads")

}