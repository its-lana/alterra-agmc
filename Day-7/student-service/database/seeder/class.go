package seeder

import (
	"log"
	"time"

	"student-service/internal/model"
	"student-service/internal/pkg/enum"

	"gorm.io/gorm"
)

func classSeeder(db *gorm.DB) {
	now := time.Now()
	var classes = []model.Class{
		{
			Name: enum.Class.String(1),
			Common: model.Common{
				ID:        1,
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
		{
			Name: enum.Class.String(2),
			Common: model.Common{
				ID:        2,
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
	}
	if err := db.Create(&classes).Error; err != nil {
		log.Printf("cannot seed data classes, with error %v\n", err)
	}
	log.Println("success seed data classes")
}
