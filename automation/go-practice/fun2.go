package main
import "fmt"
func main(){
 var res int
 n:=5
 res= mul(n)
 fmt.Println(res)
}

func mul(n int) int{

if (n==0 || n==1){
return 1
}else{
return n*mul(n-1)} 
}
