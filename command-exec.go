package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	args := os.Args
	//一つ目の引数をcmdに
	cmd := args[1]
	//実行パスを取得
	path, err := exec.LookPath(cmd)
	if err != nil {
		log.Fatal("installing %s is in your future", cmd)
	}
	log.Printf("%s is available \n", path)
	//cmd実行(結果を取得)
	/*
		out, err := exec.Command(path).Output()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%s", out)
	*/
	//cmd実行(結果未取得)
	/*
		err1 := exec.Command(path).Run()
		if err1 != nil {
			log.Fatal(err1)
		}
	*/
	//環境変数をコマンド実行に渡す
	// go run command-exec.go printenvで該当の環境変数が表示されてることを確認する
	cmddesc := exec.Command(cmd)
	cmddesc.Env = append(os.Environ(),
		"TEST=go_example",
	)
	out, err := cmddesc.Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", out)
}
