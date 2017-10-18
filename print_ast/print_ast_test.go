package print_ast

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestPrintAst(t *testing.T) {
	cases := []struct {
		filename, stmt string
		err            error
	}{
		{"simple_agg", "select sum(5) from t1", nil},
		{"one", "select sum(col_one) from table_2 where col_one > 2000", nil},
	}

	for i, c := range cases {
		textbytes, jsondata, err := PrintAst(c.stmt)
		if err != c.err {
			t.Errorf("case %d error mismatch. %s != %s", i, c.err, err)
			continue
		}

		if c.err == nil {
			err = ioutil.WriteFile(fmt.Sprintf("testdata/%s.json", c.filename), jsondata, 0777)
			if err != nil {
				t.Errorf("case %d error writing results json file: %s", i, err.Error())
			}
			err = ioutil.WriteFile(fmt.Sprintf("testdata/%s.txt", c.filename), textbytes, 0777)
			if err != nil {
				t.Errorf("case %d error writing results text file: %s", i, err.Error())
			}
		}
	}
}
