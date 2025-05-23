package main

import "fmt"

// Animal 接口定义
type Animal interface {
	Speak() string
	Move() string
}

// Dog 结构体
type Dog struct {
	Name string
}

// Cat 结构体
type Cat struct {
	Name string
}

// Bird 结构体
type Bird struct {
	Name string
}

// Dog 实现 Animal 接口的 Speak 方法
func (d Dog) Speak() string {
	// 请在此处实现狗叫的声音
	return "汪汪"
}

// Dog 实现 Animal 接口的 Move 方法
func (d Dog) Move() string {
	// 请在此处实现狗的移动方式
	return "狗跑步"
}

// Cat 实现 Animal 接口的 Speak 方法
func (c Cat) Speak() string {
	// 请在此处实现猫叫的声音
	return "喵喵"
}

// Cat 实现 Animal 接口的 Move 方法
func (c Cat) Move() string {
	// 请在此处实现猫的移动方式
	return "猫跑步"
}

// Bird 实现 Animal 接口的 Speak 方法
func (b Bird) Speak() string {
	// 请在此处实现鸟叫的声音
	return "叽叽喳喳"
}

// Bird 实现 Animal 接口的 Move 方法
func (b Bird) Move() string {
	// 请在此处实现鸟的移动方式
	return "鸟飞"
}

func main() {
	animals := []Animal{
		Dog{Name: "Buddy"},
		Cat{Name: "Whiskers"},
		Bird{Name: "Tweety"},
	}

	for _, animal := range animals {
		switch a := animal.(type) {
		case Dog:
			fmt.Printf("%s 说: %s, %s 移动方式: %s\n", a.Name, animal.Speak(), a.Name, animal.Move())
		case Cat:
			fmt.Printf("%s 说: %s, %s 移动方式: %s\n", a.Name, animal.Speak(), a.Name, animal.Move())
		case Bird:
			fmt.Printf("%s 说: %s, %s 移动方式: %s\n", a.Name, animal.Speak(), a.Name, animal.Move())
		default:
			fmt.Println("未知动物类型")
		}
	}
}
