package models

type Clas struct {
	PeriodNo       int    `json:"period_no"`
	ProfessorNo    string `json:"professor_no"`
	Subject        string `json:"subject"`
	LectureModelID int    `json:"lecture_model_id"`
}

type Lecture struct {
	ID     uint   `json:"id"`
	Stream string `json:"stream"`
	Day    string `json:"day"`
	Year   string `json:"year"`
	Clas   []Clas `json:"classes" gorm:"foreignKey:LectureModelID"`
}
