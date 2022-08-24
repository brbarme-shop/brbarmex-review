package route

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/brbarme-shop/brbarmex-review/config"
	"github.com/brbarme-shop/brbarmex-review/postgresql"
	"github.com/brbarme-shop/brbarmex-review/review"
	"github.com/gin-gonic/gin"
)

var (
	cfg        = config.NewConfiguration()
	db         = postgresql.NewSqlDB(cfg)
	repositpry = postgresql.NewPutReviewRepository(db)
)

func postReview(c *gin.Context) {

	b, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	}

	defer c.Request.Body.Close()

	var reviewInput *review.ReviewInput
	err = json.Unmarshal(b, &reviewInput)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	}

	err = review.PutReview(c.Request.Context(), reviewInput, repositpry)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}
