/*
   Copyright The containerd-compose Authors.

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

/*
   file created by mc256.com in 2020
*/

package composer

import (
	"fmt"
	"testing"
)

func Test_ComposerLoaderProvideComposeFile(t *testing.T) {
	var opts []Option

	//opts = append(opts, WithComposeFile("./test.yml"))
	opts = append(opts, WithComposeFile("./docker-compose.yml"))
	var compose *ComposeFile
	var err error
	if compose, err = loadFile(&opts); err != nil {
		t.Error(err)
	}
	if compose.Version != "2" {
		t.Error("version error")
	}
	fmt.Println(compose)
}

func Test_ComposerLoaderNonProvideComposeFile(t *testing.T) {
	var opts []Option

	opts = append(opts, WithComposeFile("./docker-compose.yml"), WithEnvFile("env.txt"))
	var compose *ComposeFile
	var err error
	if compose, err = loadFile(&opts); err != nil {
		t.Error(err)
	}
	if compose.Version != "2" {
		t.Error("version error")
	}
	fmt.Println(compose)
}
