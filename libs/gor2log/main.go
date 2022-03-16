package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func gor2log() {
	if len(os.Args) == 1 {
		log.Fatalln("dest file required")
	}

	destFile := os.Args[1]
	fin, err := os.OpenFile(destFile, os.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}
	defer fin.Close()

	fout, err := os.Create("result.log")
	if err != nil {
		panic(err)
	}
	defer fout.Close()

	sc := bufio.NewScanner(fin)
	/*default split the file use '\n'*/

	var oc, nc int
	kv := make(map[string]string)
	for sc.Scan() {
		oc++
		if strings.Contains(sc.Text(), " HTTP") {
			items := strings.Split(sc.Text(), " ")
			kv["method"] = items[0]
			kv["path"] = items[1]
			kv["proto"] = items[2]
		}
		if strings.Contains(sc.Text(), ": ") {
			items := strings.Split(sc.Text(), ": ")
			kv[items[0]] = items[1]
		}

		if sc.Text() == "🐵🙈🙉" {
			nc++
			fmt.Println(oc, nc)
			data, _ := json.Marshal(kv)
			w := bufio.NewWriter(fout)
			_, err = w.Write(data)
			if err != nil {
				log.Fatalln(err)
			}
			_, err = w.WriteString("\r\n")
			if err != nil {
				log.Fatalln(err)
			}
			err = w.Flush()
			if err != nil {
				log.Fatalln(err)
			}
			kv = make(map[string]string)
		}
	}
	if err := sc.Err(); err != nil {
		fmt.Println("An error has hippened")
	}
}
