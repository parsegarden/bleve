//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package simple

import (
	"reflect"
	"testing"

	"github.com/blevesearch/bleve/search/highlight"
)

func TestSimpleFragmenter(t *testing.T) {

	tests := []struct {
		orig      []byte
		fragments []*highlight.Fragment
		ot        highlight.TermLocations
	}{
		{
			orig: []byte("this is a test"),
			fragments: []*highlight.Fragment{
				&highlight.Fragment{
					Orig:  []byte("this is a test"),
					Start: 0,
					End:   14,
				},
			},
			ot: highlight.TermLocations{
				&highlight.TermLocation{
					Term:  "test",
					Pos:   4,
					Start: 10,
					End:   14,
				},
			},
		},
		{
			orig: []byte("0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789"),
			fragments: []*highlight.Fragment{
				&highlight.Fragment{
					Orig:  []byte("0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789"),
					Start: 0,
					End:   100,
				},
			},
			ot: highlight.TermLocations{
				&highlight.TermLocation{
					Term:  "0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789",
					Pos:   1,
					Start: 0,
					End:   100,
				},
			},
		},
		{
			orig: []byte("01234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"),
			fragments: []*highlight.Fragment{
				&highlight.Fragment{
					Orig:  []byte("01234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"),
					Start: 0,
					End:   100,
				},
				&highlight.Fragment{
					Orig:  []byte("01234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"),
					Start: 10,
					End:   101,
				},
				&highlight.Fragment{
					Orig:  []byte("01234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"),
					Start: 20,
					End:   101,
				},
				&highlight.Fragment{
					Orig:  []byte("01234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"),
					Start: 30,
					End:   101,
				},
				&highlight.Fragment{
					Orig:  []byte("01234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"),
					Start: 40,
					End:   101,
				},
				&highlight.Fragment{
					Orig:  []byte("01234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"),
					Start: 50,
					End:   101,
				},
				&highlight.Fragment{
					Orig:  []byte("01234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"),
					Start: 60,
					End:   101,
				},
				&highlight.Fragment{
					Orig:  []byte("01234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"),
					Start: 70,
					End:   101,
				},
				&highlight.Fragment{
					Orig:  []byte("01234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"),
					Start: 80,
					End:   101,
				},
				&highlight.Fragment{
					Orig:  []byte("01234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"),
					Start: 90,
					End:   101,
				},
			},
			ot: highlight.TermLocations{
				&highlight.TermLocation{
					Term:  "0123456789",
					Pos:   1,
					Start: 0,
					End:   10,
				},
				&highlight.TermLocation{
					Term:  "0123456789",
					Pos:   2,
					Start: 10,
					End:   20,
				},
				&highlight.TermLocation{
					Term:  "0123456789",
					Pos:   3,
					Start: 20,
					End:   30,
				},
				&highlight.TermLocation{
					Term:  "0123456789",
					Pos:   4,
					Start: 30,
					End:   40,
				},
				&highlight.TermLocation{
					Term:  "0123456789",
					Pos:   5,
					Start: 40,
					End:   50,
				},
				&highlight.TermLocation{
					Term:  "0123456789",
					Pos:   6,
					Start: 50,
					End:   60,
				},
				&highlight.TermLocation{
					Term:  "0123456789",
					Pos:   7,
					Start: 60,
					End:   70,
				},
				&highlight.TermLocation{
					Term:  "0123456789",
					Pos:   8,
					Start: 70,
					End:   80,
				},
				&highlight.TermLocation{
					Term:  "0123456789",
					Pos:   9,
					Start: 80,
					End:   90,
				},
				&highlight.TermLocation{
					Term:  "0123456789",
					Pos:   10,
					Start: 90,
					End:   100,
				},
			},
		},
	}

	fragmenter := NewSimpleFragmenter(100)
	for _, test := range tests {
		fragments := fragmenter.Fragment(test.orig, test.ot)
		if !reflect.DeepEqual(fragments, test.fragments) {
			t.Errorf("expected %#v, got %#v", test.fragments, fragments)
			for _, fragment := range fragments {
				t.Logf("frag: %#v", fragment)
			}
		}
	}
}

func TestSimpleFragmenterWithSize(t *testing.T) {

	tests := []struct {
		orig      []byte
		fragments []*highlight.Fragment
		ot        highlight.TermLocations
	}{
		{
			orig: []byte("this is a test"),
			fragments: []*highlight.Fragment{
				&highlight.Fragment{
					Orig:  []byte("this is a test"),
					Start: 0,
					End:   5,
				},
				&highlight.Fragment{
					Orig:  []byte("this is a test"),
					Start: 9,
					End:   14,
				},
			},
			ot: highlight.TermLocations{
				&highlight.TermLocation{
					Term:  "this",
					Pos:   1,
					Start: 0,
					End:   5,
				},
				&highlight.TermLocation{
					Term:  "test",
					Pos:   4,
					Start: 10,
					End:   14,
				},
			},
		},
	}

	fragmenter := NewSimpleFragmenter(5)
	for _, test := range tests {
		fragments := fragmenter.Fragment(test.orig, test.ot)
		if !reflect.DeepEqual(fragments, test.fragments) {
			t.Errorf("expected %#v, got %#v", test.fragments, fragments)
			for _, fragment := range fragments {
				t.Logf("frag: %#v", fragment)
			}
		}
	}
}