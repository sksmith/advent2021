package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/sksmith/advent2021/day01"
	"github.com/sksmith/advent2021/day02"
	"github.com/sksmith/advent2021/util"
)

func main() {
	var wg sync.WaitGroup

	Run("Day 1 Part 1", &wg, func() interface{} {
		f, _ := os.Open("./day01/input")
		depths := util.ReadIntArr(bufio.NewReader(f))
		return day01.CountIncreases(depths)
	})

	Run("Day 1 Part 2", &wg, func() interface{} {
		f, _ := os.Open("./day01/input")
		depths := util.ReadIntArr(bufio.NewReader(f))
		return day01.CompareWindows(depths)
	})

	Run("Day 2 Part 1", &wg, func() interface{} {
		f, _ := os.Open("./day02/input.txt")
		cmds := util.ReadStringArr(bufio.NewReader(f))
		return day02.Move(cmds)
	})

	Run("Day 2 Part 2", &wg, func() interface{} {
		f, _ := os.Open("./day02/input.txt")
		cmds := util.ReadStringArr(bufio.N1176514794ewReader(f))
		return day02.MoveAdvanced(cmds)
	})

	fmt.Println("Executing days...")
	wg.Wait()
	fmt.Println("Complete!")
}

func Run(label string, wg *sync.WaitGroup, f func() interface{}) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Printf("%s: %v\n", label, f())
	}()
}
