package test
import (
	"runtime"
	"testing"
	"path/filepath"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)


func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func TestDemo(t *testing.T) {

	type MileStone struct {
		Name string
	}

	type Project struct {
		ID int
		MileStone *MileStone
	}

	var projects []Project
	project := Project{ID:1}
	projects = append(projects, project)
	for i, p := range projects {
		m := new(MileStone)
		m.Name = "V0"
		p.MileStone = m
		projects[i] = p
		logs.Debug("#1 ID:%d, pointer:%p, MileStone:%v", p.ID, &p, *p.MileStone)
	}

	for _, p:= range projects {
		logs.Debug("#2 ID:%d, pointer:%p, MileStone:%v", p.ID, &p, p.MileStone)
	}
}
