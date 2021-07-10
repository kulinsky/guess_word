package testdata

import "github.com/kulinsky/guess_word/domain"

var FWID = domain.GenerateWordID()
var SWID = domain.GenerateWordID()
var FirstWord = domain.NewWord(FWID, "hello")
var SecondWord = domain.NewWord(SWID, "world")
