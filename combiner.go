package combiner

import (
    "math/rand"
)

func NewCombiner() (*Combiner) {
    return &Combiner{}
}

type Combiner struct {
    pieces []Piece
}

func GetPiece(name string, options map[string]interface{}) Piece {
    pieces := map[string]func(map[string]interface{}) Piece {
        "chars": NewPieceChars,
        "choice": NewPieceChoice,
        "const": NewPieceConst,
        "dict": NewPieceDict,
    }

    f := pieces[name]

    return f(options)
}

func (this *Combiner) Add(name string, options map[string]interface{}) {
    this.AddPiece(GetPiece(name, options))
}

func (this *Combiner) AddPiece(p Piece) {
    this.pieces = append(this.pieces, p)
}

func (this *Combiner) Count() uint64 {
    var total uint64 = 1
    for _, v := range this.pieces {
        total *= v.Count()
        if total == 0 {
            return 0
        }
    }

    return total
}

func (this *Combiner) Walk(callback func(string) bool) {
    length := len(this.pieces)
    var f func(string, int) bool

    f = func(prefix string, i int) bool {
        p := this.pieces[i]
        result := true
        p.Walk(func(s string) bool {
            // last piece
            if i >= length - 1 {
                r := callback(prefix + s)
                if !r {
                    result = false
                }
                return r
            } else {
                r := f(prefix + s, i + 1)
                if !r {
                    result = false
                }
                return r
            }
            return true
        })

        return result
    }

    f("", 0)
}

func (this *Combiner) RandWalk(callback func(string) bool) {
    var l []string
    this.Walk(func(s string) bool {
        l = append(l, s)
        return true
    })


    total := len(l)

    r := make([]float64, total)

    for i := 0; i < total; i++ {
        r[i] = rand.Float64()
    }

    for i := 0; i < total; i++ {
        for j := i; j < total; j++ {
            if r[i] < r[j] {
                r[i], r[j] = r[j], r[i]
                l[i], l[j] = l[j], l[i]
            }
        }
    }

    for i := 0; i < total; i++ {
        if !callback(l[i]) {
            break
        }
    }
}