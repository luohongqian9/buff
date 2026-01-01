package flags

import (
	"errors"
	"fmt"
	"os"
	"server/global"

	"github.com/urfave/cli"
	"go.uber.org/zap"
)

var (
	sqlFlag = &cli.BoolFlag{
		Name:  "sql",
		Usage: "初始化数据库",
	}
	sqlExportFlag = &cli.BoolFlag{
		Name:  "sql-export",
		Usage: "导出数据库",
	}
	sqlImportFlag = &cli.StringFlag{
		Name:  "sql-import",
		Usage: "导入数据库",
	}
	esFlag = &cli.BoolFlag{
		Name:  "es",
		Usage: "初始化ElasticSearch",
	}
	esImportFlag = &cli.StringFlag{
		Name:  "es-import",
		Usage: "导入ElasticSearch",
	}
	esExportFlag = &cli.BoolFlag{
		Name:  "es-export",
		Usage: "导出ElasticSearch",
	}
	adminFlag = &cli.BoolFlag{
		Name:  "admin",
		Usage: "创建管理员",
	}
)

func Run(c *cli.Context) {
	if c.NumFlags() > 1 {
		err := cli.NewExitError("只能输入一个参数", 1)
		global.Log.Error("Run", zap.Error(err))
		os.Exit(1)
	}

	switch {
	case c.Bool(sqlFlag.Name):
		if err := SQL(); err != nil {
			global.Log.Error("创建数据库失败", zap.Error(err))
			return
		} else {
			global.Log.Info("创建数据库成功")
		}
	case c.Bool(sqlExportFlag.Name):
		if err := SQLExport(); err != nil {
			global.Log.Error("导出数据库失败", zap.Error(err))
			return
		} else {
			global.Log.Info("导出数据库成功")
		}
	case c.IsSet(sqlImportFlag.Name):
		if errs := SQLImport(c.String(sqlImportFlag.Name)); len(errs) > 0 {
			var combinedErrors string
			for _, err := range errs {
				combinedErrors += err.Error() + "\n"
			}
			err := errors.New(combinedErrors)
			global.Log.Error("导入数据库失败", zap.Error(err))
			return
		} else {
			global.Log.Info("导入数据库成功")
		}
	case c.Bool(esFlag.Name):
		if err := Elasticsearch(); err != nil {
			global.Log.Error("创建ElasticSearch失败", zap.Error(err))
			return
		} else {
			global.Log.Info("创建ElasticSearch成功")
		}
	case c.IsSet(esImportFlag.Name):
		if num, err := ElasticSearchImport(c.String(esImportFlag.Name)); err != nil {
			global.Log.Error("导入ElasticSearch失败", zap.Error(err))
			return
		} else {
			global.Log.Info("导入ElasticSearch成功", zap.Int("num", num))
		}
	case c.Bool(esExportFlag.Name):
		if err := ElasticsearchExport(); err != nil {
			global.Log.Error("导出ElasticSearch失败", zap.Error(err))
			return
		} else {
			global.Log.Info("导出ElasticSearch成功")
		}
	case c.Bool(adminFlag.Name):
		if err := Admin(); err != nil {
			global.Log.Error("创建管理员失败", zap.Error(err))
			return
		} else {
			global.Log.Info("创建管理员成功")
		}
	default:
		err := cli.NewExitError("参数错误，请输入-h查看帮助", 1)
		global.Log.Error("Run", zap.Error(err))
	}
}

func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = "buff"
	app.Flags = []cli.Flag{
		sqlFlag,
		sqlExportFlag,
		sqlImportFlag,
		esFlag,
		esImportFlag,
		esExportFlag,
		adminFlag,
	}
	app.Action = Run
	return app
}

func InitFlag() {
	if len(os.Args) > 1 {
		app := NewApp()
		if err := app.Run(os.Args); err != nil {
			global.Log.Error("应用启动失败", zap.Error(err))
			os.Exit(1)
		}
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			fmt.Println(app.HelpName)
		}
		os.Exit(0)
	}
}
