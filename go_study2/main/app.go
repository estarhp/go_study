package main

import "fmt"

type player struct {
	name       string
	level      int
	experience int
	lifeBar    int
	strength   int
}

type Attack interface {
	attack(p *player)
}

func (p2 player) attack(p *player) {
	old := p.lifeBar
	p.lifeBar -= p2.strength
	fmt.Printf("%d --> %d", old, p.lifeBar)
}

func main() {
	p1 := player{
		name:       "p1",
		level:      1,
		experience: 30,
		lifeBar:    100,
		strength:   10,
	}

	p2 := player{
		name:       "p2",
		level:      2,
		experience: 30,
		lifeBar:    100,
		strength:   15,
	}

	var x Attack = p1
	var pp2 *player = &p2
	x.attack(pp2)

}
