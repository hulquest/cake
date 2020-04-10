package cmds

import (
	"reflect"
	"strings"
	"testing"
)

func TestCommandSuccessful(t *testing.T) {
	thisFilesName := "cmds_test.go"
	c := NewCommandLine(nil, "ls", nil, nil)
	stdout, stderr, _ := c.Program().Execute()
	if !strings.Contains(string(stdout), thisFilesName) {
		t.Errorf("expected %v to be in stdout, stdout: %v, stderr: %v", thisFilesName, string(stdout), string(stderr))
	}
}

func TestCommandNotFound(t *testing.T) {
	cmd := "im-not-a-command"
	c := NewCommandLine(nil, cmd, nil, nil)
	stdout, stderr, err := c.Program().Execute()
	if err.Error() != "exec: \"im-not-a-command\": executable file not found in $PATH" {
		t.Errorf("expected command '%v' to NOT be found, stdout: %v, stderr: %v, err: %v", cmd, string(stdout), string(stderr), err)
	}
}

func TestCmdLinkedList(t *testing.T) {
	kubectl := NewCommandLine(nil, "ls", nil, nil)
	clusterctl := NewCommandLine(nil, "pwd", nil, nil)
	tridentctl := NewCommandLine(nil, "rm", nil, nil)
	govc := NewCommandLine(nil, "cd", nil, nil)

	root := ProvisionerCommands{Name: "capv required commands"}
	root.AddCommand(kubectl.CommandName, kubectl)
	root.AddCommand(clusterctl.CommandName, clusterctl)
	root.AddCommand(tridentctl.CommandName, tridentctl)
	root.AddCommand(govc.CommandName, govc)

	all := root.GetAll()
	exist := root.Exist()

	expectedAll := []string{"ls", "pwd", "cd", "rm"}
	expectedExist := make([]string, 0)
	if reflect.DeepEqual(all, expectedAll) {
		t.Errorf("got %v, want %v", all, expectedAll)
	}
	if reflect.DeepEqual(exist, expectedExist) {
		t.Errorf("got %v, want %v", exist, expectedExist)
	}

	idontexist := NewCommandLine(nil, "idontexist", nil, nil)
	root.AddCommand(idontexist.CommandName, idontexist)
	exist = root.Exist()
	expectedExist = []string{"idontexist"}
	if reflect.DeepEqual(exist, expectedExist) {
		t.Errorf("got %v, want %v", exist, expectedExist)
	}

}
