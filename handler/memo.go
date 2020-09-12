package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-crew/Bolierplate-CRUD-Gingonic/models"
	"github.com/golang-crew/Bolierplate-CRUD-Gingonic/requests"
	"github.com/jinzhu/gorm"
)

// GetMemoList ...
// @Title Get memos
// @Description Get all memos
// @Tags Memo
// @Param	memoID    query	   string	 true	 "memoID"
// @Success 200 {object} []models.Memos
// @Failure 404
// @Failure 500
// @Router /memos [Get]
func GetMemoList(c *gin.Context) {
	response, err := models.GetMemoList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Fail to get memos"})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetMemo ...
// @Title Get memo
// @Description Get memo by memoID
// @Tags Memo
// @Param	memoID   query	string	 true	 "memoID"
// @Success 200 {object} models.Memos
// @Failure 404
// @Failure 500
// @Router /memos/:memoID [Get]
func GetMemo(c *gin.Context) {
	memoID := c.Param("memoID")
	response, err := models.GetMemo(memoID)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"message": "No record found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Fail to get memo"})
		return
	}

	c.JSON(http.StatusOK, response)
}

// CreateMemo ...
// @Title Create memo
// @Description Create memo
// @Tags Memo
// @Param	memoID    query	   string	 false	 "memoID"
// @Param	body      body     requests.CreateMemo  true "create memo"
// @Success 200 {object} responses.Message
// @Failure 404
// @Failure 500
// @Router /memos/:memoID [Post]
func CreateMemo(c *gin.Context) {
	var request requests.CreateMemo
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Fail to bind request"})
		return
	}

	err = models.CreateMemo(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Fail to create stage"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Success create a memo"})
}

// DeleteMemo ...
// @Title Delete memo
// @Description Delete memo
// @Tags Memo
// @Param	memoID    query	   string	 false	 "memoID"
// @Success 200 {object} responses.Message
// @Failure 404
// @Failure 500
// @Router /memos/:memoID [Delete]
func DeleteMemo(c *gin.Context) {
	memoID := c.Param("memoID")

	err := models.DeleteMemo(memoID)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"message": "No record found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Fail to delete memo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Memo deleted successfully"})
}

// UpdateMemo ...
// @Title Update memo
// @Description Update memo
// @Tags Memo
// @Param	memoID    query	   string	 false	 "memoID"
// @Param	body      body     requests.UpdateMemo  true "update memo"
// @Success 200 {object} responses.Message
// @Failure 404
// @Failure 500
// @Router /memos/:memoID [Delete]
func UpdateMemo(c *gin.Context) {
	memoID := c.Param("memoID")

	var memo requests.UpdateMemo
	err := c.BindJSON(&memo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Fail to bind request"})
		return
	}

	err = models.UpdateMemo(memo, memoID)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"message": "Invaild ID - record not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Fail to save memo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully saved memo infomation"})
}
