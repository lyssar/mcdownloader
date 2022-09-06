package schemas

type ApiUri string

type CoreApiStatus int

const (
	CORE_API_STATUS_PRIVATE CoreApiStatus = 1
	CORE_API_STATUS_PUBLIC  CoreApiStatus = 2
)

type CoreStatus int

const (
	CORE_STATUS_DRAFT          CoreStatus = 1
	CORE_STATUS_TEST           CoreStatus = 2
	CORE_STATUS_PENDING_REVIEW CoreStatus = 3
	CORE_STATUS_REJECTED       CoreStatus = 4
	CORE_STATUS_APPROVED       CoreStatus = 5
	CORE_STATUS_LIVE           CoreStatus = 6
)

type HashAlgo int

const (
	HASH_ALGO_SHA1 HashAlgo = 1
	HASH_ALGO_MD5  HashAlgo = 2
)

type Pagination struct {
	Index       int `json:"index"`
	PageSize    int `json:"pageSize"`
	ResultCount int `json:"resultCount"`
	TotalCount  int `json:"totalCount"`
}
