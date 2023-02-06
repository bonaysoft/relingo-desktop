// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/bonaysoft/relingo-desktop/pkg/dal/model"
)

func newWord(db *gorm.DB, opts ...gen.DOOption) word {
	_word := word{}

	_word.wordDo.UseDB(db, opts...)
	_word.wordDo.UseModel(&model.Word{})

	tableName := _word.wordDo.TableName()
	_word.ALL = field.NewAsterisk(tableName)
	_word.Id = field.NewInt(tableName, "id")
	_word.Source = field.NewString(tableName, "source")
	_word.Exposures = field.NewInt(tableName, "exposures")
	_word.Mastered = field.NewBool(tableName, "mastered")
	_word.CreatedAt = field.NewTime(tableName, "created_at")
	_word.UpdatedAt = field.NewTime(tableName, "updated_at")
	_word.DeletedAt = field.NewTime(tableName, "deleted_at")

	_word.fillFieldMap()

	return _word
}

type word struct {
	wordDo

	ALL       field.Asterisk
	Id        field.Int
	Source    field.String
	Exposures field.Int
	Mastered  field.Bool
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Time

	fieldMap map[string]field.Expr
}

func (w word) Table(newTableName string) *word {
	w.wordDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w word) As(alias string) *word {
	w.wordDo.DO = *(w.wordDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *word) updateTableName(table string) *word {
	w.ALL = field.NewAsterisk(table)
	w.Id = field.NewInt(table, "id")
	w.Source = field.NewString(table, "source")
	w.Exposures = field.NewInt(table, "exposures")
	w.Mastered = field.NewBool(table, "mastered")
	w.CreatedAt = field.NewTime(table, "created_at")
	w.UpdatedAt = field.NewTime(table, "updated_at")
	w.DeletedAt = field.NewTime(table, "deleted_at")

	w.fillFieldMap()

	return w
}

func (w *word) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *word) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 7)
	w.fieldMap["id"] = w.Id
	w.fieldMap["source"] = w.Source
	w.fieldMap["exposures"] = w.Exposures
	w.fieldMap["mastered"] = w.Mastered
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["deleted_at"] = w.DeletedAt
}

func (w word) clone(db *gorm.DB) word {
	w.wordDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w word) replaceDB(db *gorm.DB) word {
	w.wordDo.ReplaceDB(db)
	return w
}

type wordDo struct{ gen.DO }

type IWordDo interface {
	gen.SubQuery
	Debug() IWordDo
	WithContext(ctx context.Context) IWordDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWordDo
	WriteDB() IWordDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWordDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWordDo
	Not(conds ...gen.Condition) IWordDo
	Or(conds ...gen.Condition) IWordDo
	Select(conds ...field.Expr) IWordDo
	Where(conds ...gen.Condition) IWordDo
	Order(conds ...field.Expr) IWordDo
	Distinct(cols ...field.Expr) IWordDo
	Omit(cols ...field.Expr) IWordDo
	Join(table schema.Tabler, on ...field.Expr) IWordDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWordDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWordDo
	Group(cols ...field.Expr) IWordDo
	Having(conds ...gen.Condition) IWordDo
	Limit(limit int) IWordDo
	Offset(offset int) IWordDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWordDo
	Unscoped() IWordDo
	Create(values ...*model.Word) error
	CreateInBatches(values []*model.Word, batchSize int) error
	Save(values ...*model.Word) error
	First() (*model.Word, error)
	Take() (*model.Word, error)
	Last() (*model.Word, error)
	Find() ([]*model.Word, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Word, err error)
	FindInBatches(result *[]*model.Word, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Word) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWordDo
	Assign(attrs ...field.AssignExpr) IWordDo
	Joins(fields ...field.RelationField) IWordDo
	Preload(fields ...field.RelationField) IWordDo
	FirstOrInit() (*model.Word, error)
	FirstOrCreate() (*model.Word, error)
	FindByPage(offset int, limit int) (result []*model.Word, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWordDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w wordDo) Debug() IWordDo {
	return w.withDO(w.DO.Debug())
}

func (w wordDo) WithContext(ctx context.Context) IWordDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w wordDo) ReadDB() IWordDo {
	return w.Clauses(dbresolver.Read)
}

func (w wordDo) WriteDB() IWordDo {
	return w.Clauses(dbresolver.Write)
}

func (w wordDo) Session(config *gorm.Session) IWordDo {
	return w.withDO(w.DO.Session(config))
}

func (w wordDo) Clauses(conds ...clause.Expression) IWordDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w wordDo) Returning(value interface{}, columns ...string) IWordDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w wordDo) Not(conds ...gen.Condition) IWordDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w wordDo) Or(conds ...gen.Condition) IWordDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w wordDo) Select(conds ...field.Expr) IWordDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w wordDo) Where(conds ...gen.Condition) IWordDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w wordDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IWordDo {
	return w.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (w wordDo) Order(conds ...field.Expr) IWordDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w wordDo) Distinct(cols ...field.Expr) IWordDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w wordDo) Omit(cols ...field.Expr) IWordDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w wordDo) Join(table schema.Tabler, on ...field.Expr) IWordDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w wordDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWordDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w wordDo) RightJoin(table schema.Tabler, on ...field.Expr) IWordDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w wordDo) Group(cols ...field.Expr) IWordDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w wordDo) Having(conds ...gen.Condition) IWordDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w wordDo) Limit(limit int) IWordDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w wordDo) Offset(offset int) IWordDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w wordDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWordDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w wordDo) Unscoped() IWordDo {
	return w.withDO(w.DO.Unscoped())
}

func (w wordDo) Create(values ...*model.Word) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w wordDo) CreateInBatches(values []*model.Word, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w wordDo) Save(values ...*model.Word) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w wordDo) First() (*model.Word, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Word), nil
	}
}

func (w wordDo) Take() (*model.Word, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Word), nil
	}
}

func (w wordDo) Last() (*model.Word, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Word), nil
	}
}

func (w wordDo) Find() ([]*model.Word, error) {
	result, err := w.DO.Find()
	return result.([]*model.Word), err
}

func (w wordDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Word, err error) {
	buf := make([]*model.Word, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w wordDo) FindInBatches(result *[]*model.Word, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w wordDo) Attrs(attrs ...field.AssignExpr) IWordDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w wordDo) Assign(attrs ...field.AssignExpr) IWordDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w wordDo) Joins(fields ...field.RelationField) IWordDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w wordDo) Preload(fields ...field.RelationField) IWordDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w wordDo) FirstOrInit() (*model.Word, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Word), nil
	}
}

func (w wordDo) FirstOrCreate() (*model.Word, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Word), nil
	}
}

func (w wordDo) FindByPage(offset int, limit int) (result []*model.Word, count int64, err error) {
	result, err = w.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = w.Offset(-1).Limit(-1).Count()
	return
}

func (w wordDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w wordDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w wordDo) Delete(models ...*model.Word) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *wordDo) withDO(do gen.Dao) *wordDo {
	w.DO = *do.(*gen.DO)
	return w
}
