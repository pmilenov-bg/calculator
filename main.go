package main

import (
	mymath "calculator/src/helpers"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	// log.SetLevel(logrus.InfoLevel)
	log.SetLevel(logrus.DebugLevel)

	a := 6
	b := 3

	addResult := mymath.Add(a, b)
	subtractResult := mymath.Subtract(a, b)
	multiplyResult := mymath.Multiply(a, b)
	divideResult := mymath.Divide(a, b)

	log.Infof("Addition Result: %d", addResult)
	log.Infof("Subtraction Result: %d", subtractResult)
	log.Infof("Multiplication Result: %d", multiplyResult)
	log.Infof("Division Result: %d", divideResult)
}
