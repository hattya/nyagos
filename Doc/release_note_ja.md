* プロンプトが長すぎる時、強制的に改行するようにした (#104)
* ls でワイルドカードがマッチしない時のメッセージを修正 (#108)

NYAGOS 4.1.1\_0
===============

* キー入力で UTF16 のサロゲートペアをサポート
* mkdirに必要に応じて親ディレクトリを作成する /p オプションを追加

NYAGOS 4.1.0\_0
===============

* 内蔵コマンド ln を追加
* Lua コマンド lns を追加 (UACを表示後、`ln -s` を実行する)
* `ls -l` でシンボリックリンクの宛先を表示
* あるファイルでcopy/move 時に失敗した時、以降のファイルを続けるか問合せるようにした。
* 新変数: `nyagos.histchar`: ヒストリ置換文字(デフォルト「`!`」)
    - ヒストリ置換を完全に無効にする場合、`nyagos.histchar = nil`
* 新変数: `nyagos.antihistquot`: ヒストリ置換を抑制する引用符(デフォルト「`'"`」)
    - 【注意】`"!!"` は「デフォルト」では置換されなくなりました
    - 4.0互換にするには `nyagos.antihistquot = [[']]` とする
* 新変数: `nyagos.quotation`: 補完でのデリミタ文字(デフォルト「`"'`」)。
    - `nyagos.quotation` の最初の文字がデフォルトの引用符となる。
    - 二番目以降の文字は、ユーザが補完前に使用していた場合に採用される
    - `nyagos.quotation=[["']]`の場合
        - `C:\Prog[TAB]` → `"C:\Program Files\ ` (`"` が挿入される)
        - `'C:\Prog[TAB]` → `'C:\Program Files\ ` (`'` が維持される)
        - `"C:\Prog[TAB]` → `"C:\Program Files\ ` (`"` が維持される)

NYAGOS 4.1-beta
================
* クラッシュ回避のため、全てのLua のコールバック関数はそれぞれの Lua 
  インスタンスを持つようにした。
* コールバック関数と .nyagos 間で値を共有するため、テーブル share[] を作った
* `*.wsf` を cscript に関連付けた
* `nyagos[]` への不適切な代入を警告するようにした。

<!-- vim:set fenc=utf8: -->
