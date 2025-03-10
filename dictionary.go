package brotli

/* Copyright 2013 Google Inc. All Rights Reserved.

   Distributed under MIT license.
   See file LICENSE for detail or copy at https://opensource.org/licenses/MIT
*/

/* Collection of static dictionary words. */
type dictionary struct {
	size_bits_by_length [32]byte
	offsets_by_length   [32]uint32
	data_size           uint
	data                []byte
}

const minDictionaryWordLength = 4

const maxDictionaryWordLength = 24

var kBrotliDictionaryData []byte

var kBrotliDictionary = dictionary{
	/* size_bits_by_length */
	[32]byte{
		0,
		0,
		0,
		0,
		10,
		10,
		11,
		11,
		10,
		10,
		10,
		10,
		10,
		9,
		9,
		8,
		7,
		7,
		8,
		7,
		7,
		6,
		6,
		5,
		5,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
	},

	/* offsets_by_length */
	[32]uint32{
		0,
		0,
		0,
		0,
		0,
		4096,
		9216,
		21504,
		35840,
		44032,
		53248,
		63488,
		74752,
		87040,
		93696,
		100864,
		104704,
		106752,
		108928,
		113536,
		115968,
		118528,
		119872,
		121280,
		122016,
		122784,
		122784,
		122784,
		122784,
		122784,
		122784,
		122784,
	},

	/* data_size ==  sizeof(kBrotliDictionaryData) */
	122784,

	/* data */
	kBrotliDictionaryData,
}

func getDictionary() *dictionary {
	return &kBrotliDictionary
}
