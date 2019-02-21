package main

type certificate struct {
	ID        string
	Title     string
	CreatedAt int //TODO date type
	OwnerID   string
	Year      int
	Note      string
	Transfer  transfer
}

type transfer struct {
	To     string
	Status string
}

type user struct {
	Id    string
	Email string
	Name  string
}
