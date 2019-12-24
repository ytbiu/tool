package sqlut

import (
	"github.com/jinzhu/gorm"

	"fmt"
)

type SQLOperator interface {
	Get(result interface{}, setOpts ...SetOpt) error
	List(result interface{}, setOpts ...SetOpt) error
	Create(target interface{}) error
	Update(target interface{}) error
	Delete(setOpts ...SetOpt) error
}

// 持久层通用对象 提供常见curd 目前仅支持gorm 用户实现RoleOperator即可使用
type IGORMOperator interface {
	SQLOperator
	RoleOperator
}

type RoleOperator interface {
	GetDB() *gorm.DB
	GetModel() interface{}
}

type GormOperator struct {
	RoleOperator
}

// 用户注册RoleOperator的入口
func NewGORMOperator(role RoleOperator) IGORMOperator {
	return &GormOperator{RoleOperator: role}
}

func get(db *gorm.DB, model interface{}, opt *Options, result interface{}) error {
	if opt.IsORMQuery() {
		return db.Model(model).
			Where(opt.WhereQuery, opt.WhereArgs...).
			Find(result).
			Error
	}

	if opt.IsRawQuery() {
		return db.Raw(opt.RawSQL, opt.RawValue...).Scan(result).Error
	}

	return optsErr
}

func (b *GormOperator) Get(result interface{}, setOpts ...SetOpt) error {
	opt := Options{}
	for _, setOpt := range setOpts {
		setOpt(&opt)
	}

	return get(b.GetDB(), b.GetModel(), &opt, result)
}

func (b *GormOperator) List(result interface{}, setOpts ...SetOpt) error {
	return b.Get(result, setOpts...)
}

func (b *GormOperator) Create(target interface{}) error {
	return b.GetDB().Create(target).Error
}

func (b *GormOperator) Update(target interface{}) error {
	return b.GetDB().Save(target).Error
}

func (b *GormOperator) Delete(setOpts ...SetOpt) error {
	opt := Options{}
	for _, setOpt := range setOpts {
		setOpt(&opt)
	}

	return b.GetDB().Delete(
		b.GetModel(), fmt.Sprintf(opt.WhereQuery, opt.WhereArgs...),
	).Error
}
