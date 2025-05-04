package data

type AggList struct {
	Ticker       string `json:"ticker"`          // The exchange symbol that this item is traded under.
	Adjusted     bool   `json:"adjusted"`        // Whether or not this response was adjusted for splits.
	QueryCount   int    `json:"queryCount"`      // The number of aggregates (minute or day) used to generate the response.
	RequestID    string `json:"request_id"`      // A request id assigned by the server.
	ResultsCount int    `json:"resultsCount"`    // The total number of results for this request.
	Status       string `json:"status"`          // The status of this request's response.
	Results      []Agg  `json:"results"`         // An array of results containing the requested data.
	Count        int    `json:"count,omitempty"` // This appears in the response but has no documentation; may be deprecated
	NextURL      string `json:"next_url"`        // If present, this value can be used to fetch the next page of data.
}

type Agg struct {
	Close               float64 `json:"c"`             // The close price for the symbol in the given time period.
	High                float64 `json:"h"`             // The highest price for the symbol in the given time period.
	Low                 float64 `json:"l"`             // The lowest price for the symbol in the given time period.
	NumTxns             int     `json:"n"`             // The number of transactions in the aggregate window.
	Open                float64 `json:"o"`             // The open price for the symbol in the given time period.
	OTC                 bool    `json:"otc,omitempty"` // Whether or not this aggregate is for an OTC ticker. This field will be left off if false.
	Timestamp           int64   `json:"t"`             // The Unix millisecond timestamp for the start of the aggregate window.
	Volume              int     `json:"v"`             // The trading volume of the symbol in the given time period.
	VolumeWeightedPrice float64 `json:"vw"`            // The volume weighted average price.
}
