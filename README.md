# xstr #

[![构建状态](https://travis-ci.org/g-lib/xstr.svg?branch=master)](https://travis-ci.com/g-lib/xstr)
[![GoDoc](https://godoc.org/github.com/g-lib/xstr?status.svg)](https://godoc.org/github.com/g-lib/xstr)
[![Go报告](https://goreportcard.com/badge/github.com/g-lib/xstr)](https://goreportcard.com/report/github.com/g-lib/xstr)
[![覆盖状态](https://coveralls.io/repos/github/g-lib/xstr/badge.svg?branch=master)](https://coveralls.io/github/g-lib/xstr?branch=master)

[xstr](https://godoc.org/github.com/g-lib/xstr) 是一个字符串函数集合, 这些函数在其他语言中被广泛使用但是在[strings](http://golang.org/pkg/strings)中却没有.




## 安装 ##

使用 `go get` to 安装这个库

    go get -u github.com/g-lib/xstr

## API 文档 ##

前往 [GoDoc](https://godoc.org/github.com/g-lib/xstr) 查看完整文档.

## 函数列表 ##


| 函数                                                                         | Friends                                                                             | #   |
| ---------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | --- |
| [Center](https://godoc.org/github.com/g-lib/xstr#Center)                     | `str.center` in Python; `String#center` in Ruby                                     |
| [Count](https://godoc.org/github.com/g-lib/xstr#Count)                       | `String#count` in Ruby                                                              |
| [Delete](https://godoc.org/github.com/g-lib/xstr#Delete)                     | `String#delete` in Ruby                                                             |
| [ExpandTabs](https://godoc.org/github.com/g-lib/xstr#ExpandTabs)             | `str.expandtabs` in Python                                                          |
| [FirstRuneToLower](https://godoc.org/github.com/g-lib/xstr#FirstRuneToLower) | `lcfirst` in PHP or Perl                                                            |
| [FirstRuneToUpper](https://godoc.org/github.com/g-lib/xstr#FirstRuneToUpper) | `String#capitalize` in Ruby; `ucfirst` in PHP or Perl                               |
| [Insert](https://godoc.org/github.com/g-lib/xstr#Insert)                     | `String#insert` in Ruby                                                             |
| [LastPartition](https://godoc.org/github.com/g-lib/xstr#LastPartition)       | `str.rpartition` in Python; `String#rpartition` in Ruby                             |
| [LeftJustify](https://godoc.org/github.com/g-lib/xstr#LeftJustify)           | `str.ljust` in Python; `String#ljust` in Ruby                                       |
| [Len](https://godoc.org/github.com/g-lib/xstr#Len)                           | `mb_strlen` in PHP                                                                  |
| [Partition](https://godoc.org/github.com/g-lib/xstr#Partition)               | `str.partition` in Python; `String#partition` in Ruby                               |
| [Reverse](https://godoc.org/github.com/g-lib/xstr#Reverse)                   | `String#reverse` in Ruby; `strrev` in PHP; `reverse` in Perl                        |
| [RightJustify](https://godoc.org/github.com/g-lib/xstr#RightJustify)         | `str.rjust` in Python; `String#rjust` in Ruby                                       |
| [RuneWidth](https://godoc.org/github.com/g-lib/xstr#RuneWidth)               | -                                                                                   |
| [Scrub](https://godoc.org/github.com/g-lib/xstr#Scrub)                       | `String#scrub` in Ruby                                                              |
| [Shuffle](https://godoc.org/github.com/g-lib/xstr#Shuffle)                   | `str_shuffle` in PHP                                                                |
| [ShuffleSource](https://godoc.org/github.com/g-lib/xstr#ShuffleSource)       | `str_shuffle` in PHP                                                                |
| [Slice](https://godoc.org/github.com/g-lib/xstr#Slice)                       | `mb_substr` in PHP                                                                  |
| [Squeeze](https://godoc.org/github.com/g-lib/xstr#Squeeze)                   | `String#squeeze` in Ruby                                                            |
| [Successor](https://godoc.org/github.com/g-lib/xstr#Successor)               | `String#succ` or `String#next` in Ruby                                              |
| [SwapCase](https://godoc.org/github.com/g-lib/xstr#SwapCase)                 | `str.swapcase` in Python; `String#swapcase` in Ruby                                 |
| [ToCamelCase](https://godoc.org/github.com/g-lib/xstr#ToCamelCase)           | `String#camelize` in RoR                                                            |
| [ToKebab](https://godoc.org/github.com/g-lib/xstr#ToKebabCase)               | -                                                                                   |
| [ToSnakeCase](https://godoc.org/github.com/g-lib/xstr#ToSnakeCase)           | `String#underscore` in RoR                                                          |
| [Translate](https://godoc.org/github.com/g-lib/xstr#Translate)               | `str.translate` in Python; `String#tr` in Ruby; `strtr` in PHP; `tr///` in Perl     |
| [Width](https://godoc.org/github.com/g-lib/xstr#Width)                       | `mb_strwidth` in PHP                                                                |
| [WordCount](https://godoc.org/github.com/g-lib/xstr#WordCount)               | `str_word_count` in PHP                                                             |
| [WordSplit](https://godoc.org/github.com/g-lib/xstr#WordSplit)               | -                                                                                   |
| [Contains](http://golang.org/pkg/strings/#Contains)                          | `String#include?` in Ruby                                                           |
| [ContainsAny](http://golang.org/pkg/strings/#ContainsAny)                    | -                                                                                   |
| [ContainsRune](http://golang.org/pkg/strings/#ContainsRune)                  | -                                                                                   |
| [Count](http://golang.org/pkg/strings/#Count)                                | `str.count` in Python; `substr_count` in PHP                                        |
| [EqualFold](http://golang.org/pkg/strings/#EqualFold)                        | `stricmp` in PHP; `String#casecmp` in Ruby                                          |
| [Fields](http://golang.org/pkg/strings/#Fields)                              | `str.split` in Python; `split` in Perl; `String#split` in Ruby                      |
| [FieldsFunc](http://golang.org/pkg/strings/#FieldsFunc)                      | -                                                                                   |
| [HasPrefix](http://golang.org/pkg/strings/#HasPrefix)                        | `str.startswith` in Python; `String#start_with?` in Ruby                            |
| [HasSuffix](http://golang.org/pkg/strings/#HasSuffix)                        | `str.endswith` in Python; `String#end_with?` in Ruby                                |
| [Index](http://golang.org/pkg/strings/#Index)                                | `str.index` in Python; `String#index` in Ruby; `strpos` in PHP; `index` in Perl     |
| [IndexAny](http://golang.org/pkg/strings/#IndexAny)                          | -                                                                                   |
| [IndexByte](http://golang.org/pkg/strings/#IndexByte)                        | -                                                                                   |
| [IndexFunc](http://golang.org/pkg/strings/#IndexFunc)                        | -                                                                                   |
| [IndexRune](http://golang.org/pkg/strings/#IndexRune)                        | -                                                                                   |
| [Join](http://golang.org/pkg/strings/#Join)                                  | `str.join` in Python; `Array#join` in Ruby; `implode` in PHP; `join` in Perl        |
| [LastIndex](http://golang.org/pkg/strings/#LastIndex)                        | `str.rindex` in Python; `String#rindex`; `strrpos` in PHP; `rindex` in Perl         |
| [LastIndexAny](http://golang.org/pkg/strings/#LastIndexAny)                  | -                                                                                   |
| [LastIndexFunc](http://golang.org/pkg/strings/#LastIndexFunc)                | -                                                                                   |
| [Map](http://golang.org/pkg/strings/#Map)                                    | `String#each_codepoint` in Ruby                                                     |
| [Repeat](http://golang.org/pkg/strings/#Repeat)                              | operator `*` in Python and Ruby; `str_repeat` in PHP                                |
| [Replace](http://golang.org/pkg/strings/#Replace)                            | `str.replace` in Python; `String#sub` in Ruby; `str_replace` in PHP                 |
| [Split](http://golang.org/pkg/strings/#Split)                                | `str.split` in Python; `String#split` in Ruby; `explode` in PHP; `split` in Perl    |
| [SplitAfter](http://golang.org/pkg/strings/#SplitAfter)                      | -                                                                                   |
| [SplitAfterN](http://golang.org/pkg/strings/#SplitAfterN)                    | -                                                                                   |
| [SplitN](http://golang.org/pkg/strings/#SplitN)                              | `str.split` in Python; `String#split` in Ruby; `explode` in PHP; `split` in Perl    |
| [Title](http://golang.org/pkg/strings/#Title)                                | `str.title` in Python                                                               |
| [ToLower](http://golang.org/pkg/strings/#ToLower)                            | `str.lower` in Python; `String#downcase` in Ruby; `strtolower` in PHP; `lc` in Perl |
| [ToLowerSpecial](http://golang.org/pkg/strings/#ToLowerSpecial)              | -                                                                                   |
| [ToTitle](http://golang.org/pkg/strings/#ToTitle)                            | -                                                                                   |
| [ToTitleSpecial](http://golang.org/pkg/strings/#ToTitleSpecial)              | -                                                                                   |
| [ToUpper](http://golang.org/pkg/strings/#ToUpper)                            | `str.upper` in Python; `String#upcase` in Ruby; `strtoupper` in PHP; `uc` in Perl   |
| [ToUpperSpecial](http://golang.org/pkg/strings/#ToUpperSpecial)              | -                                                                                   |
| [Trim](http://golang.org/pkg/strings/#Trim)                                  | `str.strip` in Python; `String#strip` in Ruby; `trim` in PHP                        |
| [TrimFunc](http://golang.org/pkg/strings/#TrimFunc)                          | -                                                                                   |
| [TrimLeft](http://golang.org/pkg/strings/#TrimLeft)                          | `str.lstrip` in Python; `String#lstrip` in Ruby; `ltrim` in PHP                     |
| [TrimLeftFunc](http://golang.org/pkg/strings/#TrimLeftFunc)                  | -                                                                                   |
| [TrimPrefix](http://golang.org/pkg/strings/#TrimPrefix)                      | -                                                                                   |
| [TrimRight](http://golang.org/pkg/strings/#TrimRight)                        | `str.rstrip` in Python; `String#rstrip` in Ruby; `rtrim` in PHP                     |
| [TrimRightFunc](http://golang.org/pkg/strings/#TrimRightFunc)                | -                                                                                   |
| [TrimSpace](http://golang.org/pkg/strings/#TrimSpace)                        | `str.strip` in Python; `String#strip` in Ruby; `trim` in PHP                        |
| [TrimSuffix](http://golang.org/pkg/strings/#TrimSuffix)                      | `String#chomp` in Ruby; `chomp` in Perl                                             |

## 协议 ##

本库遵循[MIT协议](./LICENSE)

## 鸣谢 ##

原始代码源自 [xstrings](https://github.com/g-lib/xstrings)