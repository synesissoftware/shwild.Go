/* /////////////////////////////////////////////////////////////////////////
 * File:        flags.go
 *
 * Purpose:     Flags for shwild.Go API
 *
 * Created:     17th June 2005
 * Updated:     9th March 2019
 *
 * Home:        http://shwild.org/
 *
 * Copyright (c) 2005-2012, Matthew Wilson and Sean Kelly
 * Copyright (c) 2005-2019, Matthew Wilson and Synesis Software
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are
 * met:
 *
 * - Redistributions of source code must retain the above copyright notice,
 *   this list of conditions and the following disclaimer.
 * - Redistributions in binary form must reproduce the above copyright
 *   notice, this list of conditions and the following disclaimer in the
 *   documentation and/or other materials provided with the distribution.
 * - Neither the names of Matthew Wilson, Sean Kelly, Synesis Software nor
 *   the names of any contributors may be used to endorse or promote products
 *   derived from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS
 * IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO,
 * THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR
 * PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR
 * CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
 * EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
 * PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
 * PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
 * LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
 * NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
 * SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 * ////////////////////////////////////////////////////////////////////// */


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


