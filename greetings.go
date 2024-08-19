package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

//hello saludo para una persona especifica

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("nombre vacio ")
	}
	message := fmt.Sprintf(saludos(), name)

	return message, nil
}

func saludos() string {
	formats := []string{
		"hola, %v Bienvenido",
		"que bueno verte, %v",
		"saludo %v encantado de conocerte",
	}

	return formats[rand.Intn(len(formats))]
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}
