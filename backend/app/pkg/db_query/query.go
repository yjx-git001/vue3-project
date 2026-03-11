package db_query

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"backend/app/pkg/logger"

	"gorm.io/gorm"
)

// form:成员名
// label:标签
// filter: 查询， 包含：type/dictType/inputType/on/join
//
//	type: exact/range/contain/in/gt/gte/lt/lte
//	     range: [gte]2022-07-02 00:00:00,[lte]2022-07-29 00:00:00, [gt]45,[lt]50
//	inputType: text/radio/check_box
//
// search: 搜索
// order: 排序， 包含：default
//
//	default: +/-
//
// table: 表名
// column:表中对应的列名
const (
	// Mysql 数据库标识
	Mysql = "mysql"
	// Postgres 数据库标识
	Postgres = "postgres"
)

var DefaultMembers = []string{"Page", "PageSize", "Search", "Order"}

type SearchType int

const (
	ST_LIKE  = "like"
	ST_ILIKE = "ilike"
)

// filter: type/dictType/inputType/
type FilterType int

const (
	FT_EXACT      = "exact"
	FT_CONTAIN    = "contain"
	FT_RANGE      = "range"
	FT_in         = "in"
	FT_GT         = "gt"
	FT_GTE        = "gte"
	FT_LT         = "lt"
	FT_LTE        = "lte"
	FT_START_WITH = "startswith"
	FT_END_WITH   = "endswith"
	FT_IS_NULL    = "isnull"
)

type OrderType int

const (
	OT_ASE  = "+"
	OT_DESC = "-"
)

var CompareFlags = map[string]string{"gt": ">", "gte": ">=", "lt": "<", "lte": "<="}

type Dict struct {
	Label string      `json:"label"`
	Value interface{} `json:"value"`
}

type Filter struct {
	Name       string `json:"name"`
	Label      string `json:"label"`
	ValType    string `json:"valType"`
	FilterType string `json:"filterType"`
	DictType   string `json:"dictType"`
	InputType  string `json:"inputType"`
	Opts       []Dict `json:"opts"`
}

type Search struct {
	Name  string `json:"name"`
	Label string `json:"label"`
}

type Order struct {
	Name    string `json:"name"`
	Default string `json:"default"`
}

type DataQuery struct {
	page     int
	pageSize int
	// orders     []Order
	// searches   []Search
	// filters    []Filter
	orderList  []string
	searchList []string
	condition  *GormCondition
	//filterList []string
	searchVal   string
	orderVal    string
	isParsePage bool
}

type TypeSymbol string

const (
	gtSymbol  TypeSymbol = ">"
	gteSymbol TypeSymbol = ">="
	ltSymbol  TypeSymbol = "<"
	lteSymbol TypeSymbol = "<="
)

