package HeroClasses

AllClasses := make(map[string]class)
AllClasses["Mage"] = class{ name: "Mage", damage_type: "Magical", health: 80, mana: 100 }
AllClasses["Warrior"] = class{ name: "Warrior", damage_type: "Psysical", health: 100, mana: 0}