package handlers

import (
	"net/http"
	db "shortify/internal/db"

	"github.com/gin-gonic/gin"
)

func GetLastID(db *db.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		lastID, err := db.GetLastID()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
      			"error": "empty database",
    		})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"lastID": lastID,
			})
		}
	}
}