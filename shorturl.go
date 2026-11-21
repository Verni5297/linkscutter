package main

import "math/rand"

const symbols = "qwertyuiopasdfghjklzxcvbnm1234567890"

func generateUrl(length int) string {

	var newurl = make([]byte, length)

	for i := range length {

		newurl[i] = symbols[rand.Intn(len(symbols))]

	}

	return string(newurl)

}
