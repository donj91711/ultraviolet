package dao

type Book struct {
	ID              uint `gorm:"primaryKey"`
	Title           string
	Author          string
	PublicationDate string `gorm:"type:date"`
	ISBN            string
	PageCount       int
}
