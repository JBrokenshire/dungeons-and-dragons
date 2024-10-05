package seeders

import (
	"dnd-api/db/models"
	"log"
)

func (s *Seeder) SetItems() {
	items := []models.Item{
		{
			ID:     1,
			Name:   "Dagger",
			Meta:   "Melee Weapon",
			Weight: 1,
			Cost:   2,
			Notes:  "Simple, Finesse, Light, Thrown, Range(20/60)",
		},
		{
			ID:     2,
			Name:   "Greataxe",
			Meta:   "Melee Weapon",
			Weight: 7,
			Cost:   30,
			Notes:  "Martial, Heavy, Two-Handed",
		},
		{
			ID:     3,
			Name:   "Crossbow, Light",
			Meta:   "Ranged Weapon",
			Weight: 5,
			Cost:   25,
			Notes:  "Simple, Ammunition, Loading, Range, Two-Handed, Slow, Range(80/320)",
		},
		{
			ID:     4,
			Name:   "Studded Leather",
			Meta:   "Light Armour",
			Weight: 13,
			Cost:   45,
			Notes:  "AC 12",
		},
		{
			ID:    5,
			Name:  "Potion of Healing",
			Meta:  "Gear • Potion",
			Cost:  100,
			Notes: "Healing, Combat, Consumable",
		},
		{
			ID:     7,
			Name:   "Alchemist's Supplies",
			Meta:   "Gear • Tool",
			Cost:   50,
			Weight: 8,
			Notes:  "Utility",
		},
		{
			ID:     8,
			Name:   "Crossbow Bolts",
			Meta:   "Gear • Ammunition",
			Cost:   0.05,
			Weight: 0.075,
			Notes:  "Ammunition, Damage, Combat",
		},
		{
			ID:    9,
			Name:  "Spider Venom",
			Meta:  "Gear • Tool, Potion",
			Cost:  100,
			Notes: "Damage, Utility, Combat, Consumable",
		},
		{
			ID:     10,
			Name:   "Bedroll",
			Meta:   "Gear • Adventuring Gear",
			Cost:   1,
			Weight: 1,
			Notes:  "Utility",
		},
		{
			ID:     11,
			Name:   "Mess Kit",
			Meta:   "Gear • Adventuring Gear",
			Cost:   0.2,
			Weight: 1,
			Notes:  "Social, Utility",
		},
		{
			ID:     12,
			Name:   "Rations (1 Day)",
			Meta:   "Gear • Adventuring Gear",
			Cost:   0.5,
			Weight: 2,
			Notes:  "Social, Utility, Consumable",
		},
		{
			ID:     13,
			Name:   "Rope, Hempen (50 feet)",
			Meta:   "Gear • Adventuring Gear",
			Cost:   1,
			Weight: 10,
			Notes:  "Utility, Exploration",
		},
		{
			ID:     14,
			Name:   "Tinderbox",
			Meta:   "Gear • Adventuring Gear",
			Cost:   0.5,
			Weight: 1,
			Notes:  "Utility, Exploration",
		},
		{
			ID:     15,
			Name:   "Torch",
			Meta:   "Gear • Adventuring Gear",
			Cost:   0.01,
			Weight: 1,
			Notes:  "Damage, Utility, Exploration, Combat",
		},
		{
			ID:     16,
			Name:   "Waterskin",
			Meta:   "Gear • Adventuring Gear",
			Cost:   0.2,
			Weight: 5,
			Notes:  "Container",
		},
		{
			ID:     17,
			Name:   "Dagger of Returning",
			Meta:   "Melee Weapon • Uncommon",
			Weight: 1,
			Notes:  "Simple, Finesse, Light, Thrown, Range(20/60)",
			Rarity: "Uncommon",
		},
		{
			ID:     18,
			Name:   "Handaxe, +1",
			Meta:   "Melee Weapon • Uncommon",
			Weight: 2,
			Notes:  "Simple, Light, Thrown, Range(20/60)",
			Rarity: "Uncommon",
		},
		{
			ID:     19,
			Name:   "Warhammer, +1",
			Meta:   "Melee Weapon • Uncommon",
			Weight: 2,
			Notes:  "Martial, Versatile",
			Rarity: "Uncommon",
		},
		{
			ID:     20,
			Name:   "Clothes, Common",
			Meta:   "Gear • Adventuring Gear",
			Cost:   0.5,
			Weight: 3,
			Notes:  "Social, Outerwear",
		},
		{
			ID:    21,
			Name:  "Playing Card Set",
			Meta:  "Gear • Tool",
			Cost:  0.5,
			Notes: "Social",
		},
		{
			ID:     22,
			Name:   "Scale Mail",
			Meta:   "Medium Armour",
			Cost:   50,
			Weight: 45,
			Notes:  "AC 14, Stealth Disadvantage",
		},
		{
			ID:     23,
			Name:   "Smith's Tools",
			Meta:   "Gear • Tool",
			Cost:   20,
			Weight: 8,
			Notes:  "Utility",
		},
		{
			ID:     24,
			Name:   "Greatsword",
			Meta:   "Melee Weapon",
			Cost:   50,
			Weight: 6,
			Notes:  "Martial, Heavy, Two-Handed",
		},
		{
			ID:     25,
			Name:   "Half Plate",
			Meta:   "Medium Armour",
			Cost:   750,
			Weight: 40,
			Notes:  "AC 15, Stealth Disadvantage",
		},
		{
			ID:     26,
			Name:   "Hunting Trap",
			Meta:   "Gear • Adventuring Gear",
			Cost:   5,
			Weight: 25,
			Notes:  "Damage, Utility",
		},
		{
			ID:     27,
			Name:   "Quarterstaff",
			Meta:   "Melee Weapon",
			Cost:   0.2,
			Weight: 4,
			Notes:  "Simple, Versatile",
		},
		{
			ID:    28,
			Name:  "Sapphire",
			Meta:  "Jewel",
			Cost:  150,
			Notes: "Gem",
		},
		{
			ID:     29,
			Name:   "Cloak of Protection",
			Meta:   "Wondrous Item",
			Notes:  "Warding, Outerwear",
			Rarity: "Uncommon",
		},
		{
			ID:     30,
			Name:   "Chain Mail",
			Meta:   "Heavy Armour",
			Cost:   75,
			Weight: 55,
			Notes:  "AC 16, STR 13, Stealth Disadvantage",
		},
		{
			ID:     31,
			Name:   "Morningstar",
			Meta:   "Melee Weapon",
			Cost:   15,
			Weight: 4,
			Notes:  "Martial",
		},
		{
			ID:     32,
			Name:   "Crowbar",
			Meta:   "Gear • Adventuring Gear",
			Cost:   2,
			Weight: 5,
			Notes:  "Utility, Exploration",
		},
		{
			ID:     33,
			Name:   "Hammer",
			Meta:   "Gear • Adventuring Gear",
			Cost:   1,
			Weight: 3,
			Notes:  "Utility",
		},
		{
			ID:     34,
			Name:   "Javelin",
			Meta:   "Thrown Weapon",
			Cost:   0.5,
			Weight: 2,
			Notes:  "Simple, Thrown, Range(30/120)",
		},
		{
			ID:     35,
			Name:   "Piton",
			Meta:   "Gear • Adventuring Gear",
			Cost:   0.05,
			Weight: 0.25,
			Notes:  "Utility, Exploration",
		},
		{
			ID:     36,
			Name:   "Handaxe",
			Meta:   "Melee Weapon",
			Cost:   5,
			Weight: 2,
			Notes:  "Simple, Light, Thrown, Range(20/60)",
		},
	}

	for _, item := range items {
		err := s.DB.Where("id = ?", item.ID).FirstOrCreate(&item).Error
		if err != nil {
			log.Printf("error creating item %s: %v", item.Name, err)
		}
	}
}
