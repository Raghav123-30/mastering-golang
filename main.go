package main

func main(){
  //Explicit type declaration
  var x int
  x = 10
  println(x)

  // Shorthand declaration, go infers the type
  y := 42
  println(y)

  // multiple declaration
  var m, n = 3, "hello"
  println(m,n)

  //shorthand multiple declaration
  a, b := 3,4
  println(a,b)

  x = 5
  if true{
    x := 10 //shadows outer x
    println(x)
  }
  println(x)

  //constants
  const Pi = 3.14

  //constants are untyped ,allows assignment to int32 variable
  const Max = 1000
  var z int32 = Max

  //pointer
  g := 10;
  p := &g;

  println(z,g,p)
}