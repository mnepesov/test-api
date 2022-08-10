package nasa

import "time"

type APODResponse struct {
	Copyright      string    `json:"copyright"`
	Title          string    `json:"title"`
	Explanation    string    `json:"explanation"`
	MediaType      string    `json:"media_type"`
	URL            string    `json:"url"`
	HdUrl          string    `json:"hdurl"`
	ServiceVersion string    `json:"service_version"`
	Date           time.Time `json:"date"`
}
