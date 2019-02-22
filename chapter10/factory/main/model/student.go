package model

// 定义一个结构体
// type Student struct {
type student struct {
	// Name  string
	name  string
	score float64
}

// 因为student结构体首字母是小写的，因为只能在model包中使用。
// 我们通过工厂模式来解决。
func NewStudent(name string, score float64) *student {
	return &student{
		name:  name,
		score: score,
	}
}

// 如果score字段的首字母小写，则，在其他包不可以直接访问，我们可以提供一个方法。
func (s *student) GetScore() float64 {
	s.score = 10.1
	return s.score
}

func (s *student) GetName() string {
	s.name = "haha"
	return s.name
}

// 以上2种方式的 编写，是student结构体实例的 地址传递，而非值传递。 这种方式的好处是 在上面2个方法中 对结构体字段内容的改动 会影响函数外的结构体实例的实际值
// 还有就是 地址传递 效率要高于地址传递。

// 如果score字段的首字母小写，则，在其他包不可以直接访问，我们可以提供一个方法。
func (s student) GetScore1() float64 {
	s.score = 10.1

	return s.score
}

func (s student) GetName1() string {
	s.name = "haha"

	return s.name
}
