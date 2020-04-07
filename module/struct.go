package module

type Account struct {
	ID            uint64 `gorm:"primary_key;not null;auto_increment"`
	Email         string `gorm:"size:30;default:'';index:email"`
	Phone         string `gorm:"size:15;default:'';index:phone"`
	Username      string `gorm:"size:30;default:'';index:username"`
	Password      string `gorm:"size:32;default:''"`
	CreateAt      int64  `gorm:"default:0"`
	CreateIpAt    string `gorm:"size:12;default:''"`
	LastLoginAt   int64  `gorm:"default:0"`
	LastLoginIpAt string `gorm:"size:12;default:''"`
	LoginTimes    int32  `gorm:"default:0"`
	Status        int8   `gorm:"default:0"`
}

type AccountPlatform struct {
	ID            uint64 `gorm:"primary_key;not null;auto_increment"`
	Uid           uint64 `gorm:"not null;index:uid;"`
	PlatformId    string `gorm:"size:60;default:'';index:platform_id"`
	PlatformToken string `gorm:"size:60;default:''"`
	Type          int8   `gorm:"type:tinyint(1);default:0"`
	Nickname      string `gorm:"size:60;default:''"`
	Avatar        string `gorm:"default:''"`
	CreateAt      int64  `gorm:"default:0"`
	UpdateAt      int64  `gorm:"default:0"`
}

type AccountMember struct {
	ID       uint64 `gorm:"primary_key;not null;auto_increment"`
	Uid      uint64 `gorm:"default:0;index:uid"`
	Nickname string `gorm:"size:30"`
	Avatar   string `gorm:"default:''"`
	Gender   string
	Role     uint8 `gorm:"default:0"`
	CreateAt int64 `gorm:"default:0"`
	UpdateAt int64 `gorm:"default:0"`
}

type AccountStaff struct {
	ID       uint64 `gorm:"primary_key;not null;auto_increment"`
	Uid      uint64 `gorm:"default:0;index:uid"`
	Email    string `gorm:"size:30;default:'';index:email"`
	Phone    string `gorm:"size:15;default:'';index:phone"`
	Name     string `gorm:"size:30;default:''"`
	Nickname string `gorm:"size:30;default:''"`
	Avatar   string `gorm:"default:''"`
	Gender   string
	CreateAt int64 `gorm:"default:0"`
	UpdateAt int64 `gorm:"default:0"`
}
