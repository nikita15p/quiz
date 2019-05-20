package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func readCsvFile(file string) (q []string, ans []string) {

	r, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Errorf("csv file read error due to", err)
	}

	str := string(r)
	arr := strings.Fields(str)
	for _, v := range arr {
		qa := strings.Split(v, ",")
		q = append(q, qa[0])
		ans = append(ans, qa[1])
	}

	return q, ans
}

func main() {

	flag.Parse()

	f := flag.Lookup("file").Value.String()
	q, expected := readCsvFile(f)

	var observed string

	correct := 0
	total := len(q)
	incorrect := total
	fmt.Printf("Press Enter to start the test")
	reader := bufio.NewReader(os.Stdin)
	key, _, _ := reader.ReadRune()
	if key == rune(10) {
		timer := time.NewTimer(5 * time.Second)
		go func(){
			//If Enter Key is pressed, start the timer
			for i, ques := range q{
				fmt.Println("What is", ques, "=?")
				fmt.Scanf("%s", &observed)
				if observed == expected[i] {
					correct++
					incorrect--
				}

			}
		}(	)


			<-timer.C
		}




	fmt.Printf("Total Questions: %d Correct: %d Incorrect: %d", total, correct, incorrect)

}

func init() {
	f := flag.Lookup("file")
	if f == nil {
		flag.String("file", "problems.csv", "Set the csv file name if not default")
	}
}
