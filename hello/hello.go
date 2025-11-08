package hello

const (
	spanish = "Spanish"
	french  = "French"

	englishPrefix = "Hello, "
	spanishPrefix = "Hola, "
	frenchPrefix  = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name + "!"
}

func greetingPrefix(language string) string {
	switch language {
	case spanish:
		return spanishPrefix
	case french:
		return frenchPrefix
	default:
		return englishPrefix
	}
}
