package combiner

import (
    "github.com/ddliu/goption"
)

func NewPieceChoice(m map[string]interface{}) Piece {
    o := goption.NewOption(m)
    op := new(PieceChoiceOptions)
    choices := o.MustGet("choices")

    op.Choices = convertChoices(choices)

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

func convertChoices(choices interface{}) []string {
    if choicesArr, ok := choices.([]string); ok {
        return choicesArr
    } else if choicesArr, ok := choices.([]interface{}); ok {
        result := make([]string, len(choicesArr))

        for i := 0; i < len(choicesArr); i++ {
            if v, ok := choicesArr[i].(string); ok {
                result[i] = v
            } else {
                panic("Invalid choices")
            }
        }
        return result
    }
    
    panic("Invalid choices")
}