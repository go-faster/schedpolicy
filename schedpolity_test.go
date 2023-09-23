package schedpolicy

import (
	"os/exec"
	"testing"
)

func TestSet(t *testing.T) {
	// Spawn some process.
	t.Run("Positive", func(t *testing.T) {
		cmd := exec.Command("sleep", "100")
		if err := cmd.Start(); err != nil {
			t.Fatal(err)
		}
		t.Cleanup(func() {
			_ = cmd.Process.Kill()
		})
		pid := cmd.Process.Pid
		if err := Set(pid, Batch, 0); err != nil {
			t.Fatal(err)
		}
		v, err := Get(pid)
		if err != nil {
			t.Fatal(err)
		}
		if v != Batch {
			t.Fatalf("unexpected policy: %v", v)
		}
	})
	t.Run("Negative", func(t *testing.T) {
		if err := Set(-1, Batch, 0); err == nil {
			t.Fatal("expected error")
		}
		if _, err := Get(-1); err == nil {
			t.Fatal("expected error")
		}
	})
}
