package zigzag_conversion

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestZigZagConversion(t *testing.T) {
    tests := map[string]struct {
        s string
        numRows int
        want string
    } {
        "it_should_PAYPALISHIRING_zizac_return_PAHNAPLSIIGYIR": {
            s:    "PAYPALISHIRING",
            numRows: 3,
            want: "PAHNAPLSIIGYIR",
        },
        "it_should_PAYPALISHIRING_zizac_return_PINALSIGYAHRPI": {
            s:    "PAYPALISHIRING",
            numRows: 4,
            want: "PINALSIGYAHRPI",
        },
        "it_should_A_zizac_return_A": {
            s:    "A",
            numRows: 1,
            want: "A",
        },
        "it_should_AB_zizac_return_AB": {
            s:    "AB",
            numRows: 1,
            want: "AB",
        },
    }

    for name, tt := range tests {
        t.Run(name, func(t *testing.T){
            assert.Equal(t, tt.want, convert(tt.s, tt.numRows))
        })
    }
}
