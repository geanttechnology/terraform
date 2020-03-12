package command

import (
	"os"
	"strings"
	"testing"

	"github.com/mitchellh/cli"
)

func TestZeroThirteenUpgrade_invalidFlags(t *testing.T) {
	td := tempDir(t)
	os.MkdirAll(td, 0755)
	defer os.RemoveAll(td)
	defer testChdir(t, td)()

	ui := new(cli.MockUi)
	c := &ZeroThirteenUpgradeCommand{
		Meta: Meta{
			testingOverrides: metaOverridesForProvider(testProvider()),
			Ui:               ui,
		},
	}

	if code := c.Run([]string{"--whoops"}); code == 0 {
		t.Fatal("expected error, got:", ui.OutputWriter)
	}

	errMsg := ui.ErrorWriter.String()
	if !strings.Contains(errMsg, "Usage: terraform 0.13upgrade") {
		t.Fatal("unexpected error:", errMsg)
	}
}

func TestZeroThirteenUpgrade_empty(t *testing.T) {
	td := tempDir(t)
	os.MkdirAll(td, 0755)
	defer os.RemoveAll(td)
	defer testChdir(t, td)()

	ui := new(cli.MockUi)
	c := &ZeroThirteenUpgradeCommand{
		Meta: Meta{
			testingOverrides: metaOverridesForProvider(testProvider()),
			Ui:               ui,
		},
	}

	if code := c.Run(nil); code == 0 {
		t.Fatal("expected error, got:", ui.OutputWriter)
	}

	errMsg := ui.ErrorWriter.String()
	if !strings.Contains(errMsg, "Not a module directory") {
		t.Fatal("unexpected error:", errMsg)
	}
}

func TestZeroThirteenUpgrade_invalidProviderVersion(t *testing.T) {
	// FIXME: invalid provider version results in error exit
}

func TestZeroThirteenUpgrade_noProviders(t *testing.T) {
	// FIXME: no providers used in configuration results in successful exit
}

func TestZeroThirteenUpgrade_multipleProviders(t *testing.T) {
	// FIXME: hey it's the happy path!
}

func TestZeroThirteenUpgrade_providersFileExists(t *testing.T) {
	// FIXME: coverage for the unused filename generator
}
