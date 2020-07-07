//Package service VideoService service show have below components
// An interface with methods needed
// an struct which impelement the interface
//construction function for the interface and pointer to struct
//impelent the function define in the interface
package service

import (
	"microservices/my-project/practice/session1/model"
)

//VideoService should be created
type VideoService interface {
	//save metod accept an input of model.video and retun the same
	Save(model.Video) model.Video
	
	//Final all methods return sclice of videos
	FindAll() []model.Video
}

// an struct to implementat the inteface and will include the data to be return
type videoServiceimpl struct {
	videos []model.Video
}

//New construction function for the interface
func New() VideoService {
	return &videoServiceimpl{}
}

//Implementing the function
func (service *videoServiceimpl) Save(video model.Video) model.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoServiceimpl) FindAll() []model.Video {
	// all the object in the struct
	return service.videos
}
