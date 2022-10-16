package request


type (
	BlogFindByIDRequest struct {
		ID int `json:"id"`
	}

	BlogDeleteRequest struct {
		ID int `json:"id"`
	}
	
	BlogCreateRequest struct {
		ID int `json:"id"`
		Title string `json:"title"`
		Context string `json:"context"`
		Revision uint `json:"revision"`
	}

	BlogUpdateRequest struct {
		ID int `json:"id"`
		Title string `json:"title"`
		Context string `json:"context"`
		Revision uint `json:"revision"`
	}
)

