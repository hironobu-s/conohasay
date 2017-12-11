package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

var msg = Message{
	"Launch SSD cloud in just 25 seconds.",
	"ConoHa provides simple and powerful cloud hosting using SSD and OpenStack.",
	"It's easy to create a single server to a sophisticated development enviroment.",
}

func testMessage() Message {
	file, err := ioutil.ReadFile("main.go")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(file), "\n")
}

func TestBalloonText(t *testing.T) {
	wrapcolumn := 80
	balloon := balloonText(msg, wrapcolumn, "left")
	l := len(balloon)
	for i, line := range balloon {
		if i == 0 || i > l-5 {
			continue
		}

		if len(line) != wrapcolumn+4 {
			t.Errorf("Invalid the line width in a balloon.[%d]", len(line))
		}
	}
}

func TestFormatV(t *testing.T) {
	wrapcolumn := 80

	msg = testMessage()
	balloon := balloonText(msg, wrapcolumn, "left")

	for _, name := range ListCows() {
		for _, size := range []string{"s", "m", "l"} {
			cow, err := NewCow(name, size)
			if err != nil {
				t.Error(err)
			}

			output := formatV(msg, cow, wrapcolumn)
			c := strings.Split(output, "\n")
			if len(c) != len(balloon)+cow.Height() {
				t.Errorf("Invalid output? (cow = %s, size = %s, height = %d, expect = %d)",
					name, size, len(c), len(balloon)+cow.Height())
			}
		}
	}
}
