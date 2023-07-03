package statsrepository

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

func NewStatsRepositoryInMemory() StatsRepository {
	return &repo{
		data: make(map[string]int),
		mu:   &sync.Mutex{},
	}
}

type (
	repo struct {
		data map[string]int
		mu   *sync.Mutex
	}
)

func (r *repo) Incr(na, nb int, wa, wb string) (int, error) {
	key := r.prepareKey(na, nb, wa, wb)

	r.mu.Lock()
	defer r.mu.Unlock()

	v, _ := r.data[key] // nothing, by default 'v' equals zero

	r.data[key] += 1

	return v, nil
}

func (r *repo) Fetch(na, nb int, wa, wb string) (int, error) {
	key := r.prepareKey(na, nb, wa, wb)

	r.mu.Lock()
	defer r.mu.Unlock()

	v, exists := r.data[key]
	if !exists {
		return 0, nil
	}

	return v, nil
}

func (r *repo) prepareKey(na, nb int, wa, wb string) string {
	hash := sha256.New() // TODO: 'hash' should be a dependency of 'repo'
	hash.Write([]byte(fmt.Sprintf("%d%d%s%s", na, nb, wa, wb)))
	s := hash.Sum(nil)
	return fmt.Sprintf("%x", s)
}
