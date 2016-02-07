package hirose

import (
	"github.com/PuerkitoBio/goquery"
	"reflect"
	"strings"
	"testing"
)

func TestParseOrderListPage(t *testing.T) {
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
<div>&gt;注文情報(一覧)&lt;</div>
<hr />


<div><a href="../otc/C403.html?order_id=1603600082705700&amp;order_method=11&amp;page_index=1&amp;symbol_code=&amp;open_close_type=&amp;f_y=2015&amp;f_m=02&amp;f_d=01&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59">1603600082705700-11</a></div>
<div>
NZD/JPY&nbsp;
売
&nbsp;
区分：
指定決済
</div>
<div>
注Lot:10&nbsp;
執行:
指値
(82.300)
</div>
<div>
ﾄﾘｶﾞｰ価格:---&nbsp;
</div>
<a href="../otc/C003.html?order_id=1603600082705700&amp;order_method=11&amp;page_index=1&amp;symbol_code=&amp;open_close_type=&amp;f_y=2015&amp;f_m=02&amp;f_d=01&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;prev_page_flag=1">変更</a>&nbsp;&nbsp;
<a href="../otc/CT01.html?order_id=1603600082705700&amp;order_method=11&amp;page_index=1&amp;symbol_code=&amp;open_close_type=&amp;f_y=2015&amp;f_m=02&amp;f_d=01&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;prev_page_flag=1">取消</a>
<div><br /></div>
<div><a href="../otc/C403.html?order_id=1603600082705700&amp;order_method=12&amp;page_index=1&amp;symbol_code=&amp;open_close_type=&amp;f_y=2015&amp;f_m=02&amp;f_d=01&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59">1603600082705700-12</a></div>
<div>
NZD/JPY&nbsp;
売
&nbsp;
区分：
指定決済
</div>
<div>
注Lot:10&nbsp;
執行:
逆指
(74.300)
</div>
<div>
ﾄﾘｶﾞｰ価格:---&nbsp;
</div>
<a href="../otc/C003.html?order_id=1603600082705700&amp;order_method=12&amp;page_index=1&amp;symbol_code=&amp;open_close_type=&amp;f_y=2015&amp;f_m=02&amp;f_d=01&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;prev_page_flag=1">変更</a>&nbsp;&nbsp;
<a href="../otc/CT01.html?order_id=1603600082705700&amp;order_method=12&amp;page_index=1&amp;symbol_code=&amp;open_close_type=&amp;f_y=2015&amp;f_m=02&amp;f_d=01&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;prev_page_flag=1">取消</a>
<div><br /></div>
<div><a href="../otc/C403.html?order_id=1603600082672700&amp;order_method=11&amp;page_index=1&amp;symbol_code=&amp;open_close_type=&amp;f_y=2015&amp;f_m=02&amp;f_d=01&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59">1603600082672700-11</a></div>
<div>
NZD/JPY&nbsp;
売
&nbsp;
区分：
指定決済
</div>
<div>
注Lot:10&nbsp;
執行:
指値
(82.300)
</div>
<div>
ﾄﾘｶﾞｰ価格:---&nbsp;
</div>
<a href="../otc/C003.html?order_id=1603600082672700&amp;order_method=11&amp;page_index=1&amp;symbol_code=&amp;open_close_type=&amp;f_y=2015&amp;f_m=02&amp;f_d=01&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;prev_page_flag=1">変更</a>&nbsp;&nbsp;
<a href="../otc/CT01.html?order_id=1603600082672700&amp;order_method=11&amp;page_index=1&amp;symbol_code=&amp;open_close_type=&amp;f_y=2015&amp;f_m=02&amp;f_d=01&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;prev_page_flag=1">取消</a>
<div><br /></div>
<div><a href="../otc/C403.html?order_id=1603600082672700&amp;order_method=12&amp;page_index=1&amp;symbol_code=&amp;open_close_type=&amp;f_y=2015&amp;f_m=02&amp;f_d=01&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59">1603600082672700-12</a></div>
<div>
NZD/JPY&nbsp;
売
&nbsp;
区分：
指定決済
</div>
<div>
注Lot:10&nbsp;
執行:
逆指
(74.300)
</div>
<div>
ﾄﾘｶﾞｰ価格:---&nbsp;
</div>
<a href="../otc/C003.html?order_id=1603600082672700&amp;order_method=12&amp;page_index=1&amp;symbol_code=&amp;open_close_type=&amp;f_y=2015&amp;f_m=02&amp;f_d=01&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;prev_page_flag=1">変更</a>&nbsp;&nbsp;
<a href="../otc/CT01.html?order_id=1603600082672700&amp;order_method=12&amp;page_index=1&amp;symbol_code=&amp;open_close_type=&amp;f_y=2015&amp;f_m=02&amp;f_d=01&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;prev_page_flag=1">取消</a>
<div><br /></div>
<div><a href="../otc/C403.html?order_id=1603600082617000&amp;order_method=11&amp;page_index=1&amp;symbol_code=&amp;open_close_type=&amp;f_y=2015&amp;f_m=02&amp;f_d=01&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59">1603600082617000-11</a></div>
<div>
NZD/JPY&nbsp;
売
&nbsp;
区分：
指定決済
</div>
<div>
注Lot:10&nbsp;
執行:
指値
(82.300)
</div>
<div>
ﾄﾘｶﾞｰ価格:---&nbsp;
</div>
<a href="../otc/C003.html?order_id=1603600082617000&amp;order_method=11&amp;page_index=1&amp;symbol_code=&amp;open_close_type=&amp;f_y=2015&amp;f_m=02&amp;f_d=01&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;prev_page_flag=1">変更</a>&nbsp;&nbsp;
<a href="../otc/CT01.html?order_id=1603600082617000&amp;order_method=11&amp;page_index=1&amp;symbol_code=&amp;open_close_type=&amp;f_y=2015&amp;f_m=02&amp;f_d=01&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;prev_page_flag=1">取消</a>
<div><br /></div>
<div><a href="../otc/C403.html?order_id=1603600082617000&amp;order_method=12&amp;page_index=1&amp;symbol_code=&amp;open_close_type=&amp;f_y=2015&amp;f_m=02&amp;f_d=01&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59">1603600082617000-12</a></div>
<div>
NZD/JPY&nbsp;
売
&nbsp;
区分：
指定決済
</div>
<div>
注Lot:10&nbsp;
執行:
逆指
(74.300)
</div>
<div>
ﾄﾘｶﾞｰ価格:---&nbsp;
</div>
<a href="../otc/C003.html?order_id=1603600082617000&amp;order_method=12&amp;page_index=1&amp;symbol_code=&amp;open_close_type=&amp;f_y=2015&amp;f_m=02&amp;f_d=01&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;prev_page_flag=1">変更</a>&nbsp;&nbsp;
<a href="../otc/CT01.html?order_id=1603600082617000&amp;order_method=12&amp;page_index=1&amp;symbol_code=&amp;open_close_type=&amp;f_y=2015&amp;f_m=02&amp;f_d=01&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;prev_page_flag=1">取消</a>
<div><br /></div>
<div><a href="../otc/C403.html?order_id=1603600082542000&amp;order_method=11&amp;page_index=1&amp;symbol_code=&amp;open_close_type=&amp;f_y=2015&amp;f_m=02&amp;f_d=01&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59">1603600082542000-11</a></div>
<div>
NZD/JPY&nbsp;
売
&nbsp;
区分：
指定決済
</div>
<div>
注Lot:10&nbsp;
執行:
指値
(82.300)
</div>
<div>
ﾄﾘｶﾞｰ価格:---&nbsp;
</div>
<a href="../otc/C003.html?order_id=1603600082542000&amp;order_method=11&amp;page_index=1&amp;symbol_code=&amp;open_close_type=&amp;f_y=2015&amp;f_m=02&amp;f_d=01&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;prev_page_flag=1">変更</a>&nbsp;&nbsp;
<a href="../otc/CT01.html?order_id=1603600082542000&amp;order_method=11&amp;page_index=1&amp;symbol_code=&amp;open_close_type=&amp;f_y=2015&amp;f_m=02&amp;f_d=01&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;prev_page_flag=1">取消</a>
<div><br /></div>
<div><a href="../otc/C403.html?order_id=1603600082542000&amp;order_method=12&amp;page_index=1&amp;symbol_code=&amp;open_close_type=&amp;f_y=2015&amp;f_m=02&amp;f_d=01&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59">1603600082542000-12</a></div>
<div>
NZD/JPY&nbsp;
売
&nbsp;
区分：
指定決済
</div>
<div>
注Lot:10&nbsp;
執行:
逆指
(74.300)
</div>
<div>
ﾄﾘｶﾞｰ価格:---&nbsp;
</div>
<a href="../otc/C003.html?order_id=1603600082542000&amp;order_method=12&amp;page_index=1&amp;symbol_code=&amp;open_close_type=&amp;f_y=2015&amp;f_m=02&amp;f_d=01&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;prev_page_flag=1">変更</a>&nbsp;&nbsp;
<a href="../otc/CT01.html?order_id=1603600082542000&amp;order_method=12&amp;page_index=1&amp;symbol_code=&amp;open_close_type=&amp;f_y=2015&amp;f_m=02&amp;f_d=01&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;prev_page_flag=1">取消</a>
<div><br /></div>
<div><a href="../otc/C403.html?order_id=1603600082436100&amp;order_method=11&amp;page_index=1&amp;symbol_code=&amp;open_close_type=&amp;f_y=2015&amp;f_m=02&amp;f_d=01&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59">1603600082436100-11</a></div>
<div>
NZD/JPY&nbsp;
売
&nbsp;
区分：
指定決済
</div>
<div>
注Lot:100&nbsp;
執行:
指値
(82.200)
</div>
<div>
ﾄﾘｶﾞｰ価格:---&nbsp;
</div>
<a href="../otc/C003.html?order_id=1603600082436100&amp;order_method=11&amp;page_index=1&amp;symbol_code=&amp;open_close_type=&amp;f_y=2015&amp;f_m=02&amp;f_d=01&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;prev_page_flag=1">変更</a>&nbsp;&nbsp;
<a href="../otc/CT01.html?order_id=1603600082436100&amp;order_method=11&amp;page_index=1&amp;symbol_code=&amp;open_close_type=&amp;f_y=2015&amp;f_m=02&amp;f_d=01&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;prev_page_flag=1">取消</a>
<div><br /></div>
<div><a href="../otc/C403.html?order_id=1603600082436100&amp;order_method=12&amp;page_index=1&amp;symbol_code=&amp;open_close_type=&amp;f_y=2015&amp;f_m=02&amp;f_d=01&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59">1603600082436100-12</a></div>
<div>
NZD/JPY&nbsp;
売
&nbsp;
区分：
指定決済
</div>
<div>
注Lot:100&nbsp;
執行:
逆指
(74.400)
</div>
<div>
ﾄﾘｶﾞｰ価格:---&nbsp;
</div>
<a href="../otc/C003.html?order_id=1603600082436100&amp;order_method=12&amp;page_index=1&amp;symbol_code=&amp;open_close_type=&amp;f_y=2015&amp;f_m=02&amp;f_d=01&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;prev_page_flag=1">変更</a>&nbsp;&nbsp;
<a href="../otc/CT01.html?order_id=1603600082436100&amp;order_method=12&amp;page_index=1&amp;symbol_code=&amp;open_close_type=&amp;f_y=2015&amp;f_m=02&amp;f_d=01&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59&amp;prev_page_flag=1">取消</a>
<div><br /></div>

<hr />
<div>
<a href="../otc/C402.html?page_index=2&amp;symbol_code=&amp;open_close_type=&amp;f_y=2015&amp;f_m=02&amp;f_d=01&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=07&amp;t_h=23&amp;t_min=59" accesskey="#">[#]次へ</a>
</div>
<div><a href="C401.html">注文情報(検索)</a></div>

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
	orders, nextPage, err := parseOrderListPage(doc)
	if err != nil {
		t.Fatalf("got %v, want nil", err)
	}
	if !nextPage {
		t.Fatalf("nextPage got false, want true")
	}
	expected := []Order{
		Order{OrderId: "1603600082705700", OrderMethod: "11"},
		Order{OrderId: "1603600082705700", OrderMethod: "12"},
		Order{OrderId: "1603600082672700", OrderMethod: "11"},
		Order{OrderId: "1603600082672700", OrderMethod: "12"},
		Order{OrderId: "1603600082617000", OrderMethod: "11"},
		Order{OrderId: "1603600082617000", OrderMethod: "12"},
		Order{OrderId: "1603600082542000", OrderMethod: "11"},
		Order{OrderId: "1603600082542000", OrderMethod: "12"},
		Order{OrderId: "1603600082436100", OrderMethod: "11"},
		Order{OrderId: "1603600082436100", OrderMethod: "12"},
	}
	if !reflect.DeepEqual(expected, orders) {
		for k, p := range orders {
			t.Errorf("orders[%#v]=%#v", k, p)
		}
		t.Fatalf("wrong orders: %#v", orders)
	}
}

func TestParseOrderDetailsPage_OcoSettlement(t *testing.T) {
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
<div>&gt;注文情報(詳細)&lt;</div>
<hr />


<div>注文受付番号：</div>
<div align="right">1603600082705700</div>
<div>通貨ﾍﾟｱ：</div>
<div align="right">NZD/JPY</div>
<div>注文状況：</div>
<div align="right">
注文中
</div>
<div>注文手法：</div>
<div align="right">
<a href="../common/I123.html?order_id=1603600082705700&amp;order_method=12&amp;page_index=1&amp;symbol_code=&amp;open_close_type=&amp;f_y=2015&amp;f_m=02&amp;f_d=02&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=08&amp;t_h=23&amp;t_min=59&amp;prev_page_flag=1">OCO2</a>
</div>
<div>両建：</div>
<div align="right">
なし
</div>
<div>決済順序：</div>
<div align="right">
---
</div>
<input type="hidden" name="close_seq" value="0">
<div>決済ｵﾌﾟｼｮﾝ：</div>
<div align="right">
---
</div>
<div>売買：</div>
<div align="right">
売
</div>
<div>注文区分：</div>
<div align="right">
<a href="../common/I124.html?position_id=1603600030431300&amp;order_id=1603600082705700&amp;order_method=12&amp;page_index=1&amp;symbol_code=&amp;open_close_type=&amp;f_y=2015&amp;f_m=02&amp;f_d=02&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=08&amp;t_h=23&amp;t_min=59&amp;prev_page_flag=1">指定決済</a>
</div>
<div>執行条件：</div>
<div align="right">
逆指
</div>
<div>指定ﾚｰﾄ：</div>
<div align="right">74.300</div>
<div>ﾄﾚｰﾙ：</div>
<div align="right">---</div>
<div>注文Lot数：</div>
<div align="right">10</div>
<div>期限：</div>
<div align="right">

GTC
</div>
<div>注文受付日時：</div>
<div align="right">2016/02/05 22:57:41</div>
<div><a href="../otc/C402.html?page_index=1&amp;symbol_code=&amp;open_close_type=&amp;f_y=2015&amp;f_m=02&amp;f_d=02&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=08&amp;t_h=23&amp;t_min=59">戻る</a></div>

<hr /></a>

<a href="../otc/C003.html?order_id=1603600082705700&amp;order_method=12&amp;page_index=1&amp;symbol_code=&amp;open_close_type=&amp;f_y=2015&amp;f_m=02&amp;f_d=02&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=08&amp;t_h=23&amp;t_min=59&amp;prev_page_flag=2">変更</a>
<br />
<a href="../otc/CT01.html?order_id=1603600082705700&amp;order_method=12&amp;page_index=1&amp;symbol_code=&amp;open_close_type=&amp;f_y=2015&amp;f_m=02&amp;f_d=02&amp;f_h=00&amp;f_min=00&amp;t_y=2016&amp;t_m=02&amp;t_d=08&amp;t_h=23&amp;t_min=59&amp;prev_page_flag=2">取消</a>

<hr />
<div><a href="C401.html">注文情報(検索)</a></div>


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
	order, err := parseOrderDetailsPage(doc)
	if err != nil {
		t.Fatalf("got %v, want nil", err)
	}
	expected := Order{
		OrderId:      "1603600082705700",
		OrderMethod:  "12",
		Currency:     "NZD/JPY",
		PositionId:   "1603600030431300",
		IsSettlement: true,
		Side:         "sell",
		IsStop:       true,
		Price:        74.3,
		Amount:       10000,
	}
	if !reflect.DeepEqual(&expected, order) {
		t.Fatalf("wrong order: %#v", order)
	}
}
