package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var installStarterCmd = &cobra.Command{
	Use:   "install:starter",
	Short: "Install starter project dengan Gin dan koneksi database",
	Run: func(cmd *cobra.Command, args []string) {
		module := getModuleName()

		createFile(".env", envContent)
		createFile("main.go", mainContent(module))
		createFile("config/database.go", databaseContent)
		createFile("routes/router.go", routerContent())

		fmt.Println("✅ Starter project generated successfully without auth.")
	},
}

var envContent = `DB_USER=root
DB_PASSWORD=
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=mydb
JWT_SECRET=secret123`

func mainContent(module string) string {
	return fmt.Sprintf(`package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"

	"%s/config"
	"%s/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()

	r := gin.Default()

	routes.RegisterRoutes(r)

	r.Run(":8080")
}
`, module, module)
}

var databaseContent = `package config

import (
	"fmt"
	"log"
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	DB = database
}`

func routerContent() string {
	return `package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	// Define your routes here
}`
}

func init() {
	rootCmd.AddCommand(installStarterCmd)
}

func getModuleName() string {
	data, err := os.ReadFile("go.mod")
	if err != nil {
		return "your/module"
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module"))
		}
	}
	return "your/module"
}

func createFile(path, content string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(filepath.Dir(path), os.ModePerm)
		os.WriteFile(path, []byte(content), 0644)
		fmt.Println("✔️  Created:", path)
	}
}
