package main

import (
	"errors"

	"github.com/Raghav123-30/mastering-golang.git/packages/student"
)




func add(a int64,b int64) (sum int64) {
   sum = a+b
   return sum;
}

//same type multiple parameters
func sub(a,b int64) (difference int64) {
   difference = a-b
   return // just writing return will be enough 
}

//multiple return values
func divide(a,b float64) (float64,error){
   if(b == 0){
     return 0, errors.New("cannot devide by zero")
   }
   return a/b,nil
}

//variadic functions , accept multiple arguments

func addAll(nums ...int) int{
   sum := 0
   for _,num := range(nums){
    sum += int(num)
   }
   return sum
}

func main(){
   sum := add(3,4)
   println(sum)
   result,err := divide(3,0)
   if(err != nil){
    println(err)
   }else{
    println(result)
   }
   println(addAll(3,4,5,6,7))

  
   raghu := student.New("Raghavendra",22)
   raghu.Display()

}