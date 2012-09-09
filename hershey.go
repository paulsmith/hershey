package hershey

import (
	"fmt"
	"strings"
)

type coord struct {
	x, y int
}

type Sym struct {
	id, leftPos, rightPos int
	coords                []coord
}

func ReadSym(line string) (*Sym, error) {
	var (
		n, numVert int
		x, y       byte
		err        error
	)
	sym := new(Sym)
	r := strings.NewReader(line)
	n, err = fmt.Fscanf(r, " %4d %2d%c%c", &sym.id, &numVert, &x, &y)
	if n != 4 {
		return nil, fmt.Errorf("wanted to scan %d items, got %d", 4, n)
	}
	if err != nil {
		return nil, fmt.Errorf("scanning symbol: %s", err)
	}
	sym.leftPos = int(x) - int('R')
	sym.rightPos = int(y) - int('R')
	for i := 0; i < numVert-1; i++ {
		n, err = fmt.Fscanf(r, "%c%c", &x, &y)
		if n != 2 {
			return nil, fmt.Errorf("wanted to scan %d items, got %d", 2, n)
		}
		if err != nil {
			return nil, fmt.Errorf("scanning coords in symbol: %s", err)
		}
		sym.coords = append(sym.coords, coord{int(x) - int('R'), int(y) - int('R')})
	}
	return sym, nil
}
