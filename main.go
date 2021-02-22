package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	helpFlag    = flag.Bool("h", false, "Show this help")
	dirFlag     = flag.String("d", ".", "Directory to search for files")
	typeFlag    = flag.String("t", "go", "File suffix to parse (default 'go', others unoptimized)")
	recurseFlag = flag.Bool("r", false, "Recurse into sub-directories")
	outFlag     = flag.String("o", "", "Ouput file name (generated from package if blank)")
	stdoutFlag  = flag.Bool("so", false, "Dump output to Stdout instead of file")
)

const usage = "`parsePackage` [options]"

func main() {
	flag.Parse()
	if *helpFlag {
		fmt.Println(usage)
		flag.PrintDefaults()
		return
	}
	rootDir, err := filepath.Abs(*dirFlag)
	if err != nil {
		fmt.Println(err)
		return
	}
	var ofp string
	if *outFlag == "" {
		ofp = "pkgParse_" + filepath.Base(rootDir) + ".txt"
	} else {
		ofp, err = filepath.Abs(*outFlag)
		if err != nil {
			fmt.Println(err)
			ofp = "pkgParse_" + filepath.Base(rootDir) + ".txt"
		}
	}
	outFile, err := os.OpenFile(ofp, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outFile.Close()
	files, err := walkAtDir(rootDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(len(files))
	var gfs GoFiles
	for _, f := range files {
		gf, err := processFile(f)
		if err != nil {
			fmt.Printf("Error processing %s: %v\n", f, err)
			continue
		}
		gfs.Files = append(gfs.Files, gf)
	}
	for i := 0; i < len(gfs.Files); i++ {
		if *stdoutFlag {
			gfs.Files[i].PrintInfo()
			os.Remove(outFile.Name())
		} else {
			gfs.Files[i].WriteInfoToFile(outFile)
		}

	}
}

func walkAtDir(dirPath string) (files []string, err error) {

	err = filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if path == dirPath {
			return nil
		}
		if info.IsDir() {
			if *recurseFlag {
				//fmt.Println(path)
				sd, err := walkAtDir(path)
				if err != nil {
					return err
				}
				files = append(files, sd...)
			}
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, *typeFlag) {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

type GoFiles struct {
	Files []*GoFile
}

type GoFile struct {
	Name      string
	Package   string
	Imports   []string
	Functions []string
	Length    int
}

func (g *GoFile) PrintInfo() {
	fmt.Printf("Filename: %s\tPackagename: %s\tLength: %d\nImports: %v\nFunctions:\n", g.Name, g.Package, g.Length, g.Imports)
	for _, f := range g.Functions {
		fmt.Printf("\t%s\n", f)
	}
	fmt.Println("")
}

func (g *GoFile) WriteInfoToFile(of *os.File) {
	fmt.Fprintf(of, "Filename: %s\tPackagename: %s\tLength: %d\nImports: %v\nFunctions:\n", g.Name, g.Package, g.Length, g.Imports)
	for _, f := range g.Functions {
		fmt.Fprintf(of, "\t%s\n", f)
	}
	fmt.Fprintln(of, "")
}

// expects to receive fully qualified path
func processFile(fname string) (*GoFile, error) {
	//fmt.Println(fname)
	fi, err := os.Open(fname)
	if err != nil {
		return nil, err
	}

	counter := 0
	contBool := false
	goFile := &GoFile{
		Name: fname,
	}

	s := bufio.NewScanner(fi)
	for s.Scan() {
		counter++
		if strings.HasPrefix(s.Text(), "package ") {
			goFile.Package = strings.TrimPrefix(s.Text(), "package ")
			continue
		}
		if strings.HasPrefix(s.Text(), "import (") {
			contBool = true
			continue
		}
		if strings.HasPrefix(s.Text(), "import ") {
			goFile.Imports = append(goFile.Imports, strings.TrimPrefix(s.Text(), "import "))
			continue
		}
		if contBool {
			if strings.HasPrefix(s.Text(), ")") {
				contBool = false
				continue
			}
			goFile.Imports = append(goFile.Imports, strings.TrimSpace(s.Text()))
			continue
		}
		if strings.HasPrefix(s.Text(), "func ") {
			goFile.Functions = append(goFile.Functions, s.Text())
		}
	}
	goFile.Length = counter

	return goFile, nil
}
