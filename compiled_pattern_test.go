
package shwild_test

import (

	shwild "github.com/synesissoftware/shwild.Go"

	"fmt"
	"path"
	"runtime"
	"testing"
)

/* /////////////////////////////////////////////////////////////////////////
 * internal functions
 */

func check_CompiledPattern_Match(t *testing.T, cp shwild.CompiledPattern, s string, expectedResult bool, e error) {

	m_r, m_e := cp.Match(s)

	if expectedResult == m_r && e == m_e {

		return
	}

	_, file, line, hasCallInfo := runtime.Caller(1)

	var msg string

	if hasCallInfo {

		if expectedResult != m_r {

			msg = fmt.Sprintf("\t%s:%d: With CompiledPattern %v calling Match('%s') returned '%v'; '%v' expected", path.Base(file), line, cp, s, m_r, expectedResult);
		}
	} else {

	}

	fmt.Printf("%s\n", msg)

	t.Fail()
}

/* /////////////////////////////////////////////////////////////////////////
 * tests
 */

func Test_CompiledPattern_Match_with_empty_pattern(t *testing.T) {

	pattern	:=	""

	cp, err	:=	shwild.Compile(pattern)
	if err != nil {

		t.Errorf("Failed to compile pattern '%s'", pattern)
	}

	check_CompiledPattern_Match(t, cp, "", true, nil)
	check_CompiledPattern_Match(t, cp, "1", false, nil)
	check_CompiledPattern_Match(t, cp, "*", false, nil)
	check_CompiledPattern_Match(t, cp, ".", false, nil)
}

func Test_CompiledPattern_Match_with_wild1(t *testing.T) {

	pattern	:=	"?"

	cp, err	:=	shwild.Compile(pattern)
	if err != nil {

		t.Errorf("Failed to compile pattern '%s'", pattern)
	}

	check_CompiledPattern_Match(t, cp, "", false, nil)
	check_CompiledPattern_Match(t, cp, "?", true, nil)
}

func Test_CompiledPattern_Match_with_allstar_patterns(t *testing.T) {

	one_star	:=	"*"

	cp, err	:=	shwild.Compile(one_star)
	if err != nil {

		t.Errorf("Failed to compile pattern '%s'", one_star)
	}

	// 1 star

	check_CompiledPattern_Match(t, cp, "", true, nil)
	check_CompiledPattern_Match(t, cp, "1", true, nil)
	check_CompiledPattern_Match(t, cp, "*", true, nil)
	check_CompiledPattern_Match(t, cp, ".", true, nil)

	two_stars	:=	"**"

	cp, err	=	shwild.Compile(two_stars)
	if err != nil {

		t.Errorf("Failed to compile pattern '%s'", two_stars)
	}

	// 2 star

	check_CompiledPattern_Match(t, cp, "", true, nil)
	check_CompiledPattern_Match(t, cp, "1", true, nil)
	check_CompiledPattern_Match(t, cp, "*", true, nil)
	check_CompiledPattern_Match(t, cp, ".", true, nil)

	five_stars	:=	"*****"

	cp, err	=	shwild.Compile(five_stars)
	if err != nil {

		t.Errorf("Failed to compile pattern '%s'", five_stars)
	}

	// 5 star

	check_CompiledPattern_Match(t, cp, "", true, nil)
	check_CompiledPattern_Match(t, cp, "1", true, nil)
	check_CompiledPattern_Match(t, cp, "*", true, nil)
	check_CompiledPattern_Match(t, cp, ".", true, nil)
}

func Test_CompiledPattern_Match_with_literal(t *testing.T) {

	one_a	:=	"a"

	cp, err	:=	shwild.Compile(one_a)
	if err != nil {

		t.Errorf("Failed to compile pattern '%s'", one_a)
	}

	check_CompiledPattern_Match(t, cp, "a", true, nil)
	check_CompiledPattern_Match(t, cp, "aa", false, nil)

	two_as	:=	"aa"

	cp, err	=	shwild.Compile(two_as)
	if err != nil {

		t.Errorf("Failed to compile pattern '%s'", two_as)
	}

	check_CompiledPattern_Match(t, cp, "a", false, nil)
	check_CompiledPattern_Match(t, cp, "aa", true, nil)
}

