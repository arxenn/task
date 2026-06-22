/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/arxenn/tasks/internal/domain"
	memrepo "github.com/arxenn/tasks/internal/repository/memory"
	"github.com/arxenn/tasks/internal/service"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "Always watch your priorioties",
	Long: `tasks is a simple todo list manager, for managing your daily tasks
	with a priority first aproach.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	addCmd.Flags().StringP("priority", "p", "", "set prioritoy as string (values: block, high, med, low)")

	listCmd.Flags().StringP("priority", "p", "", "filters tasks list by priority (values: block, high, med, low)")
	listCmd.Flags().StringP("status", "s", "", "filter tasks list by status (values: todo, done)")

	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new task (doesn't need double quotes).",
	RunE: func(cmd *cobra.Command, args []string) error {
		priority, err := cmd.Flags().GetString("priority")
		if err != nil {
			return err
		}
		content := strings.Join(args, " ")

		memRepo := memrepo.NewInMemoryRepository()
		svc := service.NewService(memRepo)

		if _, err := svc.Add(content, priority); err != nil {
			return fmt.Errorf("add task failed: %w", err)
		}

		return nil
	},
}

var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "lists tasks ordered by priority (filters: status, priority).",
	RunE: func(cmd *cobra.Command, args []string) error {
		priority, err := cmd.Flags().GetString("priority")
		if err != nil {
			return err
		}

		status, err := cmd.Flags().GetString("status")
		if err != nil {
			return err
		}

		memRepo := memrepo.NewInMemoryRepository()
		svc := service.NewService(memRepo)

		list, err := svc.List("", priority, status)
		if err != nil {
			return fmt.Errorf("add task failed: %w", err)
		}
		printTasks(list)

		return nil
	},
}

var removeCmd = &cobra.Command{
	Use:   "rm",
	Short: "removes task by it's ID",
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) != 0 {
			return errors.New("invalid number of arguments (expecting one argument)")
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid task ID: %w", err)
		}

		memRepo := memrepo.NewInMemoryRepository()
		svc := service.NewService(memRepo)

		if err := svc.Delete(id); err != nil {
			return fmt.Errorf("remove task %d failed: %w", id, err)
		}

		return nil
	},
}

func printTasks(tasks []domain.Task) {}
