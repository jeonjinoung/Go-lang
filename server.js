//Post 요청 = 댓글작성 or 글작성
//http://127.0.0.1:5500/Post.html


const express = require('express');
const app = express();

app.listen(8080,function(){
    console.log('listening on 8080');
});
//누군가가 /pet으로 방문을 하면.. pet관련된 안내문을 띄어주자
app.get('/login', function(req, res){
    res.send('로그인 페이지 입니다..');
});

app.get('/', function(req, res){
    res.sendFile(__dirname + '/Post.html')
});