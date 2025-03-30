// Copyright 2021-2024 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"

	"github.com/veraison/corim/corim"
	"github.com/veraison/corim/cots"
)

// displayOptions holds output and rendering configuration.
type DisplayOptions struct {
	Writer   io.Writer
	ShowTags bool
}

var (
	corimDisplayCorimFile *string
	corimDisplayShowTags  *bool
)

var corimDisplayCmd = NewCorimDisplayCmd()

func NewCorimDisplayCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "display",
		Short: "display the content of a CoRIM as JSON",
		Long: `display the content of a CoRIM as JSON

	Display the contents of the signed CoRIM signed-corim.cbor
	
	  cocli corim display --file signed-corim.cbor

	Display the contents of the signed CoRIM yet-another-signed-corim.cbor and
	also unpack any embedded CoMID, CoSWID and CoTS
	
	  cocli corim display --file yet-another-signed-corim.cbor --show-tags
	`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := checkCorimDisplayArgs(); err != nil {
				return err
			}

			// Prepare our display options
			opts := &DisplayOptions{
				Writer:   os.Stdout,
				ShowTags: *corimDisplayShowTags,
			}
			return display(*corimDisplayCorimFile, opts)
		},
	}

	corimDisplayCorimFile = cmd.Flags().StringP("file", "f", "", "a CoRIM file (in CBOR format)")
	corimDisplayShowTags = cmd.Flags().BoolP("show-tags", "v", false, "display embedded tags")

	return cmd
}

func checkCorimDisplayArgs() error {
	if corimDisplayCorimFile == nil || *corimDisplayCorimFile == "" {
		return errors.New("no CoRIM supplied")
	}
	return nil
}

// displaySignedCorim marshals the signed CoRIM to JSON and prints it.
func displaySignedCorim(s corim.SignedCorim, corimFile string, opts *DisplayOptions) error {
	metaJSON, err := json.MarshalIndent(&s.Meta, "", "  ")
	if err != nil {
		return fmt.Errorf("error encoding CoRIM Meta from %s: %w", corimFile, err)
	}

	fmt.Fprintln(opts.Writer, "Meta:")
	fmt.Fprintln(opts.Writer, string(metaJSON))

	corimJSON, err := json.MarshalIndent(&s.UnsignedCorim, "", "  ")
	if err != nil {
		return fmt.Errorf("error encoding unsigned CoRIM from %s: %w", corimFile, err)
	}

	fmt.Fprintln(opts.Writer, "CoRIM:")
	fmt.Fprintln(opts.Writer, string(corimJSON))

	if opts.ShowTags {
		fmt.Fprintln(opts.Writer, "Tags:")
		displayTags(s.UnsignedCorim.Tags, opts)
	}

	return nil
}

// displayUnsignedCorim marshals the unsigned CoRIM to JSON and prints it.
func displayUnsignedCorim(u corim.UnsignedCorim, corimFile string, opts *DisplayOptions) error {
	corimJSON, err := json.MarshalIndent(&u, "", "  ")
	if err != nil {
		return fmt.Errorf("error encoding unsigned CoRIM from %s: %w", corimFile, err)
	}

	fmt.Fprintln(opts.Writer, "Corim:")
	fmt.Fprintln(opts.Writer, string(corimJSON))

	if opts.ShowTags {
		fmt.Fprintln(opts.Writer, "Tags:")
		displayTags(u.Tags, opts)
	}

	return nil
}

// display reads the given file, decodes it as either a signed or unsigned CoRIM,
// and then calls the appropriate display function.
func display(corimFile string, opts *DisplayOptions) error {
	corimCBOR, err := afero.ReadFile(fs, corimFile)
	if err != nil {
		return fmt.Errorf("error loading CoRIM from %s: %w", corimFile, err)
	}

	// try to decode as a signed CoRIM
	var s corim.SignedCorim
	if err = s.FromCOSE(corimCBOR); err == nil {
		// successfully decoded as signed CoRIM
		return displaySignedCorim(s, corimFile, opts)
	}

	// if decoding as signed CoRIM failed, attempt to decode as unsigned CoRIM
	var u corim.UnsignedCorim
	if err = u.FromCBOR(corimCBOR); err != nil {
		return fmt.Errorf("error decoding CoRIM (signed or unsigned) from %s: %w", corimFile, err)
	}
	// successfully decoded as unsigned CoRIM
	return displayUnsignedCorim(u, corimFile, opts)
}

// displayTags loops over each tag, identifies its type, and prints the JSON.
func displayTags(tags []corim.Tag, opts *DisplayOptions) {
	for i, t := range tags {
		if len(t) < 4 {
			fmt.Fprintf(opts.Writer, ">> skipping malformed tag at index %d\n", i)
			continue
		}

		cborTag, cborData := t[:3], t[3:]
		hdr := fmt.Sprintf(">> [ %d ]", i)

		switch {
		case bytes.Equal(cborTag, corim.ComidTag):
			if err := printComid(cborData, hdr); err != nil {
				fmt.Fprintf(opts.Writer, ">> skipping malformed CoMID tag at index %d: %v\n", i, err)
			}
		case bytes.Equal(cborTag, corim.CoswidTag):
			if err := printCoswid(cborData, hdr); err != nil {
				fmt.Fprintf(opts.Writer, ">> skipping malformed CoSWID tag at index %d: %v\n", i, err)
			}
		case bytes.Equal(cborTag, cots.CotsTag):
			if err := printCots(cborData, hdr); err != nil {
				fmt.Fprintf(opts.Writer, ">> skipping malformed CoTS tag at index %d: %v\n", i, err)
			}
		default:
			fmt.Fprintf(opts.Writer, ">> unmatched CBOR tag: %x\n", cborTag)
		}
	}
}

func init() {
	corimCmd.AddCommand(corimDisplayCmd)
}
