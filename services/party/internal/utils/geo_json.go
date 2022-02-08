package utils

type GeoJson struct {
	Type        string    `json:"-"           bson:"type,omitempty"`
	Coordinates []float64 `json:"coordinates,omitempty" bson:"coordinates,omitempty"`
}

func NewPoint(long, lat float64) GeoJson {
	return GeoJson{
		"Point",
		[]float64{long, lat},
	}
}
