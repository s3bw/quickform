package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/skip2/go-qrcode"
	"github.com/spf13/cobra"

	"quickform/pkg/typeform"
)

var (
	theme string
	name  string
)

var pollCmd = &cobra.Command{
	Use:   "poll",
	Short: "Create a poll with multiple choices",
	Long: `Create a poll with multiple choices and share it via URL and QR code.
Available themes: default, orbital, fractal, art-splash, pin-up, paper-invite, mystery, school-bell, spaceboy, desk-space

Run 'quickform themes' to see detailed theme descriptions.`,
	RunE:  runPoll,
}

func init() {
	rootCmd.AddCommand(pollCmd)
	pollCmd.Flags().StringVar(&theme, "theme", "", "Theme for the poll (e.g., 'default', 'orbital', 'fractal')")
	pollCmd.Flags().StringVar(&name, "name", "", "Name for the poll (default: 'Quick Form')")
}

func runPoll(cmd *cobra.Command, args []string) error {
	accessToken := os.Getenv("QUICKFORM_ACCESS_TOKEN")
	if accessToken == "" {
		return fmt.Errorf("QUICKFORM_ACCESS_TOKEN environment variable is not set")
	}

	client := typeform.NewClient(accessToken)

	// Get the question
	fmt.Print("Enter Question:\n$ ")
	reader := bufio.NewReader(os.Stdin)
	question, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("error reading question: %w", err)
	}
	question = strings.TrimSpace(question)

	// Get the choices
	fmt.Println("\nEnter choices (press Enter to finish):")
	var choices []typeform.Choice
	for {
		fmt.Print("$ ")
		choice, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("error reading choice: %w", err)
		}
		choice = strings.TrimSpace(choice)
		
		if choice == "" {
			break
		}
		choices = append(choices, typeform.Choice{Label: choice})
	}

	if len(choices) < 2 {
		return fmt.Errorf("at least 2 choices are required")
	}

	// Set form title
	formTitle := "Quick Form"
	if name != "" {
		formTitle = name
	}

	// Create the form
	form := typeform.Form{
		Title: formTitle,
		Fields: []typeform.FormField{
			{
				Type:  "multiple_choice",
				Title: question,
				Properties: typeform.Properties{
					AllowMultipleSelection: false,
					Choices:               choices,
				},
				Validations: typeform.Validations{
					Required: true,
				},
			},
		},
	}

	// Set theme if specified
	if theme != "" {
		themeURL, err := typeform.GetThemeURL(theme)
		if err != nil {
			return err
		}
		form.Theme = typeform.Theme{
			Href: themeURL,
		}
	}

	// Create the form
	urls, err := client.CreateForm(form)
	if err != nil {
		return fmt.Errorf("error creating form: %w", err)
	}

	fmt.Printf("\nPoll created successfully!\n")
	fmt.Printf("Share this URL: %s\n", urls.Display)
	fmt.Printf("View responses: %s\n\n", urls.Responses)

	// Generate and display QR code
	qr, err := qrcode.New(urls.Display, qrcode.Medium)
	if err != nil {
		return fmt.Errorf("error generating QR code: %w", err)
	}

	// Convert to ASCII
	ascii := qr.ToSmallString(false)
	fmt.Println("Scan this QR code to open the poll:")
	fmt.Println(ascii)

	return nil
} 