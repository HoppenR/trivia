// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Question is an object representing the database table.
type Question struct {
	ID         null.Int64  `boil:"id" json:"id,omitempty" toml:"id" yaml:"id,omitempty"`
	Question   string      `boil:"question" json:"question" toml:"question" yaml:"question"`
	Answer     string      `boil:"answer" json:"answer" toml:"answer" yaml:"answer"`
	Choices    string      `boil:"choices" json:"choices" toml:"choices" yaml:"choices"`
	Categories string      `boil:"categories" json:"categories" toml:"categories" yaml:"categories"`
	Used       int64       `boil:"used" json:"used" toml:"used" yaml:"used"`
	Source     string      `boil:"source" json:"source" toml:"source" yaml:"source"`
	Type       null.String `boil:"type" json:"type,omitempty" toml:"type" yaml:"type,omitempty"`
	Difficulty null.String `boil:"difficulty" json:"difficulty,omitempty" toml:"difficulty" yaml:"difficulty,omitempty"`

	R *questionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L questionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var QuestionColumns = struct {
	ID         string
	Question   string
	Answer     string
	Choices    string
	Categories string
	Used       string
	Source     string
	Type       string
	Difficulty string
}{
	ID:         "id",
	Question:   "question",
	Answer:     "answer",
	Choices:    "choices",
	Categories: "categories",
	Used:       "used",
	Source:     "source",
	Type:       "type",
	Difficulty: "difficulty",
}

var QuestionTableColumns = struct {
	ID         string
	Question   string
	Answer     string
	Choices    string
	Categories string
	Used       string
	Source     string
	Type       string
	Difficulty string
}{
	ID:         "questions.id",
	Question:   "questions.question",
	Answer:     "questions.answer",
	Choices:    "questions.choices",
	Categories: "questions.categories",
	Used:       "questions.used",
	Source:     "questions.source",
	Type:       "questions.type",
	Difficulty: "questions.difficulty",
}

// Generated where

type whereHelpernull_Int64 struct{ field string }

func (w whereHelpernull_Int64) EQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int64) NEQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Int64) LT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int64) LTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int64) GT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int64) GTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var QuestionWhere = struct {
	ID         whereHelpernull_Int64
	Question   whereHelperstring
	Answer     whereHelperstring
	Choices    whereHelperstring
	Categories whereHelperstring
	Used       whereHelperint64
	Source     whereHelperstring
	Type       whereHelpernull_String
	Difficulty whereHelpernull_String
}{
	ID:         whereHelpernull_Int64{field: "\"questions\".\"id\""},
	Question:   whereHelperstring{field: "\"questions\".\"question\""},
	Answer:     whereHelperstring{field: "\"questions\".\"answer\""},
	Choices:    whereHelperstring{field: "\"questions\".\"choices\""},
	Categories: whereHelperstring{field: "\"questions\".\"categories\""},
	Used:       whereHelperint64{field: "\"questions\".\"used\""},
	Source:     whereHelperstring{field: "\"questions\".\"source\""},
	Type:       whereHelpernull_String{field: "\"questions\".\"type\""},
	Difficulty: whereHelpernull_String{field: "\"questions\".\"difficulty\""},
}

// QuestionRels is where relationship names are stored.
var QuestionRels = struct {
}{}

// questionR is where relationships are stored.
type questionR struct {
}

// NewStruct creates a new relationship struct
func (*questionR) NewStruct() *questionR {
	return &questionR{}
}

// questionL is where Load methods for each relationship are stored.
type questionL struct{}

var (
	questionAllColumns            = []string{"id", "question", "answer", "choices", "categories", "used", "source", "type", "difficulty"}
	questionColumnsWithoutDefault = []string{}
	questionColumnsWithDefault    = []string{"id", "question", "answer", "choices", "categories", "used", "source", "type", "difficulty"}
	questionPrimaryKeyColumns     = []string{"id"}
)

