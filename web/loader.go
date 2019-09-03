package web

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"plugin"
)

type loader struct {
	sourceDir string
	objectDir string
}

func NewLoader() (*loader, error) {
	// The directory that will be watched for new plugins.
	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("could not find current directory: %v", err)
	}
	//pluginsDir := filepath.Join(wd, "uploads")

	// The directory where all .so files will be stored.
	tmp, err := ioutil.TempDir("", "")
	if err != nil {
		return nil, fmt.Errorf("could not create objects dir: %v", err)
	}
	return &loader{sourceDir: wd, objectDir: tmp}, nil
}

func (l *loader) CompileAndRun(name string) error {
	obj, err := l.compile(name)
	if err != nil {
		return fmt.Errorf("could not compile %s: %v", name, err)
	}
	defer os.Remove(obj)

	if err := l.call(obj); err != nil {
		return fmt.Errorf("could not run %s: %v", obj, err)
	}
	return nil
}

// goFiles lists all the files in the plugins
func (l *loader) List() []string {
	dir, err := os.Open(l.sourceDir)
	if err != nil {
		log.Fatal(err)
	}
	defer dir.Close()
	names, err := dir.Readdirnames(-1)
	if err != nil {
		log.Fatal(err)
	}

	var res []string
	for _, name := range names {
		if filepath.Ext(name) == ".go" {
			res = append(res, name)
		}
	}
	return res
}

// compile compiles the code in the given path, builds a
// plugin, and returns its path.
func (l *loader) compile(name string) (string, error) {
	// Copy the file to the objects directory with a different name
	// each time, to avoid retrieving the cached version.
	// Apparently the cache key is the path of the file compiled and
	// there's no way to invalidate it.
	f, err := ioutil.ReadFile(filepath.Join(l.sourceDir, name))
	if err != nil {
		return "", fmt.Errorf("could not read %s: %v", name, err)
	}

	name = fmt.Sprintf("%d.go", rand.Int())
	srcPath := filepath.Join(l.objectDir, name)
	if err := ioutil.WriteFile(srcPath, f, 0666); err != nil {
		return "", fmt.Errorf("could not write %s: %v", name, err)
	}

	objectPath := srcPath[:len(srcPath)-3] + ".so"

	cmd := exec.Command("go", "build", "-buildmode=plugin", "-o="+objectPath, srcPath)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("could not compile %s: %v", name, err)
	}

	return objectPath, nil
}

// call loads the plugin object in the given path and runs the Run
// function.
func (l *loader) call(object string) error {
	p, err := plugin.Open(object)
	if err != nil {
		return fmt.Errorf("could not open %s: %v", object, err)
	}
	run, err := p.Lookup("Run")
	if err != nil {
		return fmt.Errorf("could not find Run function: %v", err)
	}
	runFunc, ok := run.(func() error)
	if !ok {
		return fmt.Errorf("found Run but type is %T instead of func() error", run)
	}
	if err := runFunc(); err != nil {
		return fmt.Errorf("plugin failed with error %v", err)
	}
	return nil
}
