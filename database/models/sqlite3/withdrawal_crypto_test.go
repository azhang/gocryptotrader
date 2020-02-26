// Code generated by SQLBoiler 3.5.0-gct (https://github.com/thrasher-corp/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package sqlite3

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/thrasher-corp/sqlboiler/boil"
	"github.com/thrasher-corp/sqlboiler/queries"
	"github.com/thrasher-corp/sqlboiler/randomize"
	"github.com/thrasher-corp/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testWithdrawalCryptos(t *testing.T) {
	t.Parallel()

	query := WithdrawalCryptos()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testWithdrawalCryptosDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WithdrawalCrypto{}
	if err = randomize.Struct(seed, o, withdrawalCryptoDBTypes, true, withdrawalCryptoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalCrypto struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := WithdrawalCryptos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testWithdrawalCryptosQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WithdrawalCrypto{}
	if err = randomize.Struct(seed, o, withdrawalCryptoDBTypes, true, withdrawalCryptoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalCrypto struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := WithdrawalCryptos().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := WithdrawalCryptos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testWithdrawalCryptosSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WithdrawalCrypto{}
	if err = randomize.Struct(seed, o, withdrawalCryptoDBTypes, true, withdrawalCryptoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalCrypto struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := WithdrawalCryptoSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := WithdrawalCryptos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testWithdrawalCryptosExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WithdrawalCrypto{}
	if err = randomize.Struct(seed, o, withdrawalCryptoDBTypes, true, withdrawalCryptoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalCrypto struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := WithdrawalCryptoExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if WithdrawalCrypto exists: %s", err)
	}
	if !e {
		t.Errorf("Expected WithdrawalCryptoExists to return true, but got false.")
	}
}

func testWithdrawalCryptosFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WithdrawalCrypto{}
	if err = randomize.Struct(seed, o, withdrawalCryptoDBTypes, true, withdrawalCryptoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalCrypto struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	withdrawalCryptoFound, err := FindWithdrawalCrypto(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if withdrawalCryptoFound == nil {
		t.Error("want a record, got nil")
	}
}

func testWithdrawalCryptosBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WithdrawalCrypto{}
	if err = randomize.Struct(seed, o, withdrawalCryptoDBTypes, true, withdrawalCryptoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalCrypto struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = WithdrawalCryptos().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testWithdrawalCryptosOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WithdrawalCrypto{}
	if err = randomize.Struct(seed, o, withdrawalCryptoDBTypes, true, withdrawalCryptoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalCrypto struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := WithdrawalCryptos().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testWithdrawalCryptosAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	withdrawalCryptoOne := &WithdrawalCrypto{}
	withdrawalCryptoTwo := &WithdrawalCrypto{}
	if err = randomize.Struct(seed, withdrawalCryptoOne, withdrawalCryptoDBTypes, false, withdrawalCryptoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalCrypto struct: %s", err)
	}
	if err = randomize.Struct(seed, withdrawalCryptoTwo, withdrawalCryptoDBTypes, false, withdrawalCryptoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalCrypto struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = withdrawalCryptoOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = withdrawalCryptoTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := WithdrawalCryptos().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testWithdrawalCryptosCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	withdrawalCryptoOne := &WithdrawalCrypto{}
	withdrawalCryptoTwo := &WithdrawalCrypto{}
	if err = randomize.Struct(seed, withdrawalCryptoOne, withdrawalCryptoDBTypes, false, withdrawalCryptoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalCrypto struct: %s", err)
	}
	if err = randomize.Struct(seed, withdrawalCryptoTwo, withdrawalCryptoDBTypes, false, withdrawalCryptoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalCrypto struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = withdrawalCryptoOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = withdrawalCryptoTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := WithdrawalCryptos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func withdrawalCryptoBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *WithdrawalCrypto) error {
	*o = WithdrawalCrypto{}
	return nil
}

func withdrawalCryptoAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *WithdrawalCrypto) error {
	*o = WithdrawalCrypto{}
	return nil
}

func withdrawalCryptoAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *WithdrawalCrypto) error {
	*o = WithdrawalCrypto{}
	return nil
}

func withdrawalCryptoBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *WithdrawalCrypto) error {
	*o = WithdrawalCrypto{}
	return nil
}

func withdrawalCryptoAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *WithdrawalCrypto) error {
	*o = WithdrawalCrypto{}
	return nil
}

func withdrawalCryptoBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *WithdrawalCrypto) error {
	*o = WithdrawalCrypto{}
	return nil
}

func withdrawalCryptoAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *WithdrawalCrypto) error {
	*o = WithdrawalCrypto{}
	return nil
}

func withdrawalCryptoBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *WithdrawalCrypto) error {
	*o = WithdrawalCrypto{}
	return nil
}

func withdrawalCryptoAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *WithdrawalCrypto) error {
	*o = WithdrawalCrypto{}
	return nil
}

func testWithdrawalCryptosHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &WithdrawalCrypto{}
	o := &WithdrawalCrypto{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, withdrawalCryptoDBTypes, false); err != nil {
		t.Errorf("Unable to randomize WithdrawalCrypto object: %s", err)
	}

	AddWithdrawalCryptoHook(boil.BeforeInsertHook, withdrawalCryptoBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	withdrawalCryptoBeforeInsertHooks = []WithdrawalCryptoHook{}

	AddWithdrawalCryptoHook(boil.AfterInsertHook, withdrawalCryptoAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	withdrawalCryptoAfterInsertHooks = []WithdrawalCryptoHook{}

	AddWithdrawalCryptoHook(boil.AfterSelectHook, withdrawalCryptoAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	withdrawalCryptoAfterSelectHooks = []WithdrawalCryptoHook{}

	AddWithdrawalCryptoHook(boil.BeforeUpdateHook, withdrawalCryptoBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	withdrawalCryptoBeforeUpdateHooks = []WithdrawalCryptoHook{}

	AddWithdrawalCryptoHook(boil.AfterUpdateHook, withdrawalCryptoAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	withdrawalCryptoAfterUpdateHooks = []WithdrawalCryptoHook{}

	AddWithdrawalCryptoHook(boil.BeforeDeleteHook, withdrawalCryptoBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	withdrawalCryptoBeforeDeleteHooks = []WithdrawalCryptoHook{}

	AddWithdrawalCryptoHook(boil.AfterDeleteHook, withdrawalCryptoAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	withdrawalCryptoAfterDeleteHooks = []WithdrawalCryptoHook{}

	AddWithdrawalCryptoHook(boil.BeforeUpsertHook, withdrawalCryptoBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	withdrawalCryptoBeforeUpsertHooks = []WithdrawalCryptoHook{}

	AddWithdrawalCryptoHook(boil.AfterUpsertHook, withdrawalCryptoAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	withdrawalCryptoAfterUpsertHooks = []WithdrawalCryptoHook{}
}

func testWithdrawalCryptosInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WithdrawalCrypto{}
	if err = randomize.Struct(seed, o, withdrawalCryptoDBTypes, true, withdrawalCryptoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalCrypto struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := WithdrawalCryptos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testWithdrawalCryptosInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WithdrawalCrypto{}
	if err = randomize.Struct(seed, o, withdrawalCryptoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize WithdrawalCrypto struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(withdrawalCryptoColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := WithdrawalCryptos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testWithdrawalCryptoToOneWithdrawalHistoryUsingWithdrawalHistory(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local WithdrawalCrypto
	var foreign WithdrawalHistory

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, withdrawalCryptoDBTypes, false, withdrawalCryptoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalCrypto struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, withdrawalHistoryDBTypes, false, withdrawalHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalHistory struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.WithdrawalHistoryID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.WithdrawalHistory().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := WithdrawalCryptoSlice{&local}
	if err = local.L.LoadWithdrawalHistory(ctx, tx, false, (*[]*WithdrawalCrypto)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.WithdrawalHistory == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.WithdrawalHistory = nil
	if err = local.L.LoadWithdrawalHistory(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.WithdrawalHistory == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testWithdrawalCryptoToOneSetOpWithdrawalHistoryUsingWithdrawalHistory(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a WithdrawalCrypto
	var b, c WithdrawalHistory

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, withdrawalCryptoDBTypes, false, strmangle.SetComplement(withdrawalCryptoPrimaryKeyColumns, withdrawalCryptoColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, withdrawalHistoryDBTypes, false, strmangle.SetComplement(withdrawalHistoryPrimaryKeyColumns, withdrawalHistoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, withdrawalHistoryDBTypes, false, strmangle.SetComplement(withdrawalHistoryPrimaryKeyColumns, withdrawalHistoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*WithdrawalHistory{&b, &c} {
		err = a.SetWithdrawalHistory(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.WithdrawalHistory != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.WithdrawalCryptos[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.WithdrawalHistoryID != x.ID {
			t.Error("foreign key was wrong value", a.WithdrawalHistoryID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.WithdrawalHistoryID))
		reflect.Indirect(reflect.ValueOf(&a.WithdrawalHistoryID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.WithdrawalHistoryID != x.ID {
			t.Error("foreign key was wrong value", a.WithdrawalHistoryID, x.ID)
		}
	}
}

func testWithdrawalCryptosReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WithdrawalCrypto{}
	if err = randomize.Struct(seed, o, withdrawalCryptoDBTypes, true, withdrawalCryptoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalCrypto struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testWithdrawalCryptosReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WithdrawalCrypto{}
	if err = randomize.Struct(seed, o, withdrawalCryptoDBTypes, true, withdrawalCryptoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalCrypto struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := WithdrawalCryptoSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testWithdrawalCryptosSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WithdrawalCrypto{}
	if err = randomize.Struct(seed, o, withdrawalCryptoDBTypes, true, withdrawalCryptoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalCrypto struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := WithdrawalCryptos().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	withdrawalCryptoDBTypes = map[string]string{`ID`: `INTEGER`, `Address`: `TEXT`, `AddressTag`: `TEXT`, `Fee`: `REAL`, `WithdrawalHistoryID`: `TEXT`}
	_                       = bytes.MinRead
)

func testWithdrawalCryptosUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(withdrawalCryptoPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(withdrawalCryptoAllColumns) == len(withdrawalCryptoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &WithdrawalCrypto{}
	if err = randomize.Struct(seed, o, withdrawalCryptoDBTypes, true, withdrawalCryptoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalCrypto struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := WithdrawalCryptos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, withdrawalCryptoDBTypes, true, withdrawalCryptoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize WithdrawalCrypto struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testWithdrawalCryptosSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(withdrawalCryptoAllColumns) == len(withdrawalCryptoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &WithdrawalCrypto{}
	if err = randomize.Struct(seed, o, withdrawalCryptoDBTypes, true, withdrawalCryptoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalCrypto struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := WithdrawalCryptos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, withdrawalCryptoDBTypes, true, withdrawalCryptoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize WithdrawalCrypto struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(withdrawalCryptoAllColumns, withdrawalCryptoPrimaryKeyColumns) {
		fields = withdrawalCryptoAllColumns
	} else {
		fields = strmangle.SetComplement(
			withdrawalCryptoAllColumns,
			withdrawalCryptoPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := WithdrawalCryptoSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}
