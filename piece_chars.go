package combiner

import (
    "github.com/ddliu/goption"
    "regexp"
    "fmt"
    "strings"
    "math"
)

type PieceCharsOptions struct {
    MaxLength int
    MinLength int
    Pattern *regexp.Regexp
    Chars []string
}

// Options
// - chars
// - maxlength
// - minlength
// - pattern
func NewPieceChars(m map[string]interface{}) Piece {
    o := goption.NewOption(m)
    
    // validation...
    
    maxlength, ok := o.GetInt("maxlength")

    if !ok || maxlength < 1 {
        panic(fmt.Errorf(`Invalid "maxlength"`))
    }

    minlength, ok := o.GetInt("minlength")

    if !ok {
        minlength = maxlength
    } else {
        if minlength < 1 {
            panic(fmt.Errorf(`Invalid "minlength"`))
        }
    }
    
    chars := strings.Split(o.MustGetString("chars"), "")

    if len(chars) == 0 {
        panic(fmt.Errorf(`Invalid "chars"`))
    }

    pattern, ok := o.GetString("pattern")   

    compiledPattern, err := regexp.Compile("^" + pattern + "$")

    if err != nil{
        panic(fmt.Errorf(`Invalid "pattern"`))
    }

    po := new(PieceCharsOptions)
    po.MaxLength = maxlength
    po.MinLength = minlength
    if pattern != "" {
        po.Pattern = compiledPattern
    }
    po.Chars = chars

    return &PieceChars{po}
}

type PieceChars struct {
    options *PieceCharsOptions
}

func (this *PieceChars) Walk(callback func(string) bool) {
    var f func(string, int, int) bool
    f = func(prefix string, i int, length int) bool {
        for _, v := range this.options.Chars {
            next := prefix + v
            // last
            if i >= length - 1 {
                if this.options.Pattern == nil || this.options.Pattern.MatchString(next) {
                    if !callback(next) {
                        return false
                    }
                }
            } else {
                if !f(next, i + 1, length) {
                    return false
                }
            }
        }

        return true
    }
    for i := this.options.MinLength; i <= this.options.MaxLength; i++ {
        if !f("", 0, i) {
            break
        }
    }
}

func (this *PieceChars) Count() uint64 {
    var total uint64 = 0
    if this.options.Pattern == nil {
        length := len(this.options.Chars)
        for i := this.options.MinLength; i <= this.options.MaxLength; i++ {
            total += uint64(math.Pow(float64(length), float64(i)))
            // fmt.Println(length, i, total)
        }

        return total
    }
    

    this.Walk(func(string) bool {
        total++
        return true
    })

    return total
}