package src

import "testing"

func TestDecodeBits(t *testing.T) {
	tests := []struct {
		bits string
		want string
	}{
		{
			bits: "111",
			want: "E",
		},
		{
			bits: "101",
			want: "I",
		},
		{
			bits: "10001",
			want: "EE",
		},
		{
			bits: "111000000000111",
			want: "EE",
		},
		{
			bits: "10111",
			want: "A",
		},
		{
			bits: "1110111",
			want: "M",
		},
		{
			bits: "1111111",
			want: "E",
		},
		{
			bits: "111000111",
			want: "I",
		},
		{
			bits: "11111100111111",
			want: "M",
		},
		{
			bits: "111000111000111",
			want: "S",
		},
		{
			bits: "1100110011001100000011000000111111001100111111001111110000000000000011001111110011111100111111" +
				"000000110011001111110000001111110011001100000011",
			want: "HEY JUDE",
		},
		{
			bits: "000111000101010100010000000111011101011100010101110001010001110101110100011101011100000001110101" +
				"01000101110100011101110111000101110111000111010000000101011101000111011101110001110101011100000001" +
				"01110111011100010101110001110111000101110111010001010100000001110111011100010101011100010001011101" +
				"00000001110001010101000100000001011101010001011100011101110101000111010111011100000001110101000111" +
				"01110111000111011101000101110101110101110",
			want: "THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG.",
		},
		{
			bits: "111111111111111000000000000000111110000011111000001111100000111110000000000000001111100000000000" +
				"00000000000000000000000011111111111111100000111111111111111000001111100000111111111111111000000000" +
				"00000011111000001111100000111111111111111000000000000000111110000011111000000000000000111111111111" +
				"11100000111110000011111111111111100000111110000000000000001111111111111110000011111000001111111111" +
				"11111000000000000000000000000000000000001111111111111110000011111000001111100000111110000000000000" +
				"00111110000011111111111111100000111110000000000000001111111111111110000011111111111111100000111111" +
				"11111111100000000000000011111000001111111111111110000011111111111111100000000000000011111111111111" +
				"10000011111000000000000000000000000000000000001111100000111110000011111111111111100000111110000000" +
				"00000000111111111111111000001111111111111110000011111111111111100000000000000011111111111111100000" +
				"11111000001111100000111111111111111000000000000000000000000000000000001111100000111111111111111000" +
				"00111111111111111000001111111111111110000000000000001111100000111110000011111111111111100000000000" +
				"00001111111111111110000011111111111111100000000000000011111000001111111111111110000011111111111111" +
				"10000011111000000000000000111110000011111000001111100000000000000000000000000000000000111111111111" +
				"11100000111111111111111000001111111111111110000000000000001111100000111110000011111000001111111111" +
				"11111000000000000000111110000000000000001111100000111111111111111000001111100000000000000000000000" +
				"00000000000011111111111111100000000000000011111000001111100000111110000011111000000000000000111110" +
				"00000000000000000000000000000000001111100000111111111111111000001111100000111110000000000000001111" +
				"10000011111111111111100000000000000011111111111111100000111111111111111000001111100000111110000000" +
				"00000000111111111111111000001111100000111111111111111000001111111111111110000000000000000000000000" +
				"00000000001111111111111110000011111000001111100000000000000011111111111111100000111111111111111000" +
				"00111111111111111000000000000000111111111111111000001111111111111110000011111000000000000000111110" +
				"0000111111111111111000001111100000111111111111111000001111100000111111111111111",
			want: "THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := DecodeMorse(DecodeBits(tt.bits)); got != tt.want {
				t.Errorf("DecodeBits() = %v, want %v", got, tt.want)
			}
		})
	}
}