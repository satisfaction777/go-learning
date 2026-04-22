package main

import "fmt"



type Student struct{
	Name string
	Grades []float64
}

func (s Student) Average() float64{
	total := 0.0
	for _, value := range s.Grades { // перебор элементов слайса
		total += value
	}
	return total / float64(len(s.Grades))
}

func (s *Student) AddGrade(g float64){
	s.Grades = append(s.Grades,g)
}

func (s Student) IsExcellent() bool{
	if s.Average() >= 4.5{
	return true
	} else {
		return false
	} 
}

func (s Student) Describe() string{
	sStatus := ""
	if s.IsExcellent(){
		sStatus = " (отличник)"
	} 
	return fmt.Sprintf("%s - средний балл: %.2f%s\n",
	s.Name,s.Average(), sStatus)
	}

func bestStudent(students []Student) Student {
    best := students[0] // начинаем с первого

    for _, s := range students {
        if s.Average() > best.Average() {
            best = s // нашли лучше — обновляем
        }
    }
    return best
}

func main(){


	students := []Student{
    {Name: "Вика",  Grades: []float64{5.0, 4.5, 5.0, 4.0}},
    {Name: "Алекс", Grades: []float64{3.0, 4.0, 3.5, 4.0}},
    {Name: "Маша",  Grades: []float64{5.0, 5.0, 4.5, 5.0}},
}

students[0].AddGrade(5.0)
students[1].AddGrade(2.5)
students[2].AddGrade(2.0)

for _, grades := range students{
	fmt.Printf("%s", grades.Describe())
}

best := bestStudent(students)
fmt.Printf("Лучший студент: %s\n", best.Name)
}
