// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package validation

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCloudProfile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Provider-Local APIs Validation Suite")
}
