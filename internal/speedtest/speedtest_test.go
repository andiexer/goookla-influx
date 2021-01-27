package speedtest

import "testing"

func TestToMbit(t *testing.T) {
	var tests = []struct {
		bandwith int
		expected int
	}{
		{2000000, 16},
		{10000, 0},
		{125000, 1},
	}

	for _, v := range tests {
		m := Measurement{Bandwith: v.bandwith}
		if m.ToMbit() != v.expected {
			t.Errorf("ToMbit() returned %d, expected %d", m.ToMbit(), v.expected)
		}
	}
}
