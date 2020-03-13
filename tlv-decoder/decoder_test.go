package decoder_test

import (
	decoder "falabella-tlv/tlv-decoder"
	"testing"
)

func TestDecode(t *testing.T) {

	testCases := []struct {
		name string
		data []byte
		want map[string]string
		err  error
	}{
		{
			name: "should return an error if the input is empty",
			data: []byte{},
			want: map[string]string{},
			err:  decoder.ErrEmptyInput{},
		},
		{
			name: "should return an error if the input has the wrong format length",
			data: []byte("A0511AB3"),
			want: map[string]string{},
			err:  decoder.ErrInvalidFormat{},
		},
		{
			name: "should return an error if the input has the wrong format type",
			data: []byte("N0503AB3"),
			want: map[string]string{},
			err:  decoder.ErrInvalidFormat{},
		},
		{
			name: "should return an error if the length is not an integer",
			data: []byte("N05FFAB3"),
			want: map[string]string{},
			err:  decoder.ErrInvalidFormat{},
		},
		{
			name: "should return an error if the data does not have the minimum lenght",
			data: []byte("N051"),
			want: map[string]string{},
			err:  decoder.ErrInvalidFormat{},
		},
		{
			name: "should return the map with the values",
			data: []byte("A0511AB398765UJ1N230200"),
			want: map[string]string{
				"0": "A0511AB398765UJ1",
				"1": "N230200",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := decoder.Decode(tc.data)
			for k, v := range res {
				if v != tc.want[k] {
					t.Errorf("for %v got %v, expected %v", k, v, tc.want[k])
				}
			}
			if err != tc.err {
				t.Errorf("got %T, expected %T", err, tc.err)
			}
		})
	}
}
