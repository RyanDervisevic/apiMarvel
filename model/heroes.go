package model

type Heroes struct {
	ID         string `json:"id"`
	HeroesName string `json:"heroes_name"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Powers     string `json:"powers"`
	Equipment  string `json:"equipment"`
}

// type LoginUser struct {
// 	Email    string   `json:"email"`
// 	Password Password `json:"pass"`
// }

// type Password string

// func (p *Password) UnmarshalJSON(b []byte) error {
// 	var str string
// 	err := json.Unmarshal(b, &str)
// 	if err != nil {
// 		return err
// 	}

// 	h := sha256.New()
// 	h.Write([]byte(str))
// 	*p = Password(fmt.Sprintf("%x", h.Sum(nil)))

// 	return nil
// }

// func (p Password) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(p)
// }
