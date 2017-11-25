package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type CabinetDetail struct {
	Id        int       `orm:"column(id);auto"`
	CabinetId int       `orm:"column(cabinet_id);null" description:"柜子的id"`
	Door      int       `orm:"column(door);null" description:"门号"`
	OpenState int       `orm:"column(open_state);null" description:"开关状态，1:关，2:开"`
	Using     int       `orm:"column(using);null" description:"占用状态，1:空闲，2:占用"`
	UserID    string    `orm:"column(userID);size(255);null" description:"存物ID"`
	StoreTime time.Time `orm:"column(store_time);type(datetime);null" description:"存物时间"`
	UseState  int       `orm:"column(use_state);null" description:"启用状态，1:启用，2:停用"`
}

type Total struct {
	Doors int `json:"-" orm:"column(doors)"`
	OnUse int `json:"-" orm:"column(onUse)"`
	Close int `json:"-" orm:"column(close)"`
}

func (t *CabinetDetail) TableName() string {
	return "cabinet_detail"
}

func init() {
	orm.RegisterModel(new(CabinetDetail))
}

// 根据柜子id，获取该柜子的门的详情
func GetDetailsByCabinetId(cabinetId int) (details []CabinetDetail, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(new(CabinetDetail)).Filter("CabinetId", cabinetId).All(&details)
	return
}

// 根据柜子id，获取该柜子的总门数
func GetTotalDoors(cabinetId int) (total int) {
	tot := Total{}
	sql := "SELECT COUNT(id) AS doors FROM cabinet_detail WHERE cabinet_id=?"
	orm.NewOrm().Raw(sql, cabinetId).QueryRow(&tot)
	return tot.Doors
}

// 根据柜子id，获取该柜子的总使用中的门数
func GetTotalOnUse(cabinetId int) (onUse int) {
	tot := Total{}
	sql := "SELECT COUNT(`using`) AS onUse FROM cabinet_detail WHERE cabinet_id=? AND `using`=2"
	orm.NewOrm().Raw(sql, cabinetId).QueryRow(&tot)
	return tot.OnUse
}

// 根据柜子id，获取该柜子的总关闭状态门数
func GetTotalClose(cabinetId int) (close int) {
	tot := Total{}
	sql := "SELECT COUNT(open_state) AS `close` FROM cabinet_detail WHERE cabinet_id=? AND open_state=1"
	orm.NewOrm().Raw(sql, cabinetId).QueryRow(&tot)
	return tot.Close
}

// AddCabinetDetail insert a new CabinetDetail into database and returns
// last inserted Id on success.
func AddCabinetDetail(m *CabinetDetail) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCabinetDetailById retrieves CabinetDetail by Id. Returns error if
// Id doesn't exist
func GetCabinetDetailById(id int) (v *CabinetDetail, err error) {
	o := orm.NewOrm()
	v = &CabinetDetail{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCabinetDetail retrieves all CabinetDetail matches certain condition. Returns empty list if
// no records exist
func GetAllCabinetDetail(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CabinetDetail))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []CabinetDetail
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateCabinetDetail updates CabinetDetail by Id and returns error if
// the record to be updated doesn't exist
func UpdateCabinetDetailById(m *CabinetDetail) (err error) {
	o := orm.NewOrm()
	v := CabinetDetail{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCabinetDetail deletes CabinetDetail by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCabinetDetail(id int) (err error) {
	o := orm.NewOrm()
	v := CabinetDetail{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CabinetDetail{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
