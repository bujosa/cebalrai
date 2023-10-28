package user

import "celbalrai/database"

var UserControllerInstance *UserController

// ModuleInit initializes the user module
func ModuleInit() {
	userRepo := UserRepository{
		db: database.DB,
	}

	userService := &UserService{
		userRepository: &userRepo,
	}

	UserControllerInstance = &UserController{
		userService: userService,
	}
}
