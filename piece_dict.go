package combiner

import (
    "github.com/ddliu/goption"
    "github.com/ddliu/go-dict"
    "regexp"
)

func NewPieceDict(m map[string]interface{}) Piece {
    o := goption.NewOption(m)

    pattern, _ := o.GetString("pattern")

    p := new(PieceDict)

    if pattern != "" {
        compiledPattern := regexp.MustCompile("^" + pattern + "$")
        p.Pattern = compiledPattern
    }

    path := o.MustGetString("dict")
    d := dict.NewSimpleDict()
    d.Load(path)

    p.Dict = d

    return p
}

type PieceDict struct {
    Dict *dict.SimpleDict
    Pattern *regexp.Regexp
}

func (this *PieceDict) Walk(f func(string) bool) {
    this.Dict.Walk(func(s string) bool {
        if this.Pattern == nil || this.Pattern.MatchString(s) {
            return f(s)
        }

        return true
    })
}

func (this *PieceDict) Count() uint64 {
    var total uint64 = 0
    this.Walk(func(string) bool {
        total++
        return true
    })

    return total
}