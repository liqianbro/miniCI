package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/liqianbro/miniCI/templates"
	"github.com/liqianbro/miniCI/tool"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
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
	}
	if err := project.Create(); err != nil {
		return "", err
	}
	return project.AbsolutePath, nil
}

// generate 生成方法
func generate(args []string) error {
	var err error
	//获取当前文件路径
	str, _ := os.Getwd()
	//创建文件夹
	err = os.Mkdir(args[0], os.ModePerm)
	if err != nil {
		return errors.Wrapf(err, "创建目录错误！")
	}
	path := str + "/" + args[0]
	// init go mod
	err = InitMod(path, args[0])
	if err != nil {
		return err
	}
	//创建main文件
	err = CreateFile(path, "/main.go", "package main\n\nfunc main() {\n\tfmt.Println(\"hello,world\")\n}")
	if err != nil {
		return err
	}
	return err
}

type Mod struct {
	PackageName string
}

// CreateFile 创建文件并写入相应的内容
func CreateFile(filePath string, fileName string, content string) error {
	//创建main文件
	file, err := os.Create(filePath + fileName)
	defer file.Close()
	if err != nil {
		return errors.Errorf("创建文件发生错误：%s", err)
	}
	//写入文件
	_, err = file.WriteString(content)
	if err != nil {
		return errors.Errorf("写入文件发生错误：%s", err)
	}
	return nil
}

// ReadTemplate 读取模版文件
func ReadTemplate(path string) (string, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return "", errors.Errorf("读取文件出错:%s", err.Error())
	}
	return string(contents), nil
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

func GetPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	path = path[:index]
	return path
}
