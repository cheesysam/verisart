package main

type Certificate struct {
	ID        string
	Title     string
	CreatedAt int //TODO date type
	OwnerID   string
	Year      int
	Note      string
	//Transfer  transfer
}

type Transfer struct {
	To     string
	Status string
}

type User struct {
	Id    string
	Email string
	Name  string
}
