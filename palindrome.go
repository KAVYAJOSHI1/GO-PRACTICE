//PROGRAM TO CHECK PALINDROME
package main
import "fmt"

func main(){

	var n int
	fmt.Print("ENTER THE NUMBER")
	fmt.Scan(&n)
	
	var rev int
	var rem int
	var n1 int =n
	for n>0{
	
		rem=n%10
		rev=rev*10+rem
		n=n/10
	} 
	
	if rev==n1{
	
		fmt.Print("IT IS PALINDROME")

	}else{
		fmt.Print("IT IS NOT PALINDROME")
	}

}