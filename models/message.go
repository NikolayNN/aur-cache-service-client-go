package models

import "time"

// MessageCacheEntry mirrors the Java record used by the original project.
type MessageCacheEntry struct {
	ID              *int64    `json:"id,omitempty"`
	Time            time.Time `json:"time"`
	Latitude        float32   `json:"latitude"`
	Longitude       float32   `json:"longitude"`
	Altitude        int       `json:"altitude"`
	Speed           int       `json:"speed"`
	AmountSatellite int       `json:"amountSatellite"`
	Course          int       `json:"course"`
	GpsOdometer     float64   `json:"gpsOdometer"`
	GsmStrength     int       `json:"gsmStrength"`
	Battery         int       `json:"battery"`
	BatteryVoltage  int       `json:"batteryVoltage"`
	OnboardVoltage  float32   `json:"onboardVoltage"`
	EcoAcceleration float32   `json:"ecoAcceleration"`
	EcoBraking      float32   `json:"ecoBraking"`
	EcoCornering    float32   `json:"ecoCornering"`
	EcoBump         float32   `json:"ecoBump"`
	Params          string    `json:"params"`
	UnitID          int64     `json:"unitId"`
	IsArchive       bool      `json:"isArchive"`
	Type            string    `json:"type"`
}

// LastMessageStateCacheEntry mirrors the Java record.
type LastMessageStateCacheEntry struct {
	Last               MessageCacheEntry `json:"last"`
	LastValidLatitude  float32           `json:"lastValidLatitude"`
	LastValidLongitude float32           `json:"lastValidLongitude"`
	LastValidAltitude  int               `json:"lastValidAltitude"`
	LastValidSpeed     int               `json:"lastValidSpeed"`
	Motion             string            `json:"motion"`
	LastMotionChange   time.Time         `json:"lastMotionChange"`
	LastUpdate         time.Time         `json:"lastUpdate"`
}
