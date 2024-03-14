package pr

type Demo struct {
	Name string
}

func New(name string) *Demo {
	return &Demo{Name: name}
}

func (d *Demo) Call() string {
	return d.Name
}
