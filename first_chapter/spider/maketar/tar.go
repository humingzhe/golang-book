package main

import (
	"archive/tar"
	"io"
	"log"
	"os"
	"path/filepath"
)
	// maketar函数用来遍历目录，把目录里的文件打成tar包，tar包直接输出到标准输出上。
	// tr是以io.writer触发
func maketar(dir string, w io.Writer) error {
	base := filepath.Base(dir)
	tr := tar.NewWriter(w) // New一个tar，这个writer就是说把tar包写在那。
	defer tr.Close()
	// 指定目录做一个walk，就是把目录进行遍历，遍历每个目录每个文件都会调用这个函数。把这个目录的路径和它的info（info包括权限、ctime、mtime、属主、大小），传进FileInfo，传进去后就可以在这里面做各种各样的操作。
	return filepath.Walk(dir, func(name string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		f, err := os.Open(name)
		if err != nil {
			return err
		}
		defer f.Close()
		// tar包需要写文件的Info，需要用到FileInfoHeader把Walk那边传进去的给他传过来，形成它的一个Header，tr那里就会WriteHeader。
		h, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}
		// 因为tar包有可能里面还有目录结构，它里面就需要把传过来的name和当前的dir进行拼接，拼接成全路径，否则目录结构会发生改变。
		p, _ := filepath.Rel(dir, name)
		h.Name = filepath.Join(base, p)
		if err = tr.WriteHeader(h); err != nil {
			return err
		}
		// 上面都完事后就会用到了io.Copy，io.Copy可以把f里的内容拷贝到tr中。F就是文件。
		if info.Mode().IsRegular() {
			io.Copy(tr, f)
		}
		return nil
	})
}

func main() {
	// 调用时，直接调用os.Stdout
	err := maketar(os.Args[1], os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
