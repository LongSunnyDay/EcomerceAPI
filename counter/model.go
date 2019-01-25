package counter

const (
	collectionName = "counters"
	quoteCounter   = "quote"
)

type (
	ItemCounter struct {
		ID    string `json:"_id,omitempty" bson:"_id"`
		Value int    `json:"value,omitempty" bson:"value"`
	}

	QuoteCounter struct {
		ID    int64  `json:"id"`
		Name  string `json:"name"`
		Value int64  `json:"value"`
	}
)
