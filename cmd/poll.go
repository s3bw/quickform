package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/skip2/go-qrcode"
	"github.com/spf13/cobra"

	"quickform/pkg/typeform"
)

var (
	theme             string
	name              string
	multipleSelection bool

	blue  = color.New(color.FgBlue).SprintFunc()
	cyan  = color.New(color.FgCyan).SprintFunc()
	green = color.New(color.FgGreen).SprintFunc()
	red   = color.New(color.FgRed).SprintFunc()
)

var pollCmd = &cobra.Command{
	Use:   "poll",
	Short: "Create a poll with multiple choices",
	Long: `Create a poll with multiple choices and share it via URL and QR code.

Run 'quickform themes' to see detailed theme descriptions.`,
	RunE: runPoll,
}

func init() {
	rootCmd.AddCommand(pollCmd)
	pollCmd.Flags().StringVar(&theme, "theme", "", "Theme for the poll (e.g., 'default', 'orbital', 'fractal')")
	pollCmd.Flags().StringVar(&name, "name", "", "Name for the poll (default: 'Quick Form')")
	pollCmd.Flags().BoolVar(&multipleSelection, "multiple", false, "Allow multiple selections (default: false)")
}

func runPoll(cmd *cobra.Command, args []string) error {
	accessToken := os.Getenv("QUICKFORM_ACCESS_TOKEN")
	if accessToken == "" {
		return fmt.Errorf("%s environment variable is not set", red("QUICKFORM_ACCESS_TOKEN"))
	}

	client := typeform.NewClient(accessToken)

	// Get the question
	fmt.Printf("%s\n$ ", cyan("Enter Question:"))
	reader := bufio.NewReader(os.Stdin)
	question, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("%s: %w", red("error reading question"), err)
	}
	question = strings.TrimSpace(question)

	// Get the choices
	fmt.Printf("\n%s\n", cyan("Enter choices (press Enter to finish):"))
	var choices []typeform.Choice
	for {
		fmt.Print("$ ")
		choice, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("%s: %w", red("error reading choice"), err)
		}
		choice = strings.TrimSpace(choice)

		if choice == "" {
			break
		}
		choices = append(choices, typeform.Choice{Label: choice})
	}

	if len(choices) < 2 {
		return fmt.Errorf("%s", red("at least 2 choices are required"))
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
					AllowMultipleSelection: multipleSelection,
					Choices:                choices,
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
		return fmt.Errorf("%s: %w", red("error creating form"), err)
	}

	fmt.Printf("\n%s\n", green("Poll created successfully!"))
	fmt.Printf("Share this URL: %s\n", blue(urls.Display))
	fmt.Printf("View responses: %s\n\n", blue(urls.Responses))

	// Generate and display QR code
	qr, err := qrcode.New(urls.Display, qrcode.Medium)
	if err != nil {
		return fmt.Errorf("%s: %w", red("error generating QR code"), err)
	}

	// Convert to ASCII
	ascii := qr.ToSmallString(false)
	fmt.Println("Scan this QR code to open the poll:")
	fmt.Println(ascii)

	return nil
}
