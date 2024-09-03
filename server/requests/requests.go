package requests

type CharacterRequest struct {
	Name    string
	Level   int
	ClassID int
	RaceID  int
	isEmpty func() bool
}

func (cr CharacterRequest) IsEmpty() bool {
	if cr.Name == "" && cr.Level == 0 && cr.ClassID == 0 && cr.RaceID == 0 {
		return true
	}
	return false
}
