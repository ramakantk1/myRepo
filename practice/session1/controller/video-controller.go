// controller for the video service, include all the handlerd for the service
// an interfact similar to service
//the input parameter will be extracted from the context object
//Define the struct to implemnet the interface, which will have the reference of the inteface define in the service

package controller

import (
	"microservices/my-project/practice/session1/model"
	"microservices/my-project/practice/session1/service"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	Save(ctx *gin.Context) model.Video
	FindAll() []model.Video
}

// name of the object and the refernce to the service inteface
type controller struct {
	service service.VideoService
}

//constructor function which receive input as service object and
//return the controller struct to which we pass the service
func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

//implement the function
func (c *controller) Save(ctx *gin.Context) model.Video {
	//declare a local object to extract the video
	var video model.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}

func (c *controller) FindAll() []model.Video {
	return c.service.FindAll()
}
