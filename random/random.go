package random

var digitRunes = []rune("0123456789")
var lowerCaseRunes = []rune("abcdefghijklmnopqrstuvwxyz")
var upperCaseRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var lowerAlphanumericalRunes = []rune("abcdefghijklmnopqrstuvwxyz0123456789")
var upperAlphanumericalRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var alphanumericalRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var specialRunes []rune = []rune("!#$%&'()*+,-./:;<=>?@[]^_`{|}~") // 31 characters - without double quote and back slash
var printableRunes []rune = []rune("!#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[]^_`abcdefghijklmnopqrstuvwxyz{|}~") // 93 characters - without double quote and back slash
