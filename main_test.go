// Copyright (c) 2017, Randy Westlund. All rights reserved.
// This code is under the BSD-2-Clause license.

package gotex

import (
	"testing"
)

func TestRender(t *testing.T) {
	var document = `
        \documentclass[12pt]{article}
        \begin{document}
        This is a LaTeX document.
        \end{document}
        `
	var pdf, err = Render(document, Options{})
	if err != nil {
		t.Error(err)
	}
	if len(pdf) < 1000 {
		t.Error("Generated PDF is too short", len(pdf))
	}

	document = `\error \invalid`
	pdf, err = Render(document, Options{})
	if err == nil {
		t.Error("Should fail on invalid document")
	}
	if pdf != nil {
		t.Error("Should not product a PDF on invalid document")
	}
}
