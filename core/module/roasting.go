package module

import (
	"fmt"
	"instagram-roasting/libs"
)

type RoastingUC struct {
	scrappingIg *libs.ScrapingIGProfile
	gemini      *libs.GeminiAI
}

func NewRoastingUC(scrappingIg *libs.ScrapingIGProfile, gemini *libs.GeminiAI) *RoastingUC {
	return &RoastingUC{
		scrappingIg: scrappingIg,
		gemini:      gemini,
	}
}

func (r *RoastingUC) GetRoastedProfile(username string) (string, error) {
	// Get profile from IG
	profile, err := r.scrappingIg.GetProfile(username)
	if err != nil {
		return "", err
	}
	fmt.Println("Data Profile", profile)

	// Get roast from Gemini
	roast, err := r.gemini.RoastingWIthGemini(profile)
	if err != nil {
		return "", err
	}

	return roast, nil
}
