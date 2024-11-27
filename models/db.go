package models

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBconnect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("fail to connect db")
	}
	DB = db
}

func DBMigrate() {
	DB.AutoMigrate(&Book{}, &Unit{})
}

func DBSeed() {
	books := []Book{
		{Model: gorm.Model{ID: 1}, Title: "The Random Journey"},
		{Model: gorm.Model{ID: 2}, Title: "Shadows of the Unknown"},
		{Model: gorm.Model{ID: 3}, Title: "Whispers in the Dark"},
		{Model: gorm.Model{ID: 4}, Title: "A Tale of Two Coders"},
		{Model: gorm.Model{ID: 5}, Title: "Beyond the Horizon"},
		{Model: gorm.Model{ID: 6}, Title: "Rise of the Machines"},
		{Model: gorm.Model{ID: 7}, Title: "Echoes of Eternity"},
		{Model: gorm.Model{ID: 8}, Title: "The Final Algorithm"},
		{Model: gorm.Model{ID: 9}, Title: "Secrets of the Ancients"},
		{Model: gorm.Model{ID: 10}, Title: "Chronicles of Time"},
		{Model: gorm.Model{ID: 11}, Title: "Guardians of Light"},
		{Model: gorm.Model{ID: 12}, Title: "Winds of Change"},
		{Model: gorm.Model{ID: 13}, Title: "The Infinite Loop"},
		{Model: gorm.Model{ID: 14}, Title: "Parallel Dimensions"},
		{Model: gorm.Model{ID: 15}, Title: "The Forgotten Key"},
		{Model: gorm.Model{ID: 16}, Title: "Labyrinth of Dreams"},
		{Model: gorm.Model{ID: 17}, Title: "The Eternal Flame"},
		{Model: gorm.Model{ID: 18}, Title: "Voices from the Void"},
		{Model: gorm.Model{ID: 19}, Title: "Legends of the Lost"},
		{Model: gorm.Model{ID: 20}, Title: "The Starlit Path"},
		{Model: gorm.Model{ID: 21}, Title: "Fragments of Reality"},
		{Model: gorm.Model{ID: 22}, Title: "Through the Looking Glass"},
		{Model: gorm.Model{ID: 23}, Title: "Tales of the Unknown"},
		{Model: gorm.Model{ID: 24}, Title: "The Quantum Code"},
		{Model: gorm.Model{ID: 25}, Title: "Minds of the Future"},
		{Model: gorm.Model{ID: 26}, Title: "Beneath the Surface"},
		{Model: gorm.Model{ID: 27}, Title: "Chronicles of Shadows"},
		{Model: gorm.Model{ID: 28}, Title: "Realm of Imagination"},
		{Model: gorm.Model{ID: 29}, Title: "The Celestial Voyage"},
		{Model: gorm.Model{ID: 30}, Title: "Beyond the Stars"},
		{Model: gorm.Model{ID: 31}, Title: "In the Eye of the Storm"},
		{Model: gorm.Model{ID: 32}, Title: "Wonders of the Deep"},
		{Model: gorm.Model{ID: 33}, Title: "The Alchemist's Secret"},
		{Model: gorm.Model{ID: 34}, Title: "Pulse of the Cosmos"},
		{Model: gorm.Model{ID: 35}, Title: "Dancing with Shadows"},
		{Model: gorm.Model{ID: 36}, Title: "Unveiling the Mystery"},
		{Model: gorm.Model{ID: 37}, Title: "The Wanderer's Quest"},
		{Model: gorm.Model{ID: 38}, Title: "Falling Through Time"},
		{Model: gorm.Model{ID: 39}, Title: "The Hidden Legacy"},
		{Model: gorm.Model{ID: 40}, Title: "Flight of the Phoenix"},
		{Model: gorm.Model{ID: 41}, Title: "Reflections of the Past"},
		{Model: gorm.Model{ID: 42}, Title: "The Forbidden Forest"},
		{Model: gorm.Model{ID: 43}, Title: "Pathways to Infinity"},
		{Model: gorm.Model{ID: 44}, Title: "The Last Frontier"},
		{Model: gorm.Model{ID: 45}, Title: "Shadows and Light"},
		{Model: gorm.Model{ID: 46}, Title: "The Architect's Vision"},
		{Model: gorm.Model{ID: 47}, Title: "Chronicles of the Moon"},
		{Model: gorm.Model{ID: 48}, Title: "The Lost Artifact"},
		{Model: gorm.Model{ID: 49}, Title: "Voyage to the Unknown"},
		{Model: gorm.Model{ID: 50}, Title: "The Enchanted Garden"},
		{Model: gorm.Model{ID: 51}, Title: "The Symphony of Time"},
		{Model: gorm.Model{ID: 52}, Title: "The Gates of Eternity"},
		{Model: gorm.Model{ID: 53}, Title: "Memoirs of the Void"},
		{Model: gorm.Model{ID: 54}, Title: "The Dreamer's Awakening"},
		{Model: gorm.Model{ID: 55}, Title: "Beyond the Veil"},
		{Model: gorm.Model{ID: 56}, Title: "The Labyrinth Key"},
		{Model: gorm.Model{ID: 57}, Title: "The Silent Whisper"},
		{Model: gorm.Model{ID: 58}, Title: "Echoes of Destiny"},
		{Model: gorm.Model{ID: 59}, Title: "The Timeless Journey"},
		{Model: gorm.Model{ID: 60}, Title: "Secrets of the Deep"},
		{Model: gorm.Model{ID: 61}, Title: "Horizons of the Soul"},
		{Model: gorm.Model{ID: 62}, Title: "The Keeper's Watch"},
		{Model: gorm.Model{ID: 63}, Title: "The Edge of Reality"},
		{Model: gorm.Model{ID: 64}, Title: "The Shattered Mirror"},
		{Model: gorm.Model{ID: 65}, Title: "The Eternal Puzzle"},
		{Model: gorm.Model{ID: 66}, Title: "The Legacy of Time"},
		{Model: gorm.Model{ID: 67}, Title: "Wings of the Unknown"},
		{Model: gorm.Model{ID: 68}, Title: "Dawn of the Horizon"},
		{Model: gorm.Model{ID: 69}, Title: "The Silent Guardian"},
		{Model: gorm.Model{ID: 70}, Title: "Into the Abyss"},
		{Model: gorm.Model{ID: 71}, Title: "The Stars Align"},
		{Model: gorm.Model{ID: 72}, Title: "The Final Ascent"},
		{Model: gorm.Model{ID: 73}, Title: "Whispers of the Wind"},
		{Model: gorm.Model{ID: 74}, Title: "The Sapphire Code"},
		{Model: gorm.Model{ID: 75}, Title: "The Guardian's Quest"},
		{Model: gorm.Model{ID: 76}, Title: "Legends of the Cosmos"},
		{Model: gorm.Model{ID: 77}, Title: "A World Beyond"},
		{Model: gorm.Model{ID: 78}, Title: "The Midnight Echo"},
		{Model: gorm.Model{ID: 79}, Title: "The Path of Light"},
		{Model: gorm.Model{ID: 80}, Title: "The Enigma of Time"},
		{Model: gorm.Model{ID: 81}, Title: "Dawn of the Ancients"},
		{Model: gorm.Model{ID: 82}, Title: "The Phoenix's Call"},
		{Model: gorm.Model{ID: 83}, Title: "The Forgotten Chronicles"},
		{Model: gorm.Model{ID: 84}, Title: "The Veil of Shadows"},
		{Model: gorm.Model{ID: 85}, Title: "The Cradle of Dreams"},
		{Model: gorm.Model{ID: 86}, Title: "The Eternal Bloom"},
		{Model: gorm.Model{ID: 87}, Title: "The Labyrinth's Secret"},
		{Model: gorm.Model{ID: 88}, Title: "Wings of Eternity"},
		{Model: gorm.Model{ID: 89}, Title: "The Starlight Tale"},
		{Model: gorm.Model{ID: 90}, Title: "Beyond the Mist"},
		{Model: gorm.Model{ID: 91}, Title: "The Altar of Stars"},
		{Model: gorm.Model{ID: 92}, Title: "The Oracle's Vision"},
		{Model: gorm.Model{ID: 93}, Title: "The Voyage Home"},
		{Model: gorm.Model{ID: 94}, Title: "The Shimmering Abyss"},
		{Model: gorm.Model{ID: 95}, Title: "The Architect's Code"},
		{Model: gorm.Model{ID: 96}, Title: "The Song of Eternity"},
		{Model: gorm.Model{ID: 97}, Title: "Echoes of the Void"},
		{Model: gorm.Model{ID: 98}, Title: "The Awakening Flame"},
		{Model: gorm.Model{ID: 99}, Title: "The Path Untold"},
		{Model: gorm.Model{ID: 100}, Title: "The Final Horizon"},
	}

	DB.Create(books)

}
