package main
import "fmt"
func main(){

	fmt.Println("Lets start learning GO for automation testing...!!!")
	var age int8
	age = 19
	fmt.Println(age)
	fmt.Printf("Varible is of type %T", age)
	fmt.Println()
//declaring strings
	var myname string = "yash chouhan"
	fmt.Println(myname)	
	
	const data = 1 	// declaring interger
	fmt.Println("Data=", data)
/*	var (
		foo=2
		boo=5
	)
	fmt.Println(foo, bar)
*/	
	fmt.Println("Performing arithmetic")
	a:=100
	b:=75
	  fmt.Println("a=",a)
	  fmt.Println("b=",b)
	fmt.Println("Result=", a-b)
	fmt.Println()

	
	fmt.Printf("%d \n",100)
	fmt.Printf("%d \n",100)
	fmt.Printf("%b \n",100)
	fmt.Printf("%c \n",100)	
	fmt.Printf("%x \n",100)
	fmt.Printf("%e \n",100)
	

// Printing no. from 1 to 10

	i := 1
	for i<= 10{
	fmt.Println(i)
	i++
	}
	 fmt.Println()
	
	// Printing no. from 10 to 1
	fmt.Println("Number In descending order")
	j:=10 
	for j>0{
	 fmt.Println(j)
	j=j-1
	}
	fmt.Println()	
//Different For loop patter

	for k:=0; k<5; k++{
	fmt.Println(k)
	}

}	
