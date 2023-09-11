package vast_test

import (
	"aqwari.net/xml/xmltree"
	"bytes"
	"errors"
	"github.com/google/go-cmp/cmp"
	"go.eigsys.de/go-vast"
	"io"
	"log"
	"os"
	"path"
	"testing"
)

func TestNumericBool_MarshalText_false(t *testing.T) {
	var numericBool vast.NumericBool

	result, err := numericBool.MarshalText()
	if err != nil {
		t.Error("unexpected error")
	}

	if !bytes.Equal(result, []byte("0")) {
		t.Error("unexpected result")
	}
}

func TestNumericBool_MarshalText_true(t *testing.T) {
	var numericBool vast.NumericBool
	numericBool = true

	result, err := numericBool.MarshalText()
	if err != nil {
		t.Error("unexpected error")
	}

	if !bytes.Equal(result, []byte("1")) {
		t.Error("unexpected result")
	}
}

func mustOpenFixture(fixture string) io.ReadCloser {
	handle, err := os.Open(path.Join("testdata", fixture))
	if err != nil {
		log.Fatalf("%v", err)
	}

	return handle
}

func TestNew(t *testing.T) {
	testVAST := vast.New()
	wantVAST := &vast.VAST{
		Version: "4.2",
		XMLNS:   "http://www.iab.com/VAST",
	}

	if diff := cmp.Diff(testVAST, wantVAST); diff != "" {
		t.Errorf("wrong VAST: %s", diff)
	}
}

type NopReadCloser struct{}

func (r NopReadCloser) Read(_ []byte) (n int, err error) {
	return 0, errors.New("error")
}

func (r NopReadCloser) Close() error {
	return nil
}

func TestRead_ErrReadVAST(t *testing.T) {
	reader := NopReadCloser{}
	if _, err := vast.Read(reader); !errors.Is(err, vast.ErrReadVAST) {
		t.Error("unexpected error")
	}
}

func TestRoundTrip(t *testing.T) {
	testCases := []string{
		"iab/Ad_Verification-test.xml",
		"iab/Category-test.xml",
		"iab/Closed_Caption_Test.xml",
		"iab/Event_Tracking-test.xml",
		"iab/IconClickFallbacks.xml",
		"iab/Inline_Companion_Tag-test.xml",
		"iab/Inline_Linear_Tag-test.xml",
		"iab/Inline_Non-Linear_Tag-test.xml",
		"iab/Inline_Simple.xml",
		"iab/No_Wrapper_Tag-test.xml",
		"iab/Ready_to_serve_Media_Files_check-test.xml",
		"iab/Universal_Ad_ID-multi-test.xml",
		"iab/Video_Clicks_and_click_tracking-Inline-test.xml",
		"iab/Viewable_Impression-test.xml",
		"iab/Wrapper_Tag-test.xml",
	}

	for _, testCase := range testCases {
		t.Run(testCase, func(t *testing.T) {
			reader := mustOpenFixture(testCase)

			testVAST, err := vast.Read(reader)
			if err != nil {
				t.Error("unexpected error")
			}

			output, err := testVAST.Bytes()
			if err != nil {
				t.Error("unexpected error")
			}

			outputDoc, err := xmltree.Parse(output)
			if err != nil {
				t.Error("unexpected error")
			}

			input, err := io.ReadAll(mustOpenFixture(testCase))
			if err != nil {
				t.Error("unexpected error")
			}

			inputDoc, err := xmltree.Parse(input)
			if err != nil {
				t.Error("unexpected error")
			}

			if !xmltree.Equal(inputDoc, outputDoc) {
				inputStr := xmltree.MarshalIndent(inputDoc, "", "  ")
				outputStr := xmltree.MarshalIndent(outputDoc, "", "  ")

				diff := cmp.Diff(inputStr, outputStr)
				t.Errorf("wrong VAST: %s", diff)
			}
		})
	}
}
