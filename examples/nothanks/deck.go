package main

type deck struct {
	cards []int
}

func createDeck(first, last) *deck {
	d := &deck{}
	length := last - first
	d.cards := make([]card, length)
	for i := 0; i < length; i++ {
		d[i] = first + i
	}
	return d
}

func (d *deck) Shuffle() {
	// FIXME
}
