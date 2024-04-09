package controller

import (
	"Back_end/common"
	"Back_end/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Card struct {
}

func (ca *Card) AddCard(c *gin.Context) {
	card := service.SimplifiedCard{}
	if err := c.ShouldBind(&card); err != nil {
		c.Error(common.ErrNew(err, common.SysErr))
		return
	}
	err := srv.Card.AddCard(card)
	if err != nil {
		c.Error(common.ErrNew(err, common.SysErr))
		return
	}
	c.JSON(http.StatusOK, ResponseNew(c, "Successfully added"))
}

func (ca *Card) GetCard(c *gin.Context) {
	var CheckInfo struct {
		EncryptedId string `json:"id" form:"id"`
	}
	if err := c.ShouldBind(&CheckInfo); err != nil {
		c.Error(common.ErrNew(err, common.SysErr))
		return
	}
	err, cards := srv.Card.GetCard(CheckInfo.EncryptedId, &srv.DES)
	if err != nil {
		c.Error(common.ErrNew(err, common.SysErr))
		return
	}
	c.JSON(http.StatusOK, ResponseNew(c, cards))
}

func (ca *Card) GetEncryptedID(c *gin.Context) {
	var CheckInfo struct {
		Id string `json:"id" form:"id"`
	}
	if err := c.ShouldBind(&CheckInfo); err != nil {
		c.Error(common.ErrNew(err, common.SysErr))
		return
	}
	resp := srv.Card.GetEncryptedID(CheckInfo.Id, &srv.DES)
	c.JSON(http.StatusOK, ResponseNew(c, resp))
}

func (ca *Card) GetDecryptedID(c *gin.Context) {
	var CheckInfo struct {
		Id string `json:"id" form:"id"`
	}
	if err := c.ShouldBind(&CheckInfo); err != nil {
		c.Error(common.ErrNew(err, common.SysErr))
		return
	}
	resp := srv.Card.GetDecryptedID(CheckInfo.Id, &srv.DES)
	c.JSON(http.StatusOK, ResponseNew(c, resp))
}

func (ca *Card) HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, ResponseNew(c, "Hello World! This is HappyBug yelling!"))
}
