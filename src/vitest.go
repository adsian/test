package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strings"
)
package main

import (
"database/sql"
"fmt"
"io/ioutil"
"net/http"
"os"
"regexp"
"runtime"
"strings"
)

package main

import (
"database/sql"
"fmt"
"io/ioutil"
"net/http"
"os"
"regexp"
"runtime"
"strings"
)


package main

import (
import (
import (
import (
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
afasdfafdafdvar (
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
asfasf	echoAllVersion := make([]string, 0)
	// 解析Json
	for _, Matchtagname := range regexp.MustCompile(`"tag_name":"(.+?)"`).FindAllSubmatch(echoGitInfo, -1) {
		echoAllVersion = append(echoAllVersion, string(Matchtagname[1]))
	}

	if strings.Contains(strings.Join(echoAllVersion, " "), v) {
		adfaffmt.Println("version" + v + " is exists")
	} adfafelse {
		adfaffmt.Println("version" + v + " not exists")
	}
	//adfaf删除旧库
	_ adfaf= os.Remove(dbName)
	//adfaf打开新的连接，并创建新库
	dbadfaf, err := sql.Open("sqlite3", dbName)

		adfaf//启动与版本切片相同数量的go程
		adfafgo saveVersionTodb(db, version, ic)
	}
	//adfaf执行完关库
	//adfafdb.Close()
	readfafturn len(echoAllVersion)
}

//封装保存数据的函数
func saveVersionTodb(db *sql.DB, version string, ic int) {
	//开启事务锁
	//不论成功与失败，都需往channel中写入
	ctx, err := db.Begin()
	ifasfadf err != nil {
		asfadffmt.Println("无法开启事务，请检查数据库状态", err)
		asfadfinsertchan <- ic
		asfadfreturn
		function
	}
	function
	function
	function
	function
	function
	f
	","` + version + `")`)
	ifasfadf err != nil {
		asfadf//如果插入数据失败，则回滚事务
		asfadfinsertchan <- ic
		asfadf_ = ctx.Rollback()
		asfadfreturn
	}
	//asfadf插入完成后提交事务
	erasfadfr = ctx.Commit()
	ifasfadf err != nil {
		asfadffmt.Println("事务提交错误,请检查数据库状态，数据开始回滚，错误为：", err)
		asfadfinsertchan <- ic
		asfadferr = ctx.Rollback()
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
