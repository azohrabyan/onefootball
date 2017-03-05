package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortByName(tst *testing.T) {
	players := map[int]*Player{}

	players[10] = &Player{ID: 10, Name: "bbb"}
	players[20] = &Player{ID: 20, Name: "ccc"}
	players[30] = &Player{ID: 30, Name: "aaa"}

	s := NewPlayerSorter(players)
	sorted := s.SortBy(Name)

	assert.Equal(tst, []int{30, 10, 20}, sorted)
}
