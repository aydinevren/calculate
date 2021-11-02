package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type employee struct {
	lastName string
	name     string
	hours    int
}

// input fonksiyonu data.csv dosyasindan verileri okur
func input() []employee {
	file, err := os.Open("data.csv")
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	var dataList []employee
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		record := strings.Split(line, ",")
		hour, _ := strconv.Atoi(record[2])
		dataList = append(dataList, employee{lastName: record[0], name: record[1], hours: hour})
	}
	return dataList
}

// calculate fonksiyonu verileri hesaplar
func calculate(data []employee) map[employee]int {
	dataMap := make(map[employee]int)
	var r employee
	for _, emp := range data {
		r = employee{lastName: emp.lastName, name: emp.name}
		dataMap[r] += emp.hours
	}

	return dataMap
}

// output fonksiyonu hesaplanan verileri result.csv dosyasina yazdirir
func output(record map[employee]int) {
	fileWrite, err := os.Create("results.csv")
	defer fileWrite.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	w := bufio.NewWriter(fileWrite)
	defer w.Flush()

	var data string
	for e, h := range record {
		data = fmt.Sprintf("%s,%s,%d\n", e.lastName, e.name, h)
		_, err = w.WriteString(data)
		if err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
}

func main() {
	record := input()
	result := calculate(record)
	output(result)
}
