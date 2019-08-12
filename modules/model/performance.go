package model

type Performance struct {
	Email          string      `json:"email"`
	Date           string      `json:"date"`
	Flexibility    Flexibility `json:"flexibility"`
	Power          Power       `json:"power"`
	Flexibility_id int
	Power_id       int
}

type Flexibility struct {
	Shoulder int `json:"shoulder"`
	Wrist    int `json:"wrist"`
	Waist    int `json:"waist"`
	Leg      int `json:"leg"`
}

type Power struct {
	Strike    int `json:"strike"`
	HandSwing int `json:"handswing"`
	Spin      int `json:"spin"`
	Jump      int `json:"jump"`
	Kick      int `json:"kick"`
	LegSwing  int `json:"legswing"`
}
