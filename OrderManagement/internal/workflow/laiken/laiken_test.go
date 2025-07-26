package laiken

import (
	"OrderManagement/OrderManagement/internal/model"
	"fmt"

	"reflect"
	"sort"
	"testing"
)

func Test_generateActive(t *testing.T) {
	type args struct {
		list   []*model.InvoiceDetail
		result []InvoiceMonth
	}
	tests := []struct {
		name string
		args args
		want []InvoiceMonth
	}{
		{
			name: "Test Case 1",
			args: args{
				list: []*model.InvoiceDetail{
					{DocumentDate: "2024/12/01", Handler: "Alice", SalesRevenue: 100.0, Customer: "Shop1"},
					{DocumentDate: "2024/12/15", Handler: "Alice", SalesRevenue: 150.0, Customer: "Shop2"},
					{DocumentDate: "2025/12/01", Handler: "Alice", SalesRevenue: 200.0, Customer: "Shop2"},
					{DocumentDate: "2025/12/15", Handler: "Alice", SalesRevenue: 250.0, Customer: "Shop3"},
					{DocumentDate: "2024/12/01", Handler: "Bob", SalesRevenue: 100.0, Customer: "Shop4"},
					{DocumentDate: "2024/12/15", Handler: "Bob", SalesRevenue: 150.0, Customer: "Shop5"},
					{DocumentDate: "2025/12/01", Handler: "Bob", SalesRevenue: 200.0, Customer: "Shop5"},
					{DocumentDate: "2025/12/15", Handler: "Bob", SalesRevenue: 250.0, Customer: "Shop6"},
				},
				result: []InvoiceMonth{
					{Year: 2024, Month: 12, Handler: "Alice", SalesRevenue: 250.0},
					{Year: 2025, Month: 12, Handler: "Alice", SalesRevenue: 450.0},
					{Year: 2024, Month: 12, Handler: "Bob", SalesRevenue: 250.0},
					{Year: 2025, Month: 12, Handler: "Bob", SalesRevenue: 450.0},
				},
			},
			want: []InvoiceMonth{
				{
					Year:          2024,
					Month:         12,
					Handler:       "Alice",
					SalesRevenue:  250.0,
					ActiveShopNum: 2,
					ActiveShop:    []string{"Shop1", "Shop2"},
				},
				{
					Year:          2025,
					Month:         12,
					Handler:       "Alice",
					SalesRevenue:  450.0,
					ActiveShopNum: 2,
					ActiveShop:    []string{"Shop2", "Shop3"},
					DisActiveShop: []string{"Shop1"},
					NewActiveShop: []string{"Shop3"},
				},
				{
					Year:          2024,
					Month:         12,
					Handler:       "Bob",
					SalesRevenue:  250.0,
					ActiveShopNum: 2,
					ActiveShop:    []string{"Shop4", "Shop5"},
				},
				{
					Year:          2025,
					Month:         12,
					Handler:       "Bob",
					SalesRevenue:  450.0,
					ActiveShopNum: 2,
					ActiveShop:    []string{"Shop5", "Shop6"},
					DisActiveShop: []string{"Shop4"},
					NewActiveShop: []string{"Shop6"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateActive(tt.args.list, tt.args.result)
			for _, person := range got {
				sort.Strings(person.ActiveShop)
			}

			if !reflect.DeepEqual(got, tt.want) {
				for _, person := range got {
					fmt.Println(person)
				}
				t.Errorf("generateActive() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func Test_converInvoiceShopMonth(t *testing.T) {
	type args struct {
		list []*model.InvoiceDetail
	}
	tests := []struct {
		name string
		args args
		want []*InvoiceShopMonth
	}{
		{
			name: "Test Case 1: Basic Scenario",
			args: args{
				list: []*model.InvoiceDetail{
					{
						DocumentDate:  "2024/01/15",
						Customer:      "Customer A",
						CustomerLevel: "VIP",
						Handler:       "Handler 1",
						ProductName:   "Product 1",
						SalesRevenue:  1000.0,
						Quantity:      10,
						TotalAmount:   1000.0,
					},
					{
						DocumentDate:  "2024/01/20",
						Customer:      "Customer A",
						CustomerLevel: "VIP",
						Handler:       "Handler 1",
						ProductName:   "Product 2",
						SalesRevenue:  2000.0,
						Quantity:      20,
						TotalAmount:   2000.0,
					},
					{
						DocumentDate:  "2024/02/10",
						Customer:      "Customer B",
						CustomerLevel: "Regular",
						Handler:       "Handler 2",
						ProductName:   "Product 3",
						SalesRevenue:  1500.0,
						Quantity:      15,
						TotalAmount:   1500.0,
					},
				},
			},
			want: []*InvoiceShopMonth{
				{
					Year:          2024,
					Month:         1,
					Customer:      "Customer A",
					CustomerLevel: "VIP",
					Handler:       []string{"Handler 1"},
					SalesRevenue:  3000.0,
					SKUNum:        2,
					SKU:           []string{"Product 1", "Product 2"},
					DisSKU:        []string{},
					NewSKU:        []string{"Product 1", "Product 2"},
				},
				{
					Year:          2024,
					Month:         2,
					Customer:      "Customer B",
					CustomerLevel: "Regular",
					Handler:       []string{"Handler 2"},
					SalesRevenue:  1500.0,
					SKUNum:        1,
					SKU:           []string{"Product 3"},
					DisSKU:        []string{},
					NewSKU:        []string{"Product 3"},
				},
			},
		},
		{
			name: "Test Case 2: Multiple Handlers and Products",
			args: args{
				list: []*model.InvoiceDetail{
					{
						DocumentDate:  "2024/01/15",
						Customer:      "Customer A",
						CustomerLevel: "VIP",
						Handler:       "Handler 1",
						ProductName:   "Product 1",
						SalesRevenue:  1000.0,
						Quantity:      10,
						TotalAmount:   1000.0,
					},
					{
						DocumentDate:  "2024/01/20",
						Customer:      "Customer A",
						CustomerLevel: "VIP",
						Handler:       "Handler 2",
						ProductName:   "Product 2",
						SalesRevenue:  2000.0,
						Quantity:      20,
						TotalAmount:   2000.0,
					},
					{
						DocumentDate:  "2024/01/25",
						Customer:      "Customer A",
						CustomerLevel: "VIP",
						Handler:       "Handler 1",
						ProductName:   "Product 3",
						SalesRevenue:  3000.0,
						Quantity:      30,
						TotalAmount:   3000.0,
					},
				},
			},
			want: []*InvoiceShopMonth{
				{
					Year:          2024,
					Month:         1,
					Customer:      "Customer A",
					CustomerLevel: "VIP",
					Handler:       []string{"Handler 1", "Handler 2"},
					SalesRevenue:  6000.0,
					SKUNum:        3,
					SKU:           []string{"Product 1", "Product 2", "Product 3"},
					DisSKU:        []string{},
					NewSKU:        []string{"Product 1", "Product 2", "Product 3"},
				},
			},
		}, {
			name: "Test Case 4: Different Years",
			args: args{
				list: []*model.InvoiceDetail{
					{
						DocumentDate:  "2023/01/25",
						Customer:      "Customer A",
						CustomerLevel: "VIP",
						Handler:       "Handler 1",
						ProductName:   "Product 1",
						SalesRevenue:  500.0,
						Quantity:      5,
						TotalAmount:   500.0,
					},
					{
						DocumentDate:  "2024/01/15",
						Customer:      "Customer A",
						CustomerLevel: "VIP",
						Handler:       "Handler 1",
						ProductName:   "Product 2",
						SalesRevenue:  1000.0,
						Quantity:      10,
						TotalAmount:   1000.0,
					},
					{
						DocumentDate:  "2024/01/20",
						Customer:      "Customer A",
						CustomerLevel: "VIP",
						Handler:       "Handler 2",
						ProductName:   "Product 3",
						SalesRevenue:  2000.0,
						Quantity:      20,
						TotalAmount:   2000.0,
					},
				},
			},
			want: []*InvoiceShopMonth{
				{
					Year:          2023,
					Month:         1,
					Customer:      "Customer A",
					CustomerLevel: "VIP",
					Handler:       []string{"Handler 1"},
					SalesRevenue:  500.0,
					SKUNum:        1,
					SKU:           []string{"Product 1"},
					DisSKU:        []string{},
					NewSKU:        []string{"Product 1"},
				},
				{
					Year:          2024,
					Month:         1,
					Customer:      "Customer A",
					CustomerLevel: "VIP",
					Handler:       []string{"Handler 1", "Handler 2"},
					SalesRevenue:  3000.0,
					SKUNum:        2,
					SKU:           []string{"Product 2", "Product 3"},
					DisSKU:        []string{"Product 1"}, // 假设 Product 1 在 2023 年有记录，但在 2024 年没有
					NewSKU:        []string{"Product 2", "Product 3"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := ConverInvoiceShopMonth(tt.args.list)
			for _, person := range got {
				sort.Strings(person.SKU)
				sort.Strings(person.DisSKU)
				sort.Strings(person.NewSKU)

			}
			if !reflect.DeepEqual(got, tt.want) {
				for _, person := range got {
					fmt.Println(person)
				}
				t.Errorf("converInvoiceShopMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_converInvoiceShopSKUMonth(t *testing.T) {
	type args struct {
		list []*model.InvoiceDetail
	}
	tests := []struct {
		name string
		args args
		want []*InvoiceShopSKUMonth
	}{
		{
			name: "Test Case 1: Basic Scenario",
			args: args{
				list: []*model.InvoiceDetail{
					{
						DocumentDate:  "2024/01/15",
						Customer:      "Customer A",
						CustomerLevel: "VIP",
						Handler:       "Handler 1",
						ProductName:   "Product 1",
						SalesRevenue:  1000.0,
					},
					{
						DocumentDate:  "2024/01/20",
						Customer:      "Customer A",
						CustomerLevel: "VIP",
						Handler:       "Handler 1",
						ProductName:   "Product 2",
						SalesRevenue:  2000.0,
					},
					{
						DocumentDate:  "2024/02/10",
						Customer:      "Customer B",
						CustomerLevel: "Regular",
						Handler:       "Handler 2",
						ProductName:   "Product 3",
						SalesRevenue:  1500.0,
					},
				},
			},
			want: []*InvoiceShopSKUMonth{
				{
					Year:          2024,
					Month:         1,
					Customer:      "Customer A",
					CustomerLevel: "VIP",
					Handler:       "Handler 1",
					SKU:           "Product 1",
					SalesRevenue:  1000.0,
				},
				{
					Year:          2024,
					Month:         1,
					Customer:      "Customer A",
					CustomerLevel: "VIP",
					Handler:       "Handler 1",
					SKU:           "Product 2",
					SalesRevenue:  2000.0,
				},
				{
					Year:          2024,
					Month:         2,
					Customer:      "Customer B",
					CustomerLevel: "Regular",
					Handler:       "Handler 2",
					SKU:           "Product 3",
					SalesRevenue:  1500.0,
				},
			},
		},
		{
			name: "Test Case 2: Multiple Handlers and Products",
			args: args{
				list: []*model.InvoiceDetail{
					{
						DocumentDate:  "2024/01/15",
						Customer:      "Customer A",
						CustomerLevel: "VIP",
						Handler:       "Handler 1",
						ProductName:   "Product 1",
						SalesRevenue:  1000.0,
					},
					{
						DocumentDate:  "2024/01/20",
						Customer:      "Customer A",
						CustomerLevel: "VIP",
						Handler:       "Handler 2",
						ProductName:   "Product 2",
						SalesRevenue:  2000.0,
					},
					{
						DocumentDate:  "2024/01/25",
						Customer:      "Customer A",
						CustomerLevel: "VIP",
						Handler:       "Handler 1",
						ProductName:   "Product 3",
						SalesRevenue:  3000.0,
					},
				},
			},
			want: []*InvoiceShopSKUMonth{
				{
					Year:          2024,
					Month:         1,
					Customer:      "Customer A",
					CustomerLevel: "VIP",
					Handler:       "Handler 1",
					SKU:           "Product 1",
					SalesRevenue:  1000.0,
				},
				{
					Year:          2024,
					Month:         1,
					Customer:      "Customer A",
					CustomerLevel: "VIP",
					Handler:       "Handler 2",
					SKU:           "Product 2",
					SalesRevenue:  2000.0,
				},
				{
					Year:          2024,
					Month:         1,
					Customer:      "Customer A",
					CustomerLevel: "VIP",
					Handler:       "Handler 1",
					SKU:           "Product 3",
					SalesRevenue:  3000.0,
				},
			},
		},
		{
			name: "Test Case 4: Different Years",
			args: args{
				list: []*model.InvoiceDetail{
					{
						DocumentDate:  "2023/12/25",
						Customer:      "Customer A",
						CustomerLevel: "VIP",
						Handler:       "Handler 1",
						ProductName:   "Product 1",
						SalesRevenue:  500.0,
					},
					{
						DocumentDate:  "2024/01/15",
						Customer:      "Customer A",
						CustomerLevel: "VIP",
						Handler:       "Handler 1",
						ProductName:   "Product 2",
						SalesRevenue:  1000.0,
					},
					{
						DocumentDate:  "2024/01/20",
						Customer:      "Customer A",
						CustomerLevel: "VIP",
						Handler:       "Handler 2",
						ProductName:   "Product 3",
						SalesRevenue:  2000.0,
					},
				},
			},
			want: []*InvoiceShopSKUMonth{
				{
					Year:          2023,
					Month:         12,
					Customer:      "Customer A",
					CustomerLevel: "VIP",
					Handler:       "Handler 1",
					SKU:           "Product 1",
					SalesRevenue:  500.0,
				},
				{
					Year:          2024,
					Month:         1,
					Customer:      "Customer A",
					CustomerLevel: "VIP",
					Handler:       "Handler 1",
					SKU:           "Product 2",
					SalesRevenue:  1000.0,
				},
				{
					Year:          2024,
					Month:         1,
					Customer:      "Customer A",
					CustomerLevel: "VIP",
					Handler:       "Handler 2",
					SKU:           "Product 3",
					SalesRevenue:  2000.0,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConverInvoiceShopSKUMonth(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				for _, person := range got {
					fmt.Println(person)
				}
				t.Errorf("converInvoiceShopSKUMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}
