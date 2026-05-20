package activitylog

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"time"
)

const logFile = "activity.log"

type Action string

const (
	ActionCreate  Action = "CREATE"
	ActionEdit    Action = "EDIT"
	ActionDelete  Action = "DELETE"
	ActionArchive Action = "ARCHIVE"
)

type Entry struct {
	Timestamp   time.Time
	Action      Action
	SubjectType string
	SubjectName string
	ProjectRoot string
	User        string
	Hostname    string
	Version     string
}

// AppVersion should be set from main.go same way as Version in cmd/version.go
var AppVersion = "dev"

func Log(projectRoot string, action Action, subjectType, subjectName string) error {
	entry, err := buildEntry(projectRoot, action, subjectType, subjectName)
	if err != nil {
		return err
	}
	return write(entry)
}

func buildEntry(projectRoot string, action Action, subjectType, subjectName string) (Entry, error) {
	u, err := user.Current()
	if err != nil {
		u = &user.User{Username: "unknown"}
	}

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	return Entry{
		Timestamp:   time.Now().UTC(),
		Action:      action,
		SubjectType: subjectType,
		SubjectName: subjectName,
		ProjectRoot: projectRoot,
		User:        u.Username,
		Hostname:    hostname,
		Version:     AppVersion,
	}, nil
}

func write(e Entry) error {
	path := filepath.Join(e.ProjectRoot, logFile)

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("activitylog: failed to open log file: %w", err)
	}
	defer f.Close()

	line := fmt.Sprintf("%s\t%-8s\t%-12s\t%q\t%s@%s\t%s\n",
		e.Timestamp.Format(time.RFC3339),
		e.Action,
		e.SubjectType,
		e.SubjectName,
		e.User,
		e.Hostname,
		e.Version,
	)

	_, err = f.WriteString(line)
	return err
}
