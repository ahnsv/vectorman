package v1

import (
	"net/http"
	"strconv"

	"github.com/ahnsv/vectorman/pkg/app"
	"github.com/ahnsv/vectorman/pkg/e"
	"github.com/ahnsv/vectorman/pkg/entities"
	"github.com/gin-gonic/gin"
)

var onCallPersonnel []entities.OnCallPersonnel

// @Summary Get all on-call personnel
// @Produce json
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Router /api/v1/oncall/personnel [get]
func CreateOncallPersonnel(c *gin.Context) {
	appG := app.Gin{C: c}
	var newPersonnel entities.OnCallPersonnel
	err := c.BindJSON(&newPersonnel)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	onCallPersonnel = append(onCallPersonnel, newPersonnel)
	appG.Response(http.StatusOK, 200, gin.H{"status": "OK"})
}

// @Summary Get all on-call personnel
// @Produce json
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Router /api/v1/oncall/personnel [get]
func GetOncallPersonnel(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, 200, onCallPersonnel)
}

// @Summary Get on-call personnel by ID
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Router /api/v1/oncall/personnel/{id} [get]
func GetOncallPersonnelByID(c *gin.Context) {
	appG := app.Gin{C: c}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	for i := range onCallPersonnel {
		if onCallPersonnel[i].ID == id {
			appG.Response(http.StatusOK, 200, onCallPersonnel[i])
			return
		}
	}
	appG.Response(http.StatusNotFound, e.ERROR_NOT_EXIST_PERSONNEL, gin.H{"error": "Personnel not found"})
}

// @Summary Delete on-call personnel by ID
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Router /api/v1/oncall/personnel/{id} [put]
func UpdateOncallPersonnel(c *gin.Context) {
	appG := app.Gin{C: c}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, gin.H{"error": "Invalid ID"})
		return
	}
	var updatedPersonnel entities.OnCallPersonnel
	err = c.BindJSON(&updatedPersonnel)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, gin.H{"error": "Invalid JSON"})
		return
	}
	for i := range onCallPersonnel {
		if onCallPersonnel[i].ID == id {
			onCallPersonnel[i] = updatedPersonnel
			appG.Response(http.StatusOK, e.SUCCESS, gin.H{"status": "OK"})
			return
		}
	}
	appG.Response(http.StatusNotFound, e.ERROR_NOT_EXIST_PERSONNEL, gin.H{"error": "Personnel not found"})
}

// @Summary Delete on-call personnel by ID
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Router /api/v1/oncall/personnel/{id} [delete]
func DeleteOncallPersonnel(c *gin.Context) {
	appG := app.Gin{C: c}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, gin.H{"error": "Invalid ID"})
		return
	}
	for i := range onCallPersonnel {
		if onCallPersonnel[i].ID == id {
			onCallPersonnel = append(onCallPersonnel[:i], onCallPersonnel[i+1:]...)
			appG.Response(http.StatusOK, e.SUCCESS, gin.H{"status": "OK"})
			return
		}
	}
	appG.Response(http.StatusNotFound, e.ERROR_NOT_EXIST_PERSONNEL, gin.H{"error": "Personnel not found"})
}
