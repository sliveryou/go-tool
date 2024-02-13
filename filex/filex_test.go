package filex

import (
	"crypto/md5"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sliveryou/go-tool/mathx"
)

func TestMain(m *testing.M) {
	if err := setup(); err != nil {
		log.Fatalln(err)
	}
	code := m.Run()
	_ = teardown()
	os.Exit(code)
}

func setup() error {
	if err := os.MkdirAll("./testdata", 0o755); err != nil {
		return err
	}

	if err := ioutil.WriteFile("./testdata/test.txt",
		[]byte("Hello world!\n"),
		0o644); err != nil {
		return err
	}

	return ioutil.WriteFile("./testdata/test.md",
		[]byte("# Hello world!\n"),
		0o644)
}

func teardown() error {
	return os.RemoveAll("./testdata")
}

func TestExt(t *testing.T) {
	cases := []struct {
		fileName string
		expect   string
	}{
		{fileName: "./testdata/test.txt", expect: "txt"},
		{fileName: "./testdata/test.md", expect: "md"},
		{fileName: ".a.b.c", expect: "c"},
		{fileName: ".abc", expect: "abc"},
		{fileName: "abc", expect: ""},
		{fileName: ".", expect: ""},
		{fileName: "", expect: ""},
	}

	for _, c := range cases {
		get := Ext(c.fileName)
		assert.Equal(t, c.expect, get)
	}
}

func TestSize(t *testing.T) {
	size := Size("./testdata/test.txt")
	t.Log(size, mathx.SizeFormat(float64(size), 2))
	size = Size("./testdata/test.md")
	t.Log(size, mathx.SizeFormat(float64(size), 2))
}

func TestDirSize(t *testing.T) {
	fileNum, dirSize := DirSize("./testdata/testdir")
	t.Log(fileNum, dirSize, mathx.SizeFormat(float64(dirSize), 2))
	fileNum, dirSize = DirSize("./testdata")
	t.Log(fileNum, dirSize, mathx.SizeFormat(float64(dirSize), 2))
}

func TestHash(t *testing.T) {
	hashStr, err := Hash("./testdata/test.txt", md5.New())
	require.NoError(t, err)
	t.Log(hashStr)

	md5Str, err := MD5("./testdata/test.txt")
	require.NoError(t, err)
	t.Log(md5Str)

	sha1Str, err := SHA1("./testdata/test.txt")
	require.NoError(t, err)
	t.Log(sha1Str)

	sha256, err := SHA256("./testdata/test.txt")
	require.NoError(t, err)
	t.Log(sha256)

	sha512, err := SHA512("./testdata/test.txt")
	require.NoError(t, err)
	t.Log(sha512)
}

func TestRead(t *testing.T) {
	content, err := Read("./testdata/test.txt")
	require.NoError(t, err)
	t.Log("\n" + string(content))

	content, err = Read("./testdata/test.md")
	require.NoError(t, err)
	t.Log("\n" + string(content))
}

func TestReadToSlice(t *testing.T) {
	slice, err := ReadToSlice("./testdata/test.txt")
	require.NoError(t, err)
	t.Log(slice)
}

func TestWrite(t *testing.T) {
	data := []byte("Hello, world!")
	err := Write("./testdata/testdir/testfile.txt", data)
	require.NoError(t, err)
}

func TestAppend(t *testing.T) {
	data := []byte("\nHello, world!")
	err := Append("./testdata/testdir/testfile.txt", data)
	require.NoError(t, err)
}

func TestCopy(t *testing.T) {
	err := Copy("./testdata/test.md", "./testdata/testdir/test.md")
	require.NoError(t, err)
}

func TestRename(t *testing.T) {
	err := Rename("./testdata/testdir/testfile.txt", "./testdata/testdir/file.txt")
	require.NoError(t, err)
}

func TestRemove(t *testing.T) {
	err := Remove("./testdata/testdir/file.txt")
	require.NoError(t, err)
}

func TestMkdir(t *testing.T) {
	err := Mkdir("./testdata/testdir")
	require.NoError(t, err)
}

func TestDeldir(t *testing.T) {
	err := Deldir("./testdata/testdir")
	require.NoError(t, err)
}

func TestAbsPath(t *testing.T) {
	absPath, err := AbsPath("./testdata/test.png")
	require.NoError(t, err)
	t.Log(absPath)

	absPath, err = AbsPath("./testdata/test.txt")
	require.NoError(t, err)
	t.Log(absPath)
}

func TestIs(t *testing.T) {
	assert.False(t, IsAbsPath("./testdata/test.png"))
	assert.False(t, IsExist("./test.go"))
	assert.True(t, IsFile("./testdata/test.md"))
	assert.True(t, IsDir("./testdata"))
}
