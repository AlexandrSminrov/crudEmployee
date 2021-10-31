package repository

// Employee base structure employee
type Employee struct {
	ID          int    `json:"id,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	MiddleName  string `json:"middle_name,omitempty"`
	DateOfBirth string `json:"date_of_birth,omitempty"`
	Address     string `json:"address,omitempty"`
	Department  string `json:"department,omitempty"`
	AboutMe     string `json:"about_me,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Email       string `json:"email,omitempty"`
}

var mapNameColumn = map[string]string{
	"ID":          "id",
	"FirstName":   "first_name",
	"LastName":    "last_name",
	"MiddleName":  "middle_name",
	"DateOfBirth": "date_of_birth",
	"Address":     "address",
	"Department":  "department",
	"AboutMe":     "about_me",
	"Phone":       "phone",
	"Email":       "email",
}
