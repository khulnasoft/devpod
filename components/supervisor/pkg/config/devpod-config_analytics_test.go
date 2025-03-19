// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package config

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/khulnasoft/devpod/common-go/log"
	devpod "github.com/khulnasoft/devpod/devpod-protocol"
)

func TestAnalyzeGitpodConfig(t *testing.T) {
	tests := []struct {
		Desc    string
		Prev    *devpod.GitpodConfig
		Current *devpod.GitpodConfig
		Fields  []string
	}{
		{
			Desc: "change",
			Prev: &devpod.GitpodConfig{
				CheckoutLocation: "foo",
			},
			Current: &devpod.GitpodConfig{
				CheckoutLocation: "bar",
			},
			Fields: []string{"CheckoutLocation"},
		},
		{
			Desc: "add",
			Prev: &devpod.GitpodConfig{},
			Current: &devpod.GitpodConfig{
				CheckoutLocation: "bar",
			},
			Fields: []string{"CheckoutLocation"},
		},
		{
			Desc: "remove",
			Prev: &devpod.GitpodConfig{
				CheckoutLocation: "bar",
			},
			Current: &devpod.GitpodConfig{},
			Fields:  []string{"CheckoutLocation"},
		},
		{
			Desc: "none",
			Prev: &devpod.GitpodConfig{
				CheckoutLocation: "bar",
			},
			Current: &devpod.GitpodConfig{
				CheckoutLocation: "bar",
			},
			Fields: nil,
		},
		{
			Desc:    "fie created",
			Current: &devpod.GitpodConfig{},
			Fields:  nil,
		},
	}
	for _, test := range tests {
		t.Run(test.Desc, func(t *testing.T) {
			var fields []string
			analyzer := NewConfigAnalyzer(log.Log, 100*time.Millisecond, func(field string) {
				fields = append(fields, field)
			}, test.Prev)
			<-analyzer.Analyse(test.Current)
			if diff := cmp.Diff(test.Fields, fields); diff != "" {
				t.Errorf("unexpected output (-want +got):\n%s", diff)
			}
		})
	}
}
