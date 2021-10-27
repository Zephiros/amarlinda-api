package controllers

import (
    "github.com/Zephiros/amarlinda/database"
    "github.com/Zephiros/amarlinda/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
    "errors"
)

func GetProduct(c *gin.Context) {
    id := c.Params.ByName("id")
    product := models.Product{}
  	query := database.DB.Select("products.*")
  	query = query.Group("products.id")
  	err := query.Where("products.id = ?", id).First(&product).Error
    if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
  		  c.JSON(http.StatusInternalServerError, err.Error())
        return
  	}

  	if errors.Is(err, gorm.ErrRecordNotFound) {
  		  c.JSON(http.StatusNotFound, err.Error())
        return
  	}

  	c.JSON(http.StatusOK, product)
}

func GetProducts(c *gin.Context) {
    products := []models.Product{}
    query := database.DB.Select("products.*").Group("products.id")
    if err := query.Find(&products).Error; err != nil {
        c.JSON(http.StatusInternalServerError, err.Error())
        return
    }

    c.JSON(http.StatusOK, products)
}

func UpdateProduct(c *gin.Context) {
    id := c.Params.ByName("id")
    product := models.Product{}
    query := database.DB.Select("products.*")
    query = query.Group("products.id")
    err := query.Where("products.id = ?", id).First(&product).Error
    if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
        c.JSON(http.StatusInternalServerError, err.Error())
        return
    }

    if errors.Is(err, gorm.ErrRecordNotFound) {
        c.JSON(http.StatusNotFound, err.Error())
        return
    }

    err = c.BindJSON(&product)
    if err != nil {
        c.JSON(http.StatusInternalServerError, err.Error())
        return
    }

    if err := database.DB.Save(&product).Error; err != nil {
        c.JSON(http.StatusInternalServerError, err.Error())
        return
    }

    c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
    id := c.Params.ByName("id")
    product := models.Product{}
    query := database.DB.Select("products.*")
    query = query.Group("products.id")
    err := query.Where("products.id = ?", id).First(&product).Error
    if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
        c.JSON(http.StatusInternalServerError, err.Error())
        return
    }

    if errors.Is(err, gorm.ErrRecordNotFound) {
        c.JSON(http.StatusNotFound, err.Error())
        return
    }

    if err := database.DB.Where("id = ? ", id).Delete(&product).Error; err != nil {
        c.JSON(http.StatusInternalServerError, err.Error())
        return
    }

    c.JSON(http.StatusOK, nil)
}

func CreateProduct(c *gin.Context) {
    product := models.Product{}
    if err := c.BindJSON(&product); err != nil {
    		c.JSON(http.StatusInternalServerError, err.Error())
    		return
  	}

  	if err := database.DB.Create(&product).Error; err != nil {
    		c.JSON(http.StatusInternalServerError, err.Error())
    		return
  	}

    c.JSON(http.StatusOK, product)
}
