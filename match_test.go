
package shwild

import "fmt"
import "path"
import "runtime"
import "testing"

/* /////////////////////////////////////////////////////////////////////////
 * internal functions
 */

func check_Match(t *testing.T, pattern, s string, expectedResult bool, e error) {

	m_r, m_e := Match(pattern, s)

	if expectedResult == m_r && e == m_e {

		return
	}

	_, file, line, hasCallInfo := runtime.Caller(1)

	var msg string

	if hasCallInfo {

		if expectedResult != m_r {

			msg = fmt.Sprintf("\t%s:%d: Match('%s', '%s') returned '%v'; '%v' expected", path.Base(file), line, pattern, s, m_r, expectedResult);
		}
	} else {

	}

	fmt.Printf("%s\n", msg)

	t.Fail()
}

/* /////////////////////////////////////////////////////////////////////////
 * tests
 */

func TestMatch_with_empty_pattern(t *testing.T) {

	check_Match(t, "", "", true, nil)
	check_Match(t, "", "1", false, nil)
	check_Match(t, "", "*", false, nil)
	check_Match(t, "", ".", false, nil)
}

func TestMatch_with_wild1(t *testing.T) {

	check_Match(t, "?", "", false, nil)
	check_Match(t, "?", "?", true, nil)
}

func TestMatch_with_allstar_patterns(t *testing.T) {

	// 1 star

	check_Match(t, "*", "", true, nil)
	check_Match(t, "*", "1", true, nil)
	check_Match(t, "*", "*", true, nil)
	check_Match(t, "*", ".", true, nil)

	// 2 star

	check_Match(t, "**", "", true, nil)
	check_Match(t, "**", "1", true, nil)
	check_Match(t, "**", "*", true, nil)
	check_Match(t, "**", ".", true, nil)

	// 5 star

	check_Match(t, "*****", "", true, nil)
	check_Match(t, "*****", "1", true, nil)
	check_Match(t, "*****", "*", true, nil)
	check_Match(t, "*****", ".", true, nil)
}

func TestMatch_with_literal(t *testing.T) {

	check_Match(t, "a", "a", true, nil)
	check_Match(t, "aa", "a", false, nil)
	check_Match(t, "aa", "aa", true, nil)
	check_Match(t, "a", "aa", false, nil)
}

func TestMatch_with_literal_and_wild1(t *testing.T) {

	check_Match(t, "a?", "a", false, nil)
	check_Match(t, "a?", "a?", true, nil)
	check_Match(t, "a?", "aa", true, nil)
	check_Match(t, "a?", "aaa", false, nil)
	check_Match(t, "?a", "a", false, nil)
	check_Match(t, "?a", "a?", false, nil)
	check_Match(t, "?a", "?a", true, nil)
	check_Match(t, "?a", "aa", true, nil)
	check_Match(t, "?a", "aaa", false, nil)
}

func TestMatch_with_literal_and_wildN(t *testing.T) {

	check_Match(t, "a*", "a", true, nil)
	check_Match(t, "a*", "a*", true, nil)
	check_Match(t, "a*", "aa", true, nil)
	check_Match(t, "a*", "aaa", true, nil)
	check_Match(t, "a*", "abcdefghijklmno", true, nil)
	check_Match(t, "a*o", "abcdefghijklmno", true, nil)
	check_Match(t, "a*n", "abcdefghijklmno", false, nil)
	check_Match(t, "a*n*", "abcdefghijklmno", true, nil)
	check_Match(t, "*a", "a", true, nil)
	check_Match(t, "*a", "a*", false, nil)
	check_Match(t, "*a", "*a", true, nil)
	check_Match(t, "*a", "aa", true, nil)
	check_Match(t, "*a", "aaa", true, nil)
}

func TestMatch_with_explicit_range(t *testing.T) {

	check_Match(t, "[abc]", "a", true, nil)
	check_Match(t, "[abc]", "b", true, nil)
	check_Match(t, "[abc]", "c", true, nil)
	check_Match(t, "[abc]", "d", false, nil)

	check_Match(t, "[-abc]", "-", true, nil)
}

func TestMatch_with_forward_continuum_range(t *testing.T) {

	check_Match(t, "[a-c]", "a", true, nil)
	check_Match(t, "[a-c]", "b", true, nil)
	check_Match(t, "[a-c]", "c", true, nil)
	check_Match(t, "[a-c]", "-", false, nil)
	check_Match(t, "[a-c]", "z", false, nil)
	check_Match(t, "[a-c]", "A", false, nil)
	check_Match(t, "[a-c]", "B", false, nil)
	check_Match(t, "[a-c]", "C", false, nil)
	check_Match(t, "[a-c]", "D", false, nil)
	check_Match(t, "[a-c]", "E", false, nil)

	check_Match(t, "[-ac]", "a", true, nil)
	check_Match(t, "[-ac]", "b", false, nil)
	check_Match(t, "[-ac]", "c", true, nil)
	check_Match(t, "[-ac]", "-", true, nil)
}

func TestMatch_with_backward_continuum_range(t *testing.T) {

	check_Match(t, "[c-a]", "a", true, nil)
	check_Match(t, "[c-a]", "b", true, nil)
	check_Match(t, "[c-a]", "c", true, nil)
	check_Match(t, "[c-a]", "-", false, nil)
	check_Match(t, "[c-a]", "z", false, nil)
	check_Match(t, "[c-a]", "A", false, nil)
	check_Match(t, "[c-a]", "B", false, nil)
	check_Match(t, "[c-a]", "C", false, nil)
	check_Match(t, "[c-a]", "D", false, nil)
	check_Match(t, "[c-a]", "E", false, nil)
}

/* ///////////////////////////// end of file //////////////////////////// */


