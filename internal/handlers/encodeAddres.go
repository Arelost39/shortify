package handlers

import (
	"net/http"
	h "shortify/internal/helpers"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetEncodedAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		str := c.Query("conv_test")
		i, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
      			"error": "invalid address",
    		})
		} else {
			var encodedAddress string
			encodedAddress = h.Base62Encode(i)
			c.JSON(http.StatusOK, gin.H{
      			"encodedAddress": encodedAddress,
    		})
		}

	}
}