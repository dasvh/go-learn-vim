package components

// Title represents the structure for a title component with a main title and a subtitle
type Title struct {
	Main     string
	Subtitle string
}

// NewTitle creates a new instance of Title with the specified main title and subtitle
func NewTitle(main, sub string) *Title {
	return &Title{
		Main:     main,
		Subtitle: sub,
	}
}
