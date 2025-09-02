/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/neberson/pos-go-expert-fullcycle/modulos/CobraCli/internal/database"
	"github.com/spf13/cobra"
	_ "modernc.org/sqlite"
)

func NewCreateCmd(category database.Category) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create a new category",
		Long:  `Create a new category`,
		RunE:  runCreate(category),
	}
}

func runCreate(categoryDb database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		name, err := cmd.Flags().GetString("name")
		description, err := cmd.Flags().GetString("description")

		_, err = categoryDb.Create(name, description)
		if err != nil {
			return err
		}
		return nil
	}
}

func init() {
	createCmd := NewCreateCmd(GetCategoryDb(GetDb()))

	categoryCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("name", "n", "", "Category name")
	createCmd.Flags().StringP("description", "d", "", "Category description")
	createCmd.MarkFlagsRequiredTogether("name", "description")
}
