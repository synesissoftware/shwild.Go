/* /////////////////////////////////////////////////////////////////////////
 * File:        nodes.go
 *
 * Purpose:     Nodes (shwild.Go)
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

	"bytes"
	"fmt"
	"strings"
	"unicode"
)

/* /////////////////////////////////////////////////////////////////////////
 * API types
 */

/* /////////////////////////////////////////////////////////////////////////
 * internal types
 */

// _TokenType enumeration

type _TokenType	int

const (

	_TOK_INVALID _TokenType = -1 + iota
	_TOK_START
	_TOK_END
	_TOK_LITERAL
	_TOK_WILD_1
	_TOK_WILD_N
	_TOK_RANGE_BEG
	_TOK_NOT_RANGE
	_TOK_RANGE
	_TOK_ENOMEM
	_TOK_ESCAPED_
)

func (tt _TokenType) String() string {

	switch tt {

	case _TOK_INVALID: return "_TOK_INVALID"
	case _TOK_START: return "_TOK_START"
	case _TOK_END: return "_TOK_END"
	case _TOK_LITERAL: return "_TOK_LITERAL"
	case _TOK_WILD_1: return "_TOK_WILD_1"
	case _TOK_WILD_N: return "_TOK_WILD_N"
	case _TOK_RANGE_BEG: return "_TOK_RANGE_BEG"
	case _TOK_NOT_RANGE: return "_TOK_NOT_RANGE"
	case _TOK_RANGE: return "_TOK_RANGE"
	case _TOK_ENOMEM: return "_TOK_ENOMEM"
	case _TOK_ESCAPED_: return "_TOK_ESCAPED_"
	}

	return fmt.Sprintf("<%T %d>", tt, tt)
}

// _NodeType enumeration

type _NodeType int

const (

	_NODE_NOTHING _NodeType = iota
	_NODE_WILD_1
	_NODE_WILD_N
	_NODE_RANGE
	_NODE_NOT_RANGE
	_NODE_LITERAL
	_NODE_END
)

func (nt _NodeType) String() string {

	switch nt {

	case _NODE_NOTHING: return "_NODE_NOTHING"
	case _NODE_WILD_1: return "_NODE_WILD_1"
	case _NODE_WILD_N: return "_NODE_WILD_N"
	case _NODE_RANGE: return "_NODE_RANGE"
	case _NODE_NOT_RANGE: return "_NODE_NOT_RANGE"
	case _NODE_LITERAL: return "_NODE_LITERAL"
	case _NODE_END: return "_NODE_END"
	}

	return fmt.Sprintf("<%T %d>", nt, nt)
}


// node structure

type node struct {

	node_type	_NodeType
	flags		uint64
	data		string
}

func (n node) String() string {

	return fmt.Sprintf("<%T{ node_type=%v, flags=0x%x, data=%q}>", n, n.node_type, n.flags, n.data)
}

func make_node(node_type _NodeType, flags uint64, data string) (n node) {

	return node { node_type: node_type, flags: flags, data: data }
}

func make_range_node(node_type _NodeType, flags uint64, data string) (n node) {

	if strings.ContainsRune(data[1:], '-') {

		end_index := len(data) - 1
		var buff bytes.Buffer
		var from_rune rune
		from_index := -1

		for ix, ch := range data {

			if '-' == ch && (0 != ix && end_index != ix) {

				from_index = ix - 1

				continue
			}

			if from_index + 2 == ix {

				to_rune := ch

				if unicode.IsLetter(from_rune) && unicode.IsLetter(to_rune) && unicode.IsLower(from_rune) != unicode.IsLower(to_rune) {

					// Have to treat this differently

					var from_lower = int(unicode.ToLower(from_rune))
					var to_lower = int(unicode.ToLower(to_rune))

					var from_upper = int(unicode.ToUpper(from_rune))
					var to_upper = int(unicode.ToUpper(to_rune))

					if to_lower < from_lower {

						from_lower, to_lower = to_lower, from_lower
					}

					if to_upper < from_upper {

						from_upper, to_upper = to_upper, from_upper
					}

					write_range(&buff, from_lower, to_lower + 1)
					write_range(&buff, from_upper, to_upper + 1)

					continue
				}

				var from int = int(from_rune)
				var to int = int(to_rune)

				if to < from {

					from, to = to, from
				}

				if from < to {

					write_range(&buff, from, to)
				}
			}

			buff.WriteRune(ch)
			from_rune = ch
		}

		return make_node(node_type, flags, buff.String())
	} else {

		return make_node(node_type, flags, data)
	}
}

func write_range(buff *bytes.Buffer, from, to int) {

	for i := from; i != to; i++ {

		buff.WriteRune(rune(i))
	}
}

/* /////////////////////////////////////////////////////////////////////////
 * API functions
 */


/* /////////////////////////////////////////////////////////////////////////
 * internal functions
 */

func parse_nodes(pattern string, flags uint64) (nodes []node, err error) {

	state := _TOK_LITERAL
	prev_state := _TOK_LITERAL

	var data []rune

	for _, ch := range pattern {

		switch state {

		case _TOK_ESCAPED_:

			state = prev_state
			data = append(data, ch)
		case _TOK_LITERAL, _TOK_START:

			switch(ch) {

			case '\\':

				prev_state = state
				state = _TOK_ESCAPED_

			case '?', '*', '[':

				if 0 != len(data) {

					node	:=	make_node(_NODE_LITERAL, flags, string(data))
					nodes	=	append(nodes, node)
					data	=	make([]rune, 0)
				}

				switch(ch) {

				case '?':

					node	:=	make_node(_NODE_WILD_1, flags, "")
					nodes	=	append(nodes, node)
				case '*':

					node	:=	make_node(_NODE_WILD_N, flags, "")
					nodes	=	append(nodes, node)
				case '[':

					state = _TOK_RANGE_BEG
				}
			default:

				state = _TOK_LITERAL
				data = append(data, ch)
			}
		case _TOK_RANGE_BEG:

			switch(ch) {

			case '^':

				state = _TOK_NOT_RANGE
			default:

				state = _TOK_RANGE
				data = append(data, ch)
			}
		case _TOK_RANGE, _TOK_NOT_RANGE:

			if ']' == ch && 0 != len(data) {

				var n node

				switch state {

				case _TOK_RANGE:
					n = make_range_node(_NODE_RANGE, flags, string(data))

				case _TOK_NOT_RANGE:
					n = make_range_node(_NODE_NOT_RANGE, flags, string(data))
				}

				nodes	=	append(nodes, n)
				data	=	make([]rune, 0)
				state	=	_TOK_START
			} else {

				data = append(data, ch)
			}
		default:
		}
	}

	switch state {

	case _TOK_LITERAL:

		node	:=	make_node(_NODE_LITERAL, flags, string(data))
		nodes	=	append(nodes, node)
	case _TOK_WILD_1:

		node	:=	make_node(_NODE_WILD_1, flags, "")
		nodes	=	append(nodes, node)
	case _TOK_WILD_N:

		node	:=	make_node(_NODE_WILD_N, flags, "")
		nodes	=	append(nodes, node)
	}

	node	:=	make_node(_NODE_END, flags, "")
	nodes	=	append(nodes, node)

	return
}

/* ///////////////////////////// end of file //////////////////////////// */