type (
	// QuestionSlice is an alias for a slice of pointers to Question.
	// This should almost always be used instead of []Question.
	QuestionSlice []*Question

	questionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	questionType                 = reflect.TypeOf(&Question{})
	questionMapping              = queries.MakeStructMapping(questionType)
	questionPrimaryKeyMapping, _ = queries.BindMapping(questionType, questionMapping, questionPrimaryKeyColumns)
	questionInsertCacheMut       sync.RWMutex
	questionInsertCache          = make(map[string]insertCache)
	questionUpdateCacheMut       sync.RWMutex
	questionUpdateCache          = make(map[string]updateCache)
	questionUpsertCacheMut       sync.RWMutex
	questionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single question record from the query using the global executor.
func (q questionQuery) OneG(ctx context.Context) (*Question, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single question record from the query.
func (q questionQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Question, error) {
	o := &Question{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for questions")
	}

	return o, nil
}

// AllG returns all Question records from the query using the global executor.
func (q questionQuery) AllG(ctx context.Context) (QuestionSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Question records from the query.
func (q questionQuery) All(ctx context.Context, exec boil.ContextExecutor) (QuestionSlice, error) {
	var o []*Question

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Question slice")
	}

	return o, nil
}

// CountG returns the count of all Question records in the query, and panics on error.
func (q questionQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Question records in the query.
func (q questionQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count questions rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q questionQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q questionQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if questions exists")
	}

	return count > 0, nil
}

// Questions retrieves all the records using an executor.
func Questions(mods ...qm.QueryMod) questionQuery {
	mods = append(mods, qm.From("\"questions\""))
	return questionQuery{NewQuery(mods...)}
}

// FindQuestionG retrieves a single record by ID.
func FindQuestionG(ctx context.Context, iD null.Int64, selectCols ...string) (*Question, error) {
	return FindQuestion(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindQuestion retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindQuestion(ctx context.Context, exec boil.ContextExecutor, iD null.Int64, selectCols ...string) (*Question, error) {
	questionObj := &Question{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"questions\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, questionObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from questions")
	}

	return questionObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Question) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Question) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no questions provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(questionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	questionInsertCacheMut.RLock()
	cache, cached := questionInsertCache[key]
	questionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			questionAllColumns,
			questionColumnsWithDefault,
			questionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(questionType, questionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(questionType, questionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"questions\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"questions\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"questions\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, questionPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into questions")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for questions")
	}

CacheNoHooks:
	if !cached {
		questionInsertCacheMut.Lock()
		questionInsertCache[key] = cache
		questionInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single Question record using the global executor.
// See Update for more documentation.
func (o *Question) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Question.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Question) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	questionUpdateCacheMut.RLock()
	cache, cached := questionUpdateCache[key]
	questionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			questionAllColumns,
			questionPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update questions, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"questions\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, questionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(questionType, questionMapping, append(wl, questionPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update questions row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for questions")
	}

	if !cached {
		questionUpdateCacheMut.Lock()
		questionUpdateCache[key] = cache
		questionUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q questionQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q questionQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for questions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for questions")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o QuestionSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o QuestionSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), questionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"questions\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, questionPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in question slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all question")
	}
	return rowsAff, nil
}

// DeleteG deletes a single Question record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Question) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Question record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Question) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Question provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), questionPrimaryKeyMapping)
	sql := "DELETE FROM \"questions\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from questions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for questions")
	}

	return rowsAff, nil
}

func (q questionQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q questionQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no questionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from questions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for questions")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o QuestionSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o QuestionSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), questionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"questions\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, questionPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from question slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for questions")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Question) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no Question provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Question) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindQuestion(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *QuestionSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty QuestionSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *QuestionSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := QuestionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), questionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"questions\".* FROM \"questions\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, questionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in QuestionSlice")
	}

	*o = slice

	return nil
}

// QuestionExistsG checks if the Question row exists.
func QuestionExistsG(ctx context.Context, iD null.Int64) (bool, error) {
	return QuestionExists(ctx, boil.GetContextDB(), iD)
}

// QuestionExists checks if the Question row exists.
func QuestionExists(ctx context.Context, exec boil.ContextExecutor, iD null.Int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"questions\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if questions exists")
	}

	return exists, nil
}
