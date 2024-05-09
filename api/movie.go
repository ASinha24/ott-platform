package api

type MovieRequest struct {
	Name        string `json:"name"`
	ReleaseYear string `json:"releaseyear"`
	Director    string `json:"director"`
}

type CreateMoveieResponse struct {
	*MovieRequest
	ID string `json:"id"`
}

// type UpdateMovieRequest struct {
// 	Name        string `json:"name"`
// 	ReleaseYear string `json:"releaseyear"`
// 	Director    string `json:"director"`
// }
