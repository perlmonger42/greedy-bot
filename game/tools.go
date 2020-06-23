// Force `stringer` into go.mod, because it's needed for Piece.String().
// +build tools

package game

import _ "golang.org/x/tools/cmd/stringer"
