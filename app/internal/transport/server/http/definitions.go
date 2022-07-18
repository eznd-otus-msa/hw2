package http

import (
    "github.com/eznd-otus-msa/hw2/app/internal/domain"
)

const UserIdFieldName = "id"

type UserData struct {
    Username  string `json:"username"`
    FirstName string `json:"firstName"`
    LastName  string `json:"lastName"`
    Email     string `json:"email"`
    Phone     string `json:"phone"`
}

type User struct {
    Id domain.UserId `json:"id"`
    UserData
}

type UserUpdate struct {
    UserData
}

type UserPartialUpdate struct {
    Username  NullableString `json:"username"`
    FirstName NullableString `json:"firstName"`
    LastName  NullableString `json:"lastName"`
    Email     NullableString `json:"email"`
    Phone     NullableString `json:"phone"`
}

//
//func (u *User) ToDomain(d *domain.User) {
//    d.Id = u.Id
//    d.Username = u.Username.Value
//    d.FirstName = u.FirstName.Value
//    d.LastName = u.LastName.Value
//    d.Email = u.Email.Value
//    d.Phone = u.Phone.Value
//}
//
//func (u *User) FromDomain(d *domain.User) {
//    u.Id = d.Id
//
//    u.Username.Value = d.Username
//    u.Username.Set = true
//    u.Username.Null = false
//
//    u.FirstName.Value = d.FirstName
//    u.FirstName.Set = true
//    u.FirstName.Null = false
//
//    u.LastName.Value = d.LastName
//    u.LastName.Set = true
//    u.LastName.Null = false
//
//    u.Email.Value = d.Email
//    u.Email.Set = true
//    u.Email.Null = false
//
//    u.Phone.Value = d.Phone
//    u.Phone.Set = true
//    u.Phone.Null = false
//}
