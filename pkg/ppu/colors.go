//go:build !nesgo

package ppu

import "image/color"

var colors = [64]color.RGBA{
	{0x58, 0x58, 0x58, 0xFF},
	{0x00, 0x23, 0x7C, 0xFF},
	{0x0D, 0x10, 0x99, 0xFF},
	{0x30, 0x00, 0x92, 0xFF},
	{0x4F, 0x00, 0x6C, 0xFF},
	{0x60, 0x00, 0x35, 0xFF},
	{0x5C, 0x05, 0x00, 0xFF},
	{0x46, 0x18, 0x00, 0xFF},
	{0x27, 0x2D, 0x00, 0xFF},
	{0x09, 0x3E, 0x00, 0xFF},
	{0x00, 0x45, 0x00, 0xFF},
	{0x00, 0x41, 0x06, 0xFF},
	{0x00, 0x35, 0x45, 0xFF},
	{0x00, 0x00, 0x00, 0xFF},
	{0x00, 0x00, 0x00, 0xFF},
	{0x00, 0x00, 0x00, 0xFF},
	{0xA1, 0xA1, 0xA1, 0xFF},
	{0x0B, 0x53, 0xD7, 0xFF},
	{0x33, 0x37, 0xFE, 0xFF},
	{0x66, 0x21, 0xF7, 0xFF},
	{0x95, 0x15, 0xBE, 0xFF},
	{0xAC, 0x16, 0x6E, 0xFF},
	{0xA6, 0x27, 0x21, 0xFF},
	{0x86, 0x43, 0x00, 0xFF},
	{0x59, 0x62, 0x00, 0xFF},
	{0x2D, 0x7A, 0x00, 0xFF},
	{0x0C, 0x85, 0x00, 0xFF},
	{0x00, 0x7F, 0x2A, 0xFF},
	{0x00, 0x6D, 0x85, 0xFF},
	{0x00, 0x00, 0x00, 0xFF},
	{0x00, 0x00, 0x00, 0xFF},
	{0x00, 0x00, 0x00, 0xFF},
	{0xFF, 0xFF, 0xFF, 0xFF},
	{0x51, 0xA5, 0xFE, 0xFF},
	{0x80, 0x84, 0xFE, 0xFF},
	{0xBC, 0x6A, 0xFE, 0xFF},
	{0xF1, 0x5B, 0xFE, 0xFF},
	{0xFE, 0x5E, 0xC4, 0xFF},
	{0xFE, 0x72, 0x69, 0xFF},
	{0xE1, 0x93, 0x21, 0xFF},
	{0xAD, 0xB6, 0x00, 0xFF},
	{0x79, 0xD3, 0x00, 0xFF},
	{0x51, 0xDF, 0x21, 0xFF},
	{0x3A, 0xD9, 0x74, 0xFF},
	{0x39, 0xC3, 0xDF, 0xFF},
	{0x42, 0x42, 0x42, 0xFF},
	{0x00, 0x00, 0x00, 0xFF},
	{0x00, 0x00, 0x00, 0xFF},
	{0xFF, 0xFF, 0xFF, 0xFF},
	{0xB5, 0xD9, 0xFE, 0xFF},
	{0xCA, 0xCA, 0xFE, 0xFF},
	{0xE3, 0xBE, 0xFE, 0xFF},
	{0xF9, 0xB8, 0xFE, 0xFF},
	{0xFE, 0xBA, 0xE7, 0xFF},
	{0xFE, 0xC3, 0xBC, 0xFF},
	{0xF4, 0xD1, 0x99, 0xFF},
	{0xDE, 0xE0, 0x86, 0xFF},
	{0xC6, 0xEC, 0x87, 0xFF},
	{0xB2, 0xF2, 0x9D, 0xFF},
	{0xA7, 0xF0, 0xC3, 0xFF},
	{0xA8, 0xE7, 0xF0, 0xFF},
	{0xAC, 0xAC, 0xAC, 0xFF},
	{0x00, 0x00, 0x00, 0xFF},
	{0x00, 0x00, 0x00, 0xFF},
}
