package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"path"
	"sync"
	"github.com/rakyll/pb"
)

var (
	works = flag.Int("n", 5, "works")
)

type Block struct {
	Id    int
	Begin int64
	End   int64
	Name  string
}

func GetSize(url string) (int64, error) {
	// 通过Http.Head请求拿到Response
	resp, err := http.Head(url)
	if err != nil {
		return 0, err
	}
	// 关闭resp.Body.Close
	defer resp.Body.Close()
	// 拿到resp.ContentLength
	return resp.ContentLength, nil
}
// GenBlocks是用来切片的，类似于切大饼
func GenBlocks(total int64, works int) []*Block {
	var blocks []*Block
	// math.Ceil取整，total除以worker的个数，除完个数在求出每片有多大
	n := int64(math.Ceil(float64(total) / float64(works)))
	for i := int64(0); i < int64(works)-1; i++ {
		block := &Block{
			Id:    int(i),
			// n * i开始
			Begin: i * n,
			// i+1 * n开始
			End:   (i + 1) * n,
		}
		blocks = append(blocks, block)
	}
	blocks = append(blocks, &Block{
		Id:    works - 1,
		Begin: n * (int64(works) - 1),
		End:   total,
	})
	return blocks
}

func Download(url string, b *Block, bar *pb.ProgressBar) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", b.Begin, b.End-1))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	name := fmt.Sprintf("%s.%d", path.Base(url), b.Id)
	b.Name = name
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	w := io.MultiWriter(bar, f)
	_, err = io.Copy(w, resp.Body)
	return err
}

func Merge(name string, blocks []*Block) error {
	readers := make([]io.Reader, len(blocks))
	for i, b := range blocks {
		f, err := os.Open(b.Name)
		if err != nil {
			return err
		}
		defer f.Close()
		readers[i] = f
	}
	r := io.MultiReader(readers...)
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, r)
	if err != nil {
		return err
	}

	for _, b := range blocks {
		os.Remove(b.Name)
	}
	return nil
}

func main() {
	// 声明Flag.Parse，它是做参数解析的工具
	flag.Parse()
	// 解析URL参数
	url := flag.Arg(0)
	// 用GetSize来获取URL对应文件的大小，GetSize是通过HTTP的Head请求。
	total, err := GetSize(url)
	// 通过if进行判断
	if err != nil {
		log.Fatal(err)
	}
	// 分片下载
	log.Printf("total:%d", total)
 	//
	blocks := GenBlocks(total, *works)
	// pb是golang中的一个库，通过该库来设置进度条。
	bar := pb.New(int(total)).SetUnits(pb.U_BYTES)
	bar.Start()

	group := new(sync.WaitGroup)
	for _, b := range blocks {
		log.Printf("%v", b)
		b := b
		group.Add(1)
		go func() {
			defer group.Done()
			err := Download(url, b, bar)
			if err != nil {
				log.Fatal(err)
			}
		}()
	}

	group.Wait()
	bar.Finish()

	log.Printf("group done")
	err = Merge(path.Base(url), blocks)
	if err != nil {
		log.Fatal(err)
	}
}