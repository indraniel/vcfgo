package vcfgo

import (
	"fmt"
	"io"
	"strings"
)

// Writer allows writing VCF files.
type Writer struct {
	io.Writer
	Header *Header
}

// NewWriter returns a writer after writing the header.
func NewWriter(w io.Writer, h *Header) (*Writer, error) {
	fmt.Fprintf(w, "##fileformat=VCFv%s\n", h.FileFormat)

	for _, line := range h.Extras {
		fmt.Fprintf(w, "%s\n", line)
	}

	for _, imap := range h.Contigs {
		fmt.Fprintf(w, "##contig=<ID=%s", imap["ID"])

		for k, v := range imap {
			if k == "ID" {
				continue
			}

			fmt.Fprintf(w, ",%s=%s", k, v)
		}
		fmt.Fprintln(w, ">")
	}

	for sampleId := range h.Samples {
		fmt.Fprintln(w, h.Samples[sampleId])
	}

	for i := range h.Pedigrees {
		fmt.Fprintln(w, h.Pedigrees[i])
	}

	for k, v := range h.Filters {
		fmt.Fprintf(w, "##FILTER=<ID=%s,Description=\"%s\">\n", k, v)
	}

	for _, i := range h.Infos {
		fmt.Fprintf(w, "%s\n", i)
	}

	for _, i := range h.SampleFormats {
		fmt.Fprintf(w, "%s\n", i)
	}

	fmt.Fprint(w, "#CHROM\tPOS\tID\tREF\tALT\tQUAL\tFILTER\tINFO\tFORMAT")
	var s string
	if len(h.SampleNames) > 0 {
		s = "\t" + strings.Join(h.SampleNames, "\t")
	}

	fmt.Fprint(w, s+"\n")
	return &Writer{w, h}, nil
}

// WriteVariant writes a single variant
func (w *Writer) WriteVariant(v *Variant) {
	fmt.Fprintln(w, v)
}
