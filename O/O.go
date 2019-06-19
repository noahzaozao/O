package O

import "log"

// 实体接口
type Entity interface {
	GetID() int
	GetName() string
}

// 用户结构体
type User struct {
	ID int
	Name string
}

// 用户实现实体接口 GetID()
func (user *User) GetID() int {
	log.Println("User.GetID() => ", user.ID)
	return user.ID
}

// 用户实现实体接口 GetName()
func (user *User) GetName() string {
	log.Println("User.GetName() => ", user.Name)
	return user.Name
}

// 用户信息
func (user *User) Info() {
	log.Println("User info => ", user.ID, user.Name)
}

// 物件结构体
type Item struct {
	ID int
	Name string
	User User
}

// 物件实现实体接口 GetID()
func (item *Item) GetID() int {
	log.Println("Item.GetID() => ", item.ID)
	return item.ID
}

// 物件实现实体接口 GetName()
func (item *Item) GetName() string {
	log.Println("Item.GetName() => ", item.Name)
	return item.Name
}

// 物件信息
func (item *Item) Info() {
	log.Println("Item info => ", item.ID, item.Name)
}

// 实体接口信息
func EntityInfo(entity Entity) {
	log.Println("Entity info => ", entity.GetID(), entity.GetName())
}

// 包方法调用
func Test() {
	log.Println("O.Test()")
}

