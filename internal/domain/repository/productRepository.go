package repository

import (
	"github.com/go-pg/pg/v10"
	"github.com/labstack/gommon/log"
	"pantrycontrol-backend/internal/domain/models/entities"
)

type ProductRepository interface {
	FindAll() ([]entities.Product, error)
	FindById(int) (entities.Product, error)
	Save(entities.Product) error
	Update(int, entities.Product) error
	Delete(int) error
}

type ProductRepositoryImpl struct {
	db *pg.DB
}

func CreateProductRepository(db *pg.DB) ProductRepository {
	return &ProductRepositoryImpl{db: db}
}

func (r *ProductRepositoryImpl) FindAll() ([]entities.Product, error) {
	var products []entities.Product
	err := r.db.Model(&products).Select()
	if err != nil {
		log.Error("Error to find products.", err)
		return nil, err
	}
	return products, nil
}

func (r *ProductRepositoryImpl) FindById(id int) (entities.Product, error) {
	product := new(entities.Product)
	err := r.db.Model(product).Where("id = ?", id).Select()
	if err != nil {
		log.Error("Error to find a product.", err)
		return entities.Product{}, err
	}
	return *product, nil
}

func (r *ProductRepositoryImpl) Save(product entities.Product) error {
	_, err := r.db.Model(&product).Insert()
	if err != nil {
		log.Error("Error to save a product", err)
		return err
	}
	return err
}

func (r *ProductRepositoryImpl) Update(id int, product entities.Product) error {
	oldProduct := &entities.Product{Id: id}
	err := r.db.Model(oldProduct).WherePK().Select()
	if err != nil {
		log.Error("Product not found.", err)
		return err
	}

	product.Id = id
	_, err = r.db.Model(&product).WherePK().Update()

	return nil
}

func (r *ProductRepositoryImpl) Delete(id int) error {
	oldProduct := &entities.Product{Id: id}
	_, err := r.db.Model(oldProduct).WherePK().Delete()
	if err != nil {
		log.Error("Error to delete a product.", err)
		return err
	}
	return nil
}