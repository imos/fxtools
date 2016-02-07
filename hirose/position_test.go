package hirose

import (
	"github.com/PuerkitoBio/goquery"
	"reflect"
	"strings"
	"testing"
)

func TestParsePositionListPage(t *testing.T) {
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
<div>&gt;ﾎﾟｼﾞｼｮﾝ情報(一覧)&lt;</div>
<hr />


<div><a href="C304.html?position_id=1603600030779000&amp;symbol_code=&amp;f_y=2015&amp;f_m=02&amp;f_d=07&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;page_index=1">1603600030779000</a></div>
<div>
NZD/JPY&nbsp;&nbsp;
買
&nbsp;&nbsp;
残Lot:10&nbsp;
</div>
<div>
約定値:78.303&nbsp;&nbsp;評価:<span class="red">-7,050</span>
</div>
<a href="C002.html?position_id=1603600030779000&amp;symbol_code=&amp;f_y=2015&amp;f_m=02&amp;f_d=07&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;page_index=1&amp;next_page_flag=1&amp;prev_page_flag=1">決済<br /></a>
<br />
<div><a href="C304.html?position_id=1603600030714200&amp;symbol_code=&amp;f_y=2015&amp;f_m=02&amp;f_d=07&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;page_index=1">1603600030714200</a></div>
<div>
NZD/JPY&nbsp;&nbsp;
買
&nbsp;&nbsp;
残Lot:10&nbsp;
</div>
<div>
約定値:78.258&nbsp;&nbsp;評価:<span class="red">-6,600</span>
</div>
<a href="C002.html?position_id=1603600030714200&amp;symbol_code=&amp;f_y=2015&amp;f_m=02&amp;f_d=07&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;page_index=1&amp;next_page_flag=1&amp;prev_page_flag=1">決済<br /></a>
<br />
<div><a href="C304.html?position_id=1603600030639100&amp;symbol_code=&amp;f_y=2015&amp;f_m=02&amp;f_d=07&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;page_index=1">1603600030639100</a></div>
<div>
NZD/JPY&nbsp;&nbsp;
買
&nbsp;&nbsp;
残Lot:10&nbsp;
</div>
<div>
約定値:78.339&nbsp;&nbsp;評価:<span class="red">-7,410</span>
</div>
<a href="C002.html?position_id=1603600030639100&amp;symbol_code=&amp;f_y=2015&amp;f_m=02&amp;f_d=07&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;page_index=1&amp;next_page_flag=1&amp;prev_page_flag=1">決済<br /></a>
<br />
<div><a href="C304.html?position_id=1603600030449800&amp;symbol_code=&amp;f_y=2015&amp;f_m=02&amp;f_d=07&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;page_index=1">1603600030449800</a></div>
<div>
NZD/JPY&nbsp;&nbsp;
買
&nbsp;&nbsp;
残Lot:10&nbsp;
</div>
<div>
約定値:78.640&nbsp;&nbsp;評価:<span class="red">-10,420</span>
</div>
<a href="C002.html?position_id=1603600030449800&amp;symbol_code=&amp;f_y=2015&amp;f_m=02&amp;f_d=07&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;page_index=1&amp;next_page_flag=1&amp;prev_page_flag=1">決済<br /></a>
<br />
<div><a href="C304.html?position_id=1603600030438900&amp;symbol_code=&amp;f_y=2015&amp;f_m=02&amp;f_d=07&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;page_index=1">1603600030438900</a></div>
<div>
NZD/JPY&nbsp;&nbsp;
買
&nbsp;&nbsp;
残Lot:10&nbsp;
</div>
<div>
約定値:78.558&nbsp;&nbsp;評価:<span class="red">-9,600</span>
</div>
<a href="C002.html?position_id=1603600030438900&amp;symbol_code=&amp;f_y=2015&amp;f_m=02&amp;f_d=07&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;page_index=1&amp;next_page_flag=1&amp;prev_page_flag=1">決済<br /></a>
<br />
<div><a href="C304.html?position_id=1603600030435400&amp;symbol_code=&amp;f_y=2015&amp;f_m=02&amp;f_d=07&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;page_index=1">1603600030435400</a></div>
<div>
NZD/JPY&nbsp;&nbsp;
買
&nbsp;&nbsp;
残Lot:10&nbsp;
</div>
<div>
約定値:78.548&nbsp;&nbsp;評価:<span class="red">-9,500</span>
</div>
<a href="C002.html?position_id=1603600030435400&amp;symbol_code=&amp;f_y=2015&amp;f_m=02&amp;f_d=07&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;page_index=1&amp;next_page_flag=1&amp;prev_page_flag=1">決済<br /></a>
<br />
<div><a href="C304.html?position_id=1603600030431300&amp;symbol_code=&amp;f_y=2015&amp;f_m=02&amp;f_d=07&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;page_index=1">1603600030431300</a></div>
<div>
NZD/JPY&nbsp;&nbsp;
買
&nbsp;&nbsp;
残Lot:10&nbsp;
</div>
<div>
約定値:78.536&nbsp;&nbsp;評価:<span class="red">-9,380</span>
</div>
<a href="C002.html?position_id=1603600030431300&amp;symbol_code=&amp;f_y=2015&amp;f_m=02&amp;f_d=07&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;page_index=1&amp;next_page_flag=1&amp;prev_page_flag=1">決済<br /></a>
<br />
<div><a href="C304.html?position_id=1603600030414100&amp;symbol_code=&amp;f_y=2015&amp;f_m=02&amp;f_d=07&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;page_index=1">1603600030414100</a></div>
<div>
NZD/JPY&nbsp;&nbsp;
買
&nbsp;&nbsp;
残Lot:10&nbsp;
</div>
<div>
約定値:78.481&nbsp;&nbsp;評価:<span class="red">-8,830</span>
</div>
<a href="C002.html?position_id=1603600030414100&amp;symbol_code=&amp;f_y=2015&amp;f_m=02&amp;f_d=07&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;page_index=1&amp;next_page_flag=1&amp;prev_page_flag=1">決済<br /></a>
<br />
<div><a href="C304.html?position_id=1603600030396800&amp;symbol_code=&amp;f_y=2015&amp;f_m=02&amp;f_d=07&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;page_index=1">1603600030396800</a></div>
<div>
NZD/JPY&nbsp;&nbsp;
買
&nbsp;&nbsp;
残Lot:100&nbsp;
</div>
<div>
約定値:78.534&nbsp;&nbsp;評価:<span class="red">-93,600</span>
</div>
<a href="C002.html?position_id=1603600030396800&amp;symbol_code=&amp;f_y=2015&amp;f_m=02&amp;f_d=07&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;page_index=1&amp;next_page_flag=1&amp;prev_page_flag=1">決済<br /></a>
<br />
<div><a href="C304.html?position_id=1603600030382800&amp;symbol_code=&amp;f_y=2015&amp;f_m=02&amp;f_d=07&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;page_index=1">1603600030382800</a></div>
<div>
NZD/JPY&nbsp;&nbsp;
買
&nbsp;&nbsp;
残Lot:10&nbsp;
</div>
<div>
約定値:78.501&nbsp;&nbsp;評価:<span class="red">-9,030</span>
</div>
<a href="C002.html?position_id=1603600030382800&amp;symbol_code=&amp;f_y=2015&amp;f_m=02&amp;f_d=07&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;page_index=1&amp;next_page_flag=1&amp;prev_page_flag=1">決済<br /></a>
<br />
<div><a href="CS01.html?symbol_code=&amp;f_y=2015&amp;f_m=02&amp;f_d=07&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;page_index=1&amp;prev_page_flag=1">全決済</a></div>

<hr />
<div>
<a href="C302.html?symbol_code=&amp;f_y=2015&amp;f_m=02&amp;f_d=07&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;page_index=2" accesskey="#">[#]次へ</a>
</div>
<div><a href="C301.html">ﾎﾟｼﾞｼｮﾝ情報(検索)</a></div>
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
	positions, nextPage, err := parsePositionListPage(doc)
	if err != nil {
		t.Fatalf("got %v, want nil", err)
	}
	if !nextPage {
		t.Fatalf("nextPage got false, want true")
	}
	expected := []Position{
		Position{PositionId: "1603600030779000"},
		Position{PositionId: "1603600030714200"},
		Position{PositionId: "1603600030639100"},
		Position{PositionId: "1603600030449800"},
		Position{PositionId: "1603600030438900"},
		Position{PositionId: "1603600030435400"},
		Position{PositionId: "1603600030431300"},
		Position{PositionId: "1603600030414100"},
		Position{PositionId: "1603600030396800"},
		Position{PositionId: "1603600030382800"},
	}
	if !reflect.DeepEqual(expected, positions) {
		for k, p := range positions {
			t.Errorf("positions[%#v]=%#v", k, p)
		}
		t.Fatalf("wrong positions: %#v", positions)
	}
}

func TestParsePositionDetailsPage(t *testing.T) {
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
<div>&gt;ﾎﾟｼﾞｼｮﾝ情報(詳細)&lt;</div>
<hr />
<div align="left">NZD/JPY</div>
<hr />

<div>ﾎﾟｼﾞｼｮﾝ番号：</div>
<div align="right">1603600030779000</div>
<div>約定日時：</div>
<div align="right">2016/02/05 22:38:13</div>
<div>売買：</div>
<div align="right">
買
</div>
<div>約定価格：</div>
<div align="right">78.303</div>
<div>約定Lot数：</div>
<div align="right">10</div>
<div>残Lot数：</div>
<div align="right">10</div>
<div>ﾎﾟｼﾞｼｮﾝ損益：</div>
<div align="right">
<span class="red">-7,110</span>
</div>
<div>未実現ｽﾜｯﾌﾟ：</div>
<div align="right">
<div align="right">60</div>
</div>
<div>評価損益：</div>
<div align="right">
<span class="red">-7,050</span>
</div>
<div><a href="C302.html?symbol_code=&amp;f_y=2015&amp;f_m=02&amp;f_d=07&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;page_index=">戻る</a></div>
<hr />
<div>
<a href="C002.html?position_id=1603600030779000&amp;symbol_code=&amp;f_y=2015&amp;f_m=02&amp;f_d=07&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;page_index=&amp;next_page_flag=1&amp;prev_page_flag=2">成行決済</a>&nbsp;
<a href="C002.html?position_id=1603600030779000&amp;symbol_code=&amp;f_y=2015&amp;f_m=02&amp;f_d=07&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;page_index=&amp;next_page_flag=2&amp;prev_page_flag=2">指値(逆指)決済</a>
</div>
<div>
<a href="C002.html?position_id=1603600030779000&amp;symbol_code=&amp;f_y=2015&amp;f_m=02&amp;f_d=07&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;page_index=&amp;next_page_flag=3&amp;prev_page_flag=2">OCO決済</a>
</div>

<hr />
<div><a href="C301.html">ﾎﾟｼﾞｼｮﾝ情報(検索)</a></div>


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
	position, err := parsePositionDetailsPage(doc)
	if err != nil {
		t.Fatalf("got %v, want nil", err)
	}
	expected := Position{
		Currency:        "NZD/JPY",
		PositionId:      "1603600030779000",
		TransactionTime: "2016/02/05 22:38:13",
		Side:            "buy",
		TransactionRate: 78.303,
		Amount:          10000,
	}
	if !reflect.DeepEqual(&expected, position) {
		t.Fatalf("wrong positions: %#v", position)
	}
}
