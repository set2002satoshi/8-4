package response



type (
	ActiveBlogEntity struct {
		ID int `json:"id"`
		Title string `json:"title"`
		Context string `json:"context"`
		Option Options

	}

	HistoryBlogEntity struct {
		ID int `json:"id"`
		ActiveBlogID int `json:"active_blog_id"`
		Title string `json:"title"`
		Context string `json:"context"`
		Option Options
	}
)


type (
	FindByIDBlogResponse struct {
		Result *ActiveBlogResult `json:"result"`

		CodeErr error `json:"code_err"`
		MsgErr error `json:"msg"`
	}

	CrateBlogResponse struct {
		Result *ActiveBlogResult `json:"result"`

		CodeErr error `json:"code_err"`
		MsgErr string `json:"msg"`
	}

	DeleteBlogResponse struct {
		Result *HistoryBlogResult `json:"result"`

		CodeErr error `json:"code"`
		MsgErr string `json:"msg"`
	}

	FindAllBlogResponse struct {
		Results *ActiveBlogResults `json:"results"`

		CodeErr error `json:"code"`
		MsgErr string `json:"msg"`
	}

	UpdateBlogResponse struct {
		Result *ActiveBlogResult `json:"result"`

		CodeErr error `json:"code"`
		MsgErr string `json:"msg"`
	}
)

type (
	ActiveBlogResult struct {
		User ActiveBlogEntity `json:"user"`
	}
	ActiveBlogResults struct {
		Users []*ActiveBlogEntity `json:"users"`
	}

	HistoryBlogResult struct {
		User *HistoryBlogEntity `json:"user"`
	}

	HistoryBlogResults struct {
		Users []*HistoryBlogEntity `json:"users"`
	}

)

