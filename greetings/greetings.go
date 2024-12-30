package greetings

import ("fmt"; "errors")

func Hello(name string) (string, error) {

    if name == "" {
        return "", errors.New("No name provided")
    }

    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message, nil
}