func getQueryRule(query interface{}) ([]Filter, []Search, []Order) {
	//rule := make(map[string]interface{})
	t := reflect.TypeOf(query)

	var (
		fieldForm  string
		fieldLabel string
		fieldTag   string
		ok         bool
		filters    []Filter
		searches   []Search
		orders     []Order
	)
	for i := 0; i < t.NumField(); i++ {
		if checkDefaultMembers(t.Field(i).Name) {
			continue
		}
		fieldForm, ok = t.Field(i).Tag.Lookup("form")
		if !ok {
			continue
		}
		fieldLabel = t.Field(i).Tag.Get("label")
		fieldTag, ok = t.Field(i).Tag.Lookup("filter")
		if ok {
			f := Filter{
				Name:    fieldForm,
				Label:   fieldLabel,
				ValType: t.Field(i).Type.String(),
			}
			tagList := strings.Split(fieldTag, ",")
			for _, val := range tagList {
				subTag := strings.Split(val, ":")
				switch subTag[0] {
				case "type":
					f.FilterType = subTag[1]
				case "inputType":
					f.InputType = subTag[1]
				case "dictType":
					f.DictType = subTag[1]
				case "opts":
					// a := utils.Dict{}
					// f.Opts = subTag[1]
					opts := strings.Split(subTag[1], "#")
					for _, opt := range opts {
						sp := strings.Split(opt, "-")
						f.Opts = append(f.Opts, Dict{
							Label: sp[0],
							Value: sp[1],
						})
					}
				default:
					logger.Sugar.Warnf("subtag %v is not supported.", subTag[0])
				}
			}
			if f.InputType == "" {
				f.InputType = InputTypeText
			}
			filters = append(filters, f)
		}
		fieldTag, ok = t.Field(i).Tag.Lookup("order")
		if ok {
			var defStr string
			tagList := strings.Split(fieldTag, ",")
			for _, val := range tagList {
				subTag := strings.Split(val, ":")
				switch subTag[0] {
				case "default":
					defStr = subTag[1]
				}
			}
			orders = append(orders, Order{
				Name:    fieldForm,
				Default: defStr,
			})
		}
		fieldTag, ok = t.Field(i).Tag.Lookup("search")
		if ok {
			searches = append(searches, Search{
				Label: fieldLabel,
				Name:  fieldForm,
			})
		}
		//fmt.Printf("tag : %v, name: %v, type:%v, value:%v, tag sum: %v\n", fieldTag, fieldName, v.Field(i).EmploymentType().String(), fieldValue, v.EmploymentType().Field(i).Tag)
	}

	return filters, searches, orders
}

func GetQueryRule(query interface{}) map[string]interface{} {
	rule := make(map[string]interface{})
	filters, searches, orders := getQueryRule(query)
	rule["filter"] = filters
	rule["search"] = searches
	rule["order"] = orders
	return rule
}

func NewDataQuery(query interface{}) *DataQuery {
	q := &DataQuery{
		//filterList: make([]string, 0),
		searchList: make([]string, 0),
		orderList:  make([]string, 0),
		condition: &GormCondition{
			GormPublic: GormPublic{},
			Join:       make([]*GormJoin, 0),
		},
	}
	q.resolveQuery("mysql", query, q.condition)
	return q
}

func checkDefaultMembers(name string) bool {
	switch name {
	case "Page", "PageSize", "Order", "Search":
		return true
	default:
		return false
	}
}

func (q *DataQuery) resolveQuery(driver string, query interface{}, condition Condition) {
	t := reflect.TypeOf(query).Elem()
	v := reflect.ValueOf(query).Elem()
	if !q.isParsePage {
		q.page = int(v.FieldByName("Page").Int())
		q.pageSize = v.FieldByName("PageSize").Interface().(int)
		q.searchVal = v.FieldByName("Search").String()
		q.orderVal = v.FieldByName("Order").String()
		q.isParsePage = true
	}
	var (
		name     string
		form     string
		fieldTag string
		table    string
		column   string
		ok       bool
		tag      reflect.StructTag
	)

	for i := 0; i < t.NumField(); i++ {
		name = t.Field(i).Name
		//value = v.Field(i).Interface()
		if checkDefaultMembers(name) {
			continue
		}
		tag = t.Field(i).Tag
		form = tag.Get("form")
		//fieldLabel = t.Field(i).Tag.Get("label")
		table = tag.Get("table")
		column = tag.Get("column")
		if column == "" {
			column = form
		}
		fieldTag, ok = tag.Lookup("search")
		if ok {
			q.parserSearch(table, column)
		}
		fieldTag, ok = tag.Lookup("order")
		if ok {
			q.parseOrder(table, column, fieldTag)
		}
		if v.Field(i).IsZero() {
			continue
		}
		fieldTag, ok = tag.Lookup("filter")
		if ok {
			q.parseFilter(table, column, fieldTag, v.Field(i), condition)
		}

		//fmt.Printf("tag : %v, name: %v, type:%v, value:%v, tag sum: %v\n", fieldTag, fieldName, v.Field(i).EmploymentType().String(), fieldValue, v.EmploymentType().Field(i).Tag)
	}
}

