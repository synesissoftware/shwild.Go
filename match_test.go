
package shwild

import "fmt"
import "path"
import "runtime"
import "testing"

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

func TestMatch_with_empty_pattern(t *testing.T) {

	check_Match(t, "", "", true, nil)
	check_Match(t, "", "1", false, nil)
	check_Match(t, "", "*", false, nil)
	check_Match(t, "", ".", false, nil)
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

