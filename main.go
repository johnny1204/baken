package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/jszwec/csvutil"
)

// type Race struct {
// 	Date    string `csv:"日付(yyyy.mm.dd)"`
// 	Place   string `csv:"場所"`
// 	RaceNum int    `csv:"Ｒ"`
// 	Finish  string `csv:"着順"`
// 	Order   int    `csv:"馬番"`
// 	Id      string `csv:"レースID(新/馬番無)"`
// }

// var races []Race
type Race struct {
	Place    string `csv:"場所"`
	Distance string `csv:"距離"`
	Klass    string `csv:"クラス名"`
	Finish   string `csv:"着順"`
	Order    int    `csv:"馬番"`
	Dividend string `csv:"馬連"`
	Id       string `csv:"レースID(新/馬番無)"`
}

var races []Race

type Umaren struct {
	One   int
	Two   int
	Max   int
	Total int
	Avg   float64
	Idx   []int
}

type Compi struct {
	Id  string
	Val []string
}

func main() {
	file, _ := os.Open("csv/compi.csv")
	r := csv.NewReader(file)
	r.FieldsPerRecord = -1

	// var compis = make(map[string]string, 0)
	var targets = make(map[string][]int, 0)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}

		horse := strings.Join(record[1:], ",")

		horses := strings.Split(horse, ",")
		sortHorse := make([]int, len(horses))
		for idx, s := range horses {
			j, _ := strconv.Atoi(s)
			sortHorse[idx] = j
		}
		sort.Sort(sort.Reverse(sort.IntSlice(sortHorse)))

		targets[record[0]] = sortHorse
	}
	fmt.Println(targets)
	// r := target_races()
	// writeCsv(r, "target20220519")
	// c := race_details()
	// writeResultCsv(c, "umaren_20220519")
	// file, _ := os.Open("umaren_20220519.csv")
	// r := csv.NewReader(file)
	// r.FieldsPerRecord = -1

	// // ziku := map[int]int{1: 0, 2: 0, 3: 0}

	// i := 0
	// total := 0
	// haitou := 0
	// manbakenCnt := 0
	// max := 0
	// min := 0
	// for {
	// 	total++
	// 	record, err := r.Read()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// 	i1, _ := strconv.Atoi(record[5])
	// 	i2, _ := strconv.Atoi(record[7])
	// 	h, _ := strconv.Atoi(record[8])
	// 	c2, _ := strconv.Atoi(record[6])
	// 	c1, _ := strconv.Atoi(record[4])
	// 	if (is_ziku(i1) && is_aite(c2)) || (is_ziku(i2) && is_aite(c1)) {
	// 		// if (is_ziku2(i1) && is_aite2(i2)) || (is_ziku2(i2) && is_aite2(i1)) {
	// 		i++

	// 		if h >= 10000 {
	// 			manbakenCnt++
	// 		}

	// 		if h > max {
	// 			max = h
	// 		}

	// 		if h < min || min == 0 {
	// 			min = h
	// 		}

	// 		// if i1 == 1 || i2 == 1 {
	// 		// 	ziku[1]++
	// 		// } else if i1 == 2 || i2 == 2 {
	// 		// 	ziku[2]++
	// 		// } else {
	// 		// 	ziku[3]++
	// 		// }
	// 		haitou = haitou + h
	// 	}
	// }

	// fmt.Println("確率：", float64(i)/float64(total))
	// fmt.Print("配当：¥")
	// convert(haitou)
	// fmt.Print("掛金：¥")
	// convert(total * 1000)
	// fmt.Println("対象レース：", i)
	// fmt.Println("万馬券回数：", manbakenCnt)
	// fmt.Print("最高配当：")
	// convert(max)
	// fmt.Print("最低配当：")
	// convert(min)
	// fmt.Println(ziku)
}

func is_ziku(order int) bool {
	return order == 1 || order == 2 || order == 3
}

func is_aite(compi int) bool {
	return compi <= 50 && compi >= 46
}

func is_ziku2(order int) bool {
	return order == 1
}

func is_aite2(order int) bool {
	return order == 2 || order == 3
}

func target_races() [][]string {
	file, _ := os.Open("csv/compi.csv")
	r := csv.NewReader(file)
	r.FieldsPerRecord = -1

	// var compis = make(map[string]string, 0)
	var targets = make([][]string, 0)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}

		if len(record[1:]) >= 13 {
			horse := strings.Join(record[1:], ",")

			horses := strings.Split(horse, ",")
			sortHorse := make([]int, len(horses))
			// copyiod := make([]int, len(horses))

			ana := 0
			for idx, s := range horses {
				j, _ := strconv.Atoi(s)
				if j >= 46 && j <= 59 {
					ana++
				}

				if j != 0 {
					sortHorse[idx] = j
				}
			}

			// copy(copyiod, sortHorse)
			sort.Sort(sort.Reverse(sort.IntSlice(sortHorse)))
			oneTwoThree := (sortHorse[0] + sortHorse[1] + sortHorse[2])
			if len(sortHorse) >= 13 && sortHorse[0] <= 77 && (oneTwoThree >= 210 && oneTwoThree <= 219) && ana >= 7 {
				r := []string{record[0]}
				targets = append(targets, r)
			}
		}
		// compis[record[0]] = strings.Join(record[1:], ",")
	}

	return targets
}

