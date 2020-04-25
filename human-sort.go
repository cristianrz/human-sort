package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var items []string

	if err := getItems(&items); err != nil {
		panic(err)
	}

	//fmt.Println(items)

	items, err := sort(items)
	if err != nil {
		panic(err)
	}

	fmt.Println("sorted:", items)
}

func sort(items []string) ([]string, error) {
	n := float64(len(items))

	if n < 2 {
		return items, nil
	}

	middle := int(math.Round(n / 2))

	left := items[0:middle]
	left, err := sort(left)
	if err != nil {
		return items, err
	}

	right := items[middle:]
	right, err = sort(right)
	if err != nil {
		return items, err
	}

	l := len(left)
	r := len(right)

	var sorted []string

	for l+r != 0 {
		//fmt.Println("left:", left)
		//fmt.Println("right:", right)
		if l == 0 {
			sorted = append(sorted, right[0])
			right = right[1:]
		} else if r == 0 {
			sorted = append(sorted, left[0])
			left = left[1:]
		} else {
			fmt.Println("[+] What is better")
			fmt.Println("\t1)", left[0])
			fmt.Println("\t2)", right[0])

			var answer string

			for answer != "1\n" && answer != "2\n" {
				fmt.Printf("> ")
				reader := bufio.NewReader(os.Stdin)
				text, _ := reader.ReadString('\n')
				//if err != nil {
				//	panic(err)
				//}

				answer = text
			}

			switch answer {
			case "2\n":
				sorted = append(sorted, right[0])
				right = right[1:]
			case "1\n":
				sorted = append(sorted, left[0])
				left = left[1:]
			default:
				panic("unexpected answer")
			}
		}

		l = len(left)
		r = len(right)
	}

	//fmt.Println("returning:", sorted)
	return sorted, nil
}

func getItems(items *[]string) error {

	fmt.Println("[+] Input every option in a separate line and after the last element, press Ctrl + D")


	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		//fmt.Println("read:", scanner.Text())
		*items = append(*items, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
