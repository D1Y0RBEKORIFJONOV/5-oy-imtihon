package entity

type (
	GetAllUserReq struct {
		Field string `json:"field" redis:"field"`
		Value string `json:"value" redis:"value"`
		Page  int64  `json:"page" redis:"page"`
		Limit int64  `json:"limit" redis:"limit"`
	}

	GetAllUserRes struct {
		Users []User `json:"users" redis:"users"`
		Count int32  `json:"count" redis:"count"`
	}

	GetUserReq struct {
		Field string `json:"field" redis:"field"`
		Value string `json:"value" redis:"value"`
	}

	UpdateUserReq struct {
		UserID    string `json:"-" redis:"id"`
		FirstName string `json:"first_name" redis:"first_name"`
		LastName  string `json:"last_name" redis:"last_name"`
	}
	UpdatePasswordReq struct {
		UserID      string `json:"-" redis:"id"`
		Password    string `json:"password" redis:"password"`
		NewPassword string `json:"new_password" redis:"new_password"`
	}

	UpdateEmailReq struct {
		UserID   string `json:"-" redis:"id"`
		NewEmail string `json:"new_email" redis:"new_email"`
	}

	LoginReq struct {
		Email    string `json:"email" redis:"email"`
		Password string `json:"password" redis:"password"`
	}

	Token struct {
		AccessToken  string `json:"access_token" redis:"access_token"`
		RefreshToken string `json:"refresh_token" redis:"refresh_token"`
	}

	LoginRes struct {
		Token Token `json:"token" redis:"token"`
	}

	StatusUser struct {
		Successfully bool `json:"successfully" redis:"successfully"`
	}

	Profile struct {
		FirstName string `json:"first_name" redis:"first_name"`
		CreatedAt string `json:"created_at" redis:"created_at"`
		UpdatedAt string `json:"updated_at" redis:"updated_at"`
		DeletedAt string `json:"deleted_at" redis:"deleted_at"`
		Address   string `json:"address" redis:"address"`
	}

	User struct {
		FirstName string  `json:"first_name" redis:"first_name"`
		LastName  string  `json:"last_name" redis:"last_name"`
		Email     string  `json:"email" redis:"email"`
		Password  string  `json:"password" redis:"password"`
		ID        string  `json:"id" redis:"id"`
		Profile   Profile `json:"profile" redis:"profile"`
	}
	UserRegisterReq struct {
		FirstName string `json:"first_name" redis:"first_name"`
		LastName  string `json:"last_name" redis:"last_name"`
		Email     string `json:"email" redis:"email"`
		Address   string `json:"address" redis:"address"`
		Password  string `json:"password" redis:"password"`
		SecretKey string `json:"secret-key" redis:"secret_key"`
	}

	CreateUserReq struct {
		FirstName       string `json:"first_name" redis:"first_name"`
		LastName        string `json:"last_name" redis:"last_name"`
		Email           string `json:"email" redis:"email"`
		Address         string `json:"address" redis:"address"`
		Password        string `json:"password" redis:"password"`
		ConfirmPassword string `json:"confirm_password" redis:"confirm_password"`
	}
	StatusMessage struct {
		Message string `json:"message" redis:"message"`
	}
	VerifyUserReq struct {
		Email      string `json:"email" redis:"email"`
		SecretCode string `json:"secret_code" redis:"secret_code"`
	}
	DeleteUserReq struct {
		UserID        string `json:"-" redis:"id"`
		IsHardDeleted bool   `json:"is_hard_deleted" redis:"is_hard_deleted"`
	}
)
