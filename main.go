package main

import (
	"fmt"
	"os"
	"strconv"
)

type student struct {
	no         int
	name       string
	address    string
	occupation string
	reason     string
}

type listOfStudents struct {
	students []student
}

type data interface {
	getName(no int) string
	getAddress(no int) string
	getOccupation(no int) string
	getReason(no int) string
}

func (l listOfStudents) getName(no int) string {
	for _, student := range l.students {
		if no == student.no {
			return student.name
		}
	}
	return ""
}

func (l listOfStudents) getAddress(no int) string {
	for _, student := range l.students {
		if no == student.no {
			return student.address
		}
	}
	return ""
}

func (l listOfStudents) getOccupation(no int) string {
	for _, student := range l.students {
		if no == student.no {
			return student.occupation
		}
	}
	return ""
}

func (l listOfStudents) getReason(no int) string {
	for _, student := range l.students {
		if no == student.no {
			return student.reason
		}
	}
	return ""
}

func print(i int, d data) {
	fmt.Println("Nama\t: ", d.getName(i))
	fmt.Println("Alamat\t: ", d.getAddress(i))
	fmt.Println("Pekerjaan\t: ", d.getOccupation(i))
	fmt.Println("Alasan\t: ", d.getReason(i))
}

func main() {
	studentId := os.Args[1]

	i, _ := strconv.Atoi(studentId)

	var los listOfStudents

	var studentsData = []student{
		{
			no:         1,
			name: "Anugrah",
			address: "Jl. Test",
			occupation: "Software Engineer",
			reason: "want to earn some money",
		},
		{
			no:         2,
      name: "Kresnaya",
			address: "Jl. Ter Test",
			occupation: "Frontend Engineer",
			reason: "Like to design and code",
		},
	}

	los.students = studentsData

	print(i, los)
}
