package combiner

import (
       
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
    var f func(string, int)

    f = func(prefix string, i int) {
        p := this.pieces[i]
        p.Walk(func(s string) bool {
            // last piece
            if i >= length - 1 {
                return callback(prefix + s)
            } else {
                f(prefix + s, i + 1)
            }
            return true
        })
    }

    f("", 0)
}