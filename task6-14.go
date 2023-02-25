package main

import (
	"fmt"
	"sort"
	"time"
)

type TRAIN struct {
	NAZN string
	NUMR int
	TIME time.Time
}

func main() {
	// Створення масиву RASP та введення даних з клавіатури
	var RASP [8]TRAIN
	fmt.Println("Введіть дані про поїзди:")
	for i := 0; i < len(RASP); i++ {
		fmt.Printf("Поїзд %d:\n", i+1)
		fmt.Print("  Назва пункту призначення: ")
		fmt.Scan(&RASP[i].NAZN)
		fmt.Print("  Номер поїзда: ")
		fmt.Scan(&RASP[i].NUMR)
		var timeStr string
		fmt.Print("  Час відправлення (у форматі 15:04): ")
		fmt.Scan(&timeStr)
		t, err := time.Parse("15:04", timeStr)
		if err != nil {
			panic(err)
		}
		RASP[i].TIME = t
	}

	// Сортування за алфавітом в назвах пунктів призначення
	sort.Slice(RASP[:], func(i, j int) bool {
		return RASP[i].NAZN < RASP[j].NAZN
	})

	// Виведення інформації про поїзди, що відправляються після введеного з клавіатури часу
	var afterTime time.Time
	fmt.Print("Введіть час (у форматі 15:04), після якого шукати поїзди: ")
	fmt.Scan(&afterTime)
	found := false
	for _, train := range RASP {
		if train.TIME.After(afterTime) {
			if !found {
				fmt.Println("Поїзди, що відправляються після", afterTime.Format("15:04"), ":")
				found = true
			}
			fmt.Println(train.NAZN, train.NUMR, train.TIME.Format("15:04"))
		}
	}
	if !found {
		fmt.Println("Немає поїздів, що відправляються після", afterTime.Format("15:04"))
	}
}
