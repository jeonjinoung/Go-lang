const express = require('express');
const app = express();

app.listen(6400,function(){
    console.log('listening on 6400');
});

 app.get('/menu', function(req, res){
     res.send('메뉴페이지입니다...');
 });

 app.get('/', function(req, res){
     res.sendFile(__dirname + '/blogfollow.html')
 });