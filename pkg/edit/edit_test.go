package edit_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/raygervais/gophernotes/pkg/conf"
	"github.com/raygervais/gophernotes/pkg/edit"
	"github.com/raygervais/gophernotes/test"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestEditorEnvironmentLookup(t *testing.T) {
	db := test.SetupDatabase(t)
	defer test.TeardownDatabase(t)

	for i := 0; i < 10; i++ {
		db.Create(fmt.Sprintf("TestEditorModule%d", i))
	}

	testCases := []struct {
		desc   string
		editor string
	}{
		{
			desc:   "vim",
			editor: "vim",
		},
		{
			desc:   "nano",
			editor: "nano",
		},
		{
			desc:   "emacs",
			editor: "emacs",
		},
		{
			desc:   "VS Code",
			editor: "code",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			os.Setenv("EDITOR", tC.editor)
			test.ExpectToEqualString(t, conf.ExternalConfig.DefaultEditor, "vim")
			test.ExpectToEqualString(t, edit.GetPreferredEditorFromEnvironment(), tC.editor)
		})
	}
}
