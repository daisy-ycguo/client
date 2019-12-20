// Copyright © 2019 The Knative Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package flags

import (
	"fmt"
	"strings"
)

type ArrayList []string

func (list *ArrayList) String() string {
	str := ""
	for _, item := range *list {
		str = str + item + " "
	}
	return str
}

func (list *ArrayList) Set(value string) error {
	*list = append(*list, value)
	return nil
}

func (list *ArrayList) Type() string {
	return "[]string"
}

// ParseMapOptions to return a map type of filters
func (list *ArrayList) ParseMapOptions() (map[string]string, error) {
	mapopts := map[string]string{}
	for _, item := range *list {
		parts := strings.Split(item, "=")
		if len(parts) < 2 || parts[0] == "" || parts[1] == "" {
			return nil, fmt.Errorf("invalid filter %s", list)
		} else {
			if _, ok := mapopts[parts[0]]; ok {
				return nil, fmt.Errorf("duplicate key '%s' in filters %s", parts[0], list)
			}
			mapopts[parts[0]] = parts[1]
		}
	}
	return mapopts, nil
}

// ParseUpdateMapOptions to return a map type of filters
func (list *ArrayList) ParseUpdateMapOptions() (map[string]string, []string, error) {
	updates := map[string]string{}
	var removes []string
	for _, item := range *list {
		if strings.HasSuffix(item, "-") {
			removes = append(removes, item[:len(item)-1])
		} else {
			parts := strings.Split(item, "=")
			if len(parts) < 2 || parts[0] == "" || parts[1] == "" {
				return nil, nil, fmt.Errorf("invalid filter %s", list)
			}
			if _, ok := updates[parts[0]]; ok {
				return nil, nil, fmt.Errorf("duplicate key '%s' in filters %s", parts[0], list)
			}
			updates[parts[0]] = parts[1]
		}
	}
	return updates, removes, nil
}
