package main

import (
	"fmt"
)

/*
Фасад — это структурный паттерн проектирования, который предоставляет простой интерфейс к сложной системе классов,
библиотеке или фреймворку.

Применяется когда вам нужно представить простой или урезанный интерфейс к сложной подсистеме.

Плюсы и минусы:
+ Изолирует клиентов от компонентов сложной подсистемы.
- Фасад рискует стать божественным объектом, привязанным ко всем классам программы.
*/

type orderFacade struct {
	userService         UserService
	productService      ProductService
	paymentService      PaymentService
	notificationService NotificationService
}

func newOrderFacade() *orderFacade {
	return &orderFacade{
		userService:         UserService{},
		productService:      ProductService{},
		paymentService:      PaymentService{},
		notificationService: NotificationService{},
	}
}

func (o *orderFacade) placeOrder(userId string, productId string) {
	fmt.Println("[Facade] Начинаем добавление заказа")

	userValid := o.userService.isUserValid(userId)
	productAvailable := o.productService.productAvailable(productId)

	if userValid && productAvailable {
		o.productService.assignProductToUser(productId, userId)
		o.paymentService.makePayment(userId, productId)
		o.notificationService.notifyUser(productId)
	}
}

type UserService struct{}

func (u *UserService) isUserValid(userId string) bool {
	fmt.Println("[UserService] валидируем юзера: ", userId)
	// Complex logic for checking validity
	return true
}

type ProductService struct{}

func (p *ProductService) productAvailable(productId string) bool {
	fmt.Println("[ProductService] проверяем доступность товара: ", productId)
	// Complex logic for checking availability
	return true
}

func (p *ProductService) assignProductToUser(productId string, userId string) {
	fmt.Printf("[ProductService] добавляем товар %s к юзеру %s\n", productId, userId)
	// complex logic for product assignment
}

type PaymentService struct{}

func (p *PaymentService) makePayment(userId string, productId string) {
	fmt.Printf("[PaymentService] снимаем деньги с юзера %s за товар %s\n", userId, productId)
	// complex logic for making payment
}

type NotificationService struct{}

func (n *NotificationService) notifyUser(productId string) {
	fmt.Printf("[NotificationService] уведомляем пользователя, что товар успешно добавлен в заказы %s\n", productId)
	// complex notification logic
}

func main() {
	orderModule := newOrderFacade()

	userId := "test-user-id"
	productId := "test-product-id"

	orderModule.placeOrder(userId, productId)
}
