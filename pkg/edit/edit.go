package edit

/* Ported from:
 * https://samrapdev.com/capturing-sensitive-input-with-editor-in-golang-from-the-cli/
 */

import (
	"io/ioutil"
	"os"
	"os/exec"
)

const (
	DefaultEditor = "vim"
)

func OpenFileInEditor(filename string) error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = DefaultEditor
	}

	// Get the full executable path for the editor.
	executable, err := exec.LookPath(editor)
	if err != nil {
		return err
	}

	cmd := exec.Command(executable, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

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

	if err = OpenFileInEditor(fileName); err != nil {
		return []byte{}, err
	}

	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return []byte{}, err
	}

	return bytes, nil
}
