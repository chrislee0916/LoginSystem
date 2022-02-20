package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	CID      uint `gorm:"primarykey"`
	CName    string
	Products []Product
}

type Product struct {
	PID        uint `gorm:"primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	UserID     uint
	CategoryID uint   `json:"category_id" form:"category_id"`
	PName      string `json:"p_name" form:"p_name"`
}

func (p *Product) SaveProduct() (*Product, error) {
	err := DB.Create(&p).Error
	if err != nil {
		return nil, err
	}
	return p, err
}

func GetProductByID(pid string) (p Product, err error) {
	err = DB.Model(&Product{}).Where("p_id = ?", pid).First(&p).Error
	if err != nil {
		return Product{}, err
	}
	return
}

func CreateProduct(userID uint, category_id uint, p_name string) (err error) {
	p := &Product{
		UserID:     userID,
		CategoryID: category_id,
		PName:      p_name,
	}
	_, err = p.SaveProduct()
	if err != nil {
		return err
	}
	return

}

func DeleteProductByID(pid string) (err error) {
	err = DB.Delete(&Product{}, pid).Error
	if err != nil {
		return err
	}
	return
}

func UpdateProductByID(pid string, category_id uint, p_name string) (err error) {
	err = DB.Model(&Product{}).Where("p_id = ?", pid).Updates(Product{CategoryID: category_id, PName: p_name}).Error
	if err != nil {
		return err
	}
	return
}
