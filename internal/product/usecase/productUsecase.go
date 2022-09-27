package usecase

import "fmt"

type ProductUsecase struct {
}

func ProductUsecaseHandler(init string) *ProductUsecase {
	return &ProductUsecase{}
}

func (p *ProductUsecase) HandlingGetAll() {
	fmt.Println("handling get all")
}
