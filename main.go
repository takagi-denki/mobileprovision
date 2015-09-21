package main

import (
	"./provisioning"
	"bufio"
	"fmt"
	"flag"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	input := flag.String("i", "", "input file")
	output := flag.String("f", "content", "output format")
	flag.Parse()
	file, _ := os.Open(*input)
	reader := bufio.NewReader(file)
	buffer, _ := ioutil.ReadAll(reader)
	info := provisioning.NewContentInfo(buffer)
	
	if *output == "content" {
		fmt.Println(string(info.Content.Sequence.EncapContentInfo.Content.Content))
	} else {
		data := info.GetContent()
		certificates, _ := data.GetDeveloperCertificates()
		// fmt.Print(data)
		fmt.Print(certificates)
		fmt.Print(data.IsExpired(time.Now()))
	}
}
	

