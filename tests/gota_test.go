package test
// data frame for gota
import (
	"runtime"
	"path/filepath"
	"testing"
	"github.com/astaxie/beego"
	"hello/dataframe"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func Test_Data(t *testing.T) {
	dataframe.LoadData()
}

type Row struct {
	Data map[string]string
	Dimensions map[string]string
	DimensionFlags map[string] bool
	Metrics map[string]string
}

type Report struct {
	DimensionFlags map[string] bool
	MetricFlags map[string] bool
}

type ReportRows struct {
	Rows []Row
}

func (r Row) AddDimensions(keys []string, vals []string) {
	for i, key := range keys {
		val := vals[i]
		r.Dimensions[key] = val
		r.DimensionFlags[key] = true
	}
}

