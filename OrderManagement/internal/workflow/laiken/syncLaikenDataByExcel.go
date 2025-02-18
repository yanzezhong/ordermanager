package laiken

import (
	"OrderManagement/OrderManagement/internal/model"
	"OrderManagement/OrderManagement/internal/svc"
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/xuri/excelize/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewLaikenDataProcess(ctx context.Context, svcCtx *svc.ServiceContext) {
	// 1.查找excel位置
	f, err := excelize.OpenFile("data/nongfu/2024.xlsx")
	if err != nil {
		log.Fatalf("Failed to open Excel file: %v", err)
		return
	}
	defer f.Close()
	// 2. 提取excel数据

	details := []*model.InvoiceDetail{}

	//  2.1 获取 Sheet1 上所有单元格
	rows, err := f.GetRows("销售明细查询")
	if err != nil {
		log.Fatalf("Failed to get rows: %v", err)
		return
	}
	for index, row := range rows {
		if index < 5 {
			continue
		}
		// 2.2 生成 InvoiceDetail
		detail, err := parseRow(row)
		if err != nil {
			log.Printf("Error parsing row %d: %s", index, err.Error())
			continue
		}
		details = append(details, &detail)
	}

	//  3. 数据插入
	err = svcCtx.InvoiceDetail.InsertMany(ctx, details)
	if err != nil {
		return
	}

}

// 从 excel导入数据
func GetLaikenDataFromExcel() {
	f, err := excelize.OpenFile("data/nongfu/2024.xlsx")
	if err != nil {
		log.Fatalf("Failed to open Excel file: %v", err)
		return
	}
	defer f.Close()

	// // 获取工作表中指定单元格的值
	// cell, err := f.GetCellValue("销售明细查询", "B2")
	// if err != nil {
	// 	log.Fatalf("Failed to get cell value: %v", err)
	// 	return
	// }
	// fmt.Println(cell)

	err, collection := InitMongo()
	if err != nil {
		log.Fatalf("Failed to initialize MongoDB: %v", err)
		return
	}

	students := []interface{}{}

	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows("销售明细查询")
	if err != nil {
		log.Fatalf("Failed to get rows: %v", err)
		return
	}
	for index, row := range rows {
		if index < 5 {
			continue
		}
		detail, err := parseRow(row)
		if err != nil {
			log.Printf("Error parsing row %d: %s", index, err.Error())
			continue
		}
		students = append(students, detail)
	}
	_, err = collection.InsertMany(context.TODO(), students)
	if err != nil {
		log.Fatalf("Failed to insert data into MongoDB: %v", err)
	}
	log.Println("Data inserted successfully")
}

// 生成 InvoiceDetail
func parseRow(row []string) (model.InvoiceDetail, error) {
	if len(row) < 21 {
		for len(row) < 21 {
			row = append(row, "")
		}
	}
	var detail model.InvoiceDetail

	// 单据日期
	detail.DocumentDate = row[1]

	// 单据编号
	detail.DocumentNumber = row[2]

	// 单据类型
	detail.DocumentType = row[3]

	// 客户
	detail.Customer = row[4]

	// 客户级别
	detail.CustomerLevel = row[5]

	// 来源订单
	detail.SourceOrder = row[6]

	// 经手人
	detail.Handler = row[7]

	// 商品名称
	detail.ProductName = row[8]

	// 规格
	detail.Specification = row[9]

	// 销售数量
	salesQuantity, err := strconv.ParseFloat(row[10], 64)
	if err != nil {
		return detail, err
	}
	detail.SalesQuantity = salesQuantity

	// 销售规格
	detail.SalesSpecification = row[11]
	// 单价
	unitPrice, err := strconv.ParseFloat(row[12], 64)
	if err != nil {
		return detail, err
	}
	detail.UnitPrice = unitPrice

	// 金额
	amount, err := strconv.ParseFloat(row[13], 64)
	if err != nil {
		return detail, err
	}
	detail.Amount = amount

	// 销售收入
	salesRevenue, err := strconv.ParseFloat(row[14], 64)
	if err != nil {
		return detail, err
	}
	detail.SalesRevenue = salesRevenue

	// 重量（kg）
	weight, err := strconv.ParseFloat(row[15], 64)
	if err != nil {
		return detail, err
	}
	detail.Weight = weight

	// 商品行属性
	detail.LineItemAttribute = row[16]

	// 明细备注
	detail.DetailRemark = row[17]

	// 单据备注
	detail.DocumentRemark = row[18]

	// 数量
	if row[19] != "" {
		quantity, err := strconv.ParseFloat(row[19], 64)
		if err != nil {
			return detail, err
		}
		detail.Quantity = quantity

	}

	if row[20] != "" {
		// 总金额
		totalAmount, err := strconv.ParseFloat(row[20], 64)
		if err != nil {
			return detail, err
		}
		detail.TotalAmount = totalAmount
	}

	return detail, nil
}

var (
	ErrNotFound = errors.New("not found")
)

func InitMongo() (error, *mongo.Collection) {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	collection := client.Database("test").Collection("nongfu")
	return err, collection
}
