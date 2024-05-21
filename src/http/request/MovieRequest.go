package request

import "github.com/amirhossein2831/httpServerGo/src/model"

type MovieRequest struct {
	Name        string `json:"name"`
	Director    string `json:"director"`
	Publication string `json:"publication"`
	WatchTime   int    `json:"watch_time"`
}

func (mr *MovieRequest) ToMovie() model.Movie {
	return model.Movie{
		Name:        mr.Name,
		Director:    mr.Director,
		Publication: mr.Publication,
		WatchTime:   mr.WatchTime,
	}
}
