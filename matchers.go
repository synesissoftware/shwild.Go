/* /////////////////////////////////////////////////////////////////////////
 * File:        matchers.go
 *
 * Purpose:     Matchers (shwild.Go)
 *
 * Created:     17th June 2005
 * Updated:     9th March 2017
 *
 * Home:        http://shwild.org/
 *
 * Copyright (c) 2005-2012, Matthew Wilson and Sean Kelly
 * Copyright (c) 2005-2017, Matthew Wilson and Synesis Software
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

import "fmt"
import "os"
import "strings"

/* /////////////////////////////////////////////////////////////////////////
 * API types
 */

/* /////////////////////////////////////////////////////////////////////////
 * internal types
 */

// matcher interface

type matcher interface {

	setNext(next matcher)

	match(s string) bool;
}


// literal_matcher : matcher structure

type literal_matcher struct {
	node
	next matcher
}
func make_literal_matcher(flags uint64, value string) matcher {

	var m literal_matcher

	m.node	=	make_node(_NODE_LITERAL, flags, value)
	m.next	=	nil

	return &m
}
func (m *literal_matcher) setNext(next matcher) {

	m.next = next
}
func (m literal_matcher) match(s string) bool {

fmt.Fprintf(os.Stderr, "literal_matcher.match(s: '%v')\n", s)

	l := len(m.node.data)

	if len(s) < l {

		return false
	}

	// TODO: do case-(in)sensitive

	if m.node.data == s[0:l] {

		rest := s[l:]

		return m.next.match(rest)
	}

	return false
}


// wild1_matcher : matcher structure

type wild1_matcher struct {
	node
	next matcher
}
func make_wild1_matcher(flags uint64, value string) matcher {

	var m wild1_matcher

	m.node	=	make_node(_NODE_WILD_1, flags, value)

	return &m
}
func (m *wild1_matcher) setNext(next matcher) {

	m.next = next
}
func (m wild1_matcher) match(s string) bool {

fmt.Fprintf(os.Stderr, "wild1_matcher.match(s: '%v')\n", s)

	if len(s) < 1 {

		return false
	}

	return m.next.match(s[1:])
}


// wildN_matcher : matcher structure

type wildN_matcher struct {
	node
	next matcher
}
func make_wildN_matcher(flags uint64, value string) matcher {

	var m wildN_matcher

	m.node	=	make_node(_NODE_WILD_N, flags, value)

	return &m
}
func (m *wildN_matcher) setNext(next matcher) {

	m.next = next
}
func (m wildN_matcher) match(s string) bool {

fmt.Fprintf(os.Stderr, "wildN_matcher.match(s: '%v')\n", s)

	for i := 0; i != 1 + len(s); i++ {

		if m.next.match(s[i:]) {

			return true
		}
	}

	return false
}


// range_matcher : matcher structure

type range_matcher struct {
	node
	next matcher
}
func make_range_matcher(flags uint64, value string) matcher {

	var m range_matcher

	m.node	=	make_range_node(_NODE_RANGE, flags, value)

	return &m
}
func (m *range_matcher) setNext(next matcher) {

	m.next = next
}
func (m range_matcher) match(s string) bool {

	if len(s) < 1 {

		return false
	}

	if !strings.Contains(m.node.data, s[0:1]) {

		return false
	}

	return m.next.match(s[1:])
}


// notrange_matcher : matcher structure

type notrange_matcher struct {
	node
	next matcher
}
func make_notrange_matcher(flags uint64, value string) matcher {

	var m notrange_matcher

	m.node	=	make_node(_NODE_NOT_RANGE, flags, value)

	return &m
}
func (m *notrange_matcher) setNext(next matcher) {

	m.next = next
}
func (m notrange_matcher) match(s string) bool {

	if len(s) < 1 {

		return false
	}

	return m.next.match(s[1:])
}


// end_matcher : matcher structure

type end_matcher struct {
	node
}
func make_end_matcher(flags uint64) matcher {

	var m end_matcher

	m.node	=	make_node(_NODE_END, flags, "")

	return &m
}
func (m *end_matcher) setNext(_ matcher) {

	panic("end_matcher should never be asked to setNext()")
}
func (m end_matcher) match(s string) bool {

	return 0 == len(s)
}

/* /////////////////////////////////////////////////////////////////////////
 * API functions
 */


/* /////////////////////////////////////////////////////////////////////////
 * internal functions
 */

func parse_matchers(pattern string, flags uint64) ([]matcher, error) {

	if 0 == len(pattern) {

		return nil, nil
	}

	// create the sequence of matchers

	var matchers []matcher

	nodes, err := parse_nodes(pattern, flags)

	if nil != err {

		return nil, err
	}

	for _, n := range nodes {

		switch n.node_type {

		case _NODE_NOTHING:
			break
		case _NODE_WILD_1:
			matchers = append(matchers, make_wild1_matcher(flags, n.data))
		case _NODE_WILD_N:
			matchers = append(matchers, make_wildN_matcher(flags, n.data))
		case _NODE_RANGE:
			matchers = append(matchers, make_range_matcher(flags, n.data))
		case _NODE_NOT_RANGE:
			matchers = append(matchers, make_notrange_matcher(flags, n.data))
		case _NODE_LITERAL:
			matchers = append(matchers, make_literal_matcher(flags, n.data))
		case _NODE_END:
			matchers = append(matchers, make_end_matcher(flags))
		default:
			panic(fmt.Sprintf("VIOLATION: unexpected node type %V", n.node_type))
		}
	}

	// tie the sequence of matchers together

	for i, m := range matchers {

		if 0 != i {

			matchers[i - 1].setNext(m)
		}
	}

	return matchers, nil
}

/* ///////////////////////////// end of file //////////////////////////// */


