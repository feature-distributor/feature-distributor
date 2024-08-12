// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameToggleRule = "toggle_rule"

// ToggleRule mapped from table <toggle_rule>
type ToggleRule struct {
	RuleID          int64     `gorm:"column:rule_id;primaryKey;autoIncrement:true" json:"rule_id"`
	ToggleID        int64     `gorm:"column:toggle_id;not null" json:"toggle_id"`
	Title           string    `gorm:"column:title;not null" json:"title"`
	ListNum         int64     `gorm:"column:list_num;not null" json:"list_num"`
	RuleType        int8      `gorm:"column:rule_type;not null" json:"rule_type"`
	RuleName        string    `gorm:"column:rule_name;not null" json:"rule_name"`
	OperationType   int8      `gorm:"column:operation_type;not null" json:"operation_type"`
	RuleValue       string    `gorm:"column:rule_value;not null" json:"rule_value"`
	ReturnValue     int64     `gorm:"column:return_value;not null" json:"return_value"`
	ReturnValueType int8      `gorm:"column:return_value_type;not null" json:"return_value_type"`
	CreateTime      time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime      time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP" json:"update_time"`
}

// TableName ToggleRule's table name
func (*ToggleRule) TableName() string {
	return TableNameToggleRule
}