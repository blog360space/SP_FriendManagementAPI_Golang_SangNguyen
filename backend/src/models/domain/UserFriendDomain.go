package domain

type UserFriendDomain struct {
	Id int32 `db:"ID, primarykey, autoincrement"`
	FromUserId int32 `db:"FromUserID"`
	ToUserId int32 `db:"ToUserID"`

}
