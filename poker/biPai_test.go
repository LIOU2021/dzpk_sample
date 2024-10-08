package poker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 测试顺子接口
func Test_isStraight(t *testing.T) {
	// A-5
	hand1 := []Card{
		{Suit: Hearts, Rank: Ace},
		{Suit: Hearts, Rank: Two},
		{Suit: Hearts, Rank: Three},
		{Suit: Hearts, Rank: Four},
		{Suit: Hearts, Rank: Five},
	}
	ans, v := isStraight(hand1)
	assert.True(t, ans)
	assert.Equal(t, int(Five), v)

	// 10-A
	hand2 := []Card{
		{Suit: Hearts, Rank: Ace},
		{Suit: Hearts, Rank: Ten},
		{Suit: Hearts, Rank: Jack},
		{Suit: Hearts, Rank: Queen},
		{Suit: Hearts, Rank: King},
	}
	ans, v = isStraight(hand2)
	assert.True(t, ans)
	assert.Equal(t, int(Ace), v)

	// 7-11
	hand3 := []Card{
		{Suit: Hearts, Rank: Jack},
		{Suit: Hearts, Rank: Seven},
		{Suit: Hearts, Rank: Ten},
		{Suit: Hearts, Rank: Nine},
		{Suit: Hearts, Rank: Eight},
	}
	ans, v = isStraight(hand3)
	assert.True(t, ans)
	assert.Equal(t, int(Jack), v)

	// test error case
	hand4 := []Card{
		{Suit: Hearts, Rank: Ace},
		{Suit: Hearts, Rank: Seven},
		{Suit: Hearts, Rank: Ten},
		{Suit: Hearts, Rank: Nine},
		{Suit: Hearts, Rank: Eight},
	}
	ans, v = isStraight(hand4)
	assert.False(t, ans)
	assert.Equal(t, int(Ace), v)
}

// 测试找出散牌中第n大的牌
func Test_highCardValueV2(t *testing.T) {
	hand1 := []Card{
		{Suit: Hearts, Rank: Four},
		{Suit: Hearts, Rank: Five},
		{Suit: Hearts, Rank: Ace},
		{Suit: Hearts, Rank: Three},
		{Suit: Hearts, Rank: Two},
	}

	assert.Equal(t, int(Ace), highCardValueV2(hand1, 0))
	assert.Equal(t, int(Five), highCardValueV2(hand1, 1))
}
