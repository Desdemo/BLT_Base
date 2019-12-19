package base

type Orm struct {
	Cols        string
	Table       string
	Factor      string
	FactorValue interface{}
	Limit       []int
	SortType    string
	SortValue   string
}

func (orm *Orm) Get() (data interface{}, err error) {
	data, err = Engine.
		Table(orm.Table).
		Cols(orm.Cols).
		Where(orm.Factor, orm.FactorValue).
		Limit(orm.Limit[0], orm.Limit[1]).
		OrderBy(orm.SortValue).
		Get(data)
	return
}
