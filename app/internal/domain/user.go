package domain

type UserId int64

func (t UserId) Validate() error {
	if t <= 0 {
		return ErrInvalidUserId
	}
	return nil
}

type User struct {
	Id        UserId
	Username  string
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

type UserPartialData = map[string]string

func (pd UserPartialData) Init() UserPartialData {
	return make(map[string]string)
}
