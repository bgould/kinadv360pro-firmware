// This file is auto-generated by github.com/bgould/keyboard-firmware/hosts/usbvial/gen-def, do not edit.
package main

import "github.com/bgould/keyboard-firmware/hosts/usbvial/vial"

var VialDeviceDefinition = vial.DeviceDefinition{
	Name:      "Kinesis Advantage 360 Pro",
	VendorID:  "0x239A",
	ProductID: "0x8045",
	Matrix: vial.DeviceMatrix{
		Rows: 15,
		Cols: 7,
	},
	LzmaDefLength: uint16(len(vialLzmaDefinition)),
	LzmaDefWriter: vial.LzmaDefPageWriterFunc(vialWriteLzmaDefPage),
}

var vialLzmaDefinition = []byte{
	0x5D, 0x00, 0x00, 0x80, 0x00, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0x3D, 0x82, 0x80, 0x17, 0x1C, 0x2E, 0x8B, 0x89, 0x9F, 0x24, 0xFF, 0xE8, 0x0F, 0x31, 0x14, 0x05, 0xB2, 0x5B, 0x0A, 0xB8, 0xD2, 0xDE, 0xA9, 0xB2, 0x0B, 0x0B, 0xCF, 0x3B, 0x85, 0x91, 0x59, 0x1F, 0x3A, 0x6A, 0xD5, 0x8A, 0xA8, 0x70, 0x9B, 0x74, 0x4E, 0x1C, 0xE2, 0xD1, 0xA1, 0x8C, 0xE3, 0xB0, 0xF3, 0x52, 0x50, 0x69, 0xCB, 0xC9, 0xB3, 0x2C, 0x18, 0x54, 0x59, 0x02, 0x14, 0x97, 0xC8, 0x01, 0xE7, 0x5C, 0xFD, 0x4F, 0xFD, 0x0C, 0xC8, 0x3F, 0x7C, 0x0E, 0x14, 0x24, 0xF5, 0xAB, 0xD8, 0x28, 0xB5, 0x7D, 0xA4, 0xD5, 0x4F, 0x3F, 0x5F, 0x23, 0x2B, 0xBF, 0xFD, 0xBF, 0x49, 0x08, 0x11, 0xBE, 0x7F, 0x98, 0x02, 0xD6, 0x37, 0xDF, 0x7D, 0xA6, 0x9A, 0x24, 0xA4, 0x21, 0x66, 0x8E, 0xED, 0x2F, 0x96, 0xB9, 0xBA, 0x2A, 0x90, 0xEF, 0x18, 0xC7, 0x60, 0xB8, 0x8D, 0x98, 0x10, 0xAE, 0x5A, 0x29, 0xA9, 0xCE, 0x0D, 0xD9, 0x90, 0x78, 0x3E, 0x0B, 0x42, 0xA8, 0x2C, 0xBF, 0x61, 0x76, 0xFE, 0xC5, 0x1E, 0x65, 0x4B, 0x61, 0xD2, 0x39, 0x63, 0xC1, 0x52, 0xA1, 0x3B, 0x66, 0xDB, 0xF7, 0x8F, 0x29, 0x3A, 0x44, 0xFD, 0xF8, 0x65, 0x44, 0xB2, 0x75, 0x0F, 0x72, 0x53, 0x90, 0x67, 0xE9, 0x22, 0xA3, 0xB9, 0x33, 0x3C, 0x9C, 0xE6, 0xC9, 0x7A, 0x59, 0x92, 0x64, 0x1E, 0xC2, 0xE9, 0x2B, 0x45, 0x67, 0xB2, 0x97, 0x33, 0xE0, 0x48, 0x3C, 0x27, 0xE2, 0xEF, 0x55, 0x57, 0x52, 0x03, 0xD2, 0x38, 0x7D, 0x5F, 0xC9, 0x10, 0xD6, 0x27, 0x2B, 0x2B, 0x5E, 0xC8, 0x81, 0x6C, 0xAE, 0x98, 0x5B, 0xFA, 0x7E, 0xC3, 0x1D, 0x7F, 0xA0, 0x82, 0x84, 0x33, 0xA8, 0xBA, 0x85, 0xA8, 0xBD, 0x91, 0x02, 0xD9, 0x3D, 0xD1, 0x5C, 0x9A, 0x89, 0xE3, 0x6E, 0x17, 0x65, 0xE8, 0xD5, 0xD4, 0x66, 0x3B, 0xB4, 0x2C, 0x5A, 0xB8, 0x34, 0x3B, 0xD3, 0x25, 0x0E, 0x5B, 0xEA, 0x04, 0xCB, 0x65, 0x7A, 0x69, 0x9E, 0x2C, 0x1E, 0x25, 0x89, 0x76, 0x24, 0x84, 0x95, 0xDD, 0xB9, 0x35, 0xFA, 0x82, 0xEB, 0xE0, 0x99, 0xE0, 0x19, 0x9D, 0xA3, 0x59, 0x2A, 0x3A, 0xFF, 0x7D, 0x26, 0x97, 0x5E, 0x92, 0xC4, 0x92, 0x59, 0xA3, 0x68, 0xBE, 0xE3, 0xB7, 0x17, 0x09, 0x1B, 0x43, 0x55, 0x6C, 0x8C, 0x3A, 0x15, 0x6B, 0xFA, 0xE4, 0x4C, 0x08, 0xDB, 0xAB, 0xDB, 0xC5, 0x60, 0x3E, 0xE1, 0x70, 0x34, 0x70, 0x68, 0x9F, 0xF1, 0xF3, 0x90, 0xA7, 0x68, 0x8B, 0x3A, 0xF2, 0x4A, 0xB1, 0x21, 0x48, 0x6B, 0x00, 0x5A, 0x92, 0x17, 0x13, 0xFF, 0x89, 0xE9, 0x0A, 0x00, 0x32, 0x34, 0xE9, 0x2A, 0x71, 0x79, 0xA1, 0xDE, 0xD3, 0x88, 0xD8, 0xF3, 0xE7, 0x13, 0xB4, 0x93, 0x3E, 0x6D, 0xF1, 0x9E, 0x54, 0x6E, 0x14, 0x11, 0x4D, 0x93, 0x1F, 0x9F, 0x68, 0xFC, 0x92, 0x61, 0xA3, 0x10, 0x16, 0x6B, 0xD4, 0x31, 0x74, 0x2C, 0xFE, 0xC9, 0x0D, 0xEC, 0x00, 0x39, 0x47, 0xB2, 0x03, 0x20, 0xFD, 0x7A, 0x64, 0x35, 0x1C, 0x09, 0xB0, 0x04, 0x38, 0x18, 0x48, 0xC1, 0x5C, 0x92, 0x0C, 0xBC, 0xCA, 0x75, 0x10, 0xAA, 0xDC, 0x10, 0xB6, 0x1A, 0x57, 0xC2, 0xCC, 0x06, 0xC9, 0xF9, 0xA5, 0xAE, 0xF5, 0x28, 0xF4, 0xA4, 0x4B, 0x72, 0x78, 0xD1, 0x45, 0x13, 0xE3, 0x9B, 0xF0, 0xFD, 0xF6, 0xA3, 0x7A, 0x31, 0xE1, 0x5E, 0xFD, 0x41, 0xD8, 0x10, 0xE5, 0xAA, 0x77, 0xF6, 0x70, 0xDA, 0xB9, 0x74, 0x5A, 0x5F, 0xFB, 0xAA, 0x27, 0xF5,
}

func vialWriteLzmaDefPage(tx []byte, page uint16) bool {
	start := page * 32
	end := start + 32
	len := uint16(len(vialLzmaDefinition))
	if end < start || start >= len {
		return false
	}
	if end > len {
		end = len
	}
	copy(tx[:32], vialLzmaDefinition[start:end])
	return true
}
