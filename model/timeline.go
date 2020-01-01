package model

type Timeline struct {
	TimelineObjects []TimelineObjects `json:"timelineObjects"`
}

type TimelineObjects struct {
	ActivitySegment `json:"activitySegment"`
}

type ActivitySegment struct {
	StartLocation `json:"startLocation"`
	EndLocation   `json:"endLocation"`
	Duration      `json:"duration"`
	Distance      int         `json:"distance"`
	ActivityType  string      `json:"activityType"`
	Confidence    string      `json:"confidence"`
	Activities    []Activitie `json:"activities"`
	WaypointPath  `json:"waypointPath"`
}

type StartLocation struct {
	LatitudeE7  int `json:"latitudeE7"`
	LongitudeE7 int `json:"longitudeE7"`
}

type EndLocation struct {
	LatitudeE7  int `json:"latitudeE7"`
	LongitudeE7 int `json:"longitudeE7"`
}

type Duration struct {
	StartTimestampMs string `json:"startTimestampMs"`
	EndTimestampMs   string `json:"endTimestampMs"`
}

type Activitie struct {
	ActivityType string  `json:"activityType"`
	Probability  float64 `json:"probability"`
}

type WaypointPath struct {
	Waypoints []Waypoint `json:"waypoints"`
}

type Waypoint struct {
	LatE7 int `json:"latE7"`
	LngE7 int `json:"lngE7"`
}
