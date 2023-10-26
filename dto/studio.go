package dto

type CreateStudioRequest struct {
	Name				string		`json:"name"`
	Capacity		int				`json:"capacity"`
	CinemaID		uint			`json:"cinema_id"`
}