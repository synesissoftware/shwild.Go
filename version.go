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
	VersionMajor    int16 = 0
	VersionMinor    int16 = 2
	VersionRevision int16 = 2

	Version int64 = int64(VersionMajor)<<48 | int64(VersionMinor)<<32 | int64(VersionRevision)<<16
)

/* ///////////////////////////// end of file //////////////////////////// */
