package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repository interface {
	Select(model interface{}, fields string, where ...interface{}) *gorm.DB
	Find(model interface{}, where ...interface{}) *gorm.DB
	Create(value interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Updates(value interface{}) *gorm.DB
	Delete(value interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
	Close() error
}

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository() *GormRepository {
	constr := "root:@tcp(127.0.0.1:3306)/godemo?charset=utf8mb4&parseTime=True&loc=Local"
	dbcon, _ := gorm.Open(mysql.Open(constr), &gorm.Config{})
	return &GormRepository{db: dbcon}
}

func (r *GormRepository) Select(model interface{}, fields string, where ...interface{}) *gorm.DB {
	return r.db.Select(fields).Where(model, where...)
}

func (r *GormRepository) Find(model interface{}, where ...interface{}) *gorm.DB {
	return r.db.Find(model, where...)
}

func (r *GormRepository) Create(value interface{}) *gorm.DB {
	return r.db.Create(value)
}

func (r *GormRepository) Save(value interface{}) *gorm.DB {
	return r.db.Save(value)
}

func (r *GormRepository) Updates(value interface{}) *gorm.DB {
	return r.db.Updates(value)
}

func (r *GormRepository) Delete(value interface{}) *gorm.DB {
	return r.db.Delete(value)
}

func (r *GormRepository) Where(query interface{}, args ...interface{}) *gorm.DB {
	return r.db.Where(query, args)
}

func (r *GormRepository) Close() error {
	sqlDB, err := r.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
