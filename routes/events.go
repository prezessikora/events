package routes

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prezessikora/events/models"
)

func getEventById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not convert request param"})
		return
	}

	event, err := models.GetEventById(id)

	if err != nil {
		fmt.Println(err)
		switch {
		case errors.Is(err, sql.ErrNoRows):
			ctx.JSON(http.StatusNotFound, gin.H{"message": "event nout found"})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not execute query"})
		}
		return
	}
	ctx.JSON(http.StatusOK, event)
}

func getEvents(ctx *gin.Context) {
	events, err := models.GetAll()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not execute query"})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func createEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindBodyWithJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		fmt.Println(err)
		return

	}

	event.UserID = int(ctx.GetInt("userId"))
	err = event.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not create event"})
		fmt.Println(err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
}

func deleteEvent(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not convert request param"})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotFound, gin.H{"message": "could not find the event"})
		return
	}
	err = event.Delete()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete the event"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "event deleted"})

}

func updateEvent(ctx *gin.Context) {
	var updatedEvent models.Event
	err := ctx.ShouldBindBodyWithJSON(&updatedEvent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		fmt.Println(err)
		return
	}

	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not convert request param"})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not find the event"})
		return
	}

	event.Name = updatedEvent.Name
	event.Description = updatedEvent.Description
	err = event.Update()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not update the event"})
		fmt.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "event updated", "event": event})
}
