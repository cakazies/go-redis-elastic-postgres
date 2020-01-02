package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/cakazies/go-redis-elastic-postgres/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type (
	// RoomCtr struct for setting this file
	RoomCtr struct{}
)

// GetRoom controller function for get detail room
func (RoomCtr) GetRoom(c *gin.Context) {
	var resp Response
	var result models.Room

	idRoom := c.Param("room_id")
	if idRoom == "" {
		resp.Status = "not found"
		resp.Message = "params is not found"
		c.JSON(404, resp)
		return
	}

	keyRedis := dBName + "_" + "ROOM" + "_" + idRoom
	data, ok := checkingRedis(keyRedis)
	if ok {
		err := json.Unmarshal(data, &result)
		if err != nil {
			LogError("Unmarshal checkredis", err, "room/GetRoom")
		}
	} else {
		roomModel := models.RoomModel{}
		room, err := roomModel.GetRoom(idRoom)
		result = room
		if err != nil {
			if err.Error() == "record not found" {
				resp.Status = "fail"
				resp.Message = "Room id is not found"
				c.JSON(404, resp)
				return
			} else if err != nil {
				LogError("DB Error", err, "room/GetRoom")
				resp.Status = "error"
				resp.Message = "Unable to communicate with database"
				c.JSON(500, resp)
				return
			}
		}
		data, err := json.Marshal(&room)
		if err != nil {
			LogError("Marshal", err, "room/GetRoom")
		}
		err = setRedis(keyRedis, string(data))
		if err != nil {
			LogError("Set Redis", err, "room/GetRoom")
		}
	}

	// return success
	resp.Code = 200
	resp.Status = "Success"
	c.JSON(resp.Code, gin.H{"response": resp, "data": result})
	return
}

// InsertRoom function for create data in table room
func (RoomCtr) InsertRoom(c *gin.Context) {
	var resp Response
	var data models.Room

	c.BindJSON(&data)
	v := validator.New()
	err := v.Struct(data)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			if e != nil {
				resp.Code = 404
				resp.Message = "Data `" + e.Field() + "` doesn't exist"
				c.JSON(404, resp)
				return
			}
		}
	}

	roomModel := models.RoomModel{}
	err = roomModel.InsertRoom(data)
	if err != nil {
		resp.Status = "error"
		resp.Message = "something went wrong"
		c.JSON(404, resp)
		return
	}

	// return success
	resp.Code = 200
	resp.Status = "Success"
	c.JSON(resp.Code, gin.H{"response": resp})
	return
}

// GetRooms function for get all data
func (RoomCtr) GetRooms(c *gin.Context) {
	var resp Response

	roomModel := models.RoomModel{}
	result, err := roomModel.GetRooms()
	if err != nil {
		LogError("DB Error", err, "room/GetRoom")
		resp.Status = "error"
		resp.Message = "Unable to communicate with database"
		c.JSON(500, resp)
		return
	}

	// return success
	resp.Code = 200
	resp.Status = "Success"
	c.JSON(resp.Code, gin.H{"response": resp, "data": result})
	return
}

// UpdateRoom function for update data
func (RoomCtr) UpdateRoom(c *gin.Context) {
	var resp Response
	var data models.Room

	idRoom := c.Param("room_id")
	if idRoom == "" {
		resp.Status = "not found"
		resp.Message = "params is not found"
		c.JSON(404, resp)
		return
	}

	c.BindJSON(&data)
	v := validator.New()
	err := v.Struct(data)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			if e != nil {
				resp.Code = 404
				resp.Message = "Data `" + e.Field() + "` doesn't exist"
				c.JSON(404, resp)
				return
			}
		}
	}

	roomModel := models.RoomModel{}
	err = roomModel.UpdateRoom(idRoom, data)
	if err != nil {
		LogError("Update Room", err, "room/UpdateRoom")
		resp.Status = "error"
		resp.Message = "something went wrong"
		c.JSON(404, resp)
		return
	}

	keyRedis := dBName + "_" + "ROOM" + "_" + idRoom
	dataRedis, err := json.Marshal(&data)
	if err != nil {
		LogError("Marshal", err, "room/GetRoom")
	}
	err = setRedis(keyRedis, string(dataRedis))
	if err != nil {
		LogError("Set Redis", err, "room/GetRoom")
	}

	// return success
	resp.Code = 200
	resp.Status = "Success"
	c.JSON(resp.Code, gin.H{"response": resp})
	return
}

// DeleteRoom function for delete room
func (RoomCtr) DeleteRoom(c *gin.Context) {
	var resp Response
	var data models.Room

	idRoom := c.Param("room_id")
	if idRoom == "" {
		resp.Status = "not found"
		resp.Message = "params is not found"
		c.JSON(404, resp)
		return
	}

	roomID, err := strconv.Atoi(idRoom)
	if err != nil {
		LogError("strconv atoi", err, "room/DeleteRoom")
	}

	data.ID = (uint)(roomID)
	roomModel := models.RoomModel{}
	err = roomModel.DeleteRoom(data)

	// return success
	resp.Code = 200
	resp.Status = "Success deleted"
	c.JSON(resp.Code, gin.H{"response": resp})
	return
}
