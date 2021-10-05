package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var db *gorm.DB

type Product struct {
	gorm.Model
	Code  		string
	Price 		uint
	Category 	uint
}

type Category struct {
	gorm.Model
	Name string
}

func main() {
	Sample("Open", Open)
	Sample("Migrate", Migrate)
	Sample("Create", Create)
	Sample("CreateBatch", CreateBatch)
	Sample("CreateWithSelect", CreateWithSelect)
	Sample("CreateWithMap", CreateWithMap)
	Sample("CreateWithClauses", CreateWithClauses)
	Sample("Find", Find)
	Sample("FindWithCond", FindWithCond)
	Sample("First", First)
	Sample("FirstWithCond", FirstWithCond)
	Sample("Take", Take)
	Sample("TakeWithCond", TakeWithCond)
	Sample("Last", Last)
	Sample("LastWithCond", LastWithCond)
	Sample("ErrRecordNotFound", ErrRecordNotFound)
	Sample("WhereWithString", WhereWithString)
	Sample("WhereWithStruct", WhereWithStruct)
	Sample("WhereWithPrimary", WhereWithPrimary)
	Sample("WhereWithMap", WhereWithMap)
	Sample("WhereWithOr", WhereWithOr)
	Sample("NotWithString", NotWithString)
	Sample("NotWithStruct", NotWithStruct)
	Sample("NotWithPrimary", NotWithPrimary)
	Sample("NotWithMap", NotWithMap)
	Sample("NotWithOr", NotWithOr)
	Sample("Select", Select)
	Sample("Order", Order)
	Sample("Limit", Limit)
	Sample("Offset", Offset)
	Sample("Distinct", Distinct)
	Sample("GroupBy", GroupBy)
	Sample("Joins", Joins)
	Sample("Scan", Scan)
	Sample("Raw", Raw)
	Sample("Update", Update)
	Sample("UpdateOneColumn", UpdateOneColumn)
	Sample("UpdateMultipleColumn", UpdateMultipleColumn)
	Sample("UpdateWithSelect", UpdateWithSelect)
	Sample("UpdateBatch", UpdateBatch)
	Sample("UpdateWithExpr", UpdateWithExpr)
	Sample("Delete", Delete)
	Sample("DeleteWithWhere", DeleteWithWhere)
	Sample("DeleteWithPrimary", DeleteWithPrimary)
	Sample("DeleteBatch", DeleteBatch)
	Sample("DeleteSoft", DeleteSoft)
}

func Open() {
	var err error
	path := fmt.Sprintf("%s%v.db", os.TempDir(), time.Now().Unix())
	db, err = gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("opened", db.Name())
}

func Migrate() {
	var err error

	err = db.AutoMigrate(&Category{})
	if err != nil {
		fmt.Println("error:", err)
	}

	err = db.AutoMigrate(&Product{})
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("migrated")
}

func Create() {
	product := Product{
		Code: "P-001",
		Price: 100,
	}

	result := db.Create(&product)
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	fmt.Println("created with id", product.ID)
}

func CreateBatch() {
	product1 := Product{
		Code: "P-002",
		Price: 50,
	}

	product2 := Product{
		Code: "P-003",
		Price: 75,
	}

	var products = []Product{product1, product2}
	result := db.Create(&products)
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	for _, product := range products {
		fmt.Println("created with id", product.ID)
	}
}

func CreateWithSelect() {
	product := Product{
		Code: "P-004",
		Price: 50,
	}

	result := db.Select("Code").Create(&product)
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	fmt.Println("created with id", product.ID)
}

func CreateWithMap() {
	var product Product
	result := db.Model(&product).Create(map[string]interface{}{
		"Code": "P-005",
		"Price": 15,
	})
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	fmt.Println("created with id", product.ID)
}

func CreateWithClauses() {
	var product Product
	result := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&product)
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	fmt.Println("created with id", product.ID)
}

func Find() {
	var products []Product
	result := db.Find(&products)
	if result.Error != nil {
		fmt.Println("error:", result.Error)
	} else {
		PrintJson(products)
	}
}

func FindWithCond() {
	var products []Product
	result := db.Find(&products, "Code = ?", "P-001")
	if result.Error != nil {
		fmt.Println("error:", result.Error)
	} else {
		PrintJson(products)
	}
}

