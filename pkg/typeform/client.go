package typeform

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	accessToken string
}

func NewClient(accessToken string) *Client {
	return &Client{
		accessToken: accessToken,
	}
}

// Available themes and their IDs from themes.json
var Themes = map[string]string{
	"default":     "qHWOQ7",  // Default Theme
	"orbital":     "GKFi5U",  // Orbital
	"fractal":     "XoieTe",  // Fractal
	"art-splash":  "kjw3vL",  // Art Splash
	"pin-up":      "sFQj1y",  // Pin-up
	"paper-invite": "NqJgJG", // Paper Invite
	"mystery":     "paovnB",  // Mystery
	"school-bell": "Kd9YUR",  // School Bell
	"spaceboy":    "ZFDEUn",  // Spaceboy
	"desk-space":  "jQPvMd",  // Desk Space
}

// GetThemeURL returns the theme URL for a given theme name
func GetThemeURL(themeName string) (string, error) {
	themeID, exists := Themes[themeName]
	if !exists {
		return "", fmt.Errorf("theme '%s' not found. Available themes: %v", themeName, getThemeNames())
	}
	return fmt.Sprintf("https://api.typeform.com/themes/%s", themeID), nil
}

// getThemeNames returns a slice of available theme names
func getThemeNames() []string {
	names := make([]string, 0, len(Themes))
	for name := range Themes {
		names = append(names, name)
	}
	return names
}

type Form struct {
	Title  string      `json:"title"`
	Fields []FormField `json:"fields"`
	Theme  Theme       `json:"theme,omitempty"`
}

type Theme struct {
	Href string `json:"href,omitempty"`
}

type FormField struct {
	Type        string      `json:"type"`
	Title       string      `json:"title"`
	Properties  Properties  `json:"properties"`
	Validations Validations `json:"validations,omitempty"`
}

type Properties struct {
	Description            string   `json:"description,omitempty"`
	Choices               []Choice `json:"choices,omitempty"`
	AllowMultipleSelection bool     `json:"allow_multiple_selection,omitempty"`
}

type Validations struct {
	Required bool `json:"required,omitempty"`
}

type Choice struct {
	Label string `json:"label"`
}

type FormURLs struct {
	Display   string
	Responses string
}

func (c *Client) CreateForm(form Form) (*FormURLs, error) {
	jsonData, err := json.Marshal(form)
	if err != nil {
		return nil, fmt.Errorf("error marshaling form: %w", err)
	}

	req, err := http.NewRequest("POST", "https://api.typeform.com/forms", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("error creating form. Status: %d, Response: %s", resp.StatusCode, string(body))
	}

	var response struct {
		Links struct {
			Display   string `json:"display"`
			Responses string `json:"responses"`
		} `json:"_links"`
	}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	return &FormURLs{
		Display:   response.Links.Display,
		Responses: response.Links.Responses,
	}, nil
} 