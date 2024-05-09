package model

type Movie struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	ReleaseYear string `json:"releaseyear"`
	Director    string `json:"director"`
	OTTID       string `json:"ottid"`
}
