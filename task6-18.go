package main

import (
	"fmt"
	"sort"
)

type NOTE struct {
	NAME string
	TELE string
	BDAY []int
}

func inputBlocnote() []NOTE {
	blocnote := make([]NOTE, 8)

	for i := 0; i < 8; i++ {
		var name, tele string
		var bday []int

		fmt.Printf("Enter name for note %d: ", i+1)
		fmt.Scan(&name)

		fmt.Printf("Enter telephone number for note %d: ", i+1)
		fmt.Scan(&tele)

		fmt.Printf("Enter birth date (DD MM YYYY) for note %d: ", i+1)
		bday = make([]int, 3)
		fmt.Scan(&bday[0], &bday[1], &bday[2])

		blocnote[i] = NOTE{name, tele, bday}
	}

	sort.Slice(blocnote, func(i, j int) bool {
		return blocnote[i].BDAY[2] < blocnote[j].BDAY[2] || blocnote[i].BDAY[2] == blocnote[j].BDAY[2] && (blocnote[i].BDAY[1] < blocnote[j].BDAY[1] || blocnote[i].BDAY[1] == blocnote[j].BDAY[1] && blocnote[i].BDAY[0] < blocnote[j].BDAY[0])
	})

	return blocnote
}

func findNoteByTelephone(blocnote []NOTE, tele string) *NOTE {
	for i := 0; i < len(blocnote); i++ {
		if blocnote[i].TELE == tele {
			return &blocnote[i]
		}
	}
	return nil
}

func main() {
	blocnote := inputBlocnote()

	var tele string
	fmt.Print("Enter telephone number to search: ")
	fmt.Scan(&tele)

	note := findNoteByTelephone(blocnote, tele)
	if note != nil {
		fmt.Printf("Found note:\nName: %s\nTelephone: %s\nBirth date: %02d.%02d.%d\n", note.NAME, note.TELE, note.BDAY[0], note.BDAY[1], note.BDAY[2])
	} else {
		fmt.Println("Note not found.")
	}
}