func First() {
	var product Product
	result := db.First(&product)
	if result.Error != nil {
		fmt.Println("error:", result.Error)
	} else {
		PrintJson(product)
	}
}

func FirstWithCond() {
	var product Product

	fmt.Println("with primary field:")
	result1 := db.First(&product, 1)
	if result1.Error != nil {
		fmt.Println("error:", result1.Error)
	} else {
		PrintJson(product)
	}

	fmt.Println("with other field:")
	result2 := db.First(&product, "Code = ?", "P-001")
	if result2.Error != nil {
		fmt.Println("error:", result2.Error)
	} else {
		PrintJson(product)
	}
}

func Take() {
	var product Product
	result := db.Take(&product)
	if result.Error != nil {
		fmt.Println("error:", result.Error)
	} else {
		PrintJson(product)
	}
}

func TakeWithCond() {
	var product Product

	fmt.Println("with primary field:")
	result1 := db.Take(&product, 1)
	if result1.Error != nil {
		fmt.Println("error:", result1.Error)
	} else {
		PrintJson(product)
	}

	fmt.Println("with other field:")
	result2 := db.Take(&product, "Code = ?", "P-001")
	if result2.Error != nil {
		fmt.Println("error:", result2.Error)
	} else {
		PrintJson(product)
	}
}

func Last() {
	var product Product
	result := db.Last(&product)
	if result.Error != nil {
		fmt.Println("error:", result.Error)
	} else {
		PrintJson(product)
	}
}

func LastWithCond() {
	var product Product

	fmt.Println("with primary field:")
	result1 := db.Last(&product, 1)
	if result1.Error != nil {
		fmt.Println("error:", result1.Error)
	} else {
		PrintJson(product)
	}

	fmt.Println("with other field:")
	result2 := db.Last(&product, "Code = ?", "P-001")
	if result2.Error != nil {
		fmt.Println("error:", result2.Error)
	} else {
		PrintJson(product)
	}
}

func ErrRecordNotFound() {
	var product Product
	result := db.Last(&product, 100)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Println("not found")
		return
	}

	PrintJson(product)
}

func WhereWithString() {
	var products []Product
	result := db.Where("Code LIKE ?", "P-%").Find(&products)
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	PrintJson(products)
}

func WhereWithStruct() {
	var products []Product
	result := db.Where(&Product{Code: "P-002"}).Find(&products)
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	PrintJson(products)
}

func WhereWithPrimary() {
	var products []Product
	result := db.Where([]int64{1, 3, 5}).Find(&products)
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	PrintJson(products)
}

func WhereWithMap() {
	var product Product
	result := db.Where(map[string]interface{}{"Code": "P-004"}).First(&product)
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	PrintJson(product)
}

func WhereWithOr() {
	var products []Product
	result := db.
		Where("Code = 'P-005'").
		Or(map[string]interface{}{"Code": "P-002"}).
		Or("Price > 50").
		Find(&products)
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	PrintJson(products)
}


func NotWithString() {
	var products []Product
	result := db.Not("Code = ?", "P-003").Find(&products)
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	PrintJson(products)
}

func NotWithStruct() {
	var products []Product
	result := db.Not(&Product{Code: "P-002"}).Find(&products)
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	PrintJson(products)
}

func NotWithPrimary() {
	var products []Product
	result := db.Not([]int64{2, 4}).Find(&products)
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	PrintJson(products)
}

func NotWithMap() {
	var product Product
	result := db.Not(map[string]interface{}{"Code": "P-004", "Price": 50}).First(&product)
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	PrintJson(product)
}

func NotWithOr() {
	var products []Product
	result := db.
		Not("Code = 'P-002'").
		Or(map[string]interface{}{"Code": "P-002"}).
		Or("Price > 50").
		Find(&products)
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	PrintJson(products)
}

func Select() {
	var productCodes []string
	result := db.Table("products").Select("Code").Find(&productCodes)
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	PrintJson(productCodes)
}

func Order() {
	var products []Product
	result := db.Order("Price desc").Order("Code").Find(&products)
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	PrintJson(products)
}

