package database

import (
	"UltravioletTest/dao"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenDB(username, password, host, port, dbname string) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"

	// try to open the database several times because sometimes Docker doesn't have MySQL ready to go yet
	maxRetries := 10
	retryCount := 0
	retryInterval := 5 * time.Second
	for {
		fmt.Println("TRY: ", retryCount)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			fmt.Println("DATABASE OPEN")
			break
		}

		retryCount++

		if retryCount >= maxRetries {
			fmt.Printf("Max retries reached. Unable to connect to the database: %v\n", err)
			return nil, err
		}

		fmt.Printf("Error connecting to the database (retry %d/%d): %v\n", retryCount, maxRetries, err)

		time.Sleep(retryInterval)
	}
	return db, nil
}

func GetBookList(db *gorm.DB) ([]dao.Book, error) {
	var books []dao.Book
	result := db.Find(&books)
	if result.Error != nil {
		fmt.Println("Error in GetBookList")
		return nil, result.Error
	}
	return books, nil
}

func GetBook(db *gorm.DB, bookID string) (dao.Book, error) {
	var book dao.Book
	result := db.Where("id = ?", bookID).First(&book)
	if result.Error != nil {
		return dao.Book{}, result.Error
	}
	return book, nil
}

func AddBook(db *gorm.DB, book dao.Book) (dao.Book, error) {
	result := db.Create(&book)
	if result.Error != nil {
		return dao.Book{}, result.Error
	}
	return book, nil
}

func UpdateBook(db *gorm.DB, book dao.Book, bookID string) (dao.Book, error) {
	var b dao.Book
	result := db.Where("id = ?", bookID).First(&b)
	if result.Error != nil {
		return dao.Book{}, result.Error
	}

	b.Title = book.Title
	b.Author = book.Author
	b.ISBN = book.ISBN
	b.PageCount = book.PageCount
	b.PublicationDate = book.PublicationDate

	result = db.Save(&b)
	if result.Error != nil {
		return dao.Book{}, result.Error
	}

	return b, nil
}

func DeleteBook(db *gorm.DB, bookID string) error {
	var existingBook dao.Book
	result := db.Where("id = ?", bookID).First(&existingBook)
	if result.Error != nil {
		return result.Error
	}

	result = db.Delete(&existingBook)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
