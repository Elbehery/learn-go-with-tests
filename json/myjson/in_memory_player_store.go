package myjson

type InMemoryPlayerStore struct {
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

func (m *InMemoryPlayerStore) GetPlayerScore(player string) int {
	return m.store[player]
}

func (m *InMemoryPlayerStore) RecordWin(player string) {
	m.store[player]++
}
