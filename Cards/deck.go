package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

// create a new type of deck which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"} // slice of strings
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {

		for _, value := range cardValues {

			cards = append(cards, value+" of "+suit)

		}

	}
	return cards

}

//receiver on a function
func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}
}

//deal function
// type deck d
// handSize : here we use 3 so used int type
//returning multiple parameters and both type deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:] // return two slices (one from start to handsize(nothig before) and second UptohandSize(nothing after))

}

// join slice of string (here we have a slice of string for cards (deck)))
func (d deck) tostring() string {
	return strings.Join([]string(d), ",")
}

//save file to hard drive
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.tostring()), 0666)
}

//Read file from hard drive
func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile((fileName)) // bs : byteSlice and err is error
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)

}

//shuffle cards
// Generate random number in deck swap the position to shuffle the cards
func (d deck) shuffle() {

	for i := range d {

		newPosition := rand.Intn(len(d) - 1)        // generate random no
		d[i], d[newPosition] = d[newPosition], d[i] // swap the position

	}
}
