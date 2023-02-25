package core

import (
	"fmt"
	"os"
	"testing"
)

/*
* Reference: https://geektutu.com/post/quick-go-test.html
 */
func setup() {
	fmt.Println("Before all tests")
}

func teardown() {
	fmt.Println("After all tests")

}

func TestUrlEncode(t *testing.T) {
	cases := []struct {
		Name            string
		Input, Expected string
	}{
		{"1", "=", "%3D"},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans := UrlEncode(c.Input); ans != c.Expected {
				t.Fatalf("%s expected %s, but %s got", c.Input, c.Expected, ans)
			}
		})
	}
}

func TestBase64Encode(t *testing.T) {
	cases := []struct {
		Name            string
		Input, Expected string
	}{
		{"1", "%", "JQ=="},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans := B64Encode(c.Input); ans != c.Expected {
				t.Fatalf("%s expected %s, but %s got", c.Input, c.Expected, ans)
			}
		})
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
