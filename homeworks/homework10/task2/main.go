package main

import "time"

// 30. Отслеживание активностей пользователя
//    • Описание: Реализуйте структуру Activity с полями activityType и timestamp.
//   	Реализуйте структуру UserActivityTracker с срезом активностей и методы для добавления активности и получения
//  	всех активностей после определенного времени.
//    • Методы:
//        ◦ AddActivity(activityType string, timestamp time.Time)
//        ◦ ActivitiesAfterTime(timestamp time.Time) []Activity

type Activity struct {
	activityType string
	timestamp    time.Time
}
type UserActivityTracker struct {
	activities []Activity
}

func (u *UserActivityTracker) AddActivity(activityType string, timestamp time.Time) {

}
func (u *UserActivityTracker) ActivitiesAfterTime(timestamp time.Time) []Activity {
	return u.activities
}
func main() {

}
