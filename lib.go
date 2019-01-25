package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	unicorn "github.com/arussellsaw/unicorn-go"
)

// CheckBuildStatus checks the status of the pipeline build by looking for
// the word "passing" in the build status svg. This is the easiest way that
// doesn't involve API calls.
func CheckBuildStatus() (bool, error) {
	pipelineBuildURL := "https://travis.schibsted.io/api/finn/platform-pipeline.svg?token=123456789&branch=master"
	resp, err := http.Get(pipelineBuildURL)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	// We need to convert html to string since it's a byte array
	passing := strings.Contains(string(html), "passing")
	return passing, err
}

// Pulsate sends a full, off - on - off, slow brightness pulse to the lamp
func Pulsate(c unicorn.Client, brightness int) {
	// one full pulse is two for loops:
	for i := 0; i <= brightness; i++ {
		err := c.SetBrightness(uint(i))
		if err != nil {
			fmt.Println(err)
			return
		}
		err = c.Show()
		if err != nil {
			fmt.Println(err)
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
	for i := brightness; i >= 0; i-- {
		err := c.SetBrightness(uint(i))
		if err != nil {
			fmt.Println(err)
			return
		}
		err = c.Show()
		if err != nil {
			fmt.Println(err)
			return
		}
		time.Sleep(100 * time.Millisecond)
	}

}
