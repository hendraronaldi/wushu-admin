package model

type User struct {
	Name          string `json:"name"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Birthdate     string `json:"birthdate"`
	Birthplace    string `json:"birthplace"`
	Address       string `json:"address"`
	AddressParent string `json:"addressparent"`
	Phone         string `json:"phone"`
	NameParent    string `json:"nameparent"`
	PhoneParent   string `json:"phoneparent"`
	ExpPrev       string `json:"expprev"`
	Motivation    string `json:"motivation"`
	Status        int    `json:"status"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
