package handler

import (
	"fmt"

	Product "github.com/vincentJunior1/cleancode-vincent/internal/product"
)

type ProductHandler struct {
	ProductUsecase *Product.ProductUsecaseInteface
}

func ProductNewHandler(productUsecase Product.ProductUsecaseInteface) {
	product := &ProductHandler{
		ProductUsecase: &productUsecase,
	}
	product.GetAllProduct()
}

func (s *ProductHandler) GetAllProduct() {
	fmt.Println("init get all product")
}
