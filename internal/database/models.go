package database

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID           string     `gorm:"type:char(36);unique;not null"` // UUID、アプリケーション側で生成し、char(36)で保存
	Name           string     `gorm:"size:255;not null"`             // 名前はNULL不可
	Email          string     `gorm:"size:255;unique;not null"`      // メールアドレスは一意かつNULL不可
	HashedPassword string     `gorm:"size:255;not null"`             // ハッシュ化されたパスワード、NULL不可
	LeavedAt       *time.Time `gorm:"index"`                         // 退会日時、NULL可能でインデックスを付ける
	MyOffers       []*Offer   `gorm:"foreignKey:UserID"`             // オファー、一対多の関係
	EntryOffers    []*Offer   `gorm:"many2many:user_entry_offers"`   // オファー、多対多の関係
}

type Offer struct {
	gorm.Model
	UUID        string     `gorm:"type:char(36);unique;not null"` // UUID、アプリケーション側で生成し、char(36)で保存
	UserID      uint       `gorm:"not null"`                      // ユーザーID、NULL不可
	Giz         uint       `gorm:"not null"`                      // GIZ、NULL不可
	ChatURL     string     `gorm:"size:255;not null"`             // チャットURL、NULL不可
	Title       string     `gorm:"size:255;not null"`             // タイトル、NULL不可
	Description string     `gorm:"size:255;not null"`             // 説明、NULL不可
	IsPublic    bool       `gorm:"not null"`                      // 公開フラグ、NULL不可
	Deadline    time.Time  `gorm:"not null"`                      // 締め切り日時、NULL不可
	DoneAt      *time.Time `gorm:"index"`                         // 完了日時、NULL可能でインデックスを付ける
	EntryUsers  []*User    `gorm:"many2many:user_entry_offers"`   // ユーザー、多対多の関係
}

var models = []interface{}{
	&User{},
	&Offer{},
}
