package separation

import (
	"os"

	"github.com/GoAdminLive/go-admin/modules/config"
	adminTemplate "github.com/GoAdminLive/go-admin/template"
	"github.com/GoAdminLive/go-admin/template/components"
	"github.com/GoAdminLive/go-admin/template/types"

	"github.com/GoAdminLive/themes/common"
	"github.com/GoAdminLive/themes/sword/resource"
)

type Theme struct {
	ThemeName string
	components.Base
	*common.BaseTheme
}

var Sword = Theme{
	ThemeName: "sword_sep",
	Base: components.Base{
		Attribute: types.Attribute{
			TemplateList: common.SepTemplateList,
			Separation:   true,
		},
	},
	BaseTheme: &common.BaseTheme{
		AssetPaths:   resource.AssetPaths,
		TemplateList: common.SepTemplateList,
		Separation:   true,
	},
}

func init() {
	adminTemplate.Add("sword_sep", &Sword)
}

func Get() *Theme {
	return &Sword
}

func (t *Theme) Name() string {
	return t.ThemeName
}

func (t *Theme) GetTmplList() map[string]string {
	return common.SepTemplateList
}

func (t *Theme) GetAsset(path string) ([]byte, error) {
	return os.ReadFile(config.GetAssetRootPath() + path)
}

func (t *Theme) GetAssetList() []string {
	return resource.AssetsList
}
