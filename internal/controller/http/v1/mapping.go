package v1

import "komek/internal/usecase/user_managment"

func convertRegisterRequestToDomain(req RegisterRequest) user_managment.RegisterRequest {
	return user_managment.RegisterRequest{
		Username:     req.Username,
		Password:     req.Password,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Email:        req.Email,
		MobileNumber: req.MobileNumber,
	}
}
