/*
 * @Description: https://github.com/crazybber
 * @Author: Edward
 * @Date: 2020-05-11 10:44:52
 * @Last Modified by: Edward
 * @Last Modified time: 2020-05-11 10:44:52
 */

package parallelism

import (
	"fmt"
	"os"
	"sort"
	"testing"
)

func TestParallelism(t *testing.T) {
	// Calculate the MD5 sum of all files under the specified directory,
	// then print the results sorted by path name.
	m, err := MD5All(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	var paths []string
	for path := range m {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		fmt.Printf("%x  %s\n", m[path], path)
	}
}
