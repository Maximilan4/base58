package base58

import (
	"bytes"
	"crypto/rand"
	"strconv"
	"strings"
	"testing"
)

func prepareEncodeCases() [][]byte {
	cases := make([][]byte, 0, 11)
	cases = append(cases, []byte{})
	for power := 0; power <= 10; power++ {
		size := 2 << power
		data := make([]byte, size)
		rand.Read(data)
		cases = append(cases, data)
	}

	return cases
}

func BenchmarkEncodeBytes(b *testing.B) {
	cases := prepareEncodeCases()

	for _, tc := range cases {
		// b.ResetTimer()
		b.Run(strconv.Itoa(len(tc)), func(subB *testing.B) {
			dst := make([]byte, encodeLen(len(tc)))
			var (
				err error
			)
			subB.ResetTimer()
			for i := 0; i < subB.N; i++ {
				_, err = EncodeBytes(tc, dst)
				if err != nil {
					subB.Fatal(err)
				}
			}
		})
	}
}

func BenchmarkEncodeString(b *testing.B) {
	cases := prepareEncodeCases()
	for _, tc := range cases {
		b.Run(strconv.Itoa(len(tc)), func(subB *testing.B) {
			var (
				err error
				src = string(tc)
			)
			subB.ResetTimer()
			for i := 0; i < subB.N; i++ {
				_, err = EncodeString(src)
				if err != nil {
					subB.Fatal(err)
				}
			}
		})
	}

}

func TestEncodeBytes(t *testing.T) {
	type testCase struct {
		input  []byte
		output []byte
		err    bool
	}

	for _, tc := range []testCase{
		{input: []byte{}, output: []byte{}, err: false},
		{input: []byte("09"), output: []byte("4er"), err: false},
		{input: []byte("\\009"), output: []byte("3mF3Hg"), err: false},
		{input: []byte("\\0\\0\\0"), output: []byte("MUCrjctE"), err: false},
	} {
		t.Run(string(tc.input), func(t *testing.T) {
			// dst := make([]byte, encodeLen(len(tc.input)))
			dst := make([]byte, 20)
			var (
				err error
				n   int
			)
			if n, err = EncodeBytes(tc.input, dst); err != nil && !tc.err {
				t.Errorf("EncodeBytes(%v) got unexpected error: %v", tc.input, err)
			}

			if !bytes.EqualFold(dst[:n], tc.output) {
				t.Errorf("EncodeBytes(%v) got %v, want %v", tc.input, dst, tc.output)
			}
		})
	}
}

func TestEncodeString(t *testing.T) {
	type testCase struct {
		input  string
		output string
		err    bool
	}

	for _, tc := range []testCase{
		{input: "", output: "", err: false},
		{input: "09", output: "4er", err: false},
		{input: "\\009", output: "3mF3Hg", err: false},
		{input: "\\0\\0\\0", output: "MUCrjctE", err: false},
	} {
		t.Run(tc.input, func(t *testing.T) {
			var (
				dst string
				err error
			)

			if dst, err = EncodeString(tc.input); err != nil && !tc.err {
				t.Errorf("EncodeBytes(%v) got unexpected error: %v", tc.input, err)
			}

			if !strings.EqualFold(dst, tc.output) {
				t.Errorf("EncodeBytes(%v) got %v, want %v", tc.input, dst, tc.output)
			}
		})
	}
}