func (q *DataQuery) parserSearch(table string, column string) {
	var prefix string
	if len(table) > 0 {
		prefix = table + "."
	}

	if q.searchVal != "" {
		q.searchList = append(q.searchList, fmt.Sprintf("%s%s like '%%%s%%' ", prefix, column, q.searchVal))
	}
}

func (q *DataQuery) parseOrder(table string, column string, tag string) {
	var prefix string
	if len(table) > 0 {
		prefix = table + "."
	}

	pos := strings.Index(q.orderVal, column)
	if pos >= 0 {
		if pos == 0 {
			q.orderList = append(q.orderList, fmt.Sprintf("%v%v", prefix, column))
		} else {
			order := q.orderVal[pos-1 : pos]
			switch order {
			case "+", ",":
				q.orderList = append(q.orderList, fmt.Sprintf("%v%v", prefix, column))
			case "-":
				q.orderList = append(q.orderList, fmt.Sprintf("%v%v desc", prefix, column))
			default:
				logger.Sugar.Warnf("order input is valid. order input: %v", q.orderVal)
			}
		}
	} else {
		tagList := strings.Split(tag, ",")
		for _, val := range tagList {
			subTag := strings.Split(val, ":")
			switch subTag[0] {
			case "default":
				if subTag[1] == "+" {
					q.orderList = append(q.orderList, fmt.Sprintf("%v%v", prefix, column))
				} else if subTag[1] == "-" {
					q.orderList = append(q.orderList, fmt.Sprintf("%v%v desc", prefix, column))
				}
			}
		}
	}
}

func (q *DataQuery) buildWhereStr(t *resolveFilterTag, symbol TypeSymbol) string {
	var (
		whereStr string
	)
	switch t.ColumnType {
	case "date":
		if t.Table == "" {
			whereStr = fmt.Sprintf("DATE_FORMAT(`%s`,'%%Y-%%m-%%d') %s DATE_FORMAT(?, '%%Y-%%m-%%d')", t.Column, symbol)
		} else {
			whereStr = fmt.Sprintf("DATE_FORMAT(`%s`.`%s`,'%%Y-%%m-%%d') %s DATE_FORMAT(?, '%%Y-%%m-%%d')", t.Table, t.Column, symbol)
		}
	case "datetime":
		if t.Table == "" {
			whereStr = fmt.Sprintf("DATE_FORMAT(`%s`,'%%Y-%%m-%%d %%H:%%i:%%s') %s ?", t.Column, symbol)
		} else {
			whereStr = fmt.Sprintf("DATE_FORMAT(`%s`.`%s`,'%%Y-%%m-%%d %%H:%%i:%%s') %s ?", t.Table, t.Column, symbol)
		}
	default:
		if t.Table == "" {
			whereStr = fmt.Sprintf("`%s` %s ?", t.Column, symbol)
		} else {
			whereStr = fmt.Sprintf("`%s`.`%s` %s ?", t.Table, t.Column, symbol)
		}
	}
	return whereStr
}

