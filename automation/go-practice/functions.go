package main
import "fmt"
func main(){

var a int =10
var b int =34
var res int

res= max(a,b)
fmt.Println("Max value is ",res)
}

func max(n,m int) int{
var gret int
if (n>m){
gret=n
}else {
gret=m
}
return gret
}
