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
		fmt.Fprintf(os.Stderr, "error reading items: %v\n", err)
		os.Exit(1)
	}

	sorted, err := mergesort(items)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error sorting: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\nRanking:")
	for i, item := range sorted {
		fmt.Printf("  %d. %s\n", i+1, item)
	}
}

func mergesort(items []string) ([]string, error) {
	if len(items) < 2 {
		return items, nil
	}

	middle := int(math.Round(float64(len(items)) / 2))

	left, err := mergesort(items[:middle])
	if err != nil {
		return nil, err
	}

	right, err := mergesort(items[middle:])
	if err != nil {
		return nil, err
	}

	return merge(left, right)
}

func merge(left, right []string) ([]string, error) {
	var sorted []string

	for len(left) > 0 && len(right) > 0 {
		answer, err := compare(left[0], right[0])
		if err != nil {
			return nil, err
		}

		if answer == 1 {
			sorted = append(sorted, left[0])
			left = left[1:]
		} else {
			sorted = append(sorted, right[0])
			right = right[1:]
		}
	}

	sorted = append(sorted, left...)
	sorted = append(sorted, right...)

	return sorted, nil
}

func compare(a, b string) (int, error) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("[+] Which is better?")
		fmt.Printf("\t1) %s\n", a)
		fmt.Printf("\t2) %s\n", b)
		fmt.Printf("> ")

		text, err := reader.ReadString('\n')
		if err != nil {
			return 0, fmt.Errorf("error reading input: %w", err)
		}

		switch text {
		case "1\n":
			return 1, nil
		case "2\n":
			return 2, nil
		default:
			fmt.Println("Please enter 1 or 2.")
		}
	}
}

func getItems(items *[]string) error {
	fmt.Println("Enter each option on a separate line. Press Ctrl+D when done.")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			*items = append(*items, line)
		}
	}

	return scanner.Err()
}
