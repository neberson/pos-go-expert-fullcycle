package product

type ProductUseCase struct {
	repository ProductRepository
}

func NewProductUseCase(repository ProductRepository) *ProductUseCase {
	return &ProductUseCase{repository: repository}
}

func (uc *ProductUseCase) GetProductName(id int) (Product, error) {
	return uc.repository.GetByID(id)
}
