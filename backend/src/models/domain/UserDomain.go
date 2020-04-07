package domain

type UserDomain struct {
	Id int32 `db:"ID, primarykey, autoincrement"`
	Username string
}