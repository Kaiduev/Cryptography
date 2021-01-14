package main

import (
	"fmt"
	"os"
)

var (
	objects = []string{"o1", "o2", "o3"}
	subjects = []string{"s1", "s2"}
	rights = []string{"-", "r", "rw", "rwe", "owner"}
	matrix = [][]int{{4, 1, 4}, {0, 4, 0}}
)

func PrintMatrix()  {
	for i:=0; i<len(objects); i++ {
		fmt.Print("    ", objects[i])
	}
	fmt.Println()
	for i:=0; i<len(matrix); i++{
		fmt.Print(subjects[i], "| ")
		for j:=0; j<len(matrix[i]); j++ {
			fmt.Print("  ", matrix[i][j], "  ")
		}
		fmt.Println()
	}
	fmt.Println()
	for i:=0; i<5; i++ {
		fmt.Println(i, "-", rights[i])
	}
}

func AddRight(objName string)  {
	fmt.Println()
	fmt.Println("Добавление прав")
	var found bool
	var subjName string
	fmt.Println("Введите субъект, которому хотите разграничить права: ")
	fmt.Fscan(os.Stdin, &subjName)
	var  subjIndex int
	var objIndex int
	for i:=0; i<len(subjects); i++{
		if subjName==subjects[i] {
			found = true
			subjIndex = i
			break
		}
	}
	if found == false{
		fmt.Println("Пользователь не найден.")
	}else {
		var objFound bool = false
		for i:=0; i<len(objects); i++{
			if objName == objects[i]{
				objFound = true
				objIndex = i
			}
	}
	if objFound == false{
		fmt.Println("Объект не найден.")
	}else {
		fmt.Println("Какие права хотите выдать?")
		PrintMatrix()
		for i:=0; i<5; i++{
			if rights[i] == "-" || rights[i] == "owner"{
				continue
			}
			fmt.Println(i, ")", rights[i])
		}
		var choice int
		fmt.Fscan(os.Stdin, &choice)
		if choice != 4{
			if matrix[subjIndex][objIndex] != 4 {
				matrix[subjIndex][objIndex] = choice
			}else {
				fmt.Println("Вы не можете разграничить себе права так как вы являетесь владельцем этого объекта.")
			}
		}else{
			fmt.Println("Вы не можете полностью передать владение объектом.")
		}
	}
}

	PrintMatrix()
	fmt.Println()
	fmt.Println()
}

func NewObj()  {
	fmt.Println()
	var filename string
	var name string
	var subjIndex int
	fmt.Print("Введите имя субъекта: ")
	fmt.Fscan(os.Stdin, &name)
	for i:=0; i<len(subjects); i++{
		if name == subjects[i]{
			subjIndex = i
			break
		}
	}
	filename = "Otr"
	objects = append(objects, filename)
	for i:=0; i<len(matrix); i++{
		if subjIndex==i {
			matrix[i] = append(matrix[i], 4)
		}else {
			matrix[i] = append(matrix[i], 0)
		}
	}
	fmt.Println("Объект ", filename, " успешно создан")
	var yesorno string
	fmt.Println("Разграничить кому либо права? ")
	fmt.Println("да или нет?")
	fmt.Fscan(os.Stdin, &yesorno)
	if yesorno == "да" {
		AddRight(filename)
	}else if yesorno == "нет" {
		fmt.Println("Добавление объекта завершено")
	}
	fmt.Println()
}
func OpenTr()  {
	var choice int
	fmt.Println("Выберите объект который хотите открыть.")
	for i:=0; i<len(objects); i++ {
		fmt.Println(i,")", objects[i])
	}
	fmt.Fscan(os.Stdin, &choice)
	if objects[choice] == "Otr" {
		var name string = "Str"
		subjects = append(subjects, name)
		var newsubj []int
		for i := 0; i < len(objects); i++ {
			if matrix[0][i] == 4{
				newsubj = append(newsubj, matrix[0][i] - 1)
			}else {
				newsubj = append(newsubj, matrix[0][i])
			}
		}
		matrix = append(matrix, newsubj)
	}
	PrintMatrix()
	fmt.Println("Объект успешно открыт.")
	subjects = subjects[:len(subjects) - 1]
	matrix = matrix[:len(matrix) - 1]
	var filename string = "O"
	objects = append(objects, filename)
	for i := 0; i < len(matrix); i++ {
		if subjects[i] == "s1"{
			matrix[i] = append(matrix[i], 0)
		}else if subjects[i] == "s2"{
			matrix[i] = append(matrix[i], 1)
		}
	}
	fmt.Println()
	PrintMatrix()
}

func main(){
	var x int
	for {
		fmt.Println("1)Создание трояна")
		fmt.Println("2)Открыть объект")
		fmt.Println("3)Матрица доступа")
		fmt.Println("0)Выход")
		fmt.Fscan(os.Stdin, &x)
		if x==1{
			NewObj()
		} else if x==2 {
			OpenTr()
		}else if x == 3{
			PrintMatrix()
		}else if x == 0{
			break
		}
	}
}
