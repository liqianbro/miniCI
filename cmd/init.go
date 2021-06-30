package cmd

import (
	"fmt"
	"os"
	"text/template"

	"github.com/liqianbro/miniCI/templates"
	"github.com/liqianbro/miniCI/tool"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	GenerateType string
	// 初始化
	newCmd = &cobra.Command{
		Use:     "new [name]",
		Aliases: []string{"initialize", "initialise", "create"},
		Short:   "Initialize a Go Application",
		Long: `Initialize  will create a new application, with a license.
  * If a name is provided, a directory with that name will be created in the current directory;
  * If no name is provided, the current directory will be assumed;
`,
		Args: cobra.MinimumNArgs(1),
		Run: func(_ *cobra.Command, args []string) {
			_, err := initializeProject(args)
			tool.CheckErr(err)
			fmt.Printf("Your Go application is ready at\n%s\n", args[0])
		},
	}
)

func initializeProject(args []string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	if len(args) > 0 {
		if args[0] != "." {
			wd = fmt.Sprintf("%s/%s", wd, args[0])
		}
	}
	project := &Project{
		AbsolutePath: wd,
		PkgName:      args[0],
		Legal:        GetLicense(),
		Copyright:    CopyrightLine(),
		GenerateType: args[1],
	}
	if err := project.Create(); err != nil {
		return "", err
	}
	return project.AbsolutePath, nil
}

type Mod struct {
	PackageName string
}

// InitMod 生成go mod 相当于执行 go mod init pakcgeName
func InitMod(path string, pkgName string) error {
	//读取模版文件
	str, err := templates.TemplateFiles.ReadFile("files/gomod")
	if err != nil {
		return errors.Errorf("读取模版文件错误:%s", err)
	}
	//模版赋值
	tmpl, err := template.New("").Parse(string(str))
	if err != nil {
		return errors.Errorf("模版生成错误:%s", err)
	}
	//创建文件生成模版
	file, err := os.Create(path + "/go.mod")
	defer file.Close()
	if err != nil {
		return errors.Errorf("mod文件创建错误:%s", err)
	}
	fileName := Mod{PackageName: pkgName}
	if err := tmpl.Execute(file, fileName); err != nil {
		return errors.Errorf("模版赋值错误:%s", err)
	}
	return nil
}
