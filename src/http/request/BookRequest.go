package request

import "github.com/amirhossein2831/httpServerGo/src/model"

type BookRequest struct {
	Name        string `gorm:"type:varchar(25);not null" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func (br *BookRequest) ToBook() model.Book {
	return model.Book{
		Name:        br.Name,
		Author:      br.Author,
		Publication: br.Publication,
	}
}
