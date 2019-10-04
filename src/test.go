package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strings"
)

//定义全局变量
var (
	echoGitUrl = "https://api.github.com/repos/labstack/echo/releases"
	dbName     = "/tmp/versions.db"
	insertchan chan int
)

func checkversion(v string) int {
	resp, err := http.Get(echoGitUrl)
	if err != nil {
		fmt.Println("无法连接到指定的网址，请检查网络或网站状态！", err)
		return 0
	}
	echoGitInfo, _ := ioutil.ReadAll(resp.Body)
	echoAllVersion := make([]string, 0)
	// 解析Json
	for _, Matchtagname := range regexp.MustCompile(`"tag_name":"(.+?)"`).FindAllSubmatch(echoGitInfo, -1) {
		echoAllVersion = append(echoAllVersion, string(Matchtagname[1]))
	}

	if strings.Contains(strings.Join(echoAllVersion, " "), v) {
		fmt.Println("version" + v + " is exists")
	} else {
		fmt.Println("version" + v + " not exists")
	}
	//删除旧库
	_ = os.Remove(dbName)
	//打开新的连接，并创建新库
	db, err := sql.Open("sqlite3", dbName)

	_, err = db.Exec(`CREATE TABLE if not exists versions (repo text, version text)`)
	if err != nil {
		return 0
	}
	fmt.Println("所有的版本个数为:", len(echoAllVersion))
	//初始化与版本切边长度的channel
	insertchan = make(chan int, len(echoAllVersion))
	for ic, version := range echoAllVersion {
		//启动与版本切片相同数量的go程
		go saveVersionTodb(db, version, ic)
	}
	//执行完关库
	//db.Close()
	return len(echoAllVersion)
}

//封装保存数据的函数
func saveVersionTodb(db *sql.DB, version string, ic int) {
	//开启事务锁
	//不论成功与失败，都需往channel中写入
	ctx, err := db.Begin()
	if err != nil {
		fmt.Println("无法开启事务，请检查数据库状态", err)
		insertchan <- ic
		return
	}
	_, err = db.Exec(`insert into versions values ("echo","` + version + `")`)
	if err != nil {
		//如果插入数据失败，则回滚事务
		insertchan <- ic
		_ = ctx.Rollback()
		return
	}
	//插入完成后提交事务
	err = ctx.Commit()
	if err != nil {
		fmt.Println("事务提交错误,请检查数据库状态，数据开始回滚，错误为：", err)
		insertchan <- ic
		err = ctx.Rollback()
		//上面已经在错误判断中写入channel，所以下面判错不写
		if err != nil {
			fmt.Println("提交后失败回滚，依旧失败，请查询数据库,错误为:", err)
			return
		}
		return
	}
	insertchan <- ic
}

func main2() {
	insertchanlen := checkversion("v3.3.3")
	for i := 0; i < insertchanlen; i++ {
		<-insertchan
	}
	//主动触发垃圾回收
	runtime.GC()
}
