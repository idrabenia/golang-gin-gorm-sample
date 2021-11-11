package main

//import "fmt"
//import "rsc.io/quote"
//import "time"
import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//
//func main() {
//    message := "Hello, Go!"
//
//    fmt.Println(message)
//    fmt.Println(quote.Hello())
//    fmt.Println(time.Now())
//}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main1() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 5)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	fmt.Println(product)

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	db.Delete(&product, 1)
}
