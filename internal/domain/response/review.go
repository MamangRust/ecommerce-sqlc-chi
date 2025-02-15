package response

type ReviewResponse struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	ProductID int    `json:"product_id"`
	Name      string `json:"name"`
	Comment   string `json:"comment"`
	Rating    int    `json:"rating"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ReviewResponseDeleteAt struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	ProductID int    `json:"product_id"`
	Name      string `json:"name"`
	Comment   string `json:"comment"`
	Rating    int    `json:"rating"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type ApiResponseReview struct {
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Data    *ReviewResponse `json:"data"`
}

type ApiResponseReviewDeleteAt struct {
	Status  string                  `json:"status"`
	Message string                  `json:"message"`
	Data    *ReviewResponseDeleteAt `json:"data"`
}

type ApiResponsesReview struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Data    []*ReviewResponse `json:"data"`
}

type ApiResponseReviewDelete struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ApiResponseReviewAll struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ApiResponsePaginationReviewDeleteAt struct {
	Status     string                    `json:"status"`
	Message    string                    `json:"message"`
	Data       []*ReviewResponseDeleteAt `json:"data"`
	Pagination PaginationMeta            `json:"pagination"`
}

type ApiResponsePaginationReview struct {
	Status     string            `json:"status"`
	Message    string            `json:"message"`
	Data       []*ReviewResponse `json:"data"`
	Pagination PaginationMeta    `json:"pagination"`
}
