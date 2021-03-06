package compose_to_service_yml

import (
	"github.com/cloud66-oss/starter/common"
	"github.com/cloud66-oss/starter/definitions/docker-compose"
	"github.com/cloud66-oss/starter/packs"
	"github.com/cloud66-oss/starter/transform"
)

type Pack struct {
	packs.PackBase
	Analysis *Analysis
}

const (
	StencilTemplatePath = "" //TODO: still not implemented
)

func (p *Pack) Name() string {
	return "docker-compose"
}

func (p *Pack) LanguageVersion() string {
	return p.Analysis.LanguageVersion
}

func (p *Pack) FilesToBeAnalysed() []string {
	return []string{"docker-compose.yml"}
}

func (p *Pack) Framework() string {
	return ""
}

func (p *Pack) FrameworkVersion() string {
	return ""
}

func (p *Pack) GetSupportedLanguageVersions() []string {
	return nil
}

func (p *Pack) SetSupportedLanguageVersions(versions []string) {

}

func (p *Pack) Detector() packs.Detector {
	return &Detector{PackElement: packs.PackElement{Pack: p}}
}

func (p *Pack) Analyze(rootDir string, environment string, shouldPrompt bool, git_repo string, git_branch string) error {
	var err error
	a := Analyzer{
		AnalyzerBase: packs.AnalyzerBase{
			PackElement:  packs.PackElement{Pack: p},
			RootDir:      rootDir,
			ShouldPrompt: shouldPrompt,
			GitURL:       git_repo,
			GitBranch:    git_branch,
			Environment:  environment}}
	p.Analysis, err = a.Analyze()
	return err
}

func (p *Pack) WriteDockerfile(templateDir string, outputDir string, shouldPrompt bool) error {

	common.PrintlnWarning("You can not generate a Dockerfile using this pack. Nothing to do.")
	return nil
}

func (p *Pack) WriteServiceYAML(templateDir string, outputDir string, shouldPrompt bool) error {

	dockerBase := docker_compose.DockerCompose{}
	dockerBase.UnmarshalFromFile(outputDir + "/docker-compose.yml")

	d := transform.DockerComposeTransformer{Base: dockerBase}

	serviceYml := d.ToServiceYml(p.Analysis.GitURL, p.Analysis.GitBranch, shouldPrompt, outputDir+"/docker-compose.yml")
	serviceYml.MarshalToFile(outputDir + "/service.yml")

	return nil
}

func (p *Pack) WriteDockerComposeYAML(templateDir string, outputDir string, shouldPrompt bool) error {
	common.PrintlnWarning("There is already an existing docker-compose.yml. Nothing to do.")
	return nil
}

func (p *Pack) WriteKubesConfig(outputDir string, shouldPrompt bool) error {
	common.PrintlnWarning("You can not generate a Kubernetes configuration file using this pack. Nothing to do.")
	return nil
}

func (p *Pack) GetMessages() []string {
	return []string{}
}

func (p *Pack) GetDatabases() []string {
	return []string{}
}

func (p *Pack) GetStartCommands() []string {
	return []string{}
}

func (p *Pack) StencilRepositoryPath() string {
	return StencilTemplatePath
}

func (p *Pack) CreateSkycapFiles(outputDir string, templateDir string, branch string) error {
	common.PrintlnWarning("You can not generate the Skycap configuration files using this pack. Nothing to do.")
	return nil
}
