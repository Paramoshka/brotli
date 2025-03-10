package brotli

/* Copyright 2015 Google Inc. All Rights Reserved.

   Distributed under MIT license.
   See file LICENSE for detail or copy at https://opensource.org/licenses/MIT
*/

/* Hash table on the 4-byte prefixes of static dictionary words. */
var kStaticDictionaryHash = [32768]uint16{}
