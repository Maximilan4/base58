package base58

import (
	"bytes"
	"crypto/rand"
	"strconv"
	"testing"
)

func prepareDecodeCases() [][]byte {
	cases := make([][]byte, 0, 11)
	cases = append(cases, []byte{})
	for power := 0; power <= 10; power++ {
		size := 2 << power
		data := make([]byte, size)
		rand.Read(data)
		trg := make([]byte, encodeLen(size))
		n, _ := EncodeBytes(data, trg)

		cases = append(cases, trg[:n])
	}

	return cases
}

func BenchmarkDecodeBytes(b *testing.B) {
	cases := prepareDecodeCases()

	for _, tc := range cases {
		// b.ResetTimer()
		b.Run(strconv.Itoa(len(tc)), func(subB *testing.B) {
			dst := make([]byte, decodeLen(len(tc)))
			var (
				err error
			)
			subB.ResetTimer()
			for i := 0; i < subB.N; i++ {
				_, err = DecodeBytes(tc, dst)
				if err != nil {
					subB.Fatal(err)
				}
			}
		})
	}
}

func BenchmarkDecodeString(b *testing.B) {
	cases := prepareDecodeCases()
	for _, tc := range cases {
		b.Run(strconv.Itoa(len(tc)), func(subB *testing.B) {
			var (
				err error
				src = string(tc)
			)
			subB.ResetTimer()
			for i := 0; i < subB.N; i++ {
				_, err = DecodeString(src)
				if err != nil {
					subB.Fatal(err)
				}
			}
		})
	}

}

func TestDecodeBytes(t *testing.T) {
	type testCase struct {
		input  []byte
		output []byte
		err    bool
	}

	for _, tc := range []testCase{
		{input: []byte{}, output: nil, err: false},
		{input: []byte("4ER"), output: []byte("09"), err: false},
		{input: []byte("3mF3Hg"), output: []byte(`\009`), err: false},
		{input: []byte("MUCrjctE"), output: []byte(`\0\0\0`), err: false},
	} {
		t.Run(string(tc.input), func(t *testing.T) {
			dst := make([]byte, 20)

			var (
				n   int
				err error
			)
			if n, err = DecodeBytes(tc.input, dst); err != nil && !tc.err {
				t.Errorf("EncodeBytes(%v) got unexpected error: %v", tc.input, err)
			}

			if !bytes.EqualFold(dst[:n], tc.output) {
				t.Errorf("EncodeBytes(%v) got %v, want %v", tc.input, dst, tc.output)
			}
		})
	}
}

func TestDecodeString(t *testing.T) {
	type testCase struct {
		input  string
		output string
		err    bool
	}

	for _, tc := range []testCase{
		{input: "", output: "", err: false},
		{input: "4ER", output: "09", err: false},
		{input: "3mF3Hg", output: `\009`, err: false},
		{input: "MUCrjctE", output: `\0\0\0`, err: false},
	} {
		t.Run(tc.input, func(t *testing.T) {
			var (
				dst string
				err error
			)

			if dst, err = DecodeString(tc.input); err != nil && !tc.err {
				t.Errorf("EncodeBytes(%v) got unexpected error: %v", tc.input, err)
			}

			if dst != tc.output {
				t.Errorf("EncodeBytes(%v) got %v, want %v", tc.input, dst, tc.output)
			}
		})
	}
}
