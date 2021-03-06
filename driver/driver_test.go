// Copyright 2020 Longxiao Zhang <zhanglongx@gmail.com>.
// All rights reserved.
// Use of this source code is governed by a GPLv3-style
// license that can be found in the LICENSE file.

package driver

import (
	"testing"
)

func Test(t *testing.T) {
	test1 := map[string]interface{}{
		"root": map[string]interface{}{
			"data": 1,
		},
	}

	helperSetMap(test1, 0, "data", 20)
	if test1["root"].(map[string]interface{})["data"] != 20 {
		t.Error("failed: ", test1)
	}

	helperSetMap(test1, 0, "data1", 100)
	if test1["root"].(map[string]interface{})["data"] != 20 {
		t.Error("failed: ", test1)
	}

	helperSetMap(test1, 0, "root", 100)
	if test1["root"].(int) != 100 {
		t.Error("failed: ", test1)
	}

	helperSetMap(nil, 0, "data", 100)

	test2 := map[string]interface{}{
		"root": map[string]interface{}{
			"data": 1,
			"slice": []interface{}{
				map[string]interface{}{
					"sdata": 1,
				},
				map[string]interface{}{
					"sdata": 1,
				},
			},
		},
	}

	var ss []interface{}

	helperSetMap(test2, 0, "sdata", 100)
	ss = test2["root"].(map[string]interface{})["slice"].([]interface{})
	if ss[0].(map[string]interface{})["sdata"] != 100 {
		t.Error("failed: ", test2)
	}
	if ss[1].(map[string]interface{})["sdata"] != 1 {
		t.Error("failed: ", test2)
	}

	helperSetMap(test2, 100, "sdata", 100)
	helperSetMap(test2, 1, "", 100)
	if test2["root"].(map[string]interface{})["data"] != 1 {
		t.Error("failed: ", test2)
	}
}
