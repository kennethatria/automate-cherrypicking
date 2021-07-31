AUTOMATE EVERYTHING with Go

Problem statment:  While working with cherry picking on binary files, I have issues finding the oldest commit and those that cause conflicts

This script helps find oldest commit and prints out possible conflicts incase of any

to find if binary files return errors
    go run main.go conflicts bac5849 logs.txt
to return commit position
    go run main.go postion bac5849 logs.txt

How to Test: 
  1. run "git log --oneline >> logs.txt"
  2. provide log file as 3rd argument
  2. run Script with " go run main.go position commit-id log.txt"

additional feature
  logging script actions

