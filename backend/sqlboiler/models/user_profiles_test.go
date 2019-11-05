// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testUserProfiles(t *testing.T) {
	t.Parallel()

	query := UserProfiles()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testUserProfilesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserProfile{}
	if err = randomize.Struct(seed, o, userProfileDBTypes, true, userProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserProfile struct: %s", err)
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

	count, err := UserProfiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserProfilesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserProfile{}
	if err = randomize.Struct(seed, o, userProfileDBTypes, true, userProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserProfile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := UserProfiles().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UserProfiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserProfilesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserProfile{}
	if err = randomize.Struct(seed, o, userProfileDBTypes, true, userProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserProfile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UserProfileSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UserProfiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserProfilesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserProfile{}
	if err = randomize.Struct(seed, o, userProfileDBTypes, true, userProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserProfile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := UserProfileExists(ctx, tx, o.UserID)
	if err != nil {
		t.Errorf("Unable to check if UserProfile exists: %s", err)
	}
	if !e {
		t.Errorf("Expected UserProfileExists to return true, but got false.")
	}
}

func testUserProfilesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserProfile{}
	if err = randomize.Struct(seed, o, userProfileDBTypes, true, userProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserProfile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	userProfileFound, err := FindUserProfile(ctx, tx, o.UserID)
	if err != nil {
		t.Error(err)
	}

	if userProfileFound == nil {
		t.Error("want a record, got nil")
	}
}

func testUserProfilesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserProfile{}
	if err = randomize.Struct(seed, o, userProfileDBTypes, true, userProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserProfile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = UserProfiles().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testUserProfilesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserProfile{}
	if err = randomize.Struct(seed, o, userProfileDBTypes, true, userProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserProfile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := UserProfiles().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testUserProfilesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userProfileOne := &UserProfile{}
	userProfileTwo := &UserProfile{}
	if err = randomize.Struct(seed, userProfileOne, userProfileDBTypes, false, userProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserProfile struct: %s", err)
	}
	if err = randomize.Struct(seed, userProfileTwo, userProfileDBTypes, false, userProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserProfile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = userProfileOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = userProfileTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UserProfiles().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testUserProfilesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	userProfileOne := &UserProfile{}
	userProfileTwo := &UserProfile{}
	if err = randomize.Struct(seed, userProfileOne, userProfileDBTypes, false, userProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserProfile struct: %s", err)
	}
	if err = randomize.Struct(seed, userProfileTwo, userProfileDBTypes, false, userProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserProfile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = userProfileOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = userProfileTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserProfiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func userProfileBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *UserProfile) error {
	*o = UserProfile{}
	return nil
}

func userProfileAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *UserProfile) error {
	*o = UserProfile{}
	return nil
}

func userProfileAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *UserProfile) error {
	*o = UserProfile{}
	return nil
}

func userProfileBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *UserProfile) error {
	*o = UserProfile{}
	return nil
}

func userProfileAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *UserProfile) error {
	*o = UserProfile{}
	return nil
}

func userProfileBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *UserProfile) error {
	*o = UserProfile{}
	return nil
}

func userProfileAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *UserProfile) error {
	*o = UserProfile{}
	return nil
}

func userProfileBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *UserProfile) error {
	*o = UserProfile{}
	return nil
}

func userProfileAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *UserProfile) error {
	*o = UserProfile{}
	return nil
}

func testUserProfilesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &UserProfile{}
	o := &UserProfile{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, userProfileDBTypes, false); err != nil {
		t.Errorf("Unable to randomize UserProfile object: %s", err)
	}

	AddUserProfileHook(boil.BeforeInsertHook, userProfileBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	userProfileBeforeInsertHooks = []UserProfileHook{}

	AddUserProfileHook(boil.AfterInsertHook, userProfileAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	userProfileAfterInsertHooks = []UserProfileHook{}

	AddUserProfileHook(boil.AfterSelectHook, userProfileAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	userProfileAfterSelectHooks = []UserProfileHook{}

	AddUserProfileHook(boil.BeforeUpdateHook, userProfileBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	userProfileBeforeUpdateHooks = []UserProfileHook{}

	AddUserProfileHook(boil.AfterUpdateHook, userProfileAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	userProfileAfterUpdateHooks = []UserProfileHook{}

	AddUserProfileHook(boil.BeforeDeleteHook, userProfileBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	userProfileBeforeDeleteHooks = []UserProfileHook{}

	AddUserProfileHook(boil.AfterDeleteHook, userProfileAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	userProfileAfterDeleteHooks = []UserProfileHook{}

	AddUserProfileHook(boil.BeforeUpsertHook, userProfileBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	userProfileBeforeUpsertHooks = []UserProfileHook{}

	AddUserProfileHook(boil.AfterUpsertHook, userProfileAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	userProfileAfterUpsertHooks = []UserProfileHook{}
}

func testUserProfilesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserProfile{}
	if err = randomize.Struct(seed, o, userProfileDBTypes, true, userProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserProfile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserProfiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserProfilesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserProfile{}
	if err = randomize.Struct(seed, o, userProfileDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserProfile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(userProfileColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := UserProfiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserProfileToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local UserProfile
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, userProfileDBTypes, false, userProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserProfile struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := UserProfileSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*UserProfile)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testUserProfileToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a UserProfile
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userProfileDBTypes, false, strmangle.SetComplement(userProfilePrimaryKeyColumns, userProfileColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.UserProfile != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		if exists, err := UserProfileExists(ctx, tx, a.UserID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testUserProfilesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserProfile{}
	if err = randomize.Struct(seed, o, userProfileDBTypes, true, userProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserProfile struct: %s", err)
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

func testUserProfilesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserProfile{}
	if err = randomize.Struct(seed, o, userProfileDBTypes, true, userProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserProfile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UserProfileSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testUserProfilesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserProfile{}
	if err = randomize.Struct(seed, o, userProfileDBTypes, true, userProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserProfile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UserProfiles().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	userProfileDBTypes = map[string]string{`CreatedAt`: `datetime`, `Gender`: `varchar`, `Prefecture`: `varchar`, `UpdatedAt`: `datetime`, `UserID`: `int`}
	_                  = bytes.MinRead
)

func testUserProfilesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(userProfilePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(userProfileColumns) == len(userProfilePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UserProfile{}
	if err = randomize.Struct(seed, o, userProfileDBTypes, true, userProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserProfile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserProfiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, userProfileDBTypes, true, userProfilePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserProfile struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testUserProfilesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(userProfileColumns) == len(userProfilePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UserProfile{}
	if err = randomize.Struct(seed, o, userProfileDBTypes, true, userProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserProfile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserProfiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, userProfileDBTypes, true, userProfilePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserProfile struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(userProfileColumns, userProfilePrimaryKeyColumns) {
		fields = userProfileColumns
	} else {
		fields = strmangle.SetComplement(
			userProfileColumns,
			userProfilePrimaryKeyColumns,
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

	slice := UserProfileSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testUserProfilesUpsert(t *testing.T) {
	t.Parallel()

	if len(userProfileColumns) == len(userProfilePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLUserProfileUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := UserProfile{}
	if err = randomize.Struct(seed, &o, userProfileDBTypes, false); err != nil {
		t.Errorf("Unable to randomize UserProfile struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UserProfile: %s", err)
	}

	count, err := UserProfiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, userProfileDBTypes, false, userProfilePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserProfile struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UserProfile: %s", err)
	}

	count, err = UserProfiles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
