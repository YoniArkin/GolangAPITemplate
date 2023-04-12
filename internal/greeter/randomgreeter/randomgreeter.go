package randomgreeter

type RandomGreeter struct{}

func (RandomGreeter) Greet() string {
	return "hello world!"
}
