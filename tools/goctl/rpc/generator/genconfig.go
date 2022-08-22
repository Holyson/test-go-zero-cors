package generator

import (
	_ "embed"
	"io/ioutil"
	"os"
	"path/filepath"

	conf "github.com/Holyson/test-go-zero-cors/tools/goctl/config"
	"github.com/Holyson/test-go-zero-cors/tools/goctl/rpc/parser"
	"github.com/Holyson/test-go-zero-cors/tools/goctl/util/format"
	"github.com/Holyson/test-go-zero-cors/tools/goctl/util/pathx"
)

//go:embed config.tpl
var configTemplate string

// GenConfig generates the configuration structure definition file of the rpc service,
// which contains the zrpc.RpcServerConf configuration item by default.
// You can specify the naming style of the target file name through config.Config. For details,
// see https://github.com/Holyson/test-go-zero-cors/tree/master/tools/goctl/config/config.go
func (g *Generator) GenConfig(ctx DirContext, _ parser.Proto, cfg *conf.Config) error {
	dir := ctx.GetConfig()
	configFilename, err := format.FileNamingFormat(cfg.NamingFormat, "config")
	if err != nil {
		return err
	}

	fileName := filepath.Join(dir.Filename, configFilename+".go")
	if pathx.FileExists(fileName) {
		return nil
	}

	text, err := pathx.LoadTemplate(category, configTemplateFileFile, configTemplate)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fileName, []byte(text), os.ModePerm)
}
