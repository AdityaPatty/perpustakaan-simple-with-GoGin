package handlerbuku

import (
	"fmt"
	"net/http"
	"perpustakaan/modelbuku"
	"perpustakaan/servicebuku"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bukuHandler struct {
	Service servicebuku.BukuService
}

func NewBukuHandler(svc servicebuku.BukuService) bukuHandler {
	return bukuHandler{Service: svc}
}

func (h *bukuHandler) PerpustakaanCreateHandler(c *gin.Context) {
	var buku modelbuku.Perpustakaan

	if err := c.ShouldBind(&buku); err != nil {
		errorMassages := []string{}

		for _, e := range err.(validator.ValidationErrors) {
			errorMassage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMassages = append(errorMassages, errorMassage)

		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMassages,
		})

		return

	}

	perpustakaan, err := h.Service.Create(buku)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": fmt.Sprintf("internal server error: (%v/n)", err),
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		perpustakaan,
	)
}

func (h *bukuHandler) PerpustakaanGetHandler(c *gin.Context) {
	id := c.Param("id")
	bukuid, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err),
			},
		)
		return
	}

	perpustakaan, err := h.Service.Get(bukuid)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": fmt.Sprintf("internal server error: %v\n", err),
			},
		)
		return
	}

	c.JSON(http.StatusOK, perpustakaan)
}

func (h *bukuHandler) PerpustakaanGetsHandler(c *gin.Context) {
	perpustakaans, err := h.Service.Gets()
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": fmt.Sprintf("internal server error: %v\n", err),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		perpustakaans,
	)
}

func (h *bukuHandler) PerpustakaanUpdateHandler(c *gin.Context) {
	id := c.Param("id")
	bukuid, err := strconv.Atoi(id)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err),
			},
		)

		return
	}
	var buku modelbuku.Perpustakaan

	if err := c.ShouldBind(&buku); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err),
			},
		)

		return
	}

	user, err := h.Service.Update(bukuid, buku)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": fmt.Sprintf("internal server error: %v\n", err),
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		user,
	)
}

func (h *bukuHandler) PerpustakaanDeleteHandler(c *gin.Context) {
	id := c.Param("id")
	bukuid, err := strconv.Atoi(id)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err),
			},
		)

		return
	}

	perpustakaan, err := h.Service.Delete(bukuid)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": fmt.Sprintf("internal server error: %v\n", err),
			},
		)

		return
	}

	fmt.Sprintln("======== DATA BERHASIL DI HAPUS =========")
	c.JSON(
		http.StatusOK,

		perpustakaan,
	)
}
