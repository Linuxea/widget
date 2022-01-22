package widget

import (
	"fmt"
	"testing"
)

func TestWidget(t *testing.T) {

	rs := f(WidgetFun(func() string {
		return "interesting"
	}))
	fmt.Println(rs)
}
