package widget

type Widget interface {
	ID() string
}

type WidgetFun func() string

func (w WidgetFun) ID() string {
	return w()
}

var f = func(w Widget) string {
	return w.ID()
}
