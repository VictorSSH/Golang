package vict1

type Person struct {
	Name string
	Addr string
	Id   int
}

func Info(Name, Addr string, Id int) *Person {
	return &Person{
		Name: Name,
		Addr: Addr,
		Id:   Id,
	}
}
func (sfo Person) ShowInfo() {
	sfo.ShowInfo()
	return

}
