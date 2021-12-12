package application

type ProductService struct {
	Persistence ProductPersistenceInterface
}

func (service *ProductService) Get(id string) (ProductInterface, error) {
	product, err := service.Persistence.Get(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (service *ProductService) Create(name string, price float32) (ProductInterface, error) {
	product := NewProduct()
	product.Name = name
	product.Price = price

	_, err := product.IsValid()
	return service.saveProduct(product, err)
}

func (service *ProductService) Enable(product ProductInterface) (ProductInterface, error) {
	err := product.Enable()
	return service.saveProduct(product, err)
}

func (service *ProductService) Disable(product ProductInterface) (ProductInterface, error) {
	err := product.Disable()
	return service.saveProduct(product, err)
}

func (service *ProductService) saveProduct(product ProductInterface, err error) (ProductInterface, error) {
	if err != nil {
		return &Product{}, err
	}

	result, err := service.Persistence.Save(product)
	if err != nil {
		return &Product{}, err
	}

	return result, nil
}