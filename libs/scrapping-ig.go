package libs

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly/v2"
	"instagram-roasting/dto"
	"log"
	"time"
)

type ScrapingIGProfile struct {
}

func NewScrapingIGProfile() *ScrapingIGProfile {
	return &ScrapingIGProfile{}
}

func (s *ScrapingIGProfile) GetProfile(username string) (dto.IgScrapped, error) {
	c := colly.NewCollector(
		colly.AllowURLRevisit(),
		colly.IgnoreRobotsTxt(),
	)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*.instagram.com", // Berlaku untuk domain Instagram
		RandomDelay: 2 * time.Second,   // Random delay hingga 2 detik
		Parallelism: 1,                 // Hanya 1 request dalam satu waktu
	})

	var response dto.IgScrapped
	// Callback to get the title of the page
	c.OnHTML("title", func(e *colly.HTMLElement) {
		title := e.Text
		if title != "" {
			response.Title = title
		} else {
			response.Title = ""
		}
	})

	// Callback to get the og:image meta tag content
	c.OnHTML(`meta[property="og:image"]`, func(e *colly.HTMLElement) {
		avatar := e.Attr("content")
		if avatar != "" {
			response.Avatar = avatar
		} else {
			response.Avatar = ""
		}
	})

	// Callback to get the og:description meta tag content
	c.OnHTML(`meta[name="description"]`, func(e *colly.HTMLElement) {
		description := e.Attr("content")
		if description != "" {
			response.Description = description
		} else {
			response.Description = ""
		}
	})

	// Callback before making a request
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting " + r.URL.String())
	})

	// Callback handling errors
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Error:", r.StatusCode, err)
	})

	// Start scraping
	url := fmt.Sprintf("https://www.instagram.com/%s/", username)
	err := c.Visit(url)
	if err != nil {
		fmt.Println("Error visiting URL", err)
	}

	return response, nil
}

func (s *ScrapingIGProfile) GetProfileWithHeadless(username string) (dto.IgScrapped, error) {
	// Context untuk Chromedp
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Set timeout
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// URL profil Instagram
	url := fmt.Sprintf("https://www.instagram.com/%s/", username)

	response := dto.IgScrapped{}

	// Eksekusi tugas di Chromedp
	err := chromedp.Run(ctx,
		// Buka halaman profil
		chromedp.Navigate(url),
		// Tunggu hingga elemen 'title' tersedia
		chromedp.Title(&response.Title),
		// Ambil meta tag avatar
		chromedp.AttributeValue(`meta[property="og:image"]`, "content", &response.Avatar, nil),
		// Ambil meta tag description
		chromedp.AttributeValue(`meta[name="description"]`, "content", &response.Description, nil),
	)
	if err != nil {
		return dto.IgScrapped{}, fmt.Errorf("failed to scrape profile: %w", err)
	}

	// Kembalikan hasil
	return response, nil
}
