package combiner

import (
    "testing"
    "strings"
)

func TestChars(t *testing.T) {
    p := NewPieceChars(map[string]interface{}{
        "maxlength": 2,
        "minlength": 1,
        "chars": "abcd",
    })

    if p.Count() != 4*5 {
        t.Errorf("Count error, %d != %d", p.Count(), 4*5)
        p.Walk(func(s string) bool {
            t.Log(s)
            return true
        })
    }

    // regexp
    p = NewPieceChars(map[string]interface{}{
        "maxlength": 2,
        "minlength": 1,
        "chars": "abcd",
        "pattern": "[ab].*",
    })

    if p.Count() != 2*5 {
        t.Errorf("Count error")
    }

    // big
    p = NewPieceChars(map[string]interface{}{
        "maxlength": 5,
        "minlength": 5,
        "chars": "abcd",
    })

    if p.Count() != 4*4*4*4*4 {
        t.Errorf("Count error")
    }

    p = NewPieceChars(map[string]interface{}{
        "maxlength": 5,
        "minlength": 5,
        "chars": "abcd",
        "pattern": "b.*",
    })

    if p.Count() != 4*4*4*4 {
        t.Errorf("Count error")
    }
}

func TestChoices(t *testing.T) {
    p := NewPieceChoice(map[string]interface{} {
        "choices": []string {"choice1", "choice2", "choice3"},
    })

    if p.Count() != 3 {
        t.Errorf("Count error")
    }
}

func TestConst(t *testing.T) {
    p := NewPieceConst(map[string]interface{} {
        "const": "hello",
    })

    if p.Count() != 1 {
        t.Errorf("Count error")
    }
}

func TestDict(t *testing.T) {
    p := NewPieceDict(map[string]interface{} {
        "dict": "/usr/share/dict/words",
        "pattern": "(?i)z[a-z]*",
    })

    if p.Count() < 100 {
        t.Errorf("Count error")
    }

    p.Walk(func(s string) bool {
        if strings.ToLower(string(s[0])) != "z" {

            t.Errorf("Invalid result %s", s)
        }
        t.Log(s)

        return true
    })
}

func TestCombiner(t *testing.T) {
    c := NewCombiner()

    c.Add("const", map[string]interface{} {
        "const": "go_",
    })

    c.Add("chars", map[string]interface{} {
        "chars": "ab",
        "maxlength": 3,
    })

    c.Add("choice", map[string]interface{} {
        "choices": []string{"Mac", "Linux", "Windows"},
    })

    c.Add("dict", map[string]interface{} {
        "dict": "/usr/share/dict/words",
        "pattern": "[a-z]",
    })

    if c.Count() != 2*2*2*3*26 {
        t.Errorf("Count error")
    }

    c.Walk(func(s string) bool {
        t.Log(s)
        return true
    })
}