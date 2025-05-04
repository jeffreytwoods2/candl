package main

import (
	"bytes"
	"fmt"
	"math"
	"net/http"
	"time"

	"candl.jwoods.dev/internal/data"
)

var sampleResponse = data.AggList{
	Ticker:       "AAPl",
	Adjusted:     true,
	QueryCount:   120,
	RequestID:    "6999877fdfb5ce15f2e8c124b78d34dc",
	ResultsCount: 120,
	Status:       "OK",
	Results: []data.Agg{
		{
			Close:               100.00,
			High:                300.00,
			Low:                 60.00,
			NumTxns:             38,
			Open:                200.00,
			Timestamp:           1742284980000,
			Volume:              415,
			VolumeWeightedPrice: 213.7445,
		},
		{
			Close:               50.00,
			High:                150.00,
			Low:                 45.00,
			NumTxns:             38,
			Open:                100.00,
			Timestamp:           1742285220000,
			Volume:              115,
			VolumeWeightedPrice: 213.7413,
		},
		{
			Close:               25.00,
			High:                75.00,
			Low:                 10.00,
			NumTxns:             38,
			Open:                50.00,
			Timestamp:           1742285340000,
			Volume:              415,
			VolumeWeightedPrice: 213.7445,
		},
		{
			Close:               40.00,
			High:                50.00,
			Low:                 20.00,
			NumTxns:             38,
			Open:                25.00,
			Timestamp:           1742285400000,
			Volume:              415,
			VolumeWeightedPrice: 213.7445,
		},
		{
			Close:               80.00,
			High:                120.00,
			Low:                 30.00,
			NumTxns:             38,
			Open:                40.00,
			Timestamp:           1742285820000,
			Volume:              415,
			VolumeWeightedPrice: 213.7445,
		},
	},
	Count:   120,
	NextURL: "https://api.polygon.io/v2/aggs/ticker/AAPL/range/1/minute/1742299860000/2025-03-18?cursor=bGltaXQ9MTIwJnNvcnQ9YXNj",
}

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, data templateData) {
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, r, err)
		return
	}

	buf := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	w.WriteHeader(status)

	buf.WriteTo(w)
}

func xAxisRangeFormat(ts int64) string {
	t := time.UnixMilli(ts)
	if t.IsZero() {
		return ""
	}

	return t.Local().Format("2006-01-02 15:04")
}

func xAxisTimeFormat(ts int64) string {
	t := time.UnixMilli(ts)
	if t.IsZero() {
		return ""
	}

	return t.UTC().Format("15:04")
}

func (app *application) newTemplateData(r *http.Request) templateData {
	maxPrice := 0.0
	minPrice := math.Inf(1)
	timestamps := []string{}
	closes := []float64{}
	highs := []float64{}
	lows := []float64{}
	opens := []float64{}

	for _, agg := range sampleResponse.Results {
		stringTS := xAxisRangeFormat(agg.Timestamp)

		timestamps = append(timestamps, stringTS)
		closes = append(closes, agg.Close)
		highs = append(highs, agg.High)
		lows = append(lows, agg.Low)
		opens = append(opens, agg.Open)

		if agg.High > maxPrice {
			maxPrice = agg.High
		}
		if agg.Low < minPrice {
			minPrice = agg.Low
		}
	}

	firstTimestamp := xAxisRangeFormat(sampleResponse.Results[0].Timestamp)
	lastTimestamp := xAxisRangeFormat(sampleResponse.Results[len(sampleResponse.Results)-1].Timestamp)

	return templateData{
		CurrentYear:    time.Now().Year(),
		Ticker:         "AAPL",
		Timestamps:     timestamps,
		Closes:         closes,
		Highs:          highs,
		Lows:           lows,
		Opens:          opens,
		MaxPrice:       maxPrice,
		MinPrice:       minPrice,
		FirstTimestamp: firstTimestamp,
		LastTimestamp:  lastTimestamp,
	}
}
