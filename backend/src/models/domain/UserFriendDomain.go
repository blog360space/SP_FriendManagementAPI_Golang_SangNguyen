package domain

type UserFriendDomain struct {
	Id int `db:"ID, primarykey, autoincrement"`
	FromUserId int `db:"FromUserID"`
	ToUserId int `db:"ToUserID"`

}
