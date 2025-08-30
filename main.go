package main

import (
	"errors"

	"github.com/Raghav123-30/mastering-golang.git/packages/student"
)




func add[T int | float64](a ,b T) (sum T) {
   sum = a+b
   return sum;
}

type pair[T1 any, T2 any] struct{
   first T1
   second T2
}

func (p pair[T1,T2]) display(){
  println(p.first)
  println(p.second)
}

//same type multiple parameters
func sub[T int | float64](a,b T) (difference T) {
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

func addAll[T int|float64](nums ...T) T{
   var sum T
  
   for _,num := range(nums){
    sum += num
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
   println(add(3,4))
   println(add(3.4,5.6))
   println(addAll(4,5,6,7,8))
   println(sub(4,8.5))

   p1 := pair[int, int]{
      first:3,
      second:4,
   }

   p2 := pair[int, string]{
      first: 3,
      second:"XD",
   }

   p1.display()
   p2.display()
}