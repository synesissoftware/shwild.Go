/* /////////////////////////////////////////////////////////////////////////
 * File:        api.go
 *
 * Purpose:     Main shwild.Go API
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

import (

	"fmt"
)

const (
)

/* /////////////////////////////////////////////////////////////////////////
 * API types
 */

type patternBehaviour int

const (

	_PB_RegularPattern patternBehaviour	=	1 << iota
	_PB_EmptyPattern patternBehaviour	=	1 << iota
	_PB_AllWildPattern patternBehaviour	=	1 << iota
)

type CompiledPattern struct {

	Pattern		string
	matchers	[]matcher
	behaviour	patternBehaviour
}

func (cp CompiledPattern) Match(s string) (bool, error) {

	switch cp.behaviour {

	case _PB_EmptyPattern:

		return 0 == len(s), nil
	case _PB_AllWildPattern:

		return true, nil
	case _PB_RegularPattern:

		return match_from_compiled_(cp.matchers, s)
	default:

		msg := fmt.Sprintf("VIOLATION: unrecognised CompiledPattern.behaviour %d", cp.behaviour)

		panic(msg)
	}
}

func (cp CompiledPattern) String() string {

	switch cp.behaviour {

	case _PB_EmptyPattern:

		return fmt.Sprintf("<%T{ <empty-pattern> }>", cp)
	case _PB_AllWildPattern:

		return fmt.Sprintf("<%T{ <all-wild-pattern> }>", cp)
	case _PB_RegularPattern:

		return fmt.Sprintf("<%T{ Pattern=%q }>", cp, cp.Pattern)
	default:

		return fmt.Sprintf("<%T{ <unknown state!> }>", cp)
	}
}

/* /////////////////////////////////////////////////////////////////////////
 * internal types
 */


/* /////////////////////////////////////////////////////////////////////////
 * API functions
 */

func Match(pattern string, s string, args ...interface{}) (bool, error) {

	// An empty pattern can only match an empty string

	if 0 == len(pattern) {

		return 0 == len(s), nil
	}


	// A pattern composed entirely of '*' can match anything

	allstar := true

	for _, ch := range pattern {

		if '*' != ch {

			allstar = false
			break
		}
	}

	if allstar {

		return true, nil
	}


	// parse flags

	flags := parse_flags_(args...);

	matchers, err := parse_matchers(pattern, flags)

	if nil != err {

		return false, err
	}

	if 0 == len(matchers) {

		panic("VIOLATION: empty matchers slice")
	}

	return match_from_compiled_(matchers, s)
}

func Compile(pattern string, args ...interface{}) (CompiledPattern, error) {

	// An empty pattern can only match an empty string

	if 0 == len(pattern) {

		return CompiledPattern{ Pattern: pattern, matchers: nil, behaviour: _PB_EmptyPattern }, nil
	}


	// A pattern composed entirely of '*' can match anything

	allstar := true

	for _, ch := range pattern {

		if '*' != ch {

			allstar = false
			break
		}
	}

	if allstar {

		return CompiledPattern{ Pattern: pattern, matchers: nil, behaviour: _PB_AllWildPattern }, nil
	}


	// parse flags

	flags := parse_flags_(args...);

	matchers, err := parse_matchers(pattern, flags)

	if nil != err {

		return CompiledPattern{}, err
	}

	if 0 == len(matchers) {

		panic("VIOLATION: empty matchers slice")
	}

	return CompiledPattern{ Pattern: pattern, matchers: matchers, behaviour: _PB_RegularPattern }, nil
}

/* /////////////////////////////////////////////////////////////////////////
 * internal functions
 */

func parse_flags_(args ...interface{}) uint64 {

	var flags uint64 = 0

	for i, arg := range args {

		switch v := arg.(type) {

			case uint32:

				flags |= uint64(v)

			case uint64:

				flags |= v

			default:

				var msg = fmt.Sprintf("invalid type (%T) for argument '%v' at index %d", v, v, i)

				panic(msg)
		}
	}

	return 0;
}

func match_from_compiled_(matchers []matcher, pattern string) (bool, error) {

	return matchers[0].match(pattern), nil
}

/* ///////////////////////////// end of file //////////////////////////// */