func Test_CompiledPattern_Match_with_literal_and_wild1(t *testing.T) {

	pattern	:=	"a?"

	cp, err	:=	shwild.Compile(pattern)
	if err != nil {

		t.Errorf("Failed to compile pattern '%s'", pattern)
	}

	check_CompiledPattern_Match(t, cp, "a", false, nil)
	check_CompiledPattern_Match(t, cp, "a?", true, nil)
	check_CompiledPattern_Match(t, cp, "aa", true, nil)
	check_CompiledPattern_Match(t, cp, "aaa", false, nil)

	pattern	=	"?a"

	cp, err	=	shwild.Compile(pattern)
	if err != nil {

		t.Errorf("Failed to compile pattern '%s'", pattern)
	}

	check_CompiledPattern_Match(t, cp, "a", false, nil)
	check_CompiledPattern_Match(t, cp, "a?", false, nil)
	check_CompiledPattern_Match(t, cp, "?a", true, nil)
	check_CompiledPattern_Match(t, cp, "aa", true, nil)
	check_CompiledPattern_Match(t, cp, "aaa", false, nil)
}

func Test_CompiledPattern_Match_with_literal_and_wildN(t *testing.T) {

	pattern	:=	"a*"

	cp, err	:=	shwild.Compile(pattern)
	if err != nil {

		t.Errorf("Failed to compile pattern '%s'", pattern)
	}

	check_CompiledPattern_Match(t, cp, "a", true, nil)
	check_CompiledPattern_Match(t, cp, "a*", true, nil)
	check_CompiledPattern_Match(t, cp, "aa", true, nil)
	check_CompiledPattern_Match(t, cp, "aaa", true, nil)
	check_CompiledPattern_Match(t, cp, "abcdefghijklmno", true, nil)

	pattern	=	"*a"

	cp, err	=	shwild.Compile(pattern)
	if err != nil {

		t.Errorf("Failed to compile pattern '%s'", pattern)
	}

	check_CompiledPattern_Match(t, cp, "a", true, nil)
	check_CompiledPattern_Match(t, cp, "a*", false, nil)
	check_CompiledPattern_Match(t, cp, "*a", true, nil)
	check_CompiledPattern_Match(t, cp, "aa", true, nil)
	check_CompiledPattern_Match(t, cp, "aaa", true, nil)

	pattern	=	"a*o"

	cp, err	=	shwild.Compile(pattern)
	if err != nil {

		t.Errorf("Failed to compile pattern '%s'", pattern)
	}

	check_CompiledPattern_Match(t, cp, "abcdefghijklmno", true, nil)

	pattern	=	"a*n"

	cp, err	=	shwild.Compile(pattern)
	if err != nil {

		t.Errorf("Failed to compile pattern '%s'", pattern)
	}

	check_CompiledPattern_Match(t, cp, "abcdefghijklmno", false, nil)

	pattern	=	"a*n*"

	cp, err	=	shwild.Compile(pattern)
	if err != nil {

		t.Errorf("Failed to compile pattern '%s'", pattern)
	}

	check_CompiledPattern_Match(t, cp, "abcdefghijklmno", true, nil)
}

func Test_CompiledPattern_Match_with_explicit_range(t *testing.T) {

	pattern	:=	"[abc]"

	cp, err	:=	shwild.Compile(pattern)
	if err != nil {

		t.Errorf("Failed to compile pattern '%s'", pattern)
	}

	check_CompiledPattern_Match(t, cp, "a", true, nil)
	check_CompiledPattern_Match(t, cp, "b", true, nil)
	check_CompiledPattern_Match(t, cp, "c", true, nil)
	check_CompiledPattern_Match(t, cp, "d", false, nil)

	pattern	=	"[-abc]"

	cp, err	=	shwild.Compile(pattern)
	if err != nil {

		t.Errorf("Failed to compile pattern '%s'", pattern)
	}

	check_CompiledPattern_Match(t, cp, "-", true, nil)
}

