package status

type Status string

const (
	PENDING   Status = "PENDING"
	RUNNING          = "RUNNING"
	PAUSED           = "PAUSED"
	COMPLETED        = "COMPLETED"
	CANCELLED        = "CANCELLED"
)
