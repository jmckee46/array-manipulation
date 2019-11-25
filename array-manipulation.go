package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func arrayManipulation(n int32, queries [][]int32) int64 {
	totalArr := make([]int64, n+1)
	var total int64
	var tempTotal int64

	for _, arr := range queries {
		totalArr[arr[0]-1] += int64(arr[2])
		totalArr[arr[1]] += int64(-arr[2])
	}

	for _, value := range totalArr {
		tempTotal += value
		if tempTotal > total {
			total = tempTotal
		}
	}

	return total
}

func main() {
	file, err := os.Open("test-case-4")
	checkError(err)

	reader := bufio.NewReaderSize(file, 1024*1024)

	stdout, err := os.Create("OUTPUT-test-case-4")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	var queries [][]int32
	for i := 0; i < int(m); i++ {
		queriesRowTemp := strings.Split(readLine(reader), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != int(3) {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	result := arrayManipulation(n, queries)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// 100 0   -100 0 0
// 100 100 -100 0 0
// 100 100 0    0 -100

// 100 100 0   0   0
// 100 200 100 100 100
// 100 200 200 200 100
