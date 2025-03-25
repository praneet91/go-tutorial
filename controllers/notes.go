package controllers

import (
	"go-tutorial/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NotesController struct {
	NotesService *services.NotesService
}

func (nc *NotesController) Init(notesService *services.NotesService) *NotesController {
	nc.NotesService = notesService

	return nc
}

func (nc *NotesController) InitRoutes(router *gin.Engine) {
	router.GET("/notes", nc.GetNotes())
	router.POST("/notes", nc.CreateNotes())
	router.PUT("/notes/", nc.UpdateNotes())
	router.DELETE("/notes/:id", nc.DeleteNotes())
	router.GET("/notes/:id", nc.GetNote())
}

func (nc *NotesController) GetNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "id should be an integer"})
			return
		}

		note, error := nc.NotesService.GetNoteService(id)
		if error != nil {
			c.JSON(500, gin.H{"error": error.Error()})
			return
		}

		c.JSON(200, gin.H{
			"message": note,
		})
	}
}

func (nc *NotesController) GetNotes() gin.HandlerFunc {
	return func(c *gin.Context) {

		status := c.Query("status")
		var isCompleted *bool

		if status != "" {
			parsedStatus, err := strconv.ParseBool(status)
			if err != nil {
				c.JSON(400, gin.H{"error": "status should be a boolean value"})
				return
			}
			isCompleted = &parsedStatus
		}

		notes, error := nc.NotesService.GetNotesService(isCompleted)
		if error != nil {
			c.JSON(500, gin.H{"error": error.Error()})
			return
		}

		c.JSON(200, gin.H{
			"message": notes,
		})
	}
}

func (nc *NotesController) CreateNotes() gin.HandlerFunc {

	type CreateNoteRequest struct {
		Description string `json:"description" binding:"required"`
		IsCompleted bool   `json:"is_completed"`
	}
	return func(c *gin.Context) {
		var createNoteRequest CreateNoteRequest
		if err := c.ShouldBindJSON(&createNoteRequest); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		note, err := nc.NotesService.CreateNoteService(createNoteRequest.Description, createNoteRequest.IsCompleted)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"note": note,
		})
	}
}

func (nc *NotesController) UpdateNotes() gin.HandlerFunc {

	type CreateNoteRequest struct {
		Description string `json:"description" binding:"required"`
		IsCompleted bool   `json:"is_completed"`
		Id          int    `json:"id" binding:"required"`
	}
	return func(c *gin.Context) {
		var createNoteRequest CreateNoteRequest
		if err := c.ShouldBindJSON(&createNoteRequest); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		note, err := nc.NotesService.UpdateNoteService(createNoteRequest.Description, createNoteRequest.IsCompleted, createNoteRequest.Id)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"note": note,
		})
	}
}

func (nc *NotesController) DeleteNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "id should be an integer"})
			return
		}

		err = nc.NotesService.DeleteNoteService(id)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"message": "Note deleted successfully",
		})
	}
}
