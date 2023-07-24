#!/bin/bash

echo "============================================="
echo "|             build                         |"
echo "============================================="

go build -o tetris

echo "============================================="
echo "|             test audit                    |"
echo "============================================="

echo "badexemple00"
./tetris badexemple00.txt
echo ""


echo "badexemple01"
./tetris badexemple01.txt 
echo "" 


echo "badexemple02"
./tetris badexemple02.txt 
echo "" 


echo "badexemple03"
./tetris badexemple03.txt 
echo "" 


echo "badexemple04"
./tetris badexemple04.txt 
echo "" 


echo "badformat"
./tetris badformat.txt 
echo "" 


echo "goodexemple00"
./tetris goodexample00.txt 
echo "" 


echo "goodexemple01"
./tetris goodexample01.txt 
echo "" 


echo "goodexemple02"
./tetris goodexample02.txt 
echo "" 


echo "goodexemple03"
./tetris goodexample03.txt 
echo "" 


echo "hardexam"
./tetris hardexam.txt 
echo "" 


echo "============================================="
echo "|             clear exe                     |"
echo "============================================="

rm tetris