//TO ASK NUMBERS FROM USERS AND STORE IN SLICE
package main
import "fmt"
func main(){
	var n int
	fmt.Print("ENTER THE NUMBER")
	fmt.Scan(&n)
	a:=make([]int,0,n)
	var b int
	fmt.Print("ENTER THE NUMBER")
	fmt.Scan(&b)
	a=append(a,b)
	fmt.Print(a)
}