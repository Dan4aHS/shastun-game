package time_rater

import "iter"

func DiceIterator() iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		for i := 1; i <= 6; i++ {
			for j := 1; j <= 6; j++ {
				if !yield(i, j) {
					return
				}
			}
		}
	}
}
