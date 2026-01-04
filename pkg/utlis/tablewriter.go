package utlis

import (
	"io"
	"os"

	"github.com/olekukonko/tablewriter"
)

func NewTableWriterDefault(header []string) *tablewriter.Table {
	return NewTableWriter(os.Stdout, header)
}

func NewTableWriter(w io.Writer, header []string) *tablewriter.Table {
	t := tablewriter.NewWriter(w)
	t.Header(header)

	return t
}
