package main

import (
	"fmt"
)

//Создай мапу, где ключом будет строка (например, имя студента),
//а значением – целое число (например, оценка студента).
//	Добавь несколько пар ключ-значение в мапу, затем выведи мапу на экран.

func main() {
	groups := map[string]map[string]int{
		"A": {"Саня": 56, "Стас": 90},
		"B": {"Дэн": 34, "Ишмаев": 35},
	}
	fmt.Println(groups)
	for n, g := range groups {
		fmt.Println(n, ":", g)
		for i, d := range g {
			fmt.Println(i, ":", d)
		}

	}

	students := map[string]int{
		"Саня": 1,
		"Стас": 2,
		"Дэн":  3,
	}
	students["пащенко"] = 4
	students["ишмаев"] = 5
	students["Саня"] = 6
	fmt.Println(students)

	if ww, rr := students["Саня"]; rr {
		fmt.Println("Он есть", ww)
	} else {
		fmt.Println("Его нету")
	}
	delete(students, "ишмаев")
	fmt.Println(students)

	for name, grade := range students {
		fmt.Println(name, ":", grade)
	}

	fmt.Println(len(students))
}
