package hello

const englishPrefix = "Hello "
const spanishPrefix = "Hola "

func Hello(name, lang string) string {
	var prefix string

	if name == "" {
		name = "World"
	}

	switch lang {
	case "spanish":
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}
	return prefix + name
}
