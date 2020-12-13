package person

import "template-service/models"

func (pis *PersonInfoService) GetUserOrders(user *models.User) []models.Order {
	return pis.userRepo.GetOrdersByLogin(user.Login)
}
