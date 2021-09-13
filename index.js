// window.addEventListener('load',()=>{

// })

//이벤트 연결함수
process.on('exit',function(code){

console.log('빠이');

});

//예외가 발생할때 실행되는 이벤트
process.on('uncaughtExeption',function(error){

    console.log('예외가 발생해따!!!');
});

let count = 0;
let test = function(){
    count = count +1;
    if (count>3) return;
        
    error.error.error();
}
setTimeout(test,2000);