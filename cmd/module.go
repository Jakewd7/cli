package cmd

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

var moduleCmd = &cobra.Command{
	Use:   "create:module [name]",
	Short: "Generate module (e.g. mod_auth) dengan struktur modular",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		moduleName := args[0]
		createModule(moduleName)
	},
}

func init() {
	rootCmd.AddCommand(moduleCmd)
}

func createModule(name string) {
	base := "mod_" + strings.ToLower(name)
	dirs := []string{
		"config", "controllers", "middleware", "models",
		"repository", "routes", "service", "utils",
	}

	// Buat semua sub-folder
	for _, d := range dirs {
		path := base + "/" + d
		os.MkdirAll(path, os.ModePerm)
	}

	// Data untuk template parsing
	data := map[string]string{
		"Module": name,
		"module": strings.ToLower(name),
	}

	// Daftar file dan path-nya
	templates := map[string]string{
		"templates/controller.go.tpl": base + "/controllers/handler.go",
		"templates/route.go.tpl":      base + "/routes/" + data["module"] + "_routes.go",
		"templates/model.go.tpl":      base + "/models/" + data["module"] + ".go",
		"templates/repo.go.tpl":       base + "/repository/" + data["module"] + "_repository.go",
		"templates/service.go.tpl":    base + "/service/" + data["module"] + "_service.go",
		"templates/config.go.tpl":     base + "/config/config.go",
		"templates/middleware.go.tpl": base + "/middleware/auth.go",
		"templates/utils.go.tpl":      base + "/utils/hash.go",
	}

	for tplPath, outPath := range templates {
		content, err := os.ReadFile(tplPath)
		if err != nil {
			fmt.Println("❌ Gagal baca template:", tplPath)
			continue
		}

		tpl, err := template.New("tpl").Parse(string(content))
		if err != nil {
			fmt.Println("❌ Gagal parse:", tplPath)
			continue
		}

		f, err := os.Create(outPath)
		if err != nil {
			fmt.Println("❌ Gagal buat file:", outPath)
			continue
		}
		defer f.Close()

		tpl.Execute(f, data)
		fmt.Println("✅ Dibuat:", outPath)
	}
}
