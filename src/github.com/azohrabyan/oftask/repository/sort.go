package repository

import (
	"sort"
	"strings"
)

type Compare func(*Player, *Player) bool

type PlayerSorter struct {
	ids     []int
	players map[int]*Player
	less    Compare
}

func NewPlayerSorter(players map[int]*Player) *PlayerSorter {
	ids := make([]int, 0, len(players))
	for id := range players {
		ids = append(ids, id)
	}
	return &PlayerSorter{
		ids:     ids,
		players: players,
	}
}

func (s *PlayerSorter) Len() int {
	return len(s.ids)
}

func (s *PlayerSorter) Less(i, j int) bool {
	return s.less(s.players[s.ids[i]], s.players[s.ids[j]])
}
func (s *PlayerSorter) Swap(i, j int) {
	s.ids[i], s.ids[j] = s.ids[j], s.ids[i]
}

func (s *PlayerSorter) SortBy(less Compare) []int {
	s.less = less
	sort.Sort(s)
	return s.ids
}

func Name(p1, p2 *Player) bool {
	return strings.Compare(p1.Name, p2.Name) == -1
}
