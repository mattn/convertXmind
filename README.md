# convertXmind
macでアイディア出しにxmindのフリー版を使ってて、それを上司に見せたら

「かっこいいけど、Excelにしてよ」

と言われてしまった。

理由を聞けば、

「縦の列を見れば、概念の階層が合ってるか、
情報の粒度が合ってるかが一目でわかるでしょ」

とのことで、一理ある。

「それならEmacsのorg-modeとかどうでしょう？
　中身は単なるプレーンテキストですが、
枝の開閉やTODOとDONEのトグル、
同じ階層や違う階層のジャンプとか無駄に高機能」

とか（早口で）喋りそうになったけど、
詳しくない人にEmacsオススメしてもろくなことにならんことくらいは
さすがに学んでいるのでやめた。
大人なので。

「だから、今度からこんな感じにExcelで作ってね」

とExcelファイルの例を渡される。
慣れてるだけあってさすが見やすくてわかりやすい。

最初からこれで作るのは、行を足したりとか順番変えたりとかが面倒だが、
最終形としてここに帰着させるのはありだな。

だけどExcelにエクスポートするためにXmindの有料版を買うのも癪だ。
こんなの中身はzipで圧縮されたxmlだろ？

というわけで、goの練習問題としてパパッと作って見たのが、これというわけです。
予想以上に簡単で驚きましたが。

## インストール方法

golangがインストールされている状態で

    go get github.com/tannakaken/convertXmind

とコマンドするだけです。

## 使い方

インストール後

    convertXmind format source dist

とコマンドします。

formatはxlsx（Microdoft Excelファイル）と
org（GNU Emacsのorg-mode用ファイル）が選べます。

sourceに変換したいXMindファイルを、distに変換後の名前をコマンドします。

## 動作環境

動作確認した環境は

* macOS Sierra 10.12.6
* XMind 8 Update 4 (R3.7.4.201709040350)
* go1.9 darwin/amd64

です。

他の環境においてや、ヴァージョンが違うと動かない可能性があります。

まあ、ちょっとした練習問題なので、自分で作るといいのではないでしょうか。
