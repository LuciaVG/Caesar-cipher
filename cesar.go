package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
)

type Inputs struct {
	Inputs []Input `json:"inputs"`
}

type Input struct {
	Type     string `json:"type"`
	In_file  string `json:"in_file"`
	Out_file string `json:"out_file"`
}

func decrypt(key float64, msg_file string) (float64, string) {
	sft := 0.0
	sfd := 0.0
	zero := 33.0
	md := 94.0

	file, err := os.Open("file2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	text, err := ioutil.ReadAll(file)
	//fmt.Print(text)
	//fmt.Print(string(text))

	text2 := make([]byte, len(text))

	for i, chr := range text {
		switch {
		case chr > 32:
			if chr == byte(zero+sft) {
				text2[i] = 32
			} else {
				sft = math.Mod(key, md)

				if sft == 0 {
					key++
					sft = math.Mod(key, md)
				}
				sfd = float64(chr) - (sft + 0)
				if sfd > (zero - 1) {
					text2[i] = byte(sfd)
				} else {
					text2[i] = byte(md + sfd)
				}
				key += math.Floor(key / sft)
			}

		case chr == 32:
			text2[i] = byte(zero + sft)
		default:
			text2[i] = chr
		}
	}

	//fmt.Print(text2)
	return key, string(text2)
}

func encrypt(key float64, msg_file string) (float64, string) {
	sft := 0.0
	zero := 33.0
	md := 94.0

	file, err := os.Open(msg_file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	text, err := ioutil.ReadAll(file)
	//fmt.Print(text)
	//fmt.Print(string(text))

	text2 := make([]byte, len(text))

	for i, chr := range text {
		//fmt.Println(sft)
		switch {
		case chr > 32:
			sft = math.Mod(key, md)

			if sft == 0 {
				key++
				sft = math.Mod(key, md)
			}

			text2[i] = byte(math.Mod((float64(chr)-zero)+sft, md) + zero)
			key += math.Floor(key / sft)
		case chr == 32:
			text2[i] = byte(zero + sft)
		default:
			text2[i] = chr
		}
		//fmt.Println(chr)

	}
	return key, string(text2)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Arguments missing")
		return
	}

	jsonFile, err := os.Open("inputs.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var inputs Inputs
	json.Unmarshal(byteValue, &inputs)

	chn := make(chan string)

	for _, item := range inputs.Inputs {
		go func(t, in, out string, ch chan string) {
			public_key := 10.0
			private_key_a := 14.0
			private_key_b := 16.0

			key := public_key * private_key_a * private_key_b
			c := ""
			switch t {
			case "en":
				_, c = encrypt(key, in)

			case "de":
				_, c = decrypt(key, in)
			default:
				fmt.Println("Need a valid argument")
			}

			file, err := os.Create(out)
			if err != nil {
				log.Fatal("Cannot create file", err)
			}
			defer file.Close()

			fmt.Fprintf(file, c)
			ch <- c

		}(item.Type, item.In_file, item.Out_file, chn)

	}
	for i := 0; i < len(inputs.Inputs); i++ {
		fmt.Println(<-chn)
		fmt.Println("-----------------------------------------------------------------------------")
	}
}
