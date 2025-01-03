package main

type Service struct {
	description    string   // Описание
	durationMonths int      // Продолжительность в месяцах
	monthlyFee     float64  // Ежемесячная плата
	features       []string // Функции
}

func newService(description string, durationMonths int, price float64, features []string) Service {
	return Service{description,
		durationMonths,
		price,
		features}
}

// Методы интерфейса Expense
func (s Service) getName() string {
	return s.description
}

func (s Service) getCost(recur bool) float64 {
	if recur {
		return s.monthlyFee * float64(s.durationMonths)
	}
	return s.monthlyFee
}
