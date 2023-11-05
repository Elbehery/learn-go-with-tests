package myio

//// InMemoryPlayerStore stores player information in memory
//type InMemoryPlayerStore struct {
//	store map[string]int
//}
//
//func NewInMemoryPlayerStore() *InMemoryPlayerStore {
//	return &InMemoryPlayerStore{map[string]int{}}
//}
//
//func (i *InMemoryPlayerStore) GetPlayerWins(player string) int {
//	return i.store[player]
//}
//
//func (i *InMemoryPlayerStore) RecordWins(player string) {
//	i.store[player]++
//}
//
//func (i *InMemoryPlayerStore) GetLeague() []Player {
//	var players []Player
//	for k, v := range i.store {
//		players = append(players, Player{
//			Name: k,
//			Wins: v,
//		})
//	}
//	return players
//}
