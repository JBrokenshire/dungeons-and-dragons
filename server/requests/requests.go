package requests

type CharacterRequest struct {
	Name              string `json:"name"`
	Level             int    `json:"level"`
	ProfilePictureURL string `json:"profile_picture_url"`
	ClassID           int    `json:"class_id"`
	RaceID            int    `json:"race_id"`
	isEmpty           func() bool
}

func (cr CharacterRequest) IsEmpty() bool {
	if cr.Name == "" && cr.Level == 0 && cr.ClassID == 0 && cr.RaceID == 0 {
		return true
	}
	return false
}

type ClassRequest struct {
	Name        string
	Description string
}
