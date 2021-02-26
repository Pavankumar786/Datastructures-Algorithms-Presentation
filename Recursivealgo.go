package main  
  
import (  
 "fmt"  
)  
  
func main() {  
 var number int  
 fmt.Print("Enter Number:")  
 fmt.Scanln(&number)  
 fmt.Printf("Factorial of %d = %d\n", number, RecursiveFactorial(number))  
  
}  
func RecursiveFactorial(number int) uint64 {  
 if number >= 1 {  
  return uint64(number) * RecursiveFactorial(number-1)  
 } else {  
  return 1  
 }  
}  