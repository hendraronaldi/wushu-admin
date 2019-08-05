package model

type User struct {
	Name          string `json:"name"`
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
	Education     string `json:"education"`
	Status        int    `json:"status"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
