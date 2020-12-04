package model

//定义一个结构体
type student struct {
	Name  string
	score float64
}

//因为student结构体首字母是小写，因为是只能在model使用
//我们通过工厂模式来解决

//NewStudent ...
func NewStudent(name string, score float64) *student {
	return &student{
		Name:  name,
		score: score,
	}
}

//如果将其中的Score改成score，则会报错吗？那如何解决呢？【我们可以提供一个方法】

func (s *student) GetScore() float64 {
	return s.score
}
