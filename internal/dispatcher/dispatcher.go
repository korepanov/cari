package dispatcher

import (
	"flag"
	"fmt"
	"os"

	"github.com/korepanov/cari/internal/myerrors"
	"github.com/korepanov/cari/internal/program"
)

type dispatcher struct {
	iFlag      string
	oFlag      string
	astFlag    bool
	inputFile  *os.File
	outputFile *os.File
	stdout     *os.File
	input      program.Program
}

func newDispatcher() (dispatcher, error) {
	var d dispatcher
	d.stdout = os.Stdout

	err := d.processFlags()
	if err != nil {
		return d, err
	}

	return d, nil
}

type compileErrorT struct {
	err error
}

func Compile() error {
	d, err := newDispatcher()
	defer d.Close()

	if err != nil {
		return err
	}

	compileErr := d.compile()
	return compileErr.err
}

func (d *dispatcher) compile() compileErrorT {

	err := d.prepareInputFile()
	if err != nil {
		return d.compileError(err)
	}

	err = d.input.ReadProgram()

	if err != nil {
		return d.compileError(err)
	}

	if d.astFlag {
		d.input.Ast.Print()
		return compileErrorT{}
	}

	err = d.prepareOutputFile()

	if err != nil {
		return d.compileError(err)
	}

	d.input.WriteProgram()

	return compileErrorT{}
}

func (d *dispatcher) Close() {
	d.inputFile.Close()
	d.outputFile.Close()
	os.Stdout = d.stdout
}

func (d *dispatcher) compileError(err error) compileErrorT {
	var res compileErrorT
	res.err = fmt.Errorf("%s : %s", myerrors.ErrCompile, err)
	return res
}

var help = flag.Bool("h", false, "show help")
var astFlag = flag.Bool("ast", false, "show ast")
var iFlag = "-i"
var oFlag = "-o"

func (d *dispatcher) processFlags() error {

	flag.StringVar(&iFlag, "i", "", "input file")
	flag.StringVar(&oFlag, "o", "", "output file")

	flag.Usage = func() {
		w := flag.CommandLine.Output()
		fmt.Fprintf(w, "Usage: \n\n%s < stdin\n%s -i <file>.cari -o <file>.s\n\n", os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if *help {
		flag.Usage()
		return myerrors.ErrHelp
	}

	d.iFlag = iFlag

	if oFlag == "" {
		oFlag = "a.s"
	}

	d.oFlag = oFlag
	d.astFlag = *astFlag

	return nil
}

func (d *dispatcher) prepareInputFile() error {
	var err error

	if d.iFlag != "" {
		d.inputFile, err = os.Open(d.iFlag)
		if err != nil {
			return fmt.Errorf("%s : %s", myerrors.ErrFlagParse, err)
		}
		os.Stdin = d.inputFile
	}

	d.inputFile = os.Stdin
	return nil
}

func (d *dispatcher) prepareOutputFile() error {
	var err error

	d.outputFile, err = os.Create(d.oFlag)
	if err != nil {
		return fmt.Errorf("%s : %s", myerrors.ErrFlagParse, err)
	}

	os.Stdout = d.outputFile
	return nil
}
