package handler

import (
	"fmt"
	"github.com/Dimadetected/clientsCRUD"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getAllClients(c *gin.Context) {
	clients, err := h.services.Client.GetAll()
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "clients not found")
		return
	}

	c.JSON(http.StatusOK, map[string][]clientsCRUD.Client{
		"data": clients,
	})
}
func (h *Handler) getByIdClient(c *gin.Context) {
	clientId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid clientId in params")
		return
	}
	client, err := h.services.Client.GetById(clientId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "client not found")
		return
	}

	c.JSON(http.StatusOK, client)

}
func (h *Handler) createClient(c *gin.Context) {
	var client clientsCRUD.Client

	if err := c.BindJSON(&client); err != nil {
		fmt.Println(client)
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Client.Create(client)
	if err != nil {
		fmt.Println(err.Error())
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
func (h *Handler) deleteClient(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid client_id param")
		return
	}

	if err = h.services.Client.Delete(id); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "client not found")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"deleted": "ok",
	})

}
func (h *Handler) updateClient(c *gin.Context) {
	var input clientsCRUD.UpdateClient
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err = c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input data")
		return
	}

	if err = h.services.Client.Update(id, input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"update": "ok",
	})
}
