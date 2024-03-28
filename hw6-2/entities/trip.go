package entities

type Trips struct {
	All []TripInfo `json:"info"`
}

type TripInfo struct {
	FromTo
	Path []RoutePath `json:"path"`
}

func (r *Trips) GetTripInfo(destination string) *TripInfo {
	for _, tripInfo := range r.All {
		if tripInfo.To == destination {
			return &tripInfo
		}
	}
	return nil
}

type FromTo struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type RoutePath struct {
	FromTo
	TransportType string `json:"transport_type"`
}
