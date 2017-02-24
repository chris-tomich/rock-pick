package query

import (
	"plugin"

	rockpickplugin "github.com/chris-tomich/rock-pick/plugin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/tecbot/gorocksdb"
)

func Entry(p *plugin.Plugin, databasePath string, query string) error {
	if databasePath == "" {
		return errors.New("No database was provided.")
	}

	opts := gorocksdb.NewDefaultOptions()
	db, openErr := gorocksdb.OpenDb(opts, databasePath)

	if openErr != nil {
		return openErr
	}

	kFn, err := p.Lookup("PrintKey")

	if err != nil {
		logrus.Fatal(err)
	}

	vFn, err := p.Lookup("PrintValue")

	if err != nil {
		logrus.Fatal(err)
	}

	printKey, ok := kFn.(rockpickplugin.KeyPrinter)

	if !ok {
		logrus.Fatal(errors.New("problem with the passed in PrintKey function"))
	}

	printValue, ok := vFn.(rockpickplugin.ValuePrinter)

	if !ok {
		logrus.Fatal(errors.New("problem with the passed in PrintValue function"))
	}

	switch query {
	case "*":
		return displayAll(db, printKey, printValue)
	default:
		return nil
	}
}

func displayAll(db *gorocksdb.DB, printKey rockpickplugin.KeyPrinter, printValue rockpickplugin.ValuePrinter) error {
	ro := gorocksdb.NewDefaultReadOptions()
	iter := db.NewIterator(ro)
	iter.SeekToFirst()

	for iter.SeekToFirst(); iter.Valid(); iter.Next() {
		keyBytes := iter.Key().Data()
		valueBytes := iter.Value().Data()

		printKey(keyBytes)
		printValue(valueBytes)
	}

	return nil
}
