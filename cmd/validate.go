package cmd

import (
	"fmt"

	"github.com/7phs/coding-challenge-iban/config"
	"github.com/7phs/coding-challenge-iban/db"
	"github.com/7phs/coding-challenge-iban/helper"
	"github.com/7phs/coding-challenge-iban/model"
	"github.com/7phs/coding-challenge-iban/model/records"
	"github.com/7phs/coding-challenge-iban/model/validator"
	"github.com/spf13/cobra"
)

var ValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validating IBAN",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			helper.Fatal("IBAN an arg or more required")
		}

		conf := config.ParseConfig()

		if err := conf.Validate(); err != nil {
			helper.Fatal("config: invalid, ", err)
		}

		if err := db.Init(conf); err != nil {
			helper.Fatal("IBAN an arg or more required")
		}

		model.Init(&model.Dependencies{
			CountriesFormat: db.CountriesFmt,
		})

		valid := model.Default.Validator()

		for _, arg := range args {
			fmt.Print(arg, ": ")

			if valid.Validate(records.NewIban(arg)) == nil {
				fmt.Println(validator.Valid)
			} else {
				fmt.Println(validator.Invalid)
			}
		}
	},
}
