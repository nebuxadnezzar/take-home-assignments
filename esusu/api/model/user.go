package model

type (
	User struct {
		ID         int            `json:"id"`
		Name       string         `json:"name"`
		AuthToken  string         `json:"authToken"`
		TokenCount int            `json:"tokenCount"`
		Hits       map[string]int `json:"hits"`
	}

	UserRequest struct {
		ID        int    `query:"uid" json:"id"`
		AuthToken string `query:"token" json:"authToken"`
	}
)
