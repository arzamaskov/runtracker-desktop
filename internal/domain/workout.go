// Package domain defines the core domain concepts used by Runtracker.
package domain

type WorkoutType string

const (
	WorkoutTypeRecovery           WorkoutType = "recovery"
	WorkoutTypeLong               WorkoutType = "long"
	WorkoutTypeThresholdIntervals WorkoutType = "threshold_intervals"
	WorkoutTypeFartlek            WorkoutType = "fartlek"
	WorkoutTypeTempo              WorkoutType = "tempo"
	WorkoutTypeProgression        WorkoutType = "progression"
)

type HeartRateZone uint8

const (
	HeartRateZone1 HeartRateZone = 1
	HeartRateZone2 HeartRateZone = 2
	HeartRateZone3 HeartRateZone = 3
	HeartRateZone4 HeartRateZone = 4
	HeartRateZone5 HeartRateZone = 5
)

type HeartRateZoneRange struct {
	From HeartRateZone
	To   HeartRateZone
}
