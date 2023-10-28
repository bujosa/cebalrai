package user

import "celbalrai/database"

var UserControllerInstance *UserController
var UserServiceInstance *UserService

func ModuleInit() {
    userRepo := UserRepository{
        db: database.DB,
    }

    UserServiceInstance = &UserService{
        userRepository: &userRepo,
    }

	UserControllerInstance = &UserController{
		userService: UserServiceInstance,
	}
}



