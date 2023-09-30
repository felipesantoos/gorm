package core

type User struct {
	ID        int64
	Login     string
	Email     string
	Avatar    string
	Active    bool
	LastLogin int64
	Created   int64
	Updated   int64
}

func (u *User) Validate() bool {
	return u != &User{}
}
