package cmd

import (
	"os"
	"text/template"

	"github.com/liqianbro/mini_ci/templates"
)

// Project contains name, license and paths to projects.
type Project struct {
	// v2
	PkgName      string
	Copyright    string
	AbsolutePath string
	Legal        License
	AppName      string
}

func (p *Project) Create() error {
	// create directory
	if err := os.Mkdir(p.AbsolutePath, 0754); err != nil {
		return err
	}
	//create controller
	if err := os.Mkdir(p.AbsolutePath+"/controller", 0754); err != nil {
		return err
	}
	if err := os.Mkdir(p.AbsolutePath+"/router", 0754); err != nil {
		return err
	}
	if err := os.Mkdir(p.AbsolutePath+"/middleware", 0754); err != nil {
		return err
	}

	// create main.go
	mainFile, err := os.Create(p.AbsolutePath + "/main.go")
	if err != nil {
		return err
	}
	defer mainFile.Close()

	// create god mod
	err = InitMod(p.AbsolutePath, p.PkgName)
	if err != nil {
		return err
	}
	// create template
	mainTemplate := template.Must(template.New("main").Parse(string(templates.MainTemplate())))
	err = mainTemplate.Execute(mainFile, p)
	if err != nil {
		return err
	}

	// create license
	return p.createLicenseFile()
}

func (p *Project) createLicenseFile() error {
	data := map[string]interface{}{
		"copyright": CopyrightLine(),
	}
	licenseFile, err := os.Create(p.AbsolutePath + "/LICENSE")
	if err != nil {
		return err
	}
	defer licenseFile.Close()

	licenseTemplate := template.Must(template.New("license").Parse(p.Legal.Text))
	return licenseTemplate.Execute(licenseFile, data)
}
