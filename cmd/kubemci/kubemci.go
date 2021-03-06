// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/GoogleCloudPlatform/k8s-multicluster-ingress/cmd/kubemci/app"
	"github.com/golang/glog"
)

func main() {
	// Workaround for https://github.com/kubernetes/kubernetes/issues/17162
	flag.Parse()
	if glog.V(2) {
		fmt.Println("Turning on debug logging to STDERR")
		flag.Set("logtostderr", "true")
	}

	if err := app.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
