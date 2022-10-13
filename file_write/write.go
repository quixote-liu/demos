package write

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// A.线程A使用随机函数产生数据（数据为整形），数据写入Channel（ChannelA）。
// B.线程B从该Channel（ChannelA）中获取数据，将数据以及收到数据的时间戳（精确到秒即可）以下列格式追加存储到一个文件（./data.txt）中：
// {“data”:xxx,”time”:xxxxx}\r\n
// {“data”:xxx,”time”:xxxxx}\r\n
// {“data”:xxx,”time”:xxxxx}\r\n
// C.当文件超过1000行时，工作完成。
// 该包对外提供API：Begin，调用Begin后则启动上述流程。Begin函数在执行到步骤C时返回成功，否则返回相应的错误。

func Begin() error {
	datach := make(chan int)

	errprint := make(chan error)
	defer close(errprint)

	done := make(chan struct{})
	defer close(done)

	totalLines := 1000

	// 线程A
	go func() {
		for i := 0; i < totalLines; i++ {
			rand.Seed(time.Now().UnixMicro())
			datach <- rand.Int()
		}
		close(datach)
	}()

	// 线程B
	go func() {
		file, err := os.Create("./data.txt")
		// file, err := os.OpenFile("./data.txt", os.O_WRONLY, os.ModePerm)
		if err != nil {
			errprint <- err
			return
		}
		defer file.Close()
		for {
			data, ok := <-datach
			if !ok {
				break
			}
			d := struct {
				Data int   `json:"data"`
				Time int64 `json:"time"`
			}{
				Data: data,
				Time: time.Now().UnixMicro(),
			}
			dd, _ := json.Marshal(d)
			dd = append(dd, '\r', '\n')
			if _, err := file.Write(dd); err != nil {
				errprint <- err
				return
			}
		}
		done <- struct{}{}
	}()

	for {
		select {
		case <-time.After(time.Minute):
			// 限时一分钟
			return fmt.Errorf("time over")
		case err := <-errprint:
			return err
		case <-done:
			return nil
		}
	}
}