func Limit() {
	var products []Product
	result := db.Limit(2).Find(&products)
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	PrintJson(products)
}

func Offset() {
	var products []Product
	result := db.Limit(1).Offset(1).Find(&products)
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	PrintJson(products)
}

func Distinct() {
	var products []Product
	result := db.Distinct("Code").Order("Code").Find(&products)
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	PrintJson(products)
}

func GroupBy() {
	type GroupResult struct {
		Code  string
		Total int
	}

	var groupResult GroupResult
	result := db.
		Model(&Product{}).
		Select("Code, sum(Price) as Total").
		Group("Code").
		First(&groupResult)
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	PrintJson(groupResult)
}

func Joins() {
	type JoinResult struct {
		Code	string
		Name	string
	}

	var joinResult JoinResult
	result := db.
		Model(&Product{}).
		Select("products.Code, Categories.Name").
		Joins("left join Categories on Categories.ID = products.Category").
		Scan(&joinResult)
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	PrintJson(joinResult)
}

func Scan() {
	type ScanResult struct {
		Code 	string
		Price  	uint
	}

	var scanResult ScanResult
	result := db.
		Table("Products").
		Select("Code", "Price").
		Where("Code = ?", "P-002").
		Scan(&scanResult)
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	PrintJson(scanResult)
}

func Raw() {
	type ScanResult struct {
		Code 	string
		Price  	uint
	}

	var scanResult ScanResult
	result := db.Raw("SELECT Code, Price FROM Products WHERE Code = ?", "P-003").Scan(&scanResult)
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	PrintJson(scanResult)
}

func Update() {
	var product Product
	result := db.First(&product)
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	product.Price += 25
	result = db.Save(&product)
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	PrintJson(product)
}

func UpdateOneColumn() {
	var product Product
	result := db.
		Model(&product).
		Where("Code = ?", "P-003").
		Update("Price", 125)
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	PrintJson(product)
}

func UpdateMultipleColumn() {
	var product Product
	result := db.
		Model(&product).
		Where("Code = ?", "P-003").
		Updates(Product{Price: 150, Category: 2})
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	PrintJson(product)
}

func UpdateWithSelect() {
	var product Product
	result := db.
		Model(&product).
		Select("Price", "Category").
		Where("Code = ?", "P-003").
		Updates(Product{Price: 150, Category: 3})
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	PrintJson(product)
}

func UpdateBatch() {
	var product Product
	result := db.
		Model(&product).
		Where("Category = ?", 0).
		Updates(Product{Category: 4})
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	PrintJson(product)
}

func UpdateWithExpr() {
	var product Product
	result := db.
		Model(&product).
		Where("Code = ?", "P-004").
		Update("Price", gorm.Expr("Price * ? + ?", 2, 10))
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	PrintJson(product)
}

func Delete() {
	var product Product
	result := db.Last(&product)
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	result = db.Delete(&product)
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	fmt.Println("deleted", product.ID)
}

func DeleteWithWhere() {
	result := db.Where("Code = ?", "P-004").Delete(&Product{})
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	fmt.Printf("deleted %v product(s)\n", result.RowsAffected)
}

func DeleteWithPrimary() {
	result := db.Delete(&Product{}, 3)
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	fmt.Printf("deleted %v product(s)\n", result.RowsAffected)
}

func DeleteBatch() {
	result := db.Where("Price > ?", 100).Delete(&Product{})
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	fmt.Printf("deleted %v product(s)\n", result.RowsAffected)
}

func DeleteSoft() {
	type DeletableItem struct {
		ID      int				`gorm:"primaryKey"`
		Name    string
		Deleted gorm.DeletedAt	`gorm:"index"`
	}

	db.AutoMigrate(&DeletableItem{})

	item := DeletableItem{
		ID: 1,
		Name: "item-1",
	}

	result := db.Create(&item)
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	result = db.Where("ID = 1").Delete(&item)
	if result.Error != nil {
		fmt.Println("error:", result.Error)
		return
	}

	PrintJson(item)
}

func PrintJson(data interface{}) {
	j, err := json.Marshal(data)
	if err != nil {
		fmt.Println("json error")
		return
	}
	fmt.Printf("%s\n", j)
}

func Sample(name string, fn func())  {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}