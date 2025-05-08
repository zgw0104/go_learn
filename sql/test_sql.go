package main

import (
	"gorm.io/gorm"
)

import (
	"gorm.io/driver/mysql"
)

type User struct {
	Id    uint `gorm:"primarykey"`
	Name  string
	Age   int
	Email string `gorm:"type:varchar(255)"`
}

type Class struct {
	Id     uint `gorm:"primarykey"`
	Name   string
	counts int
}

type User2 struct {
	gorm.Model
	Name      string
	CompanyID int
	Company   Company `gorm:"foreignKey:CompanyID;references:ID"`
}

func (User) TableName() string {
	return "user2"
}

type Company struct {
	ID   int
	Name string
}

func init() {

}

func main() {
	var dsn = "root:64220392@tcp(127.0.0.1:3306)/go?charset=utf8"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			panic(err)
		}
		sqlDB.Close()
	}()

	//err = db.AutoMigrate(&Class{})
	//if err != nil {
	//	panic(err)
	//}
	//
	//var class = Class{Name: "y1", counts: 10}
	//db.Create(&class)
	//
	//var classes = []Class{
	//	{Name: "y2", counts: 12},
	//	{Name: "y3", counts: 13},
	//	{Name: "y4", counts: 14},
	//}
	//
	//db.Create(&classes)
	//
	//var classes2 []Class
	//db.Find(&classes2)
	//fmt.Println(classes2)
	//
	//var cls Class
	//db.Where("id = ?", 3).Find(&cls)
	//fmt.Println(cls)
	//
	//var user User
	//db.Model(&User{}).Where("Name = ?", "zgw").Find(&user)
	//fmt.Println(user)
	//
	//var dcls Class
	//db.Model(&Class{}).Where("id = ?", 3).Delete(&dcls)

	//err = db.AutoMigrate(&User{})
	//if err != nil {
	//	panic(err)
	//

	//插入单条数据
	//var user = User{Age: 10, Name: "jq", Email: "my@gmail.com"}
	//db.Create(&user)

	////插入批量数据
	//var users = []User{
	//	{Name: "aa", Age: 1, Email: "aa@gmail.com"},
	//	{Name: "bb", Age: 1, Email: "aa@gmail.com"},
	//	{Name: "cc", Age: 1, Email: "aa@gmail.com"},
	//	{Name: "dd", Age: 1, Email: "aa@gmail.com"},
	//	{Name: "ee", Age: 1, Email: "aa@gmail.com"},
	//}
	//db.Create(&users)

	//查询全部记录
	//var users []User
	//db.Find(&users)
	//fmt.Println(users)

	//条件查询
	//var users []User
	//db.Where("Name = ?", "jq").Find(&users)
	//fmt.Println(users)

	//查询第一条记录和最后一条记录
	//var FirstUser User
	//var FinalUser User
	//
	//db.First(&FirstUser)
	//db.Last(&FinalUser)
	//fmt.Println("FirstUser:", FirstUser)
	//fmt.Println("FinalUser:", FinalUser)

	//查询记录总数
	//var users []User
	//var totalsize int64
	//db.Find(&users).Count(&totalsize)
	//fmt.Println("total size:", totalsize)

	////按照主键修改记录
	//var user User
	//db.Where("name = ?", "gw").First(&user)
	//user.Name = "zgw"
	//db.Save(&user)

	//按照字段修改
	//var user User
	//db.Model(&user).Where("Name = ?", "my").Update("Name", "lmy")

	//修改多个字段
	//var user User
	//var field = map[string]interface{}{"name": "qwwer", "age": 123123}
	//db.Model(&user).Where("id = ?", 8).Updates(field)

	//删除记录
	//var user User
	//db.Where("id = ?", 7).Delete(&user)

	////测试关联查询
	//db.AutoMigrate(&User2{}, &Company{})
	//
	//// 因为已经指定了约束关系，User.companyId 字段会自动被 Company.ID 字段填充。
	//user := User2{
	//	Model: gorm.Model{},
	//	Name:  "amos",
	//	Company: Company{
	//		ID:   12,
	//		Name: "google",
	//	},
	//}
	//
	//// 关联插入。在插入 user 记录时，也插入 compancy 记录。
	//db.Create(&user)
	//
	//var result User2
	//db.Model(&User2{}).Preload("Company").Find(&result, "name = ?", user.Name)
	//data, _ := json.Marshal(result)
	//fmt.Println(string(data))

}
