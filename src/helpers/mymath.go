package mymath

import (
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true,
	})

	log.SetLevel(logrus.DebugLevel)

}

func Add(a, b int) int {
	result := a + b
	log.WithFields(logrus.Fields{
		"package": "mymath",
	}).Debugf("You are in the addition stage: %d", result)
	return result
}

func Subtract(a, b int) int {
	result := a - b
	log.Debugf("You are in the subtraction stage: %d", result)
	return result
}

func Multiply(a, b int) int {
	result := a * b
	log.Debugf("You are in the multiplication stage: %d", result)
	return result
}

func Divide(a, b int) int {
	result := a / b
	log.Debugf("You are in the division stage: %d", result)
	return result
}
