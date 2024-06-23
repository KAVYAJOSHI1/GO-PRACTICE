package main
import "fmt"
func main(){
	var a[2] int
	var b[2] int
	var c[2] int
	fmt.Print("ENTER THE NUMBERS:\n")
	a=append(a,0)
	fmt.Scan(&a[0],&a[1],&b[0],&b[1])
	for i:=0;i<2;i++{
		c[i]=a[i]+b[i]
		fmt.Print(c[i])
	}

}