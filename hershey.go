package hershey

import (
    "fmt"
    "strings"
)

type coord struct {
    x, y int
}

type Sym struct {
    id int
    numVert int
    leftPos, rightPos int
    coords []coord
}

func ReadSym(line string) (*Sym, error) {
    var (
        n int
        b, x, y byte
        err error
    )
    sym := new(Sym)
    r := strings.NewReader(line)
    if n, err = fmt.Fscanf(r, " %4d", &sym.id); n != 1 || err != nil {
        return nil, fmt.Errorf("couldn't scan id: %s", err)
    }
    if n, err = fmt.Fscanf(r, " %2d", &sym.numVert); n != 1 || err != nil {
        return nil, fmt.Errorf("couldn't scan numVert: %s", err)
    }
    if b, err = r.ReadByte(); err != nil {
        return nil, fmt.Errorf("couldn't scan leftPos: %s", err)
    }
    sym.leftPos = int(b) - int('R')
    if b, err = r.ReadByte(); err != nil {
        return nil, fmt.Errorf("couldn't scan rightPos: %s", err)
    }
    sym.rightPos = int(b) - int('R')
    for i := 0; i < sym.numVert-1; i++ {
        if x, err = r.ReadByte(); err != nil {
            return nil, fmt.Errorf("couldn't scan %d coord x: %s", i, err)
        }
        if y, err = r.ReadByte(); err != nil {
            return nil, fmt.Errorf("couldn't scan %d coord y: %s", i, err)
        }
        sym.coords = append(sym.coords, coord{int(x)-int('R'), int(y)-int('R')})
    }
    return sym, nil
}
