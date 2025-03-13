<div align="center">
    <h1>Quickform</h1>
</div>

<div align="center"> 
QuickForm is a command-line tool for creating beautiful Typeform forms quickly and easily.
</div>

## Installation

```bash
go install github.com/s3bw/quickform@latest
```

```bash
export QUICKFORM_ACCESS_TOKEN='your_typeform_access_token'
```

## Usage

### Creating a Poll

Create a new poll with the default theme:

```bash
$ quickform poll
```

Create a poll with a custom name:

```bash
$ quickform poll --name "Team Meeting Poll"
```

Create a poll with a specific theme:

```bash
$ quickform poll --theme mystery
```

Create a poll with both custom name and theme:

```bash
$ quickform poll --name "Product Feedback" --theme spaceboy
```

### Viewing Available Themes

List all available themes with their descriptions:

```bash
$ quickform themes
```

## Available Themes

- `default` - Clean and professional with system font
- `orbital` - Modern with Montserrat font and light purple background
- `fractal` - Bold with Lato font and dark blue background
- `art-splash` - Creative with Sniglet font and warm colors
- `pin-up` - Retro with Dancing Script font and vintage colors
- `paper-invite` - Elegant with Georgia font and paper texture
- `mystery` - Dark with Acme font and gold accents
- `school-bell` - Playful with Handlee font and green background
- `spaceboy` - Sci-fi with Arvo font and space background
- `desk-space` - Professional with Karla font and warm office colors

## Example

Here's an example of creating a poll:

```bash
$ ./quickform poll --theme mystery --name "Team Lunch Vote"

Enter Question:
$ What should we order for team lunch?

Enter choices (press Enter to finish):
$ Pizza
$ Sushi
$ Tacos
$ [Enter]

Poll created successfully!
Share this URL: https://form.typeform.com/to/xxxxx
View responses: https://admin.typeform.com/forms/xxxxx/responses

Scan this QR code to open the poll:
[QR code ASCII art]
```

## Requirements

- Go 1.16 or later
- Typeform account and access token