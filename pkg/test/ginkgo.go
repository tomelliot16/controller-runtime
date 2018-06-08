/*
Copyright 2016 The Kubernetes Authors.

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

package test

import (
	"fmt"

	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/types"
	"github.com/onsi/ginkgo"
)

var _ ginkgo.Reporter = NewlineReporter{}

// Print a newline after the default Reporter output so that the results are correctly parsed
// by test automation.
// See issue https://github.com/jstemmer/go-junit-report/issues/31
type NewlineReporter struct{}

// SpecSuiteWillBegin implements ginkgo.Reporter
func (NewlineReporter) SpecSuiteWillBegin(config config.GinkgoConfigType, summary *types.SuiteSummary) {}

// BeforeSuiteDidRun implements ginkgo.Reporter
func (NewlineReporter) BeforeSuiteDidRun(setupSummary *types.SetupSummary) {}

// AfterSuiteDidRun implements ginkgo.Reporter
func (NewlineReporter) AfterSuiteDidRun(setupSummary *types.SetupSummary) {}

// SpecWillRun implements ginkgo.Reporter
func (NewlineReporter) SpecWillRun(specSummary *types.SpecSummary) {}

// SpecDidComplete implements ginkgo.Reporter
func (NewlineReporter) SpecDidComplete(specSummary *types.SpecSummary) {}

// SpecSuiteDidEnd Prints a newline between "35 Passed | 0 Failed | 0 Pending | 0 Skipped" and "--- PASS:"
func (NewlineReporter) SpecSuiteDidEnd(summary *types.SuiteSummary) { fmt.Printf("\n") }
