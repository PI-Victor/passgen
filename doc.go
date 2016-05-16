// Package passgen is an opinionated random password generator package for go
// applications. Creates a random password that:
// * is not based on a time event
// * has at least 8 characters
// * does not contain a complete word
// * contains at least one character from each of these types:
// - ` ~ ! @ # $ % ^ & * ( )_ - + = { } [ ] \ | : ; " ' < > , . ? /
// - 0, 1, 2, 3, 4, 5, 6, 7, 8, 9
// - abcdefghijklmnopqrstuvwxyz
// - ABCDEFGHIJKLMNOPQRSTUVWXYZ
package passgen