func Test_CompiledPattern_Match_with_forward_continuum_range(t *testing.T) {

	pattern	:=	"[a-c]"

	cp, err	:=	shwild.Compile(pattern)
	if err != nil {

		t.Errorf("Failed to compile pattern '%s'", pattern)
	}

	check_CompiledPattern_Match(t, cp, "a", true, nil)
	check_CompiledPattern_Match(t, cp, "b", true, nil)
	check_CompiledPattern_Match(t, cp, "c", true, nil)
	check_CompiledPattern_Match(t, cp, "-", false, nil)
	check_CompiledPattern_Match(t, cp, "z", false, nil)
	check_CompiledPattern_Match(t, cp, "A", false, nil)
	check_CompiledPattern_Match(t, cp, "B", false, nil)
	check_CompiledPattern_Match(t, cp, "C", false, nil)
	check_CompiledPattern_Match(t, cp, "D", false, nil)
	check_CompiledPattern_Match(t, cp, "E", false, nil)

	pattern	=	"[-ac]"

	cp, err	=	shwild.Compile(pattern)
	if err != nil {

		t.Errorf("Failed to compile pattern '%s'", pattern)
	}

	check_CompiledPattern_Match(t, cp, "a", true, nil)
	check_CompiledPattern_Match(t, cp, "b", false, nil)
	check_CompiledPattern_Match(t, cp, "c", true, nil)
	check_CompiledPattern_Match(t, cp, "-", true, nil)
	check_CompiledPattern_Match(t, cp, "d", false, nil)
	check_CompiledPattern_Match(t, cp, "z", false, nil)
}

func Test_CompiledPattern_Match_with_forward_continuum_notrange(t *testing.T) {

	pattern	:=	"[^a-c]"

	cp, err	:=	shwild.Compile(pattern)
	if err != nil {

		t.Errorf("Failed to compile pattern '%s'", pattern)
	}

	check_CompiledPattern_Match(t, cp, "a", false, nil)
	check_CompiledPattern_Match(t, cp, "b", false, nil)
	check_CompiledPattern_Match(t, cp, "c", false, nil)
	check_CompiledPattern_Match(t, cp, "-", true, nil)
	check_CompiledPattern_Match(t, cp, "z", true, nil)
	check_CompiledPattern_Match(t, cp, "A", true, nil)
	check_CompiledPattern_Match(t, cp, "B", true, nil)
	check_CompiledPattern_Match(t, cp, "C", true, nil)
	check_CompiledPattern_Match(t, cp, "D", true, nil)
	check_CompiledPattern_Match(t, cp, "E", true, nil)

	pattern	=	"[^-ac]"

	cp, err	=	shwild.Compile(pattern)
	if err != nil {

		t.Errorf("Failed to compile pattern '%s'", pattern)
	}

	check_CompiledPattern_Match(t, cp, "a", false, nil)
	check_CompiledPattern_Match(t, cp, "b", true, nil)
	check_CompiledPattern_Match(t, cp, "c", false, nil)
	check_CompiledPattern_Match(t, cp, "-", false, nil)
	check_CompiledPattern_Match(t, cp, "d", true, nil)
	check_CompiledPattern_Match(t, cp, "z", true, nil)
}

func Test_CompiledPattern_Match_with_backward_continuum_range(t *testing.T) {

	pattern	:=	"[c-a]"

	cp, err	:=	shwild.Compile(pattern)
	if err != nil {

		t.Errorf("Failed to compile pattern '%s'", pattern)
	}

	check_CompiledPattern_Match(t, cp, "a", true, nil)
	check_CompiledPattern_Match(t, cp, "b", true, nil)
	check_CompiledPattern_Match(t, cp, "c", true, nil)
	check_CompiledPattern_Match(t, cp, "-", false, nil)
	check_CompiledPattern_Match(t, cp, "z", false, nil)
	check_CompiledPattern_Match(t, cp, "A", false, nil)
	check_CompiledPattern_Match(t, cp, "B", false, nil)
	check_CompiledPattern_Match(t, cp, "C", false, nil)
	check_CompiledPattern_Match(t, cp, "D", false, nil)
	check_CompiledPattern_Match(t, cp, "E", false, nil)
}

func Test_CompiledPattern_Match_with_backward_continuum_notrange(t *testing.T) {

	pattern	:=	"[^c-a]"

	cp, err	:=	shwild.Compile(pattern)
	if err != nil {

		t.Errorf("Failed to compile pattern '%s'", pattern)
	}

	check_CompiledPattern_Match(t, cp, "a", false, nil)
	check_CompiledPattern_Match(t, cp, "b", false, nil)
	check_CompiledPattern_Match(t, cp, "c", false, nil)
	check_CompiledPattern_Match(t, cp, "-", true, nil)
	check_CompiledPattern_Match(t, cp, "z", true, nil)
	check_CompiledPattern_Match(t, cp, "A", true, nil)
	check_CompiledPattern_Match(t, cp, "B", true, nil)
	check_CompiledPattern_Match(t, cp, "C", true, nil)
	check_CompiledPattern_Match(t, cp, "D", true, nil)
	check_CompiledPattern_Match(t, cp, "E", true, nil)
}

