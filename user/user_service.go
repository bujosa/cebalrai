package user

import (
	"context"
)

type UserService struct {
    userRepository *UserRepository
}

func (us *UserService) GetUserByID(ctx context.Context, userID int) (*User, error) {
    // In this case, we can add some logic before calling the repository
    return us.userRepository.GetUserByID(ctx, userID)
}