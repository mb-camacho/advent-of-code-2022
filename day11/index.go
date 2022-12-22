package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
}

func getFileContents(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func buildMonkeys(data string) []*monkey {
	monkeys := make([]*monkey, 0)
	monkeyArray := strings.Split(data, "\n\n")
	for _, mon := range monkeyArray {
		monArray := strings.Split(mon, "\n")
		monkeyObject := monkey{}
		for _, monData := range monArray {
			cleanedData := strings.TrimSpace(monData)
			if strings.HasPrefix(cleanedData, "Monkey ") {
				continue
			}
			if strings.HasPrefix(cleanedData, "Starting items: ") {
				fields := strings.Split(cleanedData, ": ")
				if len(fields) != 2 {
					panic("expecting 2 parts. string:" + cleanedData)
				}
				for _, worryItem := range strings.Split(fields[1], ", ") {
					worryItemInt, err := strconv.Atoi(worryItem)
					if err != nil {
						panic(err)
					}
					worryItemObject := item{worryNumber: worryItemInt, isInspected: false}
					monkeyObject.items = append(monkeyObject.items, &worryItemObject)
				}
				continue
			}
			if strings.HasPrefix(cleanedData, "Operation: new = old ") {
				operation := strings.Replace(cleanedData, "Operation: new = old ", "", -1)
				operations := strings.Split(operation, " ")
				if len(operations) != 2 {
					panic("expected operations len is 2. string:" + cleanedData)
				}
				monkeyObject.operationSign = operations[0]
				if operations[1] == "old" {
					monkeyObject.operationValue = -1
					continue
				}
				operationValue, err := strconv.Atoi(operations[1])
				if err != nil {
					panic(err)
				}
				monkeyObject.operationValue = operationValue
				continue
			}
			if strings.HasPrefix(cleanedData, "Test: divisible by ") {
				divisibleBy, err := strconv.Atoi((strings.Replace(cleanedData, "Test: divisible by ", "", -1)))
				if err != nil {
					panic(err)
				}
				monkeyObject.divisibleBy = divisibleBy
				continue
			}
			if strings.HasPrefix(cleanedData, "If true: throw to monkey ") {
				throwTrue, err := strconv.Atoi((strings.Replace(cleanedData, "If true: throw to monkey ", "", -1)))
				if err != nil {
					panic(err)
				}
				monkeyObject.throwTrue = throwTrue
				continue
			}
			if strings.HasPrefix(cleanedData, "If false: throw to monkey ") {
				throwFalse, err := strconv.Atoi((strings.Replace(cleanedData, "If false: throw to monkey ", "", -1)))
				if err != nil {
					panic(err)
				}
				monkeyObject.throwFalse = throwFalse
				continue
			}
			panic("unknown line. string:" + cleanedData)
		}
		monkeys = append(monkeys, &monkeyObject)
	}
	return monkeys
}
func levelOfMonkeyBusiness(contents string, rounds int) int {
	top1 := 0
	top2 := 0

	expected := 2713310158
	actual := top1 * top2
	wl := 1

	for expected != actual {
		// fmt.Println(wl)
		monkeys := buildMonkeys(contents)
		super := 0
		for _, mo := range monkeys {
			// for _, it := range mo.items {
			if super == 0 {
				super = mo.divisibleBy
				continue
			}
			super *= mo.divisibleBy
			// }
		}

		for i := 0; i < rounds; i++ {
			for _, mo := range monkeys {
				mo.ProcessItem(monkeys, super)
			}
		}

		for _, monkey := range monkeys {
			if monkey.inspectdItems > top1 {
				top2 = top1
				top1 = monkey.inspectdItems
			} else if monkey.inspectdItems > top2 {
				top2 = monkey.inspectdItems
			}
		}
		wl++
		break
	}
	return top1 * top2
}

func (mo *monkey) ProcessItem(monkeys []*monkey, wl int) {

	// wl = 0
	// for _, itemx := range mo.items {
	// 	if wl == 0 {
	// 		wl = itemx.worryNumber
	// 		continue
	// 	}
	// 	wl *= itemx.worryNumber
	// }

	for len(mo.items) > 0 {
		mo.inspectdItems++
		worryItem := mo.items[0]
		worryNumber := worryItem.worryNumber
		calc := 0
		moNumber := 0
		operationValue := worryNumber
		if mo.operationValue != -1 {
			operationValue = mo.operationValue
		}
		if mo.operationSign == "*" {
			calc = worryNumber * operationValue
		} else if mo.operationSign == "+" {
			calc = worryNumber + operationValue
		} else {
			panic("unknown operation sign:" + mo.operationSign)
		}
		// if !worryItem.isInspected {
		// 	// calc /= 3
		// 	calc = calc % 96577
		// } else {
		// }

		calc = calc % wl
		if calc%mo.divisibleBy == 0 {
			moNumber = mo.throwTrue
		} else {
			moNumber = mo.throwFalse
		}
		worryItem.worryNumber = calc
		worryItem.isInspected = true
		monkeys[moNumber].items = append(monkeys[moNumber].items, worryItem)
		if len(mo.items) == 1 {
			mo.items = make([]*item, 0)
		} else {
			mo.items = mo.items[1:]
		}
	}
}

type monkey struct {
	number         int
	inspectdItems  int
	items          []*item
	operationSign  string
	operationValue int
	divisibleBy    int
	throwTrue      int
	throwFalse     int
}

type item struct {
	worryNumber int
	isInspected bool
}
