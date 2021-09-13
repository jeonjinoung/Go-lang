
//http 메소드
// http의 요청매소드를 사용하게 된다.
//form get 요청매소드 foot 가져온다..
//GET            /index.html
//GET            /users
//POST            /user//PUT            /user/nickname (해당사용자의 닉네임)
//DELETE         /user/id

const http = require('http');

const fs = require('fs').promises;

http.createServer(async(request,response)=>{

    try{
        if(request.method==='GET'){
            if(request.url==='/'){
                const data = await fs.readFile('/index.html');
            }
            else if(request.url==='/users'){

                            }
        }   
    }catch(error){

    }
})
