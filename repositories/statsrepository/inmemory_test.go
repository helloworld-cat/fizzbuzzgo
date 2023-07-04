package statsrepository

import (
	"sync"
	"testing"
)

func TestInMemoryRepository(t *testing.T) {
	repo := NewStatsRepositoryInMemory()

	na := 3
	nb := 5
	wordA := "foo"
	wordB := "bar"

	n := 100_000
	count := 5 // number of goroutines

	var wg sync.WaitGroup
	for i := 1; i <= count; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < n; i++ {
				_, err := repo.Incr(na, nb, wordA, wordB)
				if err != nil {
					t.Fatalf("unexpected incr. error: %s", err)
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()

	v, err := repo.Fetch(na, nb, wordA, wordB)
	if err != nil {
		t.Fatalf("unexpected fetch error: %s", err)
	}

	if v != (count * n) {
		t.Errorf("unexpected value. Want: %d, but got: %d", count*n, v)
	}
}
