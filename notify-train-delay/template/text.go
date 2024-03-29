package template

// Information return string
const Information string = `⚠️
入力に誤りがあります。

地域ごとの遅延情報を閲覧したい場合は以下のいずれかの番号を入力してください。
⚠️数値は半角数字でないと正しく読み取れません⚠️
####################
2    : 北海道
3    : 東北
1, 4: 関東
5    : 中部
6    : 近畿
7    : 九州
8    : 中国
9    : 四国
####################

路線ごとの遅延情報を閲覧したい場合は路線名を入力してください(lightbulb)
⚠️以下はあくまでも例です⚠️
####################
山手線
####################
`

// NotDelay return string
const NotDelay string = `
現在、正常に運行しています。
`

// TooMuchData return string
const TooMuchData = "路線のヒット数が多すぎます。もう少し具体的な路線名を入力してください。"
