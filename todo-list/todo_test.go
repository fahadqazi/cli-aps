package todo_test

import (
	"github.com/fahadqazi/cli-apps/todo"
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	l := todo.List{}
	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("expected %q, got %q", taskName, l[0].Task)
	}
}

func TestComplete(t *testing.T) {
	l := todo.List{}
	taskName := "New task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("expected %q, got %q", taskName, l[0].Task)
	}

	if l[0].Done {
		t.Errorf("New task should not be complete")
	}

	l.Complete(1)

	if !l[0].Done {
		t.Errorf("New task should be complete")
	}
}

func TestDelete(t *testing.T) {
	l := todo.List{}

	tasks := []string{
		"Task one",
		"Task two",
		"Task three",
	}

	for _, v := range tasks {
		l.Add(v)
	}

	if l[0].Task != tasks[0] {
		t.Errorf("expected %q, got %q", tasks[0], l[0].Task)
	}

	l.Delete(2)

	if len(l) != 2 {
		t.Errorf("expected %d, got %d", 2, len(l))
	}

	if l[1].Task != tasks[2] {
		t.Errorf("expected %q, got %q", tasks[2], l[1].Task)
	}
}

func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}

	taskName := "New Task"
	l1.Add(taskName)

	if l1[0].Task != taskName {
		t.Errorf("expected %q, got %q", taskName, l1[0].Task)
	}

	tf, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("error creating temp file: %s", err)
	}
	defer os.Remove(tf.Name())

	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("error saving list to file: %s", err)
	}

	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("error getting list from file: %s", err)
	}

	if l1[0].Task != l2[0].Task {
		t.Errorf("expected %q, got %q", l1[0].Task, l2[0].Task)
	}
}
