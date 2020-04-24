package idiom

import (
	"os"
	"testing"
)

func TestFileFunctionOptions(t *testing.T) {
	err := New("empty.txt")
	if err != nil {
		panic(err)
	}
	os.Remove("empty.txt")

	err = New("file.txt", UID(1000), Contents("input some data"))
	if err != nil {
		panic(err)
	}
	os.Remove("file.txt")
}

///Options is key struct
type Options struct {
	UID         int
	GID         int
	Flags       int
	Contents    string
	Permissions os.FileMode
}

//Option func is key func
type Option func(*Options)

func UID(userID int) Option {
	return func(args *Options) {
		args.UID = userID
	}
}

func GID(groupID int) Option {
	return func(args *Options) {
		args.GID = groupID
	}
}

func Contents(c string) Option {
	return func(args *Options) {
		args.Contents = c
	}
}

func Permissions(perms os.FileMode) Option {
	return func(args *Options) {
		args.Permissions = perms
	}
}

func New(filepath string, setters ...Option) error {
	// Default Options
	args := &Options{
		UID:         os.Getuid(),
		GID:         os.Getgid(),
		Contents:    "",
		Permissions: 0666,
		Flags:       os.O_CREATE | os.O_EXCL | os.O_WRONLY,
	}

	for _, setter := range setters {
		setter(args)
	}

	f, err := os.OpenFile(filepath, args.Flags, args.Permissions)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(args.Contents); err != nil {
		return err
	}

	return err
}
