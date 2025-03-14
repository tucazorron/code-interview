# google

## problem

there is a bucket storage file explorer

for each folder, you can have many folders and many files

there are two api calls to use:

    storage.list(uri)
    returns a list of files and folders inside that folder
    if it returns [], it means its an empty folder
    if it returns nil, it means you've reached a file

    storage.size(uri)
    returns the size of a file
    return -1 if its not a file

return the top 10 largest files in size from that bucket in descending order

## input

string (root folder)

## output

10 largest files with their size and adress in descending order

## tests

    bucket.storage.google.com

    1000 bucket.storage.google.com/videos/video1.mp4
    900 bucket.storage.google.com/videos/video2.mp4
    797 bucket.storage.google.com/videos/video3.mp4
    777 bucket.storage.google.com/videos/video4.mp4
    600 bucket.storage.google.com/images/img1.jpg
    500 bucket.storage.google.com/images/img2.jpg
    100 bucket.storage.google.com/images/img3.jpg
    80 bucket.storage.google.com/images/img4.jpg
    9 bucket.storage.google.com/code/code1.go
    1 bucket.storage.google.com/code/code2.go
