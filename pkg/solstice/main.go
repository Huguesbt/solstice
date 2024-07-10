package solstice

import (
	"github.com/soniakeys/meeus/v3/julian"
	"github.com/soniakeys/meeus/v3/solstice"
	"time"
)

func Spring(year int) time.Time {
	return julian.JDToTime(solstice.March(year))
}

func Summer(year int) time.Time {
	return julian.JDToTime(solstice.June(year))
}

func Fall(year int) time.Time {
	return julian.JDToTime(solstice.September(year))
}

func Winter(year int) time.Time {
	return julian.JDToTime(solstice.December(year))
}
