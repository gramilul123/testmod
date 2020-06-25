package testmod

import (
    "fmt"
    "errors"
)

func Hello(name, lang string) (string, error) {

    switch lang {
    case "en":
        return fmt.Sprintf("Hello, %s!", name), nil
    case "pt":
        return fmt.Sprintf("Oi, %s!", name), nil
    case "es":
        return fmt.Sprintf("¡Hola, %s!", name), nil
    case "fr":
        return fmt.Sprintf("Bonjour, %s!", name), nil
    case "ru":
        return fmt.Sprintf("Привет, %s!", name), nil
    default:
        return "", errors.New("Unknown language")
    }
}
