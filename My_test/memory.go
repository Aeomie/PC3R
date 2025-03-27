package main

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{map[string]int{}, 0}
}

type MemoryStore struct {
	storeUser   map[string]int
	acc_counter int
}

func (i *MemoryStore) accCreation(username string,
	password int) {
	i.storeUser[username] = password
	i.acc_counter++
}

func (i *MemoryStore) getPw(username string) int {
	return i.storeUser[username]
}
