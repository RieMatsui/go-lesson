package main

import "testing"

func TestDie(t *testing.T) {
    // os.Exit()への関数ポインタをバックアップ
    oldExit := osExit

    // テスト終了後にosExitにバックアップしていた関数ポインタを戻す
    defer func() { osExit = oldExit }()

    // osExit()が実行された場合、終了コードが変数calledに記録される
    var status int
    exit := func(code int) {
        status = code
    }
    osExit = exit

    // テストt対象のメソッド実行および結果確認
    Die()
    if exp := 1; status != exp {
        t.Errorf("Expected exit code: %d, status: %d", exp, status)
    }
}

func TestDeferFunc(t *testing.T){
    // os.Exit()への関数ポインタをバックアップ
    oldExit := osExit

    // テスト終了後にosExitにバックアップしていた関数ポインタを戻す
    defer func() { osExit = oldExit }()

    // osExit()が実行された場合、終了コードが変数calledに記録される
    var status int
    exit := func(code int) {
        status = code
    }
    osExit = exit

    // テストt対象のメソッド実行および結果確認
	DeferFunc()
    if exp := 1; status != exp {
        t.Errorf("Expected exit code: %d, status: %d", exp, status)
    }
}

