// Copyright (c) 2021 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package external

import (
	"github.com/khulnasoft/devpod/installer/pkg/common"
	dbinit "github.com/khulnasoft/devpod/installer/pkg/components/database/init"
)

var Objects = common.CompositeRenderFunc(
	dbinit.Objects,
)
