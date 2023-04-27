package main

type print interface {
	printSelf()
}

type person struct {
	name string
	year int
}
type car struct {
	name       string
	launchYear int
}

func (p person) printSelf() {
	println("person and year=", p.year)
}
func (c car) printSelf() {
	println("car and launchYear=", c.launchYear)
}

func main() {
	var x print = person{
		name: "hhh",
		year: 12,
	}
	x.printSelf()

	var y print = car{
		name:       "三轮",
		launchYear: 1949,
	}
	y.printSelf()
}
