package data

type Lecture struct {
	SubjectCode string   `json:"subject_code"`
	SubjectName string   `json:"subject_name"`
	Lecturers   []string `json:"lecturers"`
	Date        string   `json:"date"`
	Start       string   `json:"start"`
	End         string   `json:"end"`
	Room        string   `json:"room"`
}
