package product

type ProductUseCase struct {
	repository ProductRepositoryInterface
}

func NewProductUseCase(repository ProductRepositoryInterface) *ProductUseCase {
	return &ProductUseCase{repository: repository}
}

func (uc *ProductUseCase) GetProductByID(id int) *Product {
	return uc.repository.GetProductByID(id)
}
