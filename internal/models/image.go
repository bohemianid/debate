package models

type Image struct {
	ID       int    `json:"id"`
	OriginID int    `json:"post_id"`
	Origin   string `json:"origin"`
	Path     string `json:"path"`
	Image    []byte `json:"image"`
}

type ImageService interface {
	UploadImage(file []byte, filename string) error
	GetImage(id int, originID int, origin string) (*Image, error)
	DeleteImage(id int) error
}
