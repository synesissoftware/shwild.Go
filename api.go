
package shwild

import "fmt"

const (
)

/* /////////////////////////////////////////////////////////////////////////
 * API types
 */

/* /////////////////////////////////////////////////////////////////////////
 * internal types
 */

type matcher interface {

	match(s string) bool;
}

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

	flags := parse_flags_(args);

	matchers, err := parse_matches_(pattern, flags)

	if nil != err {

		return false, err
	}

	return match_from_compiled_(matchers, s)
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


func parse_matches_(pattern string, flag uint64) ([]matcher, error) {

	if 0 == len(pattern) {

		return nil, nil
	}

	return nil, nil
}

func match_from_compiled_(matchers []matcher, pattern string) (bool, error) {

	return matchers[0].match(pattern), nil
}

/* ///////////////////////////// end of file //////////////////////////// */


