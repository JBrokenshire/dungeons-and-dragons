package requests

type CharacterRequest struct {
	Name              string `json:"name"`
	Level             int    `json:"level"`
	ProfilePictureURL string `json:"profile_picture_url"`
	ClassID           int    `json:"class_id"`
	RaceID            int    `json:"race_id"`

	Strength               int  `json:"strength"`
	Dexterity              int  `json:"dexterity"`
	Constitution           int  `json:"constitution"`
	Intelligence           int  `json:"intelligence"`
	Wisdom                 int  `json:"wisdom"`
	Charisma               int  `json:"charisma"`
	ProficientStrength     bool `json:"proficient_strength"`
	ProficientDexterity    bool `json:"proficient_dexterity"`
	ProficientConstitution bool `json:"proficient_constitution"`
	ProficientIntelligence bool `json:"proficient_intelligence"`
	ProficientWisdom       bool `json:"proficient_wisdom"`
	ProficientCharisma     bool `json:"proficient_charisma"`

	WalkingSpeedModifier int  `json:"walking_speed_modifier"`
	Inspiration          bool `json:"inspiration"`
	CurrentHitPoints     int  `json:"current_hit_points"`
	MaxHitPoints         int  `json:"max_hit_points"`
	TempHitPoints        int  `json:"temp_hit_points"`
}

func NewCharacterRequest(cr *CharacterRequest) CharacterRequest {
	if cr.Name == "" {
		cr.Name = "DEFAULT NAME"
	}
	if cr.Level == 0 {
		cr.Level = 1
	}
	if cr.ClassID == 0 {
		cr.ClassID = 1
	}
	if cr.RaceID == 0 {
		cr.RaceID = 1
	}
	if cr.Strength == 0 {
		cr.Strength = 10
	}
	if cr.Dexterity == 0 {
		cr.Dexterity = 10
	}
	if cr.Constitution == 0 {
		cr.Constitution = 10
	}
	if cr.Intelligence == 0 {
		cr.Intelligence = 10
	}
	if cr.Wisdom == 0 {
		cr.Wisdom = 10
	}
	if cr.Charisma == 0 {
		cr.Charisma = 10
	}
	if cr.CurrentHitPoints == 0 {
		cr.CurrentHitPoints = 1
	}
	if cr.MaxHitPoints == 0 {
		cr.MaxHitPoints = 1
	}

	return *cr
}

func (cr CharacterRequest) IsEmpty() bool {
	if cr.Name == "" && cr.Level == 0 && cr.ClassID == 0 && cr.RaceID == 0 && cr.Strength == 0 && cr.Dexterity == 0 && cr.Constitution == 0 && cr.Intelligence == 0 && cr.Wisdom == 0 && cr.Charisma == 0 {
		return true
	}
	return false
}
