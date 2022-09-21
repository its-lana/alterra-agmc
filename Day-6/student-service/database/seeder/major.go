package seeder

import (
	"log"
	"time"

	"student-service/internal/model"
	"student-service/internal/pkg/enum"

	"gorm.io/gorm"
)

func majorSeeder(db *gorm.DB) {
	now := time.Now()
	var majors = []model.Major{
		{
			Name: enum.Major.String(1),
			Common: model.Common{
				ID:        1,
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
		{
			Name: enum.Major.String(2),
			Common: model.Common{
				ID:        2,
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
		{
			Name: enum.Major.String(3),
			Common: model.Common{
				ID:        3,
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
	}
	if err := db.Create(&majors).Error; err != nil {
		log.Printf("cannot seed data majors, with error %v\n", err)
	}
	log.Println("success seed data majors")
}
