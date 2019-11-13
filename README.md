Opensource git go client library https://github.com/src-d/go-git/ is limited in it's merge functionality with code to 
only do fast forward merges. This is test repo in Akshata's development efforts to implement non fast forward git merging 
for go-git.

Plan for proceeding :
1. Figure out how to build and use this library. Add a function that will simply print something. 
2. Write a program to clone one of my repos with go-git. Then make changes and pull those changes with go-git. 
