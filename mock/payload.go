package mock

type ViewEvent struct {
	EventType string `json:"event_type"`
	UserID    string `json:"user_id"`
	CampignID string `json:"campign_id"`
	SessionID string `json:"session_id"`
	AdID      string `json:"ad_id"`
	Timestamp string `json:"timestamp"`
	PageURL   string `json:"page_url"`
	Device    string `json:"device"`
	Cpv       string `json:"cpv"`
}

type ClickEvent struct {
	EventType string `json:"event_type"`
	UserID    string `json:"user_id"`
	SessionID string `json:"session_id"`
	AdID      string `json:"ad_id"`
	Timestamp string `json:"timestamp"`
	ElementID string `json:"element_id"`
	PageURL   string `json:"page_url"`
	Device    string `json:"device"`
}

type ConversionEvent struct {
	EventType      string  `json:"event_type"`
	UserID         string  `json:"user_id"`
	SessionID      string  `json:"session_id"`
	ConversionID   string  `json:"conversion_id"`
	Timestamp      string  `json:"timestamp"`
	ConversionType string  `json:"conversion_type"`
	Value          float64 `json:"value"`
	Currency       string  `json:"currency"`
	Device         string  `json:"device"`
}