func (q *DataQuery) parseFilter(table string, column string, tag string, qValue reflect.Value, condition Condition) {

	driver := "mysql"
	t := makeTag(table, column, tag)
	if t == nil {
		return
	}
	var whereStr string
	//解析
	switch t.Type {
	case "left":
		//左关联
		join := condition.SetJoinOn(t.Type, fmt.Sprintf(
			"left join `%s` on `%s`.`%s` = `%s`.`%s`",
			t.Join,
			t.Join,
			t.On[0],
			t.Table,
			t.On[1],
		))
		q.resolveQuery("mysql", qValue.Interface(), join)
	case "range":
		rangeList := strings.Split(qValue.Interface().(string), ",")
		if len(rangeList) > 2 {
			logger.Sugar.Errorf("%s filter error. input value = %v", t.Column, qValue.Interface())
			return
		}

		for _, val := range rangeList {
			list := strings.Split(val, "]")
			if len(list) != 2 {
				logger.Sugar.Errorf("range of %s filter is invalid.", t.Column)
				return
			}
			whereStr = fmt.Sprintf("`%s` %s ?", t.Column, CompareFlags[list[0][1:]])
			if t.Table != "" {
				whereStr = fmt.Sprintf("`%s`.`%s` %s ?", t.Table, t.Column, CompareFlags[list[0][1:]])
			}
			condition.SetWhere(whereStr, []interface{}{list[1]})
		}
	case "exact", "iexact":
		whereStr = fmt.Sprintf("`%s` = ?", t.Column)
		if t.Table != "" {
			whereStr = fmt.Sprintf("`%s`.`%s` = ?", t.Table, t.Column)
		}
		condition.SetWhere(whereStr, []interface{}{qValue.Interface()})
	case "contain", "icontain":
		//fixme mysql不支持ilike
		qValueStr := fmt.Sprintf("%v", qValue)
		if driver == Postgres && t.Type == "icontains" {
			whereStr = fmt.Sprintf("`%s` ilike ?", t.Column)
			if t.Table != "" {
				whereStr = fmt.Sprintf("`%s`.`%s` ilike ?", t.Table, t.Column)
			}
			condition.SetWhere(whereStr, []interface{}{"%" + qValueStr + "%"})
		} else {
			if t.Table == "" {
				whereStr = fmt.Sprintf("`%s` like ?", t.Column)
			} else {
				whereStr = fmt.Sprintf("`%s`.`%s` like ?", t.Table, t.Column)
			}

			condition.SetWhere(whereStr, []interface{}{"%" + qValueStr + "%"})
		}
	case "gt":
		//if t.Table == "" {
		//	whereStr = fmt.Sprintf("`%s` > ?", t.Column)
		//} else {
		//	whereStr = fmt.Sprintf("`%s`.`%s` > ?", t.Table, t.Column)
		//}
		whereStr := q.buildWhereStr(t, gtSymbol)
		condition.SetWhere(whereStr, []interface{}{qValue.Interface()})
	case "gte":
		//if t.Table == "" {
		//	whereStr = fmt.Sprintf("`%s` >= ?", t.Column)
		//} else {
		//	whereStr = fmt.Sprintf("`%s`.`%s` >= ?", t.Table, t.Column)
		//}
		whereStr := q.buildWhereStr(t, gteSymbol)
		condition.SetWhere(whereStr, []interface{}{qValue.Interface()})
	case "lt":
		//if t.Table == "" {
		//	whereStr = fmt.Sprintf("`%s` < ?", t.Column)
		//} else {
		//	whereStr = fmt.Sprintf("`%s`.`%s` < ?", t.Table, t.Column)
		//}
		whereStr := q.buildWhereStr(t, ltSymbol)
		condition.SetWhere(whereStr, []interface{}{qValue.Interface()})
	case "lte":
		//if t.Table == "" {
		//	whereStr = fmt.Sprintf("`%s` <= ?", t.Column)
		//} else {
		//	whereStr = fmt.Sprintf("`%s`.`%s` <= ?", t.Table, t.Column)
		//}
		whereStr := q.buildWhereStr(t, lteSymbol)
		condition.SetWhere(whereStr, []interface{}{qValue.Interface()})

	case "startswith", "istartswith":
		if driver == Postgres && t.Type == "istartswith" {
			if t.Table == "" {
				whereStr = fmt.Sprintf("`%s` ilike ?", t.Column)
			} else {
				whereStr = fmt.Sprintf("`%s`.`%s` ilike ?", t.Table, t.Column)
			}
			condition.SetWhere(whereStr, []interface{}{qValue.String() + "%"})
		} else {
			if t.Table == "" {
				whereStr = fmt.Sprintf("`%s` like ?", t.Column)
			} else {
				whereStr = fmt.Sprintf("`%s`.`%s` like ?", t.Table, t.Column)
			}
			condition.SetWhere(whereStr, []interface{}{qValue.String() + "%"})
		}
	case "endswith", "iendswith":
		if driver == Postgres && t.Type == "iendswith" {
			if t.Table == "" {
				whereStr = fmt.Sprintf("`%s` ilike ?", t.Column)
			} else {
				whereStr = fmt.Sprintf("`%s`.`%s` ilike ?", t.Table, t.Column)
			}
			condition.SetWhere(whereStr, []interface{}{"%" + qValue.String()})
		} else {
			if t.Table == "" {
				whereStr = fmt.Sprintf("`%s` like ?", t.Column)
			} else {
				whereStr = fmt.Sprintf("`%s`.`%s` like ?", t.Table, t.Column)
			}
			condition.SetWhere(whereStr, []interface{}{"%" + qValue.String()})
		}
	case "in":
		if t.Table == "" {
			whereStr = fmt.Sprintf("`%s` in (?)", t.Column)
		} else {
			whereStr = fmt.Sprintf("`%s`.`%s` in (?)", t.Table, t.Column)
		}
		condition.SetWhere(whereStr, []interface{}{qValue.Interface()})
		var intList []interface{}
		list := strings.Split(qValue.Interface().(string), ",")
		for _, i := range list {
			ii, err := strconv.ParseUint(i, 10, 64)
			if err != nil {
				intList = append(intList, i)
			} else {
				intList = append(intList, ii)
			}
		}
		condition.SetWhere(whereStr, intList)
	case "isnull":
		if !(qValue.IsZero() && qValue.IsNil()) {
			if t.Table == "" {
				fmt.Sprintf("`%s` isnull", t.Column)
			} else {
				fmt.Sprintf("`%s`.`%s` isnull", t.Table, t.Column)
			}
			condition.SetWhere(whereStr, make([]interface{}, 0))
		}
	}
}

