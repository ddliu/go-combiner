package combiner

import (
    "github.com/ddliu/goption"
)

func NewPieceChoice(m map[string]interface{}) Piece {
    o := goption.NewOption(m)
    op := new(PieceChoiceOptions)
    choices := o.MustGet("choices")
    choicesArr, ok := choices.([]string)

    if !ok {
        panic("Invalid choices")
    }

    op.Choices = choicesArr

    return &PieceChoice{op}
}

type PieceChoiceOptions struct {
    Choices []string
}

type PieceChoice struct {
    options *PieceChoiceOptions
}

func (this *PieceChoice) Walk(f func(string) bool) {
    for _, v := range this.options.Choices {
        if !f(v) {
            break
        }
    }
}

func (this *PieceChoice) Count() uint64 {
    return uint64(len(this.options.Choices))
}