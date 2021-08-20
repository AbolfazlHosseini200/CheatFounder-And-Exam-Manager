package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type quiz struct {
	Answers           map[string][]int   `json:"answers"`
	Scores            map[string]float32 `json:"scores"`
	QuizName          string             `json:"quiz-name"`
	NumberOfQuestions int                `json:"numberOfQuestions"`
	Key               []int              `json:"key"`
}
type class struct {
	Students []student `json:"students"`
	Number   string    `json:"number-of-class"`
}
type student struct {
	FullName string `json:"name"`
}
type question struct {
	number int
	numberOfFaults int
	answeredCorrect []student
}
//a func to check if our input is a valid number
func checkNumValid(n int, max int, min int) int {
	for n < min || n > max {
		fmt.Println("invalid number" +
			"\n" +
			"retry!!!!!!!!")
		fmt.Scan(&n)
	}
	return n
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	var classes []class
	var quizzes []quiz
	jsonFile, _ := os.Open("C:\\Users\\AmirHossein\\Desktop\\Coding Stuff\\Code Go\\CheatFounder\\classes.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &classes)
	jsonFile2, _ := os.Open("C:\\Users\\AmirHossein\\Desktop\\Coding Stuff\\Code Go\\CheatFounder\\quizzes.json")
	byteValue2, _ := ioutil.ReadAll(jsonFile2)
	json.Unmarshal(byteValue2, &quizzes)
	for true {
		fmt.Println("welcome , Choose your action:")
		fmt.Println("1)view classes" +
			"\n" +
			"2)add class" +
			"\n" +
			"3)new quiz" +
			"\n" +
			"4)older quizzes" +
			"\n" +
			"5)save")
		var n int
		fmt.Scan(&n)
		n = checkNumValid(n, 5, 1)
		//checking old classes
		if n == 1 {
			i := 0
			for i < len(classes) {
				fmt.Print(i+1, ")"+classes[i].Number)
				i++
			}
			fmt.Print("choose class:")
			var n2 int
			fmt.Scan(&n2)
			if len(classes) > 0 {
				n2 = checkNumValid(n2, len(classes), 1)
			}
			n2--
			i = 0
			for i < len(classes[n2].Students) {
				fmt.Println(i+1, ")"+classes[n2].Students[i].FullName)
				i++
			}
		}
		//making new class
		if n == 2 {
			fmt.Print("enter class name:")
			name, _ := reader.ReadString('\n')

			fmt.Println("enter class students" +
				"\n" +
				"enter 0 at the end")
			var students2 []student
			for true {
				fmt.Print("enter student number", len(students2)+1, "'s name:")
				var studentName string
				fmt.Scan(&studentName)
				if strings.EqualFold("0", studentName) {
					break
				}
				students2 = append(students2, student{FullName: studentName})
			}
			classes = append(classes, class{Number: name, Students: students2})
		}
		//making new quiz
		if n == 3 {
			fmt.Print("choose your quiz name:")
			var quizName string
			fmt.Scan(&quizName)
			fmt.Print("how many question has this quiz:")
			var QNum int
			fmt.Scan(&QNum)
			QNum = checkNumValid(QNum, 1000, 1)
			var quizKey []int
			i := 0
			fmt.Println("enter quiz key")
			for i < QNum {
				fmt.Print(i+1, ":")
				var simpleKey int
				fmt.Scan(&simpleKey)
				simpleKey = checkNumValid(simpleKey, 4, 1)
				quizKey = append(quizKey, simpleKey)
				i++
			}
			quizzes = append(quizzes, quiz{QuizName: quizName, Key: quizKey, NumberOfQuestions: QNum})
			quizzes[len(quizzes)-1].Answers = make(map[string][]int)
			quizzes[len(quizzes)-1].Scores = make(map[string]float32)
		}
		//manage old quizzes
		if n == 4 {
			i := 0
			for i < len(quizzes) {
				fmt.Println(i+1, ")"+quizzes[i].QuizName)
				i++
			}
			fmt.Println("choose an exam or go to previous by pressing 0 button")
			var quizChosen int
			fmt.Scan(&quizChosen)
			quizChosen = checkNumValid(quizChosen, len(quizzes), 0)
			if quizChosen == 0 {
				continue
			}
			quizChosen--
			fmt.Println("choose action:" +
				"\n" +
				"1)enter answers" +
				"\n" +
				"2)show scores" +
				"\n" +
				"3)check for cheats" +
				"\n" +
				"4)hardest questions" +
				"\n" +
				"5)back")
			var n3 int
			fmt.Scan(&n3)
			n3 = checkNumValid(n3, 5, 1)
			if n3 == 1 {
				i = 0
				for i < len(classes) {
					fmt.Print(i+1, ")"+classes[i].Number)
					i++
				}
				fmt.Print("choose class:")
				var chosenClass int
				fmt.Scan(&chosenClass)
				if len(classes) > 0 {
					chosenClass = checkNumValid(chosenClass, len(classes), 1)
				}
				chosenClass--
				i = 0
				for i < len(classes[chosenClass].Students) {
					fmt.Println(i+1, ")"+classes[chosenClass].Students[i].FullName)
					i++
				}
				i = 0
				fmt.Println("Choose student:")
				var chosenStudent int
				fmt.Scan(&chosenStudent)
				chosenStudent = checkNumValid(chosenStudent, len(classes[chosenClass].Students), 1)
				chosenStudent--
				var studentAns []int
				var correct, wrong = 0.0, 0.0
				i = 0
				for i < quizzes[quizChosen].NumberOfQuestions {
					fmt.Print(i+1, ")")
					var ans int
					fmt.Scan(&ans)
					ans = checkNumValid(ans, 4, 0)
					studentAns = append(studentAns, ans)
					if ans == quizzes[quizChosen].Key[i] {
						correct++
					}
					if ans != quizzes[quizChosen].Key[i] && ans != 0 {
						wrong++
					}
					i++
				}
				i = 0
				quizzes[quizChosen].Answers[classes[chosenClass].Students[chosenStudent].FullName] = studentAns
				quizzes[quizChosen].Scores[classes[chosenClass].Students[chosenStudent].FullName] = float32(((3*correct)-wrong)/float64(quizzes[quizChosen].NumberOfQuestions*3)) * 100
			}
			if n3 == 2 {
				j := 0
				i = 0
				for i < len(classes) {
					j = 0
					for j < len(classes[i].Students) {
						fmt.Println(classes[i].Students[j].FullName+":", quizzes[quizChosen].Scores[classes[i].Students[j].FullName])
						j++
					}
					i++
				}
			}
			if n3 == 3 {
				i = 0
				j := 0
				for i < len(classes) {
					j = 0
					for j < len(classes[i].Students) {
						if quizzes[quizChosen].Answers[classes[i].Students[j].FullName] != nil {
							o := i
							oo := j + 1
							for o < len(classes) {
								oo = 0
								for oo < len(classes[o].Students) {
									if quizzes[quizChosen].Answers[classes[o].Students[oo].FullName] != nil && (o != i && oo != j) {
										k := 0
										wrongs := 0
										sameWrongs := 0
										for k < quizzes[quizChosen].NumberOfQuestions {
											if quizzes[quizChosen].Answers[classes[i].Students[j].FullName][k] != quizzes[quizChosen].Key[k] {
												wrongs++
												if quizzes[quizChosen].Answers[classes[i].Students[j].FullName][k] == quizzes[quizChosen].Answers[classes[o].Students[oo].FullName][k] {
													sameWrongs++
												}
											}
											k++
										}
										if wrongs != 0 {
											if float32(sameWrongs/wrongs) > 0.75 {
												fmt.Println(classes[i].Students[j].FullName+" has ", float32(sameWrongs/wrongs)*100, "% similarity with "+classes[o].Students[oo].FullName)
											}
										}
									}
									oo++
								}
								o++
							}
						}
						j++
					}
					i++
				}
			}
			if n3 == 4 {
				var qInfo []question
				i=0
				for i<quizzes[quizChosen].NumberOfQuestions {
					fault:=0
					var correctAnswered []student
					g := 0
					gg := 0
					for gg < len(classes) {
						g = 0
						for g < len(classes[gg].Students) {
							if len(quizzes[quizChosen].Answers[classes[gg].Students[g].FullName])!=0{
								if quizzes[quizChosen].Answers[classes[gg].Students[g].FullName][i]==quizzes[quizChosen].Key[i]{
									correctAnswered=append(correctAnswered,classes[gg].Students[g])
								}else {
									fault++
								}
							}

							g++
						}
						gg++
					}
					qInfo=append(qInfo,question{number: i+1,numberOfFaults: fault,answeredCorrect: correctAnswered})
					i++
				}
				i=0
				j:=0
				//var aux
				for j<quizzes[quizChosen].NumberOfQuestions {
					i=0
					for i<quizzes[quizChosen].NumberOfQuestions-1{
						if qInfo[i+1].numberOfFaults>qInfo[i].numberOfFaults{
							aux:=qInfo[i]
							qInfo[i]=qInfo[i+1]
							qInfo[i+1]=aux
						}
						i++
					}

					j++
				}
				i=0
				for i<quizzes[quizChosen].NumberOfQuestions {
					fmt.Println(i+1,")question number:",qInfo[i].number,"number of faults:",qInfo[i].numberOfFaults,"\n"+"guys who answered correct\n",qInfo[i].answeredCorrect)
					if i==4{
						break
					}
					i++
				}
			}
			if n3 == 5 {
				continue
			}
		}
		if n == 5 {
			file, err := json.MarshalIndent(classes, "", "  ")
			if err != nil {
				fmt.Printf("failed to marshal class json:%s", err)
			}
			err = ioutil.WriteFile("C:\\Users\\AmirHossein\\Desktop\\Coding Stuff\\Code Go\\CheatFounder\\classes.json", file, 0644)
			file2, err := json.MarshalIndent(quizzes, "", "  ")
			if err != nil {
				fmt.Printf("failed to marshal quiz json:%s", err)
			}
			err = ioutil.WriteFile("C:\\Users\\AmirHossein\\Desktop\\Coding Stuff\\Code Go\\CheatFounder\\quizzes.json", file2, 0644)

		}
	}
}
