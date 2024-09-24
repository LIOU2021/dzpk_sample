package main

import (
	"fmt"
	"poker/poker"
)

func main() {
	hand1 := poker.Hand{
		poker.Card{Suit: poker.Hearts, Rank: poker.Ace},
		poker.Card{Suit: poker.Hearts, Rank: poker.King},
		poker.Card{Suit: poker.Hearts, Rank: poker.Queen},
		poker.Card{Suit: poker.Hearts, Rank: poker.Jack},
		poker.Card{Suit: poker.Hearts, Rank: poker.Ten},
	}

	hand2 := poker.Hand{
		poker.Card{Suit: poker.Spades, Rank: poker.Ace},
		poker.Card{Suit: poker.Spades, Rank: poker.King},
		poker.Card{Suit: poker.Clubs, Rank: poker.Ace},
		poker.Card{Suit: poker.Diamonds, Rank: poker.King},
		poker.Card{Suit: poker.Hearts, Rank: poker.Queen},
	}

	result := poker.CompareHands(hand1, hand2)
	if result > 0 {
		fmt.Println("Hand 1 wins")
	} else if result < 0 {
		fmt.Println("Hand 2 wins")
	} else {
		fmt.Println("It's a tie")
	}
}
