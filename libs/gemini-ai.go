package libs

import (
	"context"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	config "instagram-roasting"
	"instagram-roasting/dto"
	"io"
	"net/http"
	"strings"
)

type GeminiAI struct {
	config *config.Config
}

func NewGeminiAI(config *config.Config) *GeminiAI {
	return &GeminiAI{
		config: config,
	}
}

func (g *GeminiAI) ImageCaptioningWIthGemini(imageURL string) (string, error) {
	textPrompt := "Buatlah deskripsi dari foto profil berikut dari pose, pakaian yang dipakai, hingga ekspresi wajahnya. Gunakan kata-kata yang kreatif dan deskriptif. Hasilkan deskripsi dalam satu kalimat."

	ctx := context.Background()

	if imageURL == "" {
		return "", fmt.Errorf("image URL is empty")
	}

	// Fetch image data from URL
	respImg, err := http.Get(imageURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch image: %v", err)
	}
	defer respImg.Body.Close()

	imgData, err := io.ReadAll(respImg.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read image data: %v", err)
	}

	// Generate text prompt
	client, err := genai.NewClient(ctx, option.WithAPIKey(g.config.GetGeminiConfig().APIKey))
	if err != nil {
		return "", fmt.Errorf("failed to create client: %v", err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx,
		genai.Text(textPrompt),
		genai.ImageData("jpg", imgData))
	if err != nil {
		return "", fmt.Errorf("failed to generate content: %v", err)
	}

	// Extract text from response
	var resultText strings.Builder
	for _, part := range resp.Candidates {
		for _, content := range part.Content.Parts {
			resultText.WriteString(fmt.Sprintf("%v", content))
		}
	}

	return resultText.String(), nil

}

func (g *GeminiAI) RoastingWIthGemini(dataProfile dto.IgScrapped) (string, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(g.config.GetGeminiConfig().APIKey))
	if err != nil {
		return "", fmt.Errorf("failed to create client: %v", err)
	}
	defer client.Close()

	// Generate Roating Image
	imageRoasted, err := g.ImageCaptioningWIthGemini(dataProfile.Avatar)
	if err != nil {
		return "", fmt.Errorf("failed to generate image roasted: %v", err)
	}

	textPrompt := fmt.Sprintf("Saya memiliki data profil Instagram seseorang. Berikut detailnya:\n\n- Deskripsi Foto Profil: %s \n- Nama: %s \n- Deskripsi: %s \n\nBuatlah roasting lucu tentang profil ini dengan gaya santai dan humor khas Indonesia. Gunakan kata-kata yang tidak ofensif, tetapi tetap menggelitik. Hasilkan roasting dalam satu paragraf pendek.", imageRoasted, dataProfile.Title, dataProfile.Description)

	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(textPrompt))
	if err != nil {
		return "", fmt.Errorf("failed to generate content: %v", err)
	}

	// Extract text from response
	var resultText strings.Builder
	for _, part := range resp.Candidates {
		for _, content := range part.Content.Parts {
			resultText.WriteString(fmt.Sprintf("%v", content))
		}
	}

	return resultText.String(), nil
}
