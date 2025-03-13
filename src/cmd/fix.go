package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var selectedFiles = make(map[string]bool)

func newFixCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fix",
		Short: "Detect and fix issues in a project",
		Long:  `An interactive tool to scan project files, identify potential problems, and suggest fixes.`,
		RunE: func(_ *cobra.Command, args []string) error {
			path := "."
			if len(args) > 0 {
				path = args[0]
			}

			if _, err := os.Stat(path); os.IsNotExist(err) {
				return errors.New("the specified path does not exist")
			}

			app := tview.NewApplication()
			treeRoot := tview.NewTreeNode(path).SetSelectable(false)
			treeView := tview.NewTreeView().SetRoot(treeRoot).SetCurrentNode(treeRoot)

			loadFilesIntoTree(treeRoot, path)

			treeView.SetSelectedFunc(func(node *tview.TreeNode) {
				path, ok := node.GetReference().(string)
				if !ok {
					return
				}

				if selectedFiles[path] {
					deselectRecursively(node)
				} else {
					selectRecursively(node)
				}
			})

			treeView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
				if event.Key() == tcell.KeyEnter {
					app.Stop()
					printSelectedFiles()
					return nil
				}
				return event
			})

			layout := tview.NewFlex().
				SetDirection(tview.FlexRow).
				AddItem(tview.NewTextView().SetText("Use ↑↓ to navigate, Space to select, Enter to confirm."), 1, 1, false).
				AddItem(treeView, 0, 1, true)

			if err := app.SetRoot(layout, true).Run(); err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}

func loadFilesIntoTree(parentNode *tview.TreeNode, path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		return
	}

	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		node := tview.NewTreeNode(file.Name()).
			SetReference(filePath).
			SetSelectable(true)

		if file.IsDir() {
			node.SetColor(tcell.ColorBlue)
			loadFilesIntoTree(node, filePath)
		} else {
			node.SetColor(tcell.ColorWhite)
		}

		parentNode.AddChild(node)
	}
}

func selectRecursively(node *tview.TreeNode) {
	path, ok := node.GetReference().(string)
	if !ok {
		return
	}

	selectedFiles[path] = true
	node.SetColor(tcell.ColorGreen)

	for _, child := range node.GetChildren() {
		selectRecursively(child)
	}
}

func deselectRecursively(node *tview.TreeNode) {
	path, ok := node.GetReference().(string)
	if !ok {
		return
	}

	delete(selectedFiles, path)
	node.SetColor(tcell.ColorWhite)

	for _, child := range node.GetChildren() {
		deselectRecursively(child)
	}
}

func printSelectedFiles() {
	if len(selectedFiles) == 0 {
		fmt.Println("No files selected.")
		return
	}
	fmt.Println("Selected files:")
	for file := range selectedFiles {
		fmt.Println(file)
	}
}
