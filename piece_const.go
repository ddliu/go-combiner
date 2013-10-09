package combiner

import (
    "github.com/ddliu/goption"
    "fmt"
)

func NewPieceConst(m map[string]interface{}) Piece {
    o := goption.NewOption(m)
    c := o.MustGetString("const")
    if c == "" {
        panic(fmt.Errorf("Invalid const"))
    }

    return &PieceConst{c}
}

type PieceConst struct {
    Const string
}

func (this *PieceConst) Walk(f func(string) bool) {
    f(this.Const)
}

func (this *PieceConst) Count() uint64 {
    return 1
}