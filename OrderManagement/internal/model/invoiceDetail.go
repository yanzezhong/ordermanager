package model

// 来肯数据 底表
type InvoiceDetail struct {
	DocumentDate       string  `json:"document_date"`       // 单据日期，记录单据的创建日期
	DocumentNumber     string  `json:"document_number"`     // 单据编号，唯一标识单据的编号
	DocumentType       string  `json:"document_type"`       // 单据类型，例如销售发票、采购订单等
	Customer           string  `json:"customer"`            // 客户名称，记录与单据相关的客户
	CustomerLevel      string  `json:"customer_level"`      // 客户级别，例如普通客户、VIP客户等
	SourceOrder        string  `json:"source_order"`        // 来源订单，记录单据的来源订单编号
	Handler            string  `json:"handler"`             // 经手人，记录处理单据的人员
	ProductName        string  `json:"product_name"`        // 商品名称，记录销售的商品名称
	Specification      string  `json:"specification"`       // 规格，记录商品的规格信息
	SalesQuantity      float64 `json:"sales_quantity"`      // 销售数量，记录销售的商品数量
	SalesSpecification string  `json:"sales_specification"` // 销售规格，记录商品的销售规格
	UnitPrice          float64 `json:"unit_price"`          // 单价，记录商品的单价
	Amount             float64 `json:"amount"`              // 金额，记录商品的总金额
	SalesRevenue       float64 `json:"sales_revenue"`       // 销售收入，记录销售的总收入
	Weight             float64 `json:"weight"`              // 重量（kg），记录商品的重量
	LineItemAttribute  string  `json:"line_item_attribute"` // 商品行属性，记录商品行的特殊属性
	DetailRemark       string  `json:"detail_remark"`       // 明细备注，记录单据明细的备注信息
	DocumentRemark     string  `json:"document_remark"`     // 单据备注，记录单据整体的备注信息
	Quantity           float64 `json:"quantity"`            // 数量，记录商品的总数量
	TotalAmount        float64 `json:"total_amount"`        // 金额，记录商品的总金额
}
