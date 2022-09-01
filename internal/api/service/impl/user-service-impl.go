package impl

import (
	"database/sql"
	"github.com/conorsheppard/user-api-go/internal/api/entity"
	i "github.com/conorsheppard/user-api-go/internal/api/service/interface"
	"github.com/conorsheppard/user-api-go/internal/api/util"
	v "github.com/conorsheppard/user-api-go/internal/api/validation"
	db "github.com/conorsheppard/user-api-go/internal/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"log"
	"net/http"
	"strconv"
)

type getUserIDFromURI struct {
	ID string `uri:"id" binding:"required,min=1"`
}

type UserServiceImpl struct {
	store db.Store
}

func NewUserService(store db.Store) i.UserService {
	return &UserServiceImpl{
		store: store,
	}
}

func (us UserServiceImpl) Create(context *gin.Context) {
	var req entity.CreateUserRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	if !validateCreateUserRequest(req, context) {
		return
	}

	userID, err := uuid.NewUUID()

	if err != nil {
		log.Fatalf("error while creating user ID: %s\n", err.Error())
	}

	hashedPassword, err := util.HashPassword(req.Password)

	if err != nil {
		log.Fatalf("failed to hash password: %s\n", err.Error())
	}

	arg := db.CreateUserParams{
		ID:             userID.String(),
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Nickname:       req.Nickname,
		HashedPassword: hashedPassword,
		Email:          req.Email,
		Country:        req.Country,
	}

	user, err := us.store.CreateUser(context, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				context.JSON(http.StatusForbidden, v.ErrorResponse(err))
				return
			}
		}
		context.JSON(http.StatusInternalServerError, v.ErrorResponse(err))
		return
	}

	if err != nil {
		log.Fatalf("error while creating new user: %s\n", err.Error())
	}

	rsp := entity.NewUserResponse(user)
	context.JSON(http.StatusOK, rsp)
}

func validateCreateUserRequest(req entity.CreateUserRequest, context *gin.Context) bool {
	if err := v.ValidateFirstname(req.FirstName); err != nil {
		context.JSON(http.StatusBadRequest, v.ErrorResponse(err))
		return false
	}

	if err := v.ValidatePassword(req.Password); err != nil {
		context.JSON(http.StatusBadRequest, v.ErrorResponse(err))
		return false
	}

	if err := v.ValidateLastName(req.LastName); err != nil {
		context.JSON(http.StatusBadRequest, v.ErrorResponse(err))
		return false
	}

	if err := v.ValidateEmail(req.Email); err != nil {
		context.JSON(http.StatusBadRequest, v.ErrorResponse(err))
		return false
	}

	return true
}

func (us UserServiceImpl) GetAll(context *gin.Context) {
	l, p, country := context.DefaultQuery("limit", "10"), context.DefaultQuery("page", "1"), context.DefaultQuery("country", "%")

	limit, err := strconv.Atoi(l)
	if limit < 1 || err != nil {
		context.JSON(http.StatusBadRequest, v.ErrorResponse(err))
	}

	page, err := strconv.Atoi(p)

	if page < 1 || err != nil {
		context.JSON(http.StatusBadRequest, v.ErrorResponse(err))
	}

	offset := (page * 10) - 10

	arg := db.GetAllUsersParams{
		Country: country,
		Limit:   int32(limit),
		Offset:  int32(offset),
	}

	users, err := us.store.GetAllUsers(context, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			context.JSON(http.StatusNotFound, v.ErrorResponse(err))
			return
		}
		context.JSON(http.StatusInternalServerError, v.ErrorResponse(err))
		return
	}

	rsp := entity.GetAllUsersResponse(users)
	context.JSON(http.StatusOK, rsp)
}

func (us UserServiceImpl) Update(context *gin.Context) {
	var req getUserIDFromURI
	if err := context.ShouldBindUri(&req); err != nil {
		context.JSON(http.StatusBadRequest, v.ErrorResponse(err))
		return
	}
	var reqBody entity.UpdateUserRequest
	if err := context.ShouldBindJSON(&reqBody); err != nil {
		context.JSON(http.StatusBadRequest, nil)
		return
	}

	arg := db.UpdateUserParams{
		FirstName: reqBody.FirstName,
		LastName:  reqBody.LastName,
		Nickname:  reqBody.Nickname,
		Email:     reqBody.Email,
		Country:   reqBody.Country,
		ID:        req.ID,
	}

	if err := us.store.UpdateUser(context, arg); err != nil {
		//context.JSON(http.StatusInternalServerError, v.ErrorResponse(err))
		log.Fatalf("error while updating user: %s\n", err.Error())
	}

	context.JSON(http.StatusNoContent, nil)
}

func (us UserServiceImpl) Delete(context *gin.Context) {
	context.JSON(http.StatusNoContent, nil)
	var req getUserIDFromURI
	if err := context.ShouldBindUri(&req); err != nil {
		context.JSON(http.StatusBadRequest, v.ErrorResponse(err))
		return
	}
	if err := us.store.DeleteUser(context, req.ID); err != nil {
		log.Fatalf("error while updating user: %s\n", err.Error())
	}
	context.JSON(http.StatusNoContent, nil)
}
