package main

type Student struct {
	age  int
	name string
	male bool
}

func main() {
	st := &Student{}
	st.age = 12
	st.name = "tuan anh"
	st.male = true
	println(st.GetAge())
	println(st.GetName())
	println(st.male)

}
func (s *Student) GetAge() int {
	return s.age
}
func (s *Student) GetName() string {
	return s.name
}
func (s *Student) IsMale() bool {
	return s.male
}
