package schemas

type ApiUri string

type CoreApiStatus int

const (
	CoreApiStatusPrivate CoreApiStatus = 1
	CoreApiStatusPublic  CoreApiStatus = 2
)

type CoreStatus int

const (
	CoreStatusDraft         CoreStatus = 1
	CoreStatusTest          CoreStatus = 2
	CoreStatusPendingReview CoreStatus = 3
	CoreStatusRejected      CoreStatus = 4
	CoreStatusApproved      CoreStatus = 5
	CoreStatusLive          CoreStatus = 6
)

type HashAlgo int

const (
	HashAlgoSha1 HashAlgo = 1
	HashAlgoMd5  HashAlgo = 2
)

type Pagination struct {
	Index       int `json:"index"`
	PageSize    int `json:"pageSize"`
	ResultCount int `json:"resultCount"`
	TotalCount  int `json:"totalCount"`
}
