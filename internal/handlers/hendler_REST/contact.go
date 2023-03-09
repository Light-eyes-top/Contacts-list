package hendler_REST

import (
	someRestApi "SomeRestApi"
	"SomeRestApi/internal/handlers/responce"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CreateContact(c *gin.Context) {
	var contact someRestApi.Contacts
	if err := c.BindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := h.services.CreateContact(contact)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, id)
}

func (h *Handler) GetContact(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		responce.NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	contact, err := h.services.GetContact(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if contact.Id == 0 && contact.Fio == "" && contact.Email == "" && contact.Phone == "" {
		responce.NewErrorResponse(c, http.StatusBadRequest, "this contact does not exist")
		return
	}

	c.JSON(http.StatusOK, contact)
}

func (h *Handler) UpdateContact(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		responce.NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var data someRestApi.Contacts
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contact, err := h.services.UpdateContact(id, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contact)
}

func (h *Handler) DeleteContact(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		responce.NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	contact, err := h.services.DeleteContact(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, responce.StatusResponse{
		Status:    "ok",
		ContactId: contact,
	})
}
