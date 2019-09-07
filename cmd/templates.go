package cmd

var templates = map[string]string{
	"ObjectPool": `package {{ .Package }}

type {{ .TypeName }}Element struct {
	value {{ .TRef }}
}

type {{ .TypeName }} struct {
	factory func() {{ .TRef }}
	firstItem {{ .TRef }}
	items []{{ .TypeName }}Element
}

func New{{ .TypeName }}(factory func() {{ .TRef}}) *{{ .TypeName }} {
	return &{{ .TypeName }}{
		factory: factory,
		firstItem: nil,
		items: make([]{{ .TypeName }}Element, 2)
	}
}

func (pool *{{ .TypeName }}) CreateInstance() {{ .TRef }} {
	return pool.factory()
}

func (pool *{{ .TypeName }}) Allocate() {{ .TRef }} {
	var inst = _firstItem;
	if (inst == null || inst != atomic.CompareAndSwapPointer(_firstItem, inst)) {
		inst = allocateSlow();
	}
	
}

	`,
}
