package domains

import "time"

type NasaAPOD struct {
	Copyright      string `json:"copyright"`
	Title          string `json:"title"`
	Explanation    string `json:"explanation"`
	MediaType      string `json:"media_type"`
	URL            string `json:"url"`
	HdUrl          string `json:"hdurl"`
	ServiceVersion string `json:"service_version"`
	Date           string `json:"date"`
}

type APOD struct {
	ID int `json:"id"`
	NasaAPOD
	CreatedAt time.Time `json:"created_at"`
}
