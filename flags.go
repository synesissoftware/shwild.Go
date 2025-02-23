// Copyright 2005-2012, Matthew Wilson and Sean Kelly. Copyright 2018-2025
// Matthew Wilson and Synesis Information Systems. All rights reserved. Use
// of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

/*
 * Created: 17th June 2005
 * Updated: 24th February 2025
 */

package shwild

const (

	// Suppresses the recognition of ranges. [ and ] are treated as literal
	// characters (and need no escaping)
	SuppressRangeSupport = 1 << iota

	// Suppresses the use of backslash interpretation as escape. \ is
	// treated as a literal character
	SuppressBackslashEscape

	// Suppresses the recognition of range continua, i.e. [0-9]
	SuppressRangeContinuumSupport

	// Suppresses the recognition of reverse range continua, i.e. [9-0],
	// [M-D]
	SuppressRangeContinuumHighlowSupport

	// Suppresses the recognition of cross-case range continua, i.e. [h-J]
	// === [hijHIJ]
	SuppressRangeContinuumCrosscaseSupport

	// Suppresses the recognition of ? and * as literal inside range
	SuppressRangeLiteralWildcard

	// Suppresses the recognition of leading/trailing hyphens as literal
	// inside range
	SuppressRangeLeadtrailLiteralHyphen

	// Suppresses the use of a leading ^ to mean not any of the following,
	// i.e. [^0-9] means do not match a digit
	SuppressRangeNot

	// Comparison is case-insensitive
	IgnoreCase

	// Treats [ and ] as literal inside range. ] only literal if immediately
	// preceeds closing ]
	AllowRangeLiteralBracket

	// Allows quantification of the wildcards, with trailing escaped
	// numbers, as in [a-Z]\2-10. All chars in 0-9- become range specifiers.
	// These are separated from actual pattern digits by []
	AllowRangeQuantification
)

/* ///////////////////////////// end of file //////////////////////////// */
