# Golang

## 환경 설정(Mac)

```
brew install go  //설치
export GOPATH=$HOME/go //Path 설정
export PATH=$PATH:$GOPATH/bin
```

환경 설정을 실패히여 코드 실행을 못하는 경우 귀찮으면 아래 온라인 컴파일러를 쓰거나  
오류를 찾아보고 해결하면 될거 같다  

https://www.programiz.com/golang/online-compiler/
---
## 프로젝트 설정

```
mkdir <프로젝트 이름>
touch <파일 이름>
go mod init <module-name>
go run main.go //파일 실행
```
---
## 목차
- [Hello](Hello/Readme.md)
- [변수](Var/Readme.md)
- [조건](If/Readme.md)
- [프로그래머스](Pro/Readme.md)
- [백준](Baejoon/Readme.md)


git push -u origin main 
