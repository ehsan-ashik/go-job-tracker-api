package common

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Paginate(ctx *fiber.Ctx) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := ctx.QueryInt("page", 1)
		limit := ctx.QueryInt("limit", 50)

		if page <= 0 {
			page = 1
		}
		if limit <= 0 {
			limit = 50
		} else if limit > 100 {
			limit = 100
		}

		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}

func Sort(ctx *fiber.Ctx) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		sort_str := ctx.Query("sort", "")
		if sort_str == "" {
			return db
		}

		var arr_str []string
		err := json.Unmarshal([]byte(sort_str), &arr_str)

		if err != nil || len(arr_str) != 2 {
			ctx.Response().Header.Add("X-Invalid-Sort", "True")
			return db
		}

		return db.Order(fmt.Sprintf("%v %v", arr_str[0], arr_str[1]))
	}
}

func FilterByVal(key string, value interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%v = ?", key), value)
	}
}

func FilterByList(key string, values interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%v IN (?)", key), values)
	}
}

func Filter(ctx *fiber.Ctx) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		filter_str := ctx.Query("filter", "")

		if filter_str == "" {
			return db
		}
		var data map[string]interface{}

		err := json.Unmarshal([]byte(filter_str), &data)
		if err != nil {
			ctx.Response().Header.Add("X-Invalid-Filter", "True")
		}

		for key, val := range data {
			if reflect.TypeOf(val).Kind() == reflect.Slice {
				db = db.Scopes(FilterByList(key, val))
			} else {
				db = db.Scopes(FilterByVal(key, val))
			}
		}
		return db
	}
}
