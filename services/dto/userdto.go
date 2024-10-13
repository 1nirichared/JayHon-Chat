package dto

type UserDTO struct {
	Username string `form:"username" binding:"required,max=16,min=2"`
	Password string `form:"password" binding:"required,max=32,min=6"`
	Email    string `form:"email" binding:"required,max=64,min=6"`
	AvatarId string `form:"avatar_id" binding:"required,max=255"`
}
