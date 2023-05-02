package csvs2

import (
	"math"

	"github.com/golang/geo/s2"
)

const earthRadius = 6371000

type DataPoint struct {
	Lat  float64
	Lng  float64
	S2id s2.Cell
}

func NewPointFromLL(lat, lng float64) *DataPoint {
	ll := s2.LatLngFromDegrees(lat, lng)
	return &DataPoint{
		Lat:  lat,
		Lng:  lng,
		S2id: s2.CellFromLatLng(ll),
	}
}

func (d *DataPoint) DistanceFrom(dp DataPoint) float64 {
	dist := haversineDistance(d.Lat, d.Lng, dp.Lat, dp.Lng)
	return dist
}

func haversineDistance(lat1, lon1, lat2, lon2 float64) float64 {
	lat1, lon1, lat2, lon2 = degreesToRadians(lat1, lon1, lat2, lon2)

	dlon := lon2 - lon1
	dlat := lat2 - lat1

	a := math.Pow(math.Sin(dlat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dlon/2), 2)
	c := 2 * math.Asin(math.Sqrt(a))

	return earthRadius * c
}

func degreesToRadians(lat1, lon1, lat2, lon2 float64) (float64, float64, float64, float64) {
	return lat1 * math.Pi / 180, lon1 * math.Pi / 180, lat2 * math.Pi / 180, lon2 * math.Pi / 180
}
