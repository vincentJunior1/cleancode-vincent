package router

import (
	productHandler "github.com/vincentJunior1/cleancode-vincent/internal/product/handler"
	productUC "github.com/vincentJunior1/cleancode-vincent/internal/product/usecase"
)

func Init() {
	productusecase := productUC.ProductUsecaseHandler("init")
	productHandler.ProductNewHandler(productusecase)
}
