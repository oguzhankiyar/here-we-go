package main

import (
	"fmt"
	"os"
)

var baseDir string

func main() {
	Setup()

	Sample("Args", Args)
	Sample("Environ", Environ)
	Sample("Getenv", Getenv)
	Sample("Setenv", Setenv)
	Sample("Unsetenv", Unsetenv)
	Sample("Expand", Expand)
	Sample("Executable", Executable)
	Sample("Getpid", Getpid)
	Sample("Getwd", Getwd)
	Sample("Chdir", Chdir)
	Sample("Hostname", Hostname)
	Sample("UserHomeDir", UserHomeDir)
	Sample("UserCacheDir", UserCacheDir)
	Sample("UserConfigDir", UserConfigDir)
	Sample("TempDir", TempDir)
	Sample("DirFS", DirFS)
	Sample("Stat", Stat)
	Sample("Stat", Lstat)
	Sample("ReadDir", ReadDir)
	Sample("Mkdir", Mkdir)
	Sample("MkdirAll", MkdirAll)
	Sample("Create", Create)
	Sample("Open", Open)
	Sample("OpenFile", OpenFile)
	Sample("ReadFile", ReadFile)
	Sample("SameFile", SameFile)
	Sample("WriteFile", WriteFile)
	Sample("MkdirTemp", MkdirTemp)
	Sample("CreateTemp", CreateTemp)
	Sample("Rename", Rename)
	Sample("Truncate", Truncate)
	Sample("Remove", Remove)
	Sample("RemoveAll", RemoveAll)

	Clear()

	Sample("Exit", Exit)
}

func Setup() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	dir := wd + "/06-libraries/01-standard/19-os/01-os/temp"

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	err = os.WriteFile(dir+"/temp.txt", []byte("Hello, there!"), os.ModePerm)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	baseDir = dir
}

func Clear() {
	err := os.RemoveAll(baseDir)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
}

func Args() {
	// Try "go run main.go.go test 123"
	// Writes
	// ...main.go.go
	// test
	// 123

	args := os.Args

	for _, v := range args {
		fmt.Println(v)
	}
}

func Environ() {
	envs := os.Environ()

	for _, v := range envs {
		fmt.Printf("%s, ", v)
	}

	fmt.Println()
}

func Getenv() {
	fmt.Printf("$GOPATH -> %s\n", os.Getenv("GOPATH"))
}

func Setenv() {
	err := os.Setenv("TESTGO", "test")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("Set TESTGO as test")
}

func Unsetenv() {
	err := os.Unsetenv("TESTGO")
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("Unset TESTGO")
}

func Expand() {
	fn := func(str string, mapper func(string) string) {
		fmt.Printf("%s -> %s\n", str, os.Expand(str, mapper))
	}

	fn("${GOPATH}", os.Getenv)
	fn("$GREETING there!", func(str string) string {
		if str == "GREETING" {
			return "Hello"
		}

		return ""
	})
}

func Executable() {
	result, err := os.Executable()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println(result)
}

func Getpid() {
	fmt.Println(os.Getpid())
}

func Getwd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println(dir)
}

func Chdir() {
	err := os.Chdir("01-get-started")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println(dir)
}

func Hostname() {
	dir, err := os.Hostname()
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(dir)
}

func UserHomeDir() {
	dir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(dir)
}

func UserCacheDir() {
	dir, err := os.UserCacheDir()
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(dir)
}

func UserConfigDir() {
	dir, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(dir)
}

func TempDir() {
	dir := os.TempDir()

	fmt.Println(dir)
}

func DirFS() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fs := os.DirFS(wd)
	fmt.Printf("%#v\n", fs)
}

func Stat() {
	path := baseDir + "/temp.txt"

	info, err := os.Stat(path) // if file not exist, panics
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("Name:", info.Name())
	fmt.Println("Mode:", info.Mode())
	fmt.Println("Size:", info.Size())
	fmt.Println("ModTime:", info.ModTime())
}

func Lstat() {
	path := baseDir + "/temp2.txt"

	info, err := os.Lstat(path) // if file not exist, returns err
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("Name:", info.Name())
	fmt.Println("Mode:", info.Mode())
	fmt.Println("Size:", info.Size())
	fmt.Println("ModTime:", info.ModTime())
}

func ReadDir() {
	entries, err := os.ReadDir(baseDir)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	for _, v := range entries {
		fmt.Println(v.Name())
	}
}

func Mkdir() {
	// If exists, returns error

	path := baseDir + "/mkdir"

	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("Created", path)
}

func MkdirAll() {
	// Creates directories with parents if not exist
	// If exists, do nothing

	path := baseDir + "/mkdirall/sub"
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("Created", path)
}

func Create() {
	file, err := os.Create(baseDir + "/hey.json")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	_, _ = file.WriteString("{}")
	_ = file.Close()

	fmt.Println("Created", file.Name())
}

func Open() {
	file, err := os.Open(baseDir + "/hey.json")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	content := make([]byte, 16)

	_, err = file.Read(content)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	_ = file.Close()

	fmt.Printf("Content: %s\n", content)
}

func OpenFile() {
	file, err := os.OpenFile(baseDir + "/hey.json", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	content := make([]byte, 16)

	_, err = file.Read(content)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	_ = file.Close()

	fmt.Printf("Content: %s\n", content)
}

func ReadFile() {
	content, err := os.ReadFile(baseDir + "/hey.json")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("Content: %s\n", content)
}

func SameFile() {
	file1, err := os.Stat(baseDir + "/hey.json")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	file2, err := os.Stat(baseDir + "/hey.json")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	same := os.SameFile(file1, file2)

	fmt.Println(same)
}

func WriteFile() {
	err := os.MkdirAll(baseDir + "/writefile", os.ModePerm)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	err = os.WriteFile(baseDir + "/writefile/hello.json", []byte("{}"), os.ModePerm)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	content, err := os.ReadFile(baseDir + "/writefile/hello.json")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("Content: %s\n", content)
}

func MkdirTemp() {
	// Creates random folder in temp folder

	name, err := os.MkdirTemp("", "*-tmp")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println(name)

	err = os.Remove(name)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
}

func CreateTemp() {
	file, err := os.CreateTemp("", "createtemp-*.txt")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	file.WriteString("Hello")
	file.Close()

	fmt.Println(file.Name())
}

func Rename() {
	dir := baseDir + "/rename"

	oldPath := dir + "/old"
	newPath := dir + "/new"

	err := os.MkdirAll(oldPath, os.ModePerm)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	err = os.Rename(oldPath, newPath)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	for _, v := range entries {
		fmt.Printf("%s ", v.Name())
	}

	fmt.Println()
}

func Truncate() {
	dir := baseDir + "/truncate"

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	path := dir + "/truncate.txt"

	err = os.WriteFile(path, []byte("Hello there!"), os.ModePerm)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	err = os.Truncate(path, 5)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%s\n", content) // Hello
}

func Remove() {
	// Removes file or empty directory

	dir := baseDir + "/remove"

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	err = os.Remove(dir)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("Removed", dir)
}

func RemoveAll() {
	// Removes all even if directory is not empty

	dir := baseDir + "/removeall/sub"

	err := os.MkdirAll(dir + "/nested", os.ModePerm)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	err = os.RemoveAll(dir)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("Removed", dir)
}

func Exit() {
	fmt.Print("Exiting with code 100")
	os.Exit(100)
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}