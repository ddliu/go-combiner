package combiner

import (
)

type Piece interface {
    Walk(func(string) bool)
    Count() uint64
}