func race_details() [][]string {
	file, _ := os.Open("targets20220519.csv")
	r := csv.NewReader(file)
	r.FieldsPerRecord = -1

	b, _ := ioutil.ReadFile("csv/umaren20220519.csv")
	csvutil.Unmarshal(b, &races)

	var compis = make([][]string, 0)
	header := []string{"場所", "コース", "距離", "クラス名", "1着指数", "1着指数順位", "2着指数", "2着指数順位", "馬連"}
	compis = append(compis, header)

	var id string
	var one string
	var oneOrd string
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}

		for _, race := range races {
			if rec[0] == race.Id {
				result := make([]string, 9)
				result[0] = race.Place

				course := race.Distance[:3]
				if course == "ダ" {
					result[1] = "ダート"
				} else {
					result[1] = course
				}

				result[2] = race.Distance[3:]

				if race.Klass == "500万" {
					result[3] = "1勝"
				} else if race.Klass == "1000万" {
					result[3] = "2勝"
				} else {
					result[3] = race.Klass
				}

				result[8] = race.Dividend
				horse := strings.Join(rec[1:], ",")

				horses := strings.Split(horse, ",")
				sortHorse := make([]int, len(horses))
				for idx, s := range horses {
					j, _ := strconv.Atoi(s)

					if j != 0 {
						sortHorse[idx] = j
					}
				}

				sort.Sort(sort.Reverse(sort.IntSlice(sortHorse)))

				if id != race.Id {
					one = horses[race.Order-1]
					for k, v := range sortHorse {
						i, _ := strconv.Atoi(horses[race.Order-1])
						if v == i {
							oneOrd = strconv.Itoa(k + 1)
							break
						}
					}
					id = race.Id
				} else {
					for k, v := range sortHorse {
						i, _ := strconv.Atoi(horses[race.Order-1])
						if v == i {
							result[6] = horses[race.Order-1]
							result[7] = strconv.Itoa(k + 1)
							break
						}
					}
					result[4] = one
					result[5] = oneOrd
					compis = append(compis, result)
				}

				continue
			}
		}
	}

	return compis
}

func zi_to_compi() {
	file, _ := os.Open("csv/manbaken_zi.csv")
	r := csv.NewReader(file)
	r.FieldsPerRecord = -1

	var ids = make([]string, 0)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		ids = append(ids, record[0])
	}

	file2, _ := os.Open("csv/compi.csv")
	r2 := csv.NewReader(file2)
	r2.FieldsPerRecord = -1

	var compis = make([]Compi, 0)
	for {
		record, err := r2.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		for _, v := range ids {
			if record[0] == v {
				var compi Compi
				compi.Id = record[0]
				compi.Val = record[1:]
				compis = append(compis, compi)
			}
		}
		// ids = append(ids, record[0])
	}
	fmt.Println(len(compis))

	wfile, _ := os.OpenFile("target_ids.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer wfile.Close()
	// w := csv.NewWriter(os.Stdout)
	w := csv.NewWriter(wfile)
	// r := make([]string, 0, 1)
	for _, value := range compis {
		r := make([]string, 0, 1+len(value.Val))
		r = append(r, value.Id)
		for i := 0; i < len(value.Val); i++ {
			r = append(r, value.Val[i])
		}
		if err := w.Write(r); err != nil {
			fmt.Println("error writing record to csv:", err)
		}
	}
	defer w.Flush()
}

func writeCsv(records [][]string, file_name string) {
	file, _ := os.OpenFile(file_name+".csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	// w := csv.NewWriter(os.Stdout)
	w := csv.NewWriter(file)
	// r := make([]string, 0, 1)
	for _, value := range records {
		for i := 0; i < len(value); i++ {
			if err := w.Write(value); err != nil {
				fmt.Println("error writing record to csv:", err)
			}
		}
	}
	defer w.Flush()
}

func writeResultCsv(records [][]string, file_name string) {
	file, _ := os.OpenFile(file_name+".csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	// w := csv.NewWriter(os.Stdout)
	w := csv.NewWriter(file)
	// r := make([]string, 0, 1)
	for _, value := range records {
		if err := w.Write(value); err != nil {
			fmt.Println("error writing record to csv:", err)
		}
	}
	defer w.Flush()
}

func convert(integer int) {
	arr := strings.Split(fmt.Sprintf("%d", integer), "")
	cnt := len(arr) - 1
	res := ""
	i2 := 0
	for i := cnt; i >= 0; i-- {
		if i2 > 2 && i2%3 == 0 {
			res = fmt.Sprintf(",%s", res)
		}
		res = fmt.Sprintf("%s%s", arr[i], res)
		i2++
	}
	fmt.Println(res)
}
