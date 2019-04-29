package cc

type variable struct {
	name   string
	offset int
}

type variables struct {
	offset int
	vars   map[string]variable
}

func (vars *variables) add(key string) {
	vars.offset += 8
	vars.vars[key] = variable{
		name:   key,
		offset: vars.offset,
	}
}

func (vars *variables) exist(key string) bool {
	_, ok := vars.vars[key]
	return ok
}

func (vars *variables) get(key string) *variable {
	v, _ := vars.vars[key]
	return &v
}
