package cmd

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"

	"github.com/spf13/cobra"
)

//go:embed templates/*.tpl
var templatesFS embed.FS

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
		"handlers", "models",
		"repository", "service",
	}

	// Buat semua sub-folder
	for _, d := range dirs {
		path := filepath.Join(base, d)
		_ = os.MkdirAll(path, os.ModePerm)
	}

	// Data untuk template parsing
	data := map[string]string{
		"Module":  name,
		"module":  strings.ToLower(name),
		"ModRoot": getModulesName(),
	}

	// Daftar file dan output path-nya
	templates := map[string]string{
		"handler.go.tpl": base + "/handlers/handler.go",
		"model.go.tpl":   base + "/models/" + data["module"] + ".go",
		"repo.go.tpl":    base + "/repository/" + data["module"] + "_repository.go",
		"service.go.tpl": base + "/service/" + data["module"] + "_service.go",
	}

	for tplName, outPath := range templates {
		tplContent, err := templatesFS.ReadFile("templates/" + tplName)
		if err != nil {
			fmt.Println("❌ Gagal baca template:", tplName)
			continue
		}

		tpl, err := template.New(tplName).Funcs(template.FuncMap{
			"capitalize": capitalize,
		}).Parse(string(tplContent))
		if err != nil {
			fmt.Println("❌ Gagal parse:", tplName)
			continue
		}

		f, err := os.Create(outPath)
		if err != nil {
			fmt.Println("❌ Gagal buat file:", outPath)
			continue
		}
		defer f.Close()

		_ = tpl.Execute(f, data)
		fmt.Println("✅ Dibuat:", outPath)
	}
}

// getModulesName mengambil nama module dari go.mod
func getModulesName() string {
	content, err := os.ReadFile("go.mod")
	if err != nil {
		return "your_module" // fallback
	}
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "module") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module"))
		}
	}
	return "your_module"
}

// capitalize huruf pertama string
func capitalize(str string) string {
	if str == "" {
		return ""
	}
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
