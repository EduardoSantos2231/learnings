package datatypes

type UserData struct {
	Name string `json:"name"`
}

type UserReqData struct {
	Ids []int `json:"ids"`
}

func (u *UserData) IsValid() bool {
	return u.Name != ""
}

func (u *UserReqData) IsValid() bool {
	return len(u.Ids) > 0
}
