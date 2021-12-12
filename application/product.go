package application

import (
	"errors"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type ProductServiceInterface interface {
	Get(id string) (ProductInterface, error)
	Create(name string, price float64) (ProductInterface, error)
	Enable(product ProductInterface) (ProductInterface, error)
	Disable(product ProductInterface) (ProductInterface, error)
}

type ProductReader interface {
	Get(id string) (ProductInterface, error)
}

type ProductWriter interface {
	Save(product ProductInterface) (ProductInterface, error)
}

type ProductPersistenceInterface interface {
	ProductReader
	ProductWriter
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetId() string
	GetName() string
	GetStatus() string
	GetPrice() float32	
}

const (
	DISABLED = "DISABLED"
	ENABLED = "ENABLED"
)

type Product struct {
	ID string `valid:"uuidv4"`
	Name string `valid:"required"`
	Price float32 `valid:"float,optional"`
	Status string `valid:"required"`
}

func NewProduct() *Product {
	product := Product{
		ID: uuid.NewV4().String(),
		Status: DISABLED,
	}
	
	return &product
}

func (product *Product) IsValid() (bool, error)  {
	if product.Status == "" {
		product.Status = DISABLED
	}

	if product.Status != ENABLED && product.Status != DISABLED {
		return false, errors.New("the status must be enabled or disabled")
	}

	if product.Price < 0 {
		return false, errors.New("the price must be greather or equal zero")
	}

	_, err := govalidator.ValidateStruct(product)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (product *Product) Enable() error {
	if product.Price <= 0 {
		return errors.New("the price must be greater than zero to enable the product")
	}

	product.Status = ENABLED
	return nil
}

func (product *Product) Disable() error  {
	if product.Price != 0 {
		return errors.New("the price must be zero in order to have the product disabled")
	}

	product.Status = DISABLED
	return nil
}

func (product *Product) GetId() string  {
	return product.ID
}

func (product *Product) GetName() string  {
	return product.Name
}

func (product *Product) GetStatus() string  {
	return product.Status
}

func (product *Product) GetPrice() float32 {
	return product.Price
}