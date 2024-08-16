package models

// создание моделей для описания сущностей			
type User struct {
	Username string
	Password string
	Habits   []Habit
}

type Habit struct {
	Name   string
	UserId int
	User   User
	Marks  []Mark
}

type Mark struct {
	Date    string
	HabitID int
	Habit   Habit
}
