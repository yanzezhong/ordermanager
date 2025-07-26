package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func randomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func randomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	db := client.Database("ordermanager")

	// 1. 部门
	deptCol := db.Collection("dept")
	var depts []interface{}
	for i := 1; i <= 10; i++ {
		var parentId int64
		if i == 1 {
			parentId = 1 // 或 0，按你的业务需求
		} else {
			parentId = int64(rand.Intn(i-1) + 1)
		}
		depts = append(depts, bson.M{
			"_id":      int64(i),
			"code":     fmt.Sprintf("DPT%03d", i),
			"name":     fmt.Sprintf("部门%d", i),
			"parentId": parentId,
			"status":   int64(1),
			"sort":     int64(i),
			"updateAt": time.Now(),
			"createAt": time.Now().Add(-time.Duration(rand.Intn(1000)) * time.Hour),
		})
	}
	_, err = deptCol.InsertMany(ctx, depts)
	if err != nil {
		log.Fatal("插入部门失败:", err)
	}
	fmt.Println("部门测试数据插入完成")

	// 2. 菜单
	menuCol := db.Collection("menu")
	var menus []interface{}
	for i := 1; i <= 10; i++ {
		var parentId int64
		if i == 1 {
			parentId = 1 // 或 0，按你的业务需求
		} else {
			parentId = int64(rand.Intn(i-1) + 1)
		}
		menus = append(menus, bson.M{
			"id":         int64(i),
			"parentId":   parentId,
			"name":       fmt.Sprintf("菜单%d", i),
			"routeName":  fmt.Sprintf("route%d", i),
			"routePath":  fmt.Sprintf("/path%d", i),
			"component":  "Layout",
			"perm":       "*",
			"visible":    int32(1),
			"sort":       int32(i),
			"icon":       "el-icon-menu",
			"redirect":   "",
			"keepAlive":  int32(0),
			"alwaysShow": int32(0),
			"params":     []bson.M{{"key": "k", "value": "v"}},
			"updateAt":   time.Now(),
			"createAt":   time.Now().Add(-time.Duration(rand.Intn(1000)) * time.Hour),
		})
	}
	_, err = menuCol.InsertMany(ctx, menus)
	if err != nil {
		log.Fatal("插入菜单失败:", err)
	}
	fmt.Println("菜单测试数据插入完成")

	// 3. 商品
	productCol := db.Collection("product")
	var products []interface{}
	for i := 1; i <= 10; i++ {
		products = append(products, bson.M{
			"_id":           fmt.Sprintf("P%03d", i),
			"name":          fmt.Sprintf("商品%d", i),
			"price":         bson.M{"terminal": randomFloat(10, 100), "wholesale": randomFloat(5, 50), "cost": randomFloat(3, 30), "srp": randomFloat(10, 120), "warning": randomFloat(1, 10)},
			"specification": rand.Intn(5) + 1,
			"isActive":      rand.Intn(2) == 0,
			"image":         "",
			"tag":           "标签",
			"brandId":       fmt.Sprintf("B%02d", rand.Intn(5)+1),
			"nickName":      randomString(6),
			"barCode":       fmt.Sprintf("BC%06d", rand.Intn(100000)),
			"updateAt":      time.Now(),
			"createAt":      time.Now().Add(-time.Duration(rand.Intn(1000)) * time.Hour),
		})
	}
	_, err = productCol.InsertMany(ctx, products)
	if err != nil {
		log.Fatal("插入商品失败:", err)
	}
	fmt.Println("商品测试数据插入完成")

	// 4. 订单
	orderCol := db.Collection("order")
	var orders []interface{}
	for i := 1; i <= 10; i++ {
		var orderProducts []bson.M
		for j := 1; j <= rand.Intn(3)+1; j++ {
			orderProducts = append(orderProducts, bson.M{
				"productId":   fmt.Sprintf("P%03d", rand.Intn(10)+1),
				"productName": fmt.Sprintf("商品%d", rand.Intn(10)+1),
				"price":       randomFloat(10, 100),
				"count":       rand.Intn(5) + 1,
			})
		}
		orders = append(orders, bson.M{
			"_id":         fmt.Sprintf("O%03d", i),
			"products":    orderProducts,
			"shopId":      fmt.Sprintf("S%02d", rand.Intn(5)+1),
			"shopName":    fmt.Sprintf("门店%d", rand.Intn(5)+1),
			"address":     fmt.Sprintf("地址%d号", i),
			"state":       rand.Intn(4),
			"payment":     rand.Intn(3),
			"purchaserId": fmt.Sprintf("U%02d", rand.Intn(10)+1),
			"driverId":    fmt.Sprintf("D%02d", rand.Intn(5)+1),
			"picture":     "",
			"updateAt":    time.Now(),
			"createAt":    time.Now().Add(-time.Duration(rand.Intn(1000)) * time.Hour),
		})
	}
	_, err = orderCol.InsertMany(ctx, orders)
	if err != nil {
		log.Fatal("插入订单失败:", err)
	}
	fmt.Println("订单测试数据插入完成")
}
