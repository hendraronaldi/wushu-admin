package model

type User struct {
	Name            string `json:"name"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	Birthdate       string `json:"birthdate"`
	Birthplace      string `json:"birthplace"`
	AddressNow      string `json:"addressnow"`
	AddressHometown string `json:"addresshometown"`
	Phone           string `json:"phone"`
	PhoneHome       string `json:"phonehome"`
	NameParent      string `json:"nameparent"`
	PhoneParent     string `json:"phoneparent"`
	ExpPrev         string `json:"expprev"`
	Motivation      string `json:"motivation"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
