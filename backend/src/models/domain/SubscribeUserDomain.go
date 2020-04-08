package domain

type SubscribeUserDomain struct {
	Id int `db:"ID, primarykey, autoincrement"`
	Requestor int `db:"Requestor"`
	Target int `db:"Target"`
	Status string `db:"Status"`
}