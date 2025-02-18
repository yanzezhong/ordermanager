package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SortType string


const (
	SortTypeDesc SortType = "desc" // 降序
	SortTypeAsc  SortType = "asc"  // 升序
)

type (
	PageParam struct {
		Page            int         // 页码, 和Mark互斥，二选一
		Size            int         // 每页数量
		SortKey         string      // 排序字段，多个字段排序使用英文,分隔；如果是Mark翻页，SortKey即Mark对应的字段
		SortType        SortType    // 排序方向；如果是Mark翻页，SortType和Mark实现逻辑有相关性
		IsSortByChinese bool
		// 如果查询参数仅有 _id，且无排序需求，同时表中存在 _id&排序字段的联合索引；那推荐不要 sort，直接筛选 _id，如果加了 sort，会导致优先匹配上联合索引
		// 所以在明确无排序需求时，推荐设置 NoSort 为 true
		NoDefaultSort bool // 不需要排序，默认需要
	}

)


func (page PageParam) GeneratePageOption(opts ...*options.FindOptions) *options.FindOptions {
	var option = options.Find()
	// 使用传入的options做默认值
	if len(opts) > 0 {
		option = opts[0]
	}
	// sort
	switch page.SortType {
	case SortTypeDesc:
		option.SetSort(bson.M{page.SortKey: -1})
	case SortTypeAsc:
		option.SetSort(bson.M{page.SortKey: 1})
	}
	if page.IsSortByChinese {
		option.SetCollation(&options.Collation{Locale: "zh", CaseLevel: true})
	}

	if page.IsPage() {
		option.SetSkip(page.Skip())
		option.SetLimit(page.Limit())
	}

	return option
}

// IsPage 是否配置了基于page和size翻页
func (page PageParam) IsPage() bool {
	if page.Page > 0 && page.Size > 0 {
		return true
	}

	return false
}

func (page PageParam) Skip() int64 {
	if page.IsPage() {
		return int64((page.Page - 1) * page.Size)
	}
	return 0
}

func (page PageParam) Limit() int64 {
	if page.IsPage() {
		return int64(page.Size)
	}
	return 1000
}

// TimeRange 时间区间
type TimeRange struct {
	Start string `json:"start"`
	End   string `json:"end"`
}