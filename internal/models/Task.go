package models

type Task struct {
	ID       int    `json:"id"`
	Status   string `json:"status"`
	Title    string `json:"title"`
	ActiveAt string `json:"activeAt"`
}

func (t *Task) String() string {
	//return t.Title + " " + t.ActiveAt.String()
	return t.Title + " " + t.ActiveAt
}

func NewTask(id int, status string, title string, activeAt string) *Task {
	//layout := "2006-01-02"
	//activeAtTime, err := time.Parse(layout, activeAt)
	//if err != nil {
	//	log.Printf("Error parsing date: %v", err)
	//	return nil
	//}
	return &Task{
		ID:     id,
		Status: status,
		Title:  title,
		//ActiveAt: activeAtTime,
		ActiveAt: activeAt,
	}
}
