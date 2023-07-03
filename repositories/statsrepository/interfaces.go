package statsrepository

type (
	StatsRepository interface {
		// Incr increments stats from parameters,
		// returns current stats about parameters.
		Incr(numberA, numberB int, wordA, wordB string) (int, error)

		// Fetch returns current stats from parameters.
		Fetch(numberA, numberB int, wordA, wordB string) (int, error)
	}
)
