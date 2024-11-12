package model

type User struct {
	ID      int
	Name    string
	Surname string
}

func GetAllUsers() (users []User, err error) {
	users = []User{
		{1, "Вася", "Пупкин"},
		{2, "Саша", "Грэй"},
		{3, "Джэк", "Воробей"},
		{4, "Джейн", "Эйр"},
		{5, "Тарзан", "Обезьянович"},
		{6, "Дориан", "Грэй"},
		{7, "Бильбо", "Бэггинз"},
		{8, "Фродо", "Сумкин"},
		{9, "Гермиона", "Грейнджер"},
	}
	return
}
