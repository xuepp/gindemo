package models

// FontSize 字体大小设置
type FontSize struct {
	ArticleForeignFontSize   int `json:"articleForeignFontSize"`
	ArticleTranslateFontSize int `json:"articleTranslateFontSize"`
	WordForeignFontSize      int `json:"wordForeignFontSize"`
	WordTranslateFontSize    int `json:"wordTranslateFontSize"`
}

// Settings 用户设置
type Settings struct {
	// 发音类型（us=美式，uk=英式）
	SoundType string `json:"soundType"`

	// 是否朗读单词
	WordSound bool `json:"wordSound"`
	// 单词朗读音量（0~100）
	WordSoundVolume int `json:"wordSoundVolume"`
	// 单词朗读语速（1=正常）
	WordSoundSpeed float64 `json:"wordSoundSpeed"`

	// 是否朗读文章
	ArticleSound bool `json:"articleSound"`
	// 朗读文章时是否自动播放下一段
	ArticleAutoPlayNext bool `json:"articleAutoPlayNext"`
	// 文章朗读音量
	ArticleSoundVolume int `json:"articleSoundVolume"`
	// 文章朗读速度
	ArticleSoundSpeed float64 `json:"articleSoundSpeed"`

	// 键盘音效开关
	KeyboardSound bool `json:"keyboardSound"`
	// 键盘音效音量
	KeyboardSoundVolume int `json:"keyboardSoundVolume"`
	// 键盘音效文件名称
	KeyboardSoundFile string `json:"keyboardSoundFile"`

	// 是否启用效果音（提示音、翻页音等）
	EffectSound bool `json:"effectSound"`
	// 效果音音量
	EffectSoundVolume int `json:"effectSoundVolume"`

	// 单词重复次数（默认）
	RepeatCount int `json:"repeatCount"`
	// 自定义重复次数（null = 未指定）
	RepeatCustomCount *int `json:"repeatCustomCount"`

	// 是否开启听写模式
	Dictation bool `json:"dictation"`

	// 是否显示翻译
	Translate bool `json:"translate"`

	// 是否显示相近单词（拼写相似）
	ShowNearWord bool `json:"showNearWord"`

	// 匹配单词时是否忽略大小写
	IgnoreCase bool `json:"ignoreCase"`

	// 是否允许显示单词提示（鼠标悬停等）
	AllowWordTip bool `json:"allowWordTip"`

	// 切换单词前等待的毫秒数
	WaitTimeForChangeWord int `json:"waitTimeForChangeWord"`

	// 字体大小设置
	FontSize FontSize `json:"fontSize"`

	// 是否显示工具栏
	ShowToolbar bool `json:"showToolbar"`
	// 是否显示侧栏
	ShowPanel bool `json:"showPanel"`

	// 侧栏是否展开
	SideExpand bool `json:"sideExpand"`

	// 主题模式（auto/light/dark）
	Theme string `json:"theme"`

	// 快捷键映射
	ShortcutKeyMap map[string]string `json:"shortcutKeyMap"`

	// 是否第一次使用
	First bool `json:"first"`
	// 第一次使用时间戳（毫秒）
	FirstTime int64 `json:"firstTime"`

	// 是否加载完成
	Load bool `json:"load"`

	// 快捷键冲突提醒
	ConflictNotice bool `json:"conflictNotice"`

	// 是否忽略简单词（如 a / the / to）
	IgnoreSimpleWord bool `json:"ignoreSimpleWord"`

	// 单词练习模式（枚举值）
	WordPracticeMode int `json:"wordPracticeMode"`
	// 单词练习类型（枚举值）
	WordPracticeType int `json:"wordPracticeType"`

	// 禁用练习设置弹窗
	DisableShowPracticeSettingDialog bool `json:"disableShowPracticeSettingDialog"`

	// 自动跳到下一个单词
	AutoNextWord bool `json:"autoNextWord"`

	// 输入错误时自动清空输入框
	InputWrongClear bool `json:"inputWrongClear"`
}
