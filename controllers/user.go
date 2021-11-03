package controllers

import (
	"errors"
	"golang_gorm_practice/database"
	"golang_gorm_practice/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepo struct {
	Db *gorm.DB
}

func New() *UserRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.User{}) // migrate schema

	return &UserRepo{Db: db}
}

// create a user
func (repo *UserRepo) CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)                        // bind request
	err := models.CreateUser(repo.Db, &user) // get function CreateUser from models
	if err != nil {                          // error thrown
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err},
		)

		return
	}

	c.JSON(http.StatusOK, user) // return with data created user
}

// get users
func (repo *UserRepo) GetUsers(c *gin.Context) {
	var user []models.User                 // declare user as slice models.User
	err := models.GetUsers(repo.Db, &user) // get function GetUsers from models

	if err != nil { // error thrown
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err},
		)

		return
	}
	c.JSON(http.StatusOK, user) // return with data users
}

// get user by id
func (repo *UserRepo) GetUserById(c *gin.Context) {
	id, _ := c.Params.Get("id")                   // set param
	var user models.User                          // declare user as models.User
	err := models.GetUserById(repo.Db, &user, id) // get function GetUserById from models

	if err != nil { // error thrown
		if errors.Is(err, gorm.ErrRecordNotFound) { // if error because no id
			c.AbortWithStatus(http.StatusNotFound)

			return
		}

		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err},
		)

		return
	}

	c.JSON(http.StatusOK, user) // return with data user
}

// update user
func (repo *UserRepo) UpdateUser(c *gin.Context) {
	var user models.User
	id, _ := c.Params.Get("id")                   // set param
	err := models.GetUserById(repo.Db, &user, id) // get function GetUserById from models

	if err != nil { // error thrown
		if errors.Is(err, gorm.ErrRecordNotFound) { // if error because no id
			c.AbortWithStatus(http.StatusNotFound)

			return
		}

		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err},
		)

		return
	}

	c.BindJSON(&user)                       // bind param
	err = models.UpdateUser(repo.Db, &user) // get function UpdateUser from models
	if err != nil {                         // error thrown
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err},
		)

		return
	}

	c.JSON(http.StatusOK, user) // return with data user
}

// delete user
func (repo *UserRepo) DeleteUser(c *gin.Context) {
	var user models.User
	id, _ := c.Params.Get("id")                  // set param
	err := models.DeleteUser(repo.Db, &user, id) // get function DeleteUser from models

	if err != nil { // error thrown
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)

			return
		}

		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err},
		)

		return
	}

	c.JSON( // return with message
		http.StatusOK,
		gin.H{"message": "user deleted successfully"},
	)
}
