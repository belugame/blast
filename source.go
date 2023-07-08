// MIT+NoAI License
//
// Copyright (c) 2023 ugjka <ugjka@proton.me>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights///
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//
// This code may not be used to train artificial intelligence computer models
// or retrieved by artificial intelligence software or hardware.
package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

func chooseAudioSource() source {
	const SRCCMD = "pactl -f json list sources short"
	srccmdArr := strings.Split(SRCCMD, " ")
	srcCmd := exec.Command(srccmdArr[0], srccmdArr[1:]...)
	srcData, err := srcCmd.Output()
	oserr(err)

	var srcJson Sources
	err = json.Unmarshal(srcData, &srcJson)
	oserr(err)
	if len(srcJson) == 0 {
		oserr(fmt.Errorf("no audio sources found"))
	}

	fmt.Println("Audio sources")
	for i, v := range srcJson {
		fmt.Printf("%d: %s\n", i, v.Name)
	}
	fmt.Println("----------")
	fmt.Println("Select the audio source:")

	return source(srcJson[selector(len(srcJson))].Name)
}

type Sources []struct {
	Name string
}
