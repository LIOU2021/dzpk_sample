package poker

import (
	"sort"
)

// Suit represents the suit of a card
type Suit int

// Rank represents the rank of a card
type Rank int

// Constants for card suits
const (
	Spades Suit = iota
	Hearts
	Diamonds
	Clubs
)

// Constants for card ranks
const (
	Two Rank = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

// Card represents a single playing card
type Card struct {
	Suit Suit
	Rank Rank
}

// Hand represents a poker hand
type Hand []Card

// HandRank represents the rank of a poker hand
type HandRank int

// Constants for hand ranks
const (
	HighCard HandRank = iota
	OnePair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
	RoyalFlush
)

// CompareHands compares two poker hands and returns:
// 1 if hand1 wins, -1 if hand2 wins, 0 if it's a tie
func CompareHands(hand1, hand2 Hand) int {
	rank1, value1 := evaluateHand(hand1)
	rank2, value2 := evaluateHand(hand2)

	if rank1 != rank2 {
		return int(rank1) - int(rank2)
	}
	return value1 - value2
}

// evaluateHand determines the rank and value of a poker hand
func evaluateHand(hand Hand) (HandRank, int) {
	if rank, value := isRoyalFlush(hand); rank {
		return RoyalFlush, value
	}
	if rank, value := isStraightFlush(hand); rank {
		return StraightFlush, value
	}
	if rank, value := isFourOfAKind(hand); rank {
		return FourOfAKind, value
	}
	if rank, value := isFullHouse(hand); rank {
		return FullHouse, value
	}
	if rank, value := isFlush(hand); rank {
		return Flush, value
	}
	if rank, value := isStraight(hand); rank {
		return Straight, value
	}
	if rank, value := isThreeOfAKind(hand); rank {
		return ThreeOfAKind, value
	}
	if rank, value := isTwoPair(hand); rank {
		return TwoPair, value
	}
	if rank, value := isOnePair(hand); rank {
		return OnePair, value
	}
	return HighCard, highCardValue(hand)
}

// isRoyalFlush checks if the hand is a royal flush
// 皇家同花顺
func isRoyalFlush(hand Hand) (bool, int) {
	if straightFlush, value := isStraightFlush(hand); straightFlush && value == int(Ace) {
		return true, value
	}
	return false, 0
}

// isStraightFlush checks if the hand is a straight flush
// 同花顺
func isStraightFlush(hand Hand) (bool, int) {
	if straight, straightValue := isStraight(hand); straight {
		if flush, _ := isFlush(hand); flush {
			return true, straightValue
		}
	}
	return false, 0
}

// isFourOfAKind checks if the hand contains four of a kind
func isFourOfAKind(hand Hand) (bool, int) {
	rankCount := make(map[Rank]int)
	for _, card := range hand {
		rankCount[card.Rank]++
		if rankCount[card.Rank] == 4 {
			return true, int(card.Rank)
		}
	}
	return false, 0
}

// isFullHouse checks if the hand is a full house
func isFullHouse(hand Hand) (bool, int) {
	rankCount := make(map[Rank]int)
	for _, card := range hand {
		rankCount[card.Rank]++
	}
	var threeOfAKindRank, pairRank Rank
	for rank, count := range rankCount {
		if count == 3 {
			threeOfAKindRank = rank
		} else if count == 2 {
			pairRank = rank
		}
	}
	if threeOfAKindRank != 0 && pairRank != 0 {
		return true, int(threeOfAKindRank)*13 + int(pairRank)
	}
	return false, 0
}

// isFlush checks if the hand is a flush
// 检查是否为同花
func isFlush(hand Hand) (bool, int) {
	suitCount := make(map[Suit]int)
	for _, card := range hand {
		suitCount[card.Suit]++
		if suitCount[card.Suit] == 5 {
			return true, highCardValue(hand)
		}
	}
	return false, 0
}

// isStraight checks if the hand is a straight
// 检查是不是顺子
// 第一个回传值true表示为顺子
// 第二个回传值为顺子的尾，如果是A-5就回传5，如果是10-A就回传A
func isStraight(hand Hand) (bool, int) {
	ranks := make([]int, len(hand))
	for i, card := range hand {
		ranks[i] = int(card.Rank)
	}
	sort.Ints(ranks)
	if ranks[len(ranks)-1] == int(Ace) && ranks[0] == int(Two) { // 判断如果是A-5的顺子处理
		return isConsecutive(ranks[:4]), int(Five)
	}
	return isConsecutive(ranks), ranks[len(ranks)-1]
}

// isConsecutive checks if a slice of integers is consecutive
func isConsecutive(ranks []int) bool {
	for i := 1; i < len(ranks); i++ {
		if ranks[i] != ranks[i-1]+1 {
			return false
		}
	}
	return true
}

// isThreeOfAKind checks if the hand contains three of a kind
func isThreeOfAKind(hand Hand) (bool, int) {
	rankCount := make(map[Rank]int)
	for _, card := range hand {
		rankCount[card.Rank]++
		if rankCount[card.Rank] == 3 {
			return true, int(card.Rank)
		}
	}
	return false, 0
}

// isTwoPair checks if the hand contains two pairs
func isTwoPair(hand Hand) (bool, int) {
	rankCount := make(map[Rank]int)
	pairs := []Rank{}
	for _, card := range hand {
		rankCount[card.Rank]++
		if rankCount[card.Rank] == 2 {
			pairs = append(pairs, card.Rank)
		}
	}
	if len(pairs) == 2 {
		return true, int(pairs[1])*13 + int(pairs[0])
	}
	return false, 0
}

// isOnePair checks if the hand contains one pair
func isOnePair(hand Hand) (bool, int) {
	rankCount := make(map[Rank]int)
	for _, card := range hand {
		rankCount[card.Rank]++
		if rankCount[card.Rank] == 2 {
			return true, int(card.Rank)
		}
	}
	return false, 0
}

// highCardValue returns the value of the highest card in the hand
func highCardValue(hand Hand) int {
	maxRank := 0
	for _, card := range hand {
		if int(card.Rank) > maxRank {
			maxRank = int(card.Rank)
		}
	}
	return maxRank
}
