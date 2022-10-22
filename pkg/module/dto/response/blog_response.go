package response

type (
	ActiveBlogEntity struct {
		ID      int    `json:"id"`
		Title   string `json:"title"`
		Context string `json:"context"`
		Option  Options
	}

	HistoryBlogEntity struct {
		ID           int    `json:"id"`
		ActiveBlogID int    `json:"active_blog_id"`
		Title        string `json:"title"`
		Context      string `json:"context"`
		Option       Options
	}
)

type (
	FindByIDBlogResponse struct {
		Result *ActiveBlogResult `json:"result"`

		CodeErr error `json:"code_err"`
		MsgErr  error `json:"msg"`
	}

	CreateBlogResponse struct {
		Result *ActiveBlogResult `json:"result"`

		CodeErr error  `json:"code_err"`
		MsgErr  string `json:"msg"`
	}

	DeleteBlogResponse struct {
		Result *HistoryBlogResult `json:"result"`

		CodeErr error  `json:"code"`
		MsgErr  string `json:"msg"`
	}

	FindAllBlogResponse struct {
		Results *ActiveBlogResults `json:"results"`

		CodeErr error  `json:"code"`
		MsgErr  string `json:"msg"`
	}

	UpdateBlogResponse struct {
		Result *ActiveBlogResult `json:"result"`

		CodeErr error  `json:"code"`
		MsgErr  string `json:"msg"`
	}
)

type (
	ActiveBlogResult struct {
		Blog *ActiveBlogEntity `json:"blog"`
	}
	ActiveBlogResults struct {
		Blogs []*ActiveBlogEntity `json:"blog"`
	}

	HistoryBlogResult struct {
		Blog *HistoryBlogEntity `json:"blog"`
	}

	HistoryBlogResults struct {
		Blogs []*HistoryBlogEntity `json:"blogs"`
	}
)
