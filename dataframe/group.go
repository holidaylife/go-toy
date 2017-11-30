package dataframe

import (
	"github.com/kniren/gota/dataframe"
	"github.com/kniren/gota/series"
	"fmt"
)

func LoadData() {
	df := dataframe.New(
		series.New([]string{"b", "a"}, series.String, "A"),
		series.New([]int{1, 2}, series.Int, "B"),
		series.New([]float64{3.0, 4.0}, series.Float, "F"),
	)
	fmt.Println(df)

	df2 := dataframe.LoadRecords(
		[][]string{
			[]string{"A", "F", "D"},
			[]string{"1", "1", "true"},
			[]string{"4", "2", "false"},
			[]string{"2", "8", "false"},
			[]string{"5", "9", "false"},
		},
	)
	fmt.Println(df2.Nrow())
	fmt.Println(df2.Records()[1])
	fmt.Println(df2)

}
