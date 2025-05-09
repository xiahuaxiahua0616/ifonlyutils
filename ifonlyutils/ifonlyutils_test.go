package ifonlyutils

import (
	"reflect"
	"testing"
)

func TestConv1to14(t *testing.T) {
	tests := []struct {
		name         string
		cards        []byte
		joker        byte
		wantStraight []byte
	}{
		{
			name:         "方片 0x01",
			cards:        []byte{0x01, 0x02, 0x03, 0x04},
			joker:        0,
			wantStraight: []byte{0x02, 0x03, 0x04, 0x0e},
		},
		{
			name:         "梅花 0x11",
			cards:        []byte{0x11, 0x12, 0x13, 0x14},
			joker:        0,
			wantStraight: []byte{0x12, 0x13, 0x14, 0x1e},
		},
		{
			name:         "红桃 0x21",
			cards:        []byte{0x21, 0x22, 0x23, 0x24},
			joker:        0,
			wantStraight: []byte{0x22, 0x23, 0x24, 0x2e},
		},
		{
			name:         "黑桃 0x31",
			cards:        []byte{0x31, 0x32, 0x33, 0x34},
			joker:        0,
			wantStraight: []byte{0x32, 0x33, 0x34, 0x3e},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := conv1to14(tt.cards)
			if !reflect.DeepEqual(result, tt.wantStraight) {
				t.Errorf("expected %v, got %v", tt.wantStraight, result)
			}
		})
	}
}

func TestConv14to1(t *testing.T) {
	tests := []struct {
		name         string
		cards        []byte
		joker        byte
		wantStraight []byte
	}{
		{
			name:         "方片 0x01",
			cards:        []byte{0x02, 0x03, 0x04, 0x0e},
			joker:        0,
			wantStraight: []byte{0x01, 0x02, 0x03, 0x04},
		},
		{
			name:         "梅花 0x11",
			cards:        []byte{0x12, 0x13, 0x14, 0x1e},
			joker:        0,
			wantStraight: []byte{0x11, 0x12, 0x13, 0x14},
		},
		{
			name:         "红桃 0x21",
			cards:        []byte{0x22, 0x23, 0x24, 0x2e},
			joker:        0,
			wantStraight: []byte{0x21, 0x22, 0x23, 0x24},
		},
		{
			name:         "黑桃 0x31",
			cards:        []byte{0x32, 0x33, 0x34, 0x3e},
			joker:        0,
			wantStraight: []byte{0x31, 0x32, 0x33, 0x34},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := conv14to1(tt.cards)
			if !reflect.DeepEqual(result, tt.wantStraight) {
				t.Errorf("expected %v, got %v", tt.wantStraight, result)
			}
		})
	}
}
