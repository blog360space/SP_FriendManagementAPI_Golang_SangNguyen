package domain

type UserDomain struct {
	Id int `db:"ID, primarykey, autoincrement"`
	Username string
}