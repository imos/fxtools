package hirose

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
	"testing"
)

func TestParseStatus(t *testing.T) {
	input := `
<?xml version="1.0" encoding="Shift_JIS"?>
<!DOCTYPE html PUBLIC "-//OPENWAVE//DTD XHTML 1.0//EN" "http://www.openwave.com/DTD/xhtml-basic.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=shift_jis" />
<meta name="viewport" content="width=320; initial-scale=1.0; maximum-scale=1.0; user-scalable=0;" />
<meta name="disparea" content="vga" />
<title>LION FX</title>
<link href="../css/style.css" rel="stylesheet" type="text/css" />
</head>
<body>
<div class="container">
<div>&gt;証拠金状況照会&lt;</div>
<hr />

<div>証拠金預託額：</div>
<div align="right">
1,850,000
</div>
<div>ﾎﾟｼﾞｼｮﾝ損益：</div>
<div align="right">
<span class="red">-257,710</span>
</div>
<div>未実現ｽﾜｯﾌﾟ：</div>
<div align="right">
5,400
</div>
<div>評価損益：</div>
<div align="right">
<span class="red">-252,310</span>
</div>
<div>必要証拠金額：</div>
<div align="right">
960,000
</div>
<div>発注証拠金額：</div>
<div align="right">
0
</div>
<div>有効証拠金額：</div>
<div align="right">
1,597,690
</div>
<div>有効比率：</div>
<div align="right">166.42 %</div>
<div>ｱﾗｰﾄ基準額：</div>
<div align="right">
1,920,000
</div>
<div>ﾛｽｶｯﾄ基準額：</div>
<div align="right">
960,000
</div>
<div>発注可能額：</div>
<div align="right">
637,690
</div>
<div>出金可能額：</div>
<div align="right">
637,690
</div>
<div>出金依頼額：</div>
<div align="right">
0
</div>
<div>金額指定全決済：</div>
<div align="right">
使わない
</div>
<div>金額指定全決済判定基準：</div>
<div align="right">
評価損益
</div>
<div>金額指定全決済(上限)：</div>
<div align="right">
---
</div>
<div>金額指定全決済(下限)：</div>
<div align="right">
---
</div>
<div>新規注文取消(金額指定)：</div>
<div align="right">
取消しない
</div>
<div>時間指定全決済：</div>
<div align="right">
使わない
</div>
<div>全決済指定時間：</div>
<div align="right">
---
</div>
<div>新規注文取消(時間指定)：</div>
<div align="right">
取消しない
</div>

<hr />
<div><a href="../M002.html" accesskey="1">[1]お知らせ</a></div> 
<div><a href="../common/I101.html" accesskey="2">[2]ﾚｰﾄ</a></div> 
<div><a href="../common/I311.html" accesskey="3">[3]ﾁｬｰﾄ</a></div> 
<div><a href="../M003.html" accesskey="4">[4]取引</a></div> 
<div><a href="../I005.html" accesskey="5">[5]ﾆｭｰｽ</a></div> 
<div><a href="../M005.html" accesskey="6">[6]照会</a></div> 
<div><a href="../M006.html" accesskey="7">[7]入出金</a></div> 
<div><a href="../M004.html" accesskey="8">[8]設定</a></div> 
<div><a href="../M008.html" accesskey="9">[9]小林芳彦のﾏｰｹｯﾄﾅﾋﾞ</a></div> 
<div><a href="../M001.html" accesskey="0">[0]ﾒｲﾝﾒﾆｭｰ</a></div>
<hr />
</div>
</body>
</html>
`
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(input))
	status, err := parseStatusPage(doc)
	if err != nil {
		t.Fatalf("got %v, want nil", err)
	}
	if *status.ActualDeposit != 1597690 {
		t.Errorf("ActualDeposit should be 1597690, but %d.", *status.ActualDeposit)
	}
	if *status.NecessaryDeposit != 960000 {
		t.Errorf("NecessaryDeposit should be 960000, but %d.", *status.NecessaryDeposit)
	}
}
