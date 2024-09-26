package requests

type CharacterRequest struct {
	Name              string `json:"name"`
	Level             int    `json:"level"`
	ProfilePictureURL string `json:"profile_picture_url"`
	ClassID           int    `json:"class_id"`
	RaceID            int    `json:"race_id"`

	Strength     int `json:"strength"`
	Dexterity    int `json:"dexterity"`
	Constitution int `json:"constitution"`
	Intelligence int `json:"intelligence"`
	Wisdom       int `json:"wisdom"`
	Charisma     int `json:"charisma"`

	WalkingSpeedModifier int  `json:"walking_speed_modifier"`
	Inspiration          bool `json:"inspiration"`
	CurrentHitPoints     int  `json:"current_hit_points"`
	MaxHitPoints         int  `json:"max_hit_points"`
	TempHitPoints        int  `json:"temp_hit_points"`
}

func (cr CharacterRequest) IsEmpty() bool {
	if cr.Name == "" && cr.Level == 0 && cr.ClassID == 0 && cr.RaceID == 0 && cr.Strength == 0 && cr.Dexterity == 0 && cr.Constitution == 0 && cr.Intelligence == 0 && cr.Wisdom == 0 && cr.Charisma == 0 {
		return true
	}
	return false
}

type ClassRequest struct {
	Name             string
	ShortDescription string
	LongDescription  string
}
