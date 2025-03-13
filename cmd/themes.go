package cmd

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"
	"quickform/pkg/typeform"
)

var themesCmd = &cobra.Command{
	Use:   "themes",
	Short: "List available themes",
	Long:  `List all available themes with their descriptions and styling information.`,
	RunE:  runThemes,
}

func init() {
	rootCmd.AddCommand(themesCmd)
}

func runThemes(cmd *cobra.Command, args []string) error {
	fmt.Println("Available Themes:")
	fmt.Println("---------------")

	// Get theme names and sort them for consistent display
	themeNames := make([]string, 0, len(typeform.Themes))
	for name := range typeform.Themes {
		themeNames = append(themeNames, name)
	}
	sort.Strings(themeNames)

	// Theme descriptions
	descriptions := map[string]string{
		"default":     "Clean and professional with system font",
		"orbital":     "Modern with Montserrat font and light purple background",
		"fractal":     "Bold with Lato font and dark blue background",
		"art-splash":  "Creative with Sniglet font and warm colors",
		"pin-up":      "Retro with Dancing Script font and vintage colors",
		"paper-invite": "Elegant with Georgia font and paper texture",
		"mystery":     "Dark with Acme font and gold accents",
		"school-bell": "Playful with Handlee font and green background",
		"spaceboy":    "Sci-fi with Arvo font and space background",
		"desk-space":  "Professional with Karla font and warm office colors",
	}

	// Print themes in sorted order
	for _, name := range themeNames {
		desc := descriptions[name]
		fmt.Printf("%-12s - %s\n", name, desc)
	}

	fmt.Println("\nUse --theme <name> with the poll command to apply a theme.")
	return nil
} 