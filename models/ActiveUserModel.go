package golibraries/models

type ActiveUserModel struct {
	ID        string     `gorm:"column:id;primaryKey" json:"id"`       //
	UserID    NullString `gorm:"column:user_id" json:"user_id,string"` //
	Token     NullString `gorm:"column:token" json:"token"`            //
	ExpiredOn NullString `gorm:"column:expired_on" json:"expired_on"`  //
}

// TableName sets the insert table name for this struct type
func (u *ActiveUserModel) TableName() string {
	return "active_user"
}
