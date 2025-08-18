// Copyright 2005-2012, Matthew Wilson and Sean Kelly. Copyright 2018-2025
// Matthew Wilson and Synesis Information Systems. All rights reserved. Use
// of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

/*
 * Created: 17th June 2005
 * Updated: 24th February 2025
 */

package shwild

import (
	"fmt"
)

/* /////////////////////////////////////////////////////////////////////////
 * API types
 */

type patternBehaviour int

const (
	_PB_RegularPattern patternBehaviour = 1 << iota
	_PB_EmptyPattern   patternBehaviour = 1 << iota
	_PB_AllWildPattern patternBehaviour = 1 << iota
)

type CompiledPattern struct {
	Pattern   string
	matchers  []matcher
	behaviour patternBehaviour
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

func Match(pattern string, s string, args ...any) (bool, error) {

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

	flags := parse_flags_(args...)

	matchers, err := parse_matchers(pattern, flags)

	if nil != err {

		return false, err
	}

	if 0 == len(matchers) {

		panic("VIOLATION: empty matchers slice")
	}

	return match_from_compiled_(matchers, s)
}

func Compile(pattern string, args ...any) (CompiledPattern, error) {

	// An empty pattern can only match an empty string

	if 0 == len(pattern) {

		return CompiledPattern{Pattern: pattern, matchers: nil, behaviour: _PB_EmptyPattern}, nil
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

		return CompiledPattern{Pattern: pattern, matchers: nil, behaviour: _PB_AllWildPattern}, nil
	}

	// parse flags

	flags := parse_flags_(args...)

	matchers, err := parse_matchers(pattern, flags)

	if nil != err {

		return CompiledPattern{}, err
	}

	if 0 == len(matchers) {

		panic("VIOLATION: empty matchers slice")
	}

	return CompiledPattern{Pattern: pattern, matchers: matchers, behaviour: _PB_RegularPattern}, nil
}

/* /////////////////////////////////////////////////////////////////////////
 * internal functions
 */

func parse_flags_(args ...any) uint64 {

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

	return 0
}

func match_from_compiled_(matchers []matcher, pattern string) (bool, error) {

	return matchers[0].match(pattern), nil
}

/* ///////////////////////////// end of file //////////////////////////// */
