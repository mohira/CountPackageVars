package CountPackageVars

import (
	"fmt"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/analysis"
)

const doc = "CountPackageVars is ..."

var Analyzer = &analysis.Analyzer{
	Name: "CountPackageVars",
	Doc:  doc,
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	var count int
	for _, file := range pass.Files {
		// ASTの確認: ASTをみればパッケージ変数がどう扱われているか検討がつく
		// ast.Print(pass.Fset, file)

		// ファイル名
		// fmt.Println(file.Name) // a

		for _, decl := range file.Decls {
			// 各Declのチェック(VARもFUNCもある)
			// fmt.Printf("%T %v\n", decl, decl)

			genDecl, ok := decl.(*ast.GenDecl)
			if !ok {
				continue // 早期Returnな感じでフィルタリングするほうが読みやすい
			}

			// token.Tokenをチェック
			if genDecl.Tok != token.VAR {
				continue
			}

			for _, spec := range genDecl.Specs {
				// 詳細を表すノードに変換(ast.Specと実装している型)
				s, ok := spec.(*ast.ValueSpec)
				if !ok {
					continue
				}
				count += len(s.Names)
				fmt.Printf("%v\n", s.Names)
				// [name]
				// [age]
				// [address]
				// [_]
				// [_]
				// [x y]
				// [a b c]
				for _, name := range s.Names {
					// _ は パッケージ変数とはみなさないものとしておく
					if name.Name != "_" {
						count++
					}
				}
			}
		}

		fmt.Printf("パッケージ変数の数は: %d\n", count)
	}
	return nil, nil
}
