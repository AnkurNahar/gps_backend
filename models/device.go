package models

type Device struct {
    DeviceID   string  `json:"device_id"`
    DisplayName string  `json:"display_name"`
    ActiveState string  `json:"active_state"`
    Lat         float64 `json:"lat"`
    Lng         float64 `json:"lng"`
    Speed       float64 `json:"speed"`
    DriveStatus string  `json:"drive_status"`
}