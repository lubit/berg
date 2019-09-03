# berg

# 提交job
curl -F 'fileX=@/path/to/fileX' -F 'fileY=@/path/to/fileY'  http://localhost:8080

curl -F 'myFile=@/Users/peilin/Downloads/gaoyue.jpg' http://localhost:8080/addjob


curl -F 'myFile[]=@/Users/peilin/dev/lubit/berg/client/gaoyue.jpg' -F 'myFile[]=@/Users/peilin/dev/lubit/berg/client/ggg.tmp' http://localhost:8080/addjob
