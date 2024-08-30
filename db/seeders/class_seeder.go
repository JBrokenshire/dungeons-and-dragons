package seeders

import (
	"dungeons-and-dragons/models"
	"log"
)

func (s *Seeder) SetClasses() {
	classes := []models.Class{
		{ID: 1, Name: "Artificer"},
		{ID: 2, Name: "Barbarian", Description: "A fierce warrior who can enter a battle rage"},
		{ID: 3, Name: "Bard", Description: "An inspiring magician whose power echoes the music of creation"},
		{ID: 4, Name: "Blood Hunter", Description: "Willing to suffer whatever it takes to achieve victory, these adept warriors have forged themselves into a potent force dedicated to protecting the innocent."},
		{ID: 5, Name: "Cleric", Description: "A priestly champion who wields divine magic in service of a higher power"},
		{ID: 6, Name: "Druid", Description: "A priest of the Old Faith, wielding the powers of nature and adopting animal forms"},
		{ID: 7, Name: "Fighter", Description: "A master of martial combat, skilled with a variety of weapons and armor"},
		{ID: 8, Name: "Monk", Description: "A master of martial arts, harnessing the power of the body in pursuit of physical and spiritual perfection"},
		{ID: 9, Name: "Paladin", Description: "A holy warrior bound to a sacred oath"},
		{ID: 10, Name: "Ranger", Description: "A warrior who combats threats on the edges of civilization"},
		{ID: 11, Name: "Rogue", Description: "A scoundrel who uses stealth and trickery to overcome obstacles and enemies"},
		{ID: 12, Name: "Sorcerer", Description: "A spellcaster who draws on inherent magic from a gift or bloodline"},
		{ID: 13, Name: "Warlock", Description: "A wielder of magic that is derived from a bargain with an extraplanar entity"},
		{ID: 14, Name: "Wizard", Description: "A scholarly magic-user capable of manipulating the structures of reality"},
	}

	for _, class := range classes {
		err := s.DB.Where("id = ?", class.ID).FirstOrCreate(&class).Error
		if err != nil {
			log.Printf("error creating class %s: %v", class.Name, err.Error())
		}
	}
}