func Test_CompiledPattern_Match_with_forward_crosscase_continuum_range(t *testing.T) {

	pattern	:=	"[a-C]"

	cp, err	:=	shwild.Compile(pattern)
	if err != nil {

		t.Errorf("Failed to compile pattern '%s'", pattern)
	}

	check_CompiledPattern_Match(t, cp, "-", false, nil)
	check_CompiledPattern_Match(t, cp, "a", true, nil)
	check_CompiledPattern_Match(t, cp, "b", true, nil)
	check_CompiledPattern_Match(t, cp, "c", true, nil)
	check_CompiledPattern_Match(t, cp, "d", false, nil)
	check_CompiledPattern_Match(t, cp, "z", false, nil)
	check_CompiledPattern_Match(t, cp, "A", true, nil)
	check_CompiledPattern_Match(t, cp, "B", true, nil)
	check_CompiledPattern_Match(t, cp, "C", true, nil)
	check_CompiledPattern_Match(t, cp, "D", false, nil)
	check_CompiledPattern_Match(t, cp, "E", false, nil)
}

func Test_CompiledPattern_Match_with_forward_crosscase_continuum_notrange(t *testing.T) {

	pattern	:=	"[^a-C]"

	cp, err	:=	shwild.Compile(pattern)
	if err != nil {

		t.Errorf("Failed to compile pattern '%s'", pattern)
	}

	check_CompiledPattern_Match(t, cp, "-", true, nil)
	check_CompiledPattern_Match(t, cp, "a", false, nil)
	check_CompiledPattern_Match(t, cp, "b", false, nil)
	check_CompiledPattern_Match(t, cp, "c", false, nil)
	check_CompiledPattern_Match(t, cp, "d", true, nil)
	check_CompiledPattern_Match(t, cp, "z", true, nil)
	check_CompiledPattern_Match(t, cp, "A", false, nil)
	check_CompiledPattern_Match(t, cp, "B", false, nil)
	check_CompiledPattern_Match(t, cp, "C", false, nil)
	check_CompiledPattern_Match(t, cp, "D", true, nil)
	check_CompiledPattern_Match(t, cp, "E", true, nil)
}

func Test_CompiledPattern_Match_with_escaped_special_characters(t *testing.T) {

	pattern	:=	"a\\*c"

	cp, err	:=	shwild.Compile(pattern)
	if err != nil {

		t.Errorf("Failed to compile pattern '%s'", pattern)
	}

	check_CompiledPattern_Match(t, cp, "a_c", false, nil)
	check_CompiledPattern_Match(t, cp, "a*c", true, nil)

	pattern	=	"a\\?c"

	cp, err	=	shwild.Compile(pattern)
	if err != nil {

		t.Errorf("Failed to compile pattern '%s'", pattern)
	}

	check_CompiledPattern_Match(t, cp, "a_c", false, nil)
	check_CompiledPattern_Match(t, cp, "a?c", true, nil)

	pattern	=	"a\\[c"

	cp, err	=	shwild.Compile(pattern)
	if err != nil {

		t.Errorf("Failed to compile pattern '%s'", pattern)
	}

	check_CompiledPattern_Match(t, cp, "a_c", false, nil)
	check_CompiledPattern_Match(t, cp, "a[c", true, nil)

	pattern	=	"a\\]c"

	cp, err	=	shwild.Compile(pattern)
	if err != nil {

		t.Errorf("Failed to compile pattern '%s'", pattern)
	}

	check_CompiledPattern_Match(t, cp, "a_c", false, nil)
	check_CompiledPattern_Match(t, cp, "a]c", true, nil)

	pattern	=	"a]c"

	cp, err	=	shwild.Compile(pattern)
	if err != nil {

		t.Errorf("Failed to compile pattern '%s'", pattern)
	}

	check_CompiledPattern_Match(t, cp, "a_c", false, nil)
	check_CompiledPattern_Match(t, cp, "a]c", true, nil)
}

/* ///////////////////////////// end of file //////////////////////////// */


