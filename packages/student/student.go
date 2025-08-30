package student

type Student struct{
   name string 
   age int
}

func New(name string,age int) *Student{
  return &Student{
	name,
	age,
  }
}

func(s *Student) Display(){
   println(s.name)
   println(s.age)
}