package versions

import (
	"gorm.io/gorm"

	"github.com/nekizz/telegram-bot/internal/model"
)

func Version20230808080810(tx *gorm.DB) error {
	type User struct {
		model.Model
		FirstName     string `gorm:"type:varchar(128)"`
		LastName      string `gorm:"type:varchar(128)"`
		DateOfBirth   string `gorm:"type:varchar(64)"`
		Email         string `gorm:"unique;type:varchar(255)"`
		ContactNumber string `gorm:"type:varchar(16)"`
		Notes         string `gorm:"type:varchar(1024)"`
		Status        int    `gorm:"type:tinyint;not null"`
	}

	type JobInfo struct {
		model.Model
		Company  string `gorm:"type:varchar(128)"`
		Checkin  string `gorm:"type:varchar(32)"`
		Checkout string `gorm:"type:varchar(32)"`
		Status   int    `gorm:"type:tinyint;not null"`
	}

	return tx.AutoMigrate(&User{}, &JobInfo{})
}

func Rollback20230808080810(tx *gorm.DB) error {
	if err := tx.Migrator().DropTable(&model.User{}); err != nil {
		return err
	}

	if err := tx.Migrator().DropTable(&model.User{}); err != nil {
		return err
	}

	return nil
}
