package controller

import (
	"net/http"

	"github.com/daudfauzy98/Pengenalan-Konsep-MVC/app/model"
	"github.com/gin-gonic/gin"
)

// Di Golang, jika nama fungsi/variabel berawalan huruf besar (bisa diakses public),
// disarankan memberikan komentar diatas fungsi/variabel tersebut dan kata pertama
// yang tulis yaitu nama funsi/variabel itu sendiri

// AddAntrianHandler is a function to add queue
func AddAntrianHandler(c *gin.Context) {
	flag, err := model.AddAntrian()

	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success!",
		})
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error":  err,
		})
	}
}

// GetAntrianHandler is a function to get all queue
func GetAntrianHandler(c *gin.Context) {
	flag, resp, err := model.GetAntrian()

	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success!",
			"data":   resp,
		})
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error":  err,
		})
	}
}

// UpdateAntrianHandler is a function to update a queue
func UpdateAntrianHandler(c *gin.Context) {
	idAntrian := c.Param("idAntrian")
	flag, err := model.UpdateAntrian(idAntrian)

	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success!",
		})
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error":  err,
		})
	}
}

// DeleteAntrianHandler is a function to delete a queue
func DeleteAntrianHandler(c *gin.Context) {
	idAntrian := c.Param("idAntrian")
	flag, err := model.DeleteAntrian(idAntrian)

	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success!",
		})
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error":  err,
		})
	}
}

// PageAntrianHandler menampilkan antrian yang sedang berlangsung di index.html
func PageAntrianHandler(c *gin.Context) {
	flag, result, err := model.GetAntrian()
	var currentAntrian map[string]interface{} // Tipe data JSON

	for _, item := range result {
		if item != nil {
			currentAntrian = item
			break
		}
	}

	if flag && len(result) > 0 {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"antrian": currentAntrian["id"],
		})
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error":  err,
		})
	}
}
