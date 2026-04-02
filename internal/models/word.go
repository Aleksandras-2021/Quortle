package models

type Word struct {
	ID   uint   `gorm:"primaryKey"`
	Word string `gorm:"size:4;not null;unique"`
}

func (Word) TableName() string {
	return "quortle.words"
}
