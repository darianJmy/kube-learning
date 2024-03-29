package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/hpcloud/tail"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

/***
编译时需要安装以下依赖：
go get github.com/gorilla/websocket
go get github.com/hpcloud/tail
*/
const (
	// Time allowed to write the file to the client.
	//writeWait = 1 * time.Second
	writeWait = 100 * time.Millisecond

	// Time allowed to read the next pong message from the client.
	//pongWait = 24 * time.Hour
	pongWait = 60 * time.Second

	// Send pings to client with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Poll file for changes with this period.
	filePeriod = 1 * time.Second
)

var (
	homeTempl = template.Must(template.New("").Parse(homeHTML))
	filename  string
	addr      string
	start     bool
	upgrader  = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func init() {
	flag.StringVar(&filename, "f", "/Users/jimingyu_1/Documents/stu/test/1.txt", "指定一个文件的绝对路径")
	flag.StringVar(&addr, "a", ":8080", "http 服务地址")
	flag.Usage = usage
}
func usage() {
	fmt.Fprintf(os.Stderr, `
filewatch v1.0.0:检测文件变化，读取指定的文件，启动一个websocket页面实时读取，类似web版本的tail -f xxxx.log

Usage: filewatch [-f 文件绝对路径] [-a 监听的地址]

example:
filewatch -f /var/log/message -a :8080
`)
	flag.PrintDefaults()
}
func readFileIfModified(lastMod time.Time) ([]byte, time.Time, error) {
	fi, err := os.Stat(filename)
	if err != nil {
		return nil, lastMod, err
	}
	if !fi.ModTime().After(lastMod) {
		return nil, lastMod, nil
	}
	p, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fi.ModTime(), err
	}
	return p, fi.ModTime(), nil
}

func reader(ws *websocket.Conn) {
	defer ws.Close()
	ws.SetReadLimit(512)
	ws.SetReadDeadline(time.Now().Add(pongWait))
	ws.SetPongHandler(func(string) error { ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			break
		}
	}
}
func tailFile() *tail.Tail {
	tailfs, err := tail.TailFile(filename, tail.Config{
		ReOpen:    true,                                 // 文件被移除或被打包，需要重新打开
		Follow:    true,                                 // 实时跟踪
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 如果程序出现异常，保存上次读取的位置，避免重新读取。
		MustExist: false,                                // 如果文件不存在，是否推出程序，false是不退出
		Poll:      true,
	})

	if err != nil {
		fmt.Println("tailf failed, err:", err)
		return nil
	}
	return tailfs
}
func writer(ws *websocket.Conn, lastMod time.Time) {
	tailfs := tailFile()
	pingTicker := time.NewTicker(pingPeriod)
	fileTicker := time.NewTicker(filePeriod)
	defer func() {
		pingTicker.Stop()
		fileTicker.Stop()
		ws.Close()
	}()

	for {
		select {
		case msg, ok := <-tailfs.Lines:
			if ok {
				ws.SetWriteDeadline(time.Now().Add(writeWait))
				fmt.Printf("read file content： %s\n", msg)
				if err := ws.WriteMessage(websocket.TextMessage, []byte(msg.Text)); err != nil {
					return
				}
			}
		case <-pingTicker.C:
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}
		return
	}

	var lastMod time.Time
	if n, err := strconv.ParseInt(r.FormValue("lastMod"), 16, 64); err == nil {
		lastMod = time.Unix(0, n)
	}

	go writer(ws, lastMod)
	reader(ws)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	p, lastMod, err := readFileIfModified(time.Time{})
	if err != nil {
		p = []byte(err.Error())
		lastMod = time.Unix(0, 0)
	}
	var v = struct {
		Host    string
		Data    string
		LastMod string
	}{
		r.Host,
		string(p),
		strconv.FormatInt(lastMod.UnixNano(), 16),
	}
	homeTempl.Execute(w, &v)
}

func main() {
	flag.Parse()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", serveWs)
	fmt.Printf("文件读取路径：%s  \n", filename)
	fmt.Println("Listening and serving HTTP on " + addr)
	if !strings.Contains(addr, ":") {
		addr = ":" + addr
	}
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

const homeHTML = `<!DOCTYPE html>
<html lang="en">
    <head>
		<meta http-equiv="content-type" content="text/html;charset=utf-8">
        <title>WebSocket Example</title>
		<style>
				body{
					background-color: #0e1012;color: #ffffff;
				}
				*{
					
				}
				#msg{
					overflow:auto; border:10px solid #303030; color:#ffffff; background-color: #2b2b2b; font-size: 13px; position: absolute; left: 8px; right: 8px; bottom: 8px; top: 40px; word-break:
		break-all;
				}
				#log{
					position: fixed; top: 0; left: 0; width: 100%; height: 40px; text-align: left; margin: 4px 0 0 8px;
				}
				#log b{
					font-size: 26px;
				}
				#msgBtn{
					padding: 5px 10px; border: none; background: #777; float: right; margin: 0 16px 0 0;
				}
			</style>
    </head>
    <body>
	<div id="log">
		<span>
			<b>实时日志</b>
		</span>
		<input id="msgBtn" type="button" id="button1" οnclick="clearData()" value="清空" />
	</div>

    <div id="msg">
		<ul class="list">
        	<pre id="fileData"></pre>
		</ul>
	</div>
        <script type="text/javascript">
            (function() {
				var consoleDiv = document.getElementById('msg');
				consoleDiv.scrollTop = consoleDiv.scrollHeight;
                var dataElement = document.getElementById("fileData");
                var conn = new WebSocket("ws://{{.Host}}/ws?lastMod={{.LastMod}}");
                conn.onclose = function(evt) {
                    dataElement.textContent = 'Connection closed';
                }
                conn.onmessage = function(evt) {
                    console.log('file updated ---> ' + evt.data );
                    dataElement.textContent = dataElement.textContent + "\n" + evt.data ;
					consoleDiv.scrollTop = consoleDiv.scrollHeight;
                }
            })();

			function clearData()
			{
				var dataElement = document.getElementById("fileData");
				dataElement.innerText="";
			}
        </script>
    </body>
</html>
`
