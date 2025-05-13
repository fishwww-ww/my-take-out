package common

import (
	"gorm.io/gorm"
	"my-take-out/common/enum"
)

type Result struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type PageResult struct {
	Total   int64       `json:"total"`   //总记录数
	Records interface{} `json:"records"` //当前页数据集合
}

// Pageverify 分页查询 防止非法参数
func PageVerify(page *int, pageSize *int) {
	if *page < 1 {
		*page = 1
	}
	switch {
	case *pageSize > 100:
		*pageSize = enum.MaxPageSize
	case *pageSize < 0:
		*pageSize = 0
	}
}

func (p *PageResult) Paginate(page *int, pageSize *int) func(db *gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		PageVerify(page, pageSize)
		d.Offset((*page - 1) * *pageSize).Limit(*pageSize)
		return d
	}
}
