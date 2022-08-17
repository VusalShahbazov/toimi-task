package announcement

type AnnSearch struct {
	Page      int    `json:"page" form:"page"`
	SortField string `json:"sort_field" form:"sort_field"`
	SortDesc  string `json:"sort_desc" form:"sort_desc"`
}

const PerPage = 10

var AllowedOrderFields = []string{"created_at", "price"}
var AllowedOrderDesc = []string{"desc", "asc"}

func (s *AnnSearch) GetLimit() int {
	return PerPage
}

func (s *AnnSearch) GetOffset() int {
	if s.Page > 1 {
		return (s.Page - 1) * PerPage
	}
	return 0
}

func (s *AnnSearch) Sort() string {
	if s.SortField == "" {
		s.SortField = "created_at"
	}

	if s.SortDesc == "" {
		s.SortDesc = "asc"
	}
	return s.SortField + " " + s.SortDesc
}
