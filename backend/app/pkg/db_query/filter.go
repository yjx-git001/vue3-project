package db_query

import "strings"

const (
	InputTypeText          = "text"
	InputTypeRadio         = "radio"
	InputTypeCheckBox      = "check_box"
	InputTypeDateRange     = "date_range"
	InputTypeDateTimeRange = "date_time_range"
)

type Condition interface {
	SetWhere(k string, v []interface{})
	SetOr(k string, v []interface{})
	SetOrder(k string)
	SetJoinOn(t, on string) Condition
	SetWhereExact(k string, v interface{})
}

type GormCondition struct {
	GormPublic
	Join []*GormJoin
}

type GormPublic struct {
	Where      map[string][]interface{}
	Order      []string
	Or         map[string][]interface{}
	WhereExact map[string]interface{}
}

type GormJoin struct {
	Type   string
	JoinOn string
	GormPublic
}

func (e *GormJoin) SetJoinOn(t, on string) Condition {
	return nil
}

func (e *GormPublic) SetWhere(k string, v []interface{}) {
	if e.Where == nil {
		e.Where = make(map[string][]interface{})
	}
	e.Where[k] = v
}

func (e *GormPublic) SetWhereExact(k string, v interface{}) {
	if e.WhereExact == nil {
		e.WhereExact = make(map[string]interface{})
	}
	e.WhereExact[k] = v
}

func (e *GormPublic) SetOr(k string, v []interface{}) {
	if e.Or == nil {
		e.Or = make(map[string][]interface{})
	}
	e.Or[k] = v
}

func (e *GormPublic) SetOrder(k string) {
	if e.Order == nil {
		e.Order = make([]string, 0)
	}
	e.Order = append(e.Order, k)
}

func (e *GormCondition) SetJoinOn(t, on string) Condition {
	if e.Join == nil {
		e.Join = make([]*GormJoin, 0)
	}
	join := &GormJoin{
		Type:       t,
		JoinOn:     on,
		GormPublic: GormPublic{},
	}
	e.Join = append(e.Join, join)
	return join
}

type resolveFilterTag struct {
	Type       string
	ColumnType string
	Column     string
	Table      string
	On         []string
	Join       string
}

// makeTag 解析search的tag标签
func makeTag(table, column, tag string) *resolveFilterTag {
	if strings.Contains(tag, "ignore") {
		return nil
	}
	r := &resolveFilterTag{
		Table:  table,
		Column: column,
	}
	tags := strings.Split(tag, ",")
	var ts []string
	for _, t := range tags {
		ts = strings.Split(t, ":")
		if len(ts) == 0 {
			continue
		}
		switch ts[0] {
		case "type":
			if len(ts) > 1 {
				r.Type = ts[1]
			}
		case "ct":
			if len(ts) > 1 {
				r.ColumnType = ts[1]
			}
		case "on":
			if len(ts) > 1 {
				r.On = ts[1:]
			}
		case "join":
			if len(ts) > 1 {
				r.Join = ts[1]
			}
		}
	}
	return r
}
