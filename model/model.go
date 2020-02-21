package model
//import ""

var Num1 string ="a" 

type student struct{
	Name string
	score float64
}

func (s *student) Newscore()float64{
	return s.score
}
func Newstudent (n string,s float64) *student {
	return &student{
		Name :n,
		score :s,
	}
}