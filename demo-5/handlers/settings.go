package handlers

import (
	"net/http"

	"demo-5/config"
	"demo-5/models"

	"github.com/gin-gonic/gin"
)

// Response 统一响应结构
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

// GetSettings 获取设置信息
func GetSettings(c *gin.Context) {
	// 检查数据库连接是否可用
	if config.DB == nil {
		// 数据库连接不可用，返回默认数据
		responseData := getDefaultSettings()
		response := Response{
			Success: true,
			Data:    responseData,
			Message: "获取设置成功（使用默认数据）",
		}
		c.JSON(http.StatusOK, response)
		return
	}

	var settings config.UserSettings

	// 使用GORM从数据库查询第一条记录
	result := config.DB.First(&settings)
	if result.Error != nil {
		// 如果查询失败，返回默认数据
		responseData := getDefaultSettings()
		response := Response{
			Success: true,
			Data:    responseData,
			Message: "获取设置成功（使用默认数据）",
		}
		c.JSON(http.StatusOK, response)
		return
	}

	// 构建返回数据，保持与原有结构兼容
	responseData := models.Settings{
		SoundType:             settings.SoundType,
		WordSound:             settings.WordSound,
		WordSoundVolume:       settings.WordSoundVolume,
		WordSoundSpeed:        settings.WordSoundSpeed,
		ArticleSound:          settings.ArticleSound,
		ArticleAutoPlayNext:   settings.ArticleAutoPlayNext,
		ArticleSoundVolume:    settings.ArticleSoundVolume,
		ArticleSoundSpeed:     settings.ArticleSoundSpeed,
		KeyboardSound:         settings.KeyboardSound,
		KeyboardSoundVolume:   settings.KeyboardSoundVolume,
		KeyboardSoundFile:     settings.KeyboardSoundFile,
		EffectSound:           settings.EffectSound,
		EffectSoundVolume:     settings.EffectSoundVolume,
		RepeatCount:           settings.RepeatCount,
		RepeatCustomCount:     settings.RepeatCustomCount,
		Dictation:             settings.Dictation,
		Translate:             settings.Translate,
		ShowNearWord:          settings.ShowNearWord,
		IgnoreCase:            settings.IgnoreCase,
		AllowWordTip:          settings.AllowWordTip,
		WaitTimeForChangeWord: settings.WaitTimeForChangeWord,
		FontSize: models.FontSize{
			ArticleForeignFontSize:   settings.ArticleForeignFontSize,
			ArticleTranslateFontSize: settings.ArticleTranslateFontSize,
			WordForeignFontSize:      settings.WordForeignFontSize,
			WordTranslateFontSize:    settings.WordTranslateFontSize,
		},
		ShowToolbar:                      settings.ShowToolbar,
		ShowPanel:                        settings.ShowPanel,
		SideExpand:                       settings.SideExpand,
		Theme:                            settings.Theme,
		First:                            settings.First,
		FirstTime:                        settings.FirstTime,
		Load:                             settings.Load,
		ConflictNotice:                   settings.ConflictNotice,
		IgnoreSimpleWord:                 settings.IgnoreSimpleWord,
		WordPracticeMode:                 settings.WordPracticeMode,
		WordPracticeType:                 settings.WordPracticeType,
		DisableShowPracticeSettingDialog: settings.DisableShowPracticeSettingDialog,
		AutoNextWord:                     settings.AutoNextWord,
		InputWrongClear:                  settings.InputWrongClear,
		ShortcutKeyMap: map[string]string{
			"EditArticle":           "Ctrl+E",
			"ShowWord":              "Escape",
			"Previous":              "Alt+⬅",
			"Next":                  "Tab",
			"ToggleSimple":          "",
			"ToggleCollect":         "Enter",
			"PreviousChapter":       "Ctrl+⬅",
			"NextChapter":           "Ctrl+➡",
			"RepeatChapter":         "Ctrl+Enter",
			"DictationChapter":      "Alt+Enter",
			"PlayWordPronunciation": "Ctrl+P",
			"ToggleShowTranslate":   "Ctrl+Z",
			"ToggleDictation":       "Ctrl+I",
			"ToggleTheme":           "Ctrl+Q",
			"ToggleConciseMode":     "Ctrl+M",
			"TogglePanel":           "Ctrl+L",
			"RandomWrite":           "Ctrl+R",
			"NextRandomWrite":       "Ctrl+Shift+R",
			"KnowWord":              "1",
			"UnknownWord":           "2",
		},
	}

	// 返回新的JSON结构
	response := Response{
		Success: true,
		Data:    responseData,
		Message: "获取设置成功",
	}

	c.JSON(http.StatusOK, response)
}

// getDefaultSettings 获取默认设置数据
func getDefaultSettings() models.Settings {
	return models.Settings{
		SoundType:             "uk",
		WordSound:             true,
		WordSoundVolume:       80,
		WordSoundSpeed:        1,
		ArticleSound:          true,
		ArticleAutoPlayNext:   false,
		ArticleSoundVolume:    100,
		ArticleSoundSpeed:     1,
		KeyboardSound:         true,
		KeyboardSoundVolume:   100,
		KeyboardSoundFile:     "笔记本键盘",
		EffectSound:           true,
		EffectSoundVolume:     100,
		RepeatCount:           1,
		RepeatCustomCount:     nil,
		Dictation:             false,
		Translate:             true,
		ShowNearWord:          true,
		IgnoreCase:            false,
		AllowWordTip:          true,
		WaitTimeForChangeWord: 300,
		FontSize: models.FontSize{
			ArticleForeignFontSize:   48,
			ArticleTranslateFontSize: 20,
			WordForeignFontSize:      48,
			WordTranslateFontSize:    20,
		},
		ShowToolbar: true,
		ShowPanel:   true,
		SideExpand:  false,
		Theme:       "auto",
		ShortcutKeyMap: map[string]string{
			"EditArticle":           "Ctrl+E",
			"ShowWord":              "Escape",
			"Previous":              "Alt+⬅",
			"Next":                  "Tab",
			"ToggleSimple":          "",
			"ToggleCollect":         "Enter",
			"PreviousChapter":       "Ctrl+⬅",
			"NextChapter":           "Ctrl+➡",
			"RepeatChapter":         "Ctrl+Enter",
			"DictationChapter":      "Alt+Enter",
			"PlayWordPronunciation": "Ctrl+P",
			"ToggleShowTranslate":   "Ctrl+Z",
			"ToggleDictation":       "Ctrl+I",
			"ToggleTheme":           "Ctrl+Q",
			"ToggleConciseMode":     "Ctrl+M",
			"TogglePanel":           "Ctrl+L",
			"RandomWrite":           "Ctrl+R",
			"NextRandomWrite":       "Ctrl+Shift+R",
			"KnowWord":              "1",
			"UnknownWord":           "2",
		},
		First:                            true,
		FirstTime:                        1731582332678,
		Load:                             false,
		ConflictNotice:                   true,
		IgnoreSimpleWord:                 false,
		WordPracticeMode:                 0,
		WordPracticeType:                 0,
		DisableShowPracticeSettingDialog: false,
		AutoNextWord:                     true,
		InputWrongClear:                  false,
	}
}
