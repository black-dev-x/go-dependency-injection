package product

type ProductUseCase struct {
	repository *ProductRepository
}

func NewProductUseCase(repository *ProductRepository) *ProductUseCase {
	return &ProductUseCase{repository: repository}
}

func (uc *ProductUseCase) GetProductByID(id int) *Product {
	return uc.repository.GetProductByID(id)
}
