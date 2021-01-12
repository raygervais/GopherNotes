package edit

/* Ported from:
 * https://samrapdev.com/capturing-sensitive-input-with-editor-in-golang-from-the-cli/
 */

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

const (
	// DefaultEditor currently only supports Linux and TUI semantics
	DefaultEditor = "vim"
)

// PreferredEditorResolver is a function that returns an editor that the user
// prefers to use, such as the configured `$EDITOR` environment variable.
type PreferredEditorResolver func() string

// GetPreferredEditorFromEnvironment returns the user's editor as defined by the
// `$EDITOR` environment variable, or the `DefaultEditor` if it is not set.
func GetPreferredEditorFromEnvironment() string {
	editor := os.Getenv("EDITOR")

	if editor == "" {
		return DefaultEditor
	}

	return editor
}

func resolveEditorArguments(executable string, filename string) []string {
	args := []string{filename}

	if strings.Contains(executable, "Visual Studio Code.app") || strings.Contains(executable, "code") {
		args = append([]string{"--wait"}, args...)
	}

	// Other common editors

	return args
}

// OpenFileInEditor allows us to edit the temp file in default $EDITOR
func OpenFileInEditor(filename string, resolveEditor PreferredEditorResolver) error {
	// Get the full executable path for the editor.
	executable, err := exec.LookPath(resolveEditor())
	if err != nil {
		return err
	}

	cmd := exec.Command(executable, resolveEditorArguments(executable, filename)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// CaptureInputFromEditor opens the temp file created and reads changes
// Returns changes found in the tempfile prior to cleaning up tempdir
func CaptureInputFromEditor(id int, note, date string) ([]byte, error) {
	file, err := ioutil.TempFile(os.TempDir(), "*")
	if err != nil {
		return []byte{}, err
	}
	fileName := file.Name()
	file.Write([]byte(note))

	defer os.Remove(fileName)

	if err = file.Close(); err != nil {
		return []byte{}, err
	}

	if err = OpenFileInEditor(fileName, GetPreferredEditorFromEnvironment); err != nil {
		return []byte{}, err
	}

	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return []byte{}, err
	}

	return bytes, nil
}
