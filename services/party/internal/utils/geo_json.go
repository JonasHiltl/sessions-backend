package utils

type GeoJson struct {
	Type        string    `json:"-"           bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}

func NewPoint(long, lat float64) GeoJson {
	return GeoJson{
		"Point",
		[]float64{long, lat},
	}
}
