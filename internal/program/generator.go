package program

import (
	"fmt"

	"github.com/korepanov/cari/internal/sysinfo"
)

func (p *Program) makeComment() {
	fmt.Printf("# This code was made by %s version %s\n", sysinfo.Name, sysinfo.Version)
}

func (p *Program) makeData() {

}
