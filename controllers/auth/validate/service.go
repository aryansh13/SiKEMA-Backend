package validateAuth

import (
	model "attendance-is/models"
	util "attendance-is/utils"
)

type Service interface {
	ValidateAuth(token string) (*model.User, string)
}

type service struct {
	repository Repository
}

func NewValidateAuthService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ValidateAuth(token string) (*model.User, string) {
	if util.ValidateToken(token) != nil {
		return nil, "AUTH_UNAUTHENTICATED_401"
	}

	claim, err := util.ExtractTokenClaim(token)
	if err != nil {
		return nil, "AUTH_UNEXPECTED_500 : " + err.Error()
	}

	res, err2 := s.repository.GetUserByID(claim.Uid)
	// fmt.Println(claim.Uid)
	if err2 != "" {
		return nil, err2
	}
	return res, ""
}
