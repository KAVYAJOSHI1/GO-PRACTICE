 package main 
 import "fmt"

 func main(){
	var n int
	fmt.Print("ENTER THE NUMBER\n")
	fmt.Scan(&n)
	if n%4==0{
		fmt.Print("IT IS A LEAP YEAR\n")
	}else{
		fmt.Print("IT IS NOT A LEAP YEAR\n")
	}
	

	
 }