//	func (q *DataQuery) Filter() []string {
//		logger.Sugar.Infof("query filter: %+v", q.filterList)
//		return q.filterList
//	}
//
//	func (q *DataQuery) Search() []string {
//		logger.Sugar.Infof("query search: %+v", q.searchList)
//		return q.searchList
//	}
//
//	func (q *DataQuery) Order() []string {
//		logger.Sugar.Infof("query order: %+v", q.orderList)
//		return q.orderList
//	}
func (q *DataQuery) Pagination() (int, int) {
	logger.Sugar.Infof("query pagination: page = %v, page size = %v", q.page, q.pageSize)
	return q.page, q.pageSize
}

func (q *DataQuery) GenerateQueryDB(db *gorm.DB, hasCount bool) (*gorm.DB, int64, error) {
	// filter
	condition := q.condition
	for _, join := range condition.Join {
		if join == nil {
			continue
		}
		db = db.Joins(join.JoinOn)
		for k, v := range join.Where {
			db = db.Where(k, v)
		}

		for k, v := range join.WhereExact {
			db = db.Where(k, v)
		}

		// for k, v := range join.Or {
		// 	db = db.Or(k, v...)
		// }
		// for _, o := range join.Order {
		// 	db = db.Order(o)
		// }
	}
	for k, v := range condition.Where {
		db = db.Where(k, v)
	}
	for k, v := range condition.WhereExact {
		db = db.Where(k, v)
	}
	//search
	var searchStr string
	for key, s := range q.searchList {
		if key == 0 {
			//db = db.Where(s)
			searchStr = s
		} else {
			//db = db.Or(s)
			searchStr += " or " + s
		}
	}
	if searchStr != "" {
		db = db.Where(searchStr)
	}
	var count int64
	if hasCount {
		if err := db.Count(&count).Error; err != nil {
			return nil, 0, err
		}
	}

	for _, o := range q.orderList {
		db = db.Order(o)
	}
	//if q.page < 1 {
	//	q.page = 1
	//}
	//if q.pageSize < 1 {
	//	q.pageSize = 10
	//}
	if q.page != 0 && q.pageSize != 0 {
		db = db.Limit(q.pageSize).Offset((q.page - 1) * q.pageSize)
	}

	return db, count, nil
}
