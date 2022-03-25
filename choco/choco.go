package choco

type Choco struct {
	Question string
	Answer   string
}

var Xlist = []Choco{
	{
		Question: "   雑学クイズ\n\n  始める・・Enter",
		Answer:   "Enter",
	},
	{
		Question: "4+5=",
		Answer:   "9",
	},
	{
		Question: "私の誕生月は？",
		Answer:   "0",
	},
	{
		Question: "私は2022年4月に\n中学何年生になる？",
		Answer:   "3",
	},
	{
		Question: "全問正解！！",
		Answer:   "Enter",
	},
}
