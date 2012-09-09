package hershey

import (
    "testing"
    "reflect"
)

type ReadSymTest struct {
    in string
    sym Sym
}

var readSymTests = []ReadSymTest{
    {
        "    2 16MWOMOV ROMSMUNUPSQ ROQSQURUUSVOV",
        Sym{2, 16, -5, 5, []coord{
            {-3,-5}, {-3,4}, {-50,0}, {-3,-5}, {1,-5}, {3,-4}, {3,-2},
            {1,-1}, {-50,0}, {-3,-1}, {1,-1}, {3,0}, {3,3}, {1,4}, {-3,4},
        }},
    },
    {
        "    8  9MWOMOV RUMUV ROQUQ",
        Sym{8, 9, -5, 5, []coord{
            {-3,-5}, {-3,4}, {-50,0}, {3,-5}, {3,4}, {-50,0}, {-3,-1}, {3,-1},
        }},
    },
    {
        "  223  3PTRLRW",
        Sym{223, 3, -2, 2, []coord{{0,-6}, {0,5},}},
    },
    {
        " 2123 24F^JMN[ RKMNX RRMN[ RRMV[ RSMVX RZMV[ RGMNM RWM]M",
        Sym{2123, 24, -12, 12, []coord{
            {-8,-5}, {-4,9}, {-50,0}, {-7,-5}, {-4,6}, {-50,0}, {0,-5},
            {-4,9}, {-50,0}, {0,-5}, {4,9}, {-50,0}, {1,-5}, {4,6},
            {-50,0}, {8,-5}, {4,9}, {-50,0}, {-11,-5}, {-4,-5}, {-50,0},
            {5,-5}, {11,-5},
        }},
    },
}

func TestInput(t *testing.T) {
    for _, test := range readSymTests {
        sym, err := ReadSym(test.in)
        if err != nil {
            t.Error(err)
        } else {
            if !reflect.DeepEqual(test.sym, *sym) {
                t.Errorf("scanning %s: expected %+v, got %+v", test.in, test.sym, sym)
            }
        }
    }
}
