package database

import "gorm.io/gorm"

type SqlHandler interface {
	Exec(string, ...interface{}) *gorm.DB
	Find(interface{}, ...interface{}) *gorm.DB
	First(interface{}, ...interface{}) *gorm.DB
	Raw(string, ...interface{}) *gorm.DB
	Create(interface{}) *gorm.DB
	Save(interface{}) *gorm.DB
	Delete(interface{}) *gorm.DB
	Where(interface{}, ...interface{}) *gorm.DB
	Joins(query string, args ...interface{}) *gorm.DB
	Model(value interface{}) *gorm.DB
	Preload(query string, args ...interface{}) *gorm.DB
	Table(query string, args ...interface{}) *gorm.DB
}
