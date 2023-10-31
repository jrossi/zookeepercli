/*
   Copyright 2014 Outbrain Inc.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

// output provides with controlled printing functions
package output

import (
	"encoding/json"
	"fmt"
	"strings"
)

// PrintString prints a single string to standard output, in either plaintext format or JSON format
func PrintString(data []byte, formatType string) {

	switch formatType {
	case "txt":
		fmt.Println(fmt.Sprintf("%s", data))
	case "json":
		s, _ := json.Marshal(string(data))
		fmt.Println(string(s))
	}
}

// PrintString prints a string array to standard output, in either plaintext format (one string per row) or JSON format
func PrintStringArray(stringArray []string, formatType string) {

	switch formatType {
	case "txt":
		fmt.Println(strings.Join(stringArray, "\n"))
	case "json":
		s, _ := json.Marshal(stringArray)
		fmt.Println(string(s))
	}
}

// Pair is a simple struct used to print map values. We use pairs
// instead of maps, so we can order the entries.
type Pair struct {
	Key   string
	Value string
}

// PrintMap prints out a map, given an ordered set of key, value pairs.
func PrintMap(items []Pair, formatType string) {
	switch formatType {
	case "txt":
		for _, item := range items {
			fmt.Printf("%s: %s\n", item.Key, item.Value)
		}
	case "json":
		fmt.Println("{")
		for i, item := range items {
			maybeComma := ","
			if i == len(items)-1 {
				maybeComma = ""
			}
			fmt.Printf("  %q: %q%s\n", item.Key, item.Value, maybeComma)
		}
		fmt.Println("}")
	}
}
