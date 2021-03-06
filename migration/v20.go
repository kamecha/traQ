package migration

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

// v20 パーミッション周りの調整
func v20() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20",
		Migrate: func(db *gorm.DB) error {
			deletedPermissions := []string{
				"get_heartbeat",
				"post_heartbeat",
			}
			for _, v := range deletedPermissions {
				if err := db.Delete(v20RolePermission{}, v20RolePermission{Permission: v}).Error; err != nil {
					return err
				}
			}
			return nil
		},
	}
}

type v20RolePermission struct {
	Role       string `gorm:"type:varchar(30);not null;primary_key"`
	Permission string `gorm:"type:varchar(30);not null;primary_key"`
}

func (*v20RolePermission) TableName() string {
	return "user_role_permissions"
}
