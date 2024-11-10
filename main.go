package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"math/rand"
	"os"
	"time"
)

func myWeightOnMars() {
	fmt.Printf("My WeightOn Mars: %v", 100*2)
}

func stringFormat() {
	fmt.Printf("%-15v $%6v\n", "Porche", "155") // %v用来表示占位符，之间可以插入间距值
	fmt.Printf("%-15v $%6v\n", "BMW", "2000000000")
}

func lightSpeed() {
	const c = 299792 // km/s
	var distance = 56000000
	fmt.Println(distance/c, "sec")
}

func spaceXToMars() {
	const spaceXSpeed = 100800 // km/h
	var distance = 96300000    // km
	fmt.Printf("SapceX reach mars should take %v days", distance/spaceXSpeed/24)
}

// Restaurant 定义一个结构来接收 YAML 数据
type Restaurant struct {
	Items []string `yaml:"restaurants"`
}

func readRestaurants() {
	data, err := os.ReadFile("data/restaurants.yaml")
	if err != nil {
		log.Fatalf("读取YAML文件失败: %v", err)
	}
	var restList Restaurant
	err = yaml.Unmarshal(data, &restList)
	if err != nil {
		log.Fatalf("解析YAML文件失败: %v", err)
	}
	// 检查切片是否为空
	if len(restList.Items) == 0 {
		log.Fatal("YAML文件中的items为空")
	}

	// 随机选择一个元素
	rng := rand.New(rand.NewSource(time.Now().UnixNano())) // 使用新的随机数生成器
	randomItem := restList.Items[rng.Intn(len(restList.Items))]

	fmt.Printf("随机选择的餐馆: %s\n", randomItem)
}

func main() {
	readRestaurants()
}
