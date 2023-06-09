# 인증

## [ 인증서 저장하기 ]

 - https://git-scm.com/book/ko/v2/Git-%EB%8F%84%EA%B5%AC-Credential-%EC%A0%80%EC%9E%A5%EC%86%8C

```
git config --global credential.helper cache # 15분까지 유지
git config --global credential.helper store # 디스크에 인증서 정보 저장
cat ~/.git-credentials                      # 인증서 확인하기
```

# Commit

## [ Editor 변경하기 ]

```
git config --global core.editor "vim"
```

## [ Author 변경하기 ]

```
git commit --amend --reset-author
```
```
git commit --amend --author="Author Name <email@example.com>"
```

## [ 여러 Commit들 변경하기 - rebase ]

1. perfrom rebase

> ```
> git rebase -i <commit-position>
> ```

2. change `pick` --> `edit` 

>```
> edit 5ga39z1
>```

3. add files that have been changed to stage area

>```
>git add sample.code
>```

4. reflect changed status

>```
>git commit --amend
>git rebase --continu
>```


# Pull

## [ 원격 브랜치에 대하여 로컬 브랜치 생성하기 ]

```
git switch origin_branch
git checkout --track origin/origin_branch
```

# README.md

## C4 Diagrams
 - https://docs.github.com/en/get-started/writing-on-github/working-with-advanced-formatting/creating-diagrams
 - https://mermaid.js.org/syntax/c4c.html
