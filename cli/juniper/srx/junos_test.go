// NetD makes network device operations easy.
// Copyright (C) 2019  sky-cloud.net
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package srx

import (
	"testing"

	"github.com/sky-cloud-tec/netd/cli"
	. "github.com/smartystreets/goconvey/convey"
)

func TestJunosOp(t *testing.T) {

	Convey("srx op", t, func() {
		op := createOpJunos()
		So(
			cli.Match(op.GetPrompts("login"), "admin@hostname> "),
			ShouldBeTrue,
		)
		So(
			cli.Match(op.GetPrompts("configure"), "admin# "),
			ShouldBeTrue,
		)
	})
}