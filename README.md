

# 프로젝트팀명:HeavySmoker 


## 👜👢🧢('HeavySmoker') - 특정 동호회 소셜 커뮤니티
## 프로젝트 시작일 :2021-11-05

## 👜👢🧢[**서비스로 이동**](http://54.180.155.42:6002/)



p.s. 현재 서버를 닫은 관계로 실행이 되지 않습니다

--  
<img width="700" alt="스크린샷 2021-10-15 오후 3 51 02" src="https://i.esdrop.com/d/cl3pewp2aooj/G6lV8sxt2F.jpg">

---   

Member. 



이동녘:[깃허브](https://github.com/dongnycklee)  
이성현:[깃허브](https://github.com/coolmarvel)  
전진영:[깃허브](https://github.com/jeonjinoung)  
김현호:[깃허브](https://github.com/alsrhkd101)  



  
  

# 목차
[1.개요](#개요)

[2.목적](#목적)

- 기존 서비스와의 차별점

[3.전체 소스 코드](#전체-소스-코드-click)

[4.사용한 기술](#사용한-기술)

[5.주요 기능](#주요-기능)

[6.발생한 이슈 & 해결 방법](#발생한-이슈--해결-방법)

[7.상세 설명](#상세-설명)

 - DB 구조 (ERD)

 - 전체 흐름도
 
 - 프로젝트 설명 PPT

[8.시연 영상](#시연-영상)

[9.참여인원](#참여-인원-4명)

***

### 개요

특정 커뮤니티 활성화를 위한 동호회SNS 구성

### 목적

>  특정 커뮤니티 활성화를 위한 동호회SNS 구성
> 
> 1. 커뮤니티 활성화
> 2. 프로젝트 경험
> 3. React 구현 숙달
> 

- **기존 서비스와의 차별점**

   - 소켓 I.O 를 활용한 실시간 채팅
   - 커뮤니티 서비스 제공
   

### 전체 소스 코드 [소스](https://github.com/3eteam/3eteamproject)


### 사용한 기술

- 웹 화면 구성 : `HTML5` `CSS3` `bootstrap` `JavaScript``Scss``React`
- 검색 결과 데이터 요청 및 출력 : `jQuery` `Ajax` `JSON``axios`
- DB 액션 처리 : `MongoDB`
- DBMS : `MongoDB`
- 개발 Tool :`PostMan``Visual Studio Code` `Atom`
- AWS 배포 : `EC2` `MongoDB`
- 로그인 Api : `KakaoTalk API``Google API`
- 채팅 Api : `Gitpull`
- 프레임워크 : `NodeJs`
- 프로젝트 관리 Tool : `Google Drive` `GitHub`
- 사용 모듈(v) : ` "axios": "^0.22.0",
                "bcrypt": "^5.0.1",
                "bootpay-js": "^3.3.3",
                "cookie-parser": "^1.4.5",
                "dotenv": "^10.0.0",
                "express": "^4.17.1",
                "express-session": "^1.17.2",
                "moment": "^2.29.1",
                "morgan": "^1.10.0",
                "multer": "^1.4.3",
                "mysql": "^2.18.1",
                "mysql2": "^2.3.0",
                "nodemon": "^2.0.13",
                "nunjucks": "^3.2.3",
                "passport": "^0.5.0",
                "passport-kakao": "^1.0.1",
                "passport-local": "^1.0.0",
                "sequelize": "^6.6.5",
                "sequelize-cli": "^6.2.0",
                "socket.io": "^4.2.0"`
### 주요 기능

- 로그인 : 일반 로그인, `MongoDB``passport``Compass`
- 회원가입 : `MongoDB` `passport``Compass`
- 회원정보변경 :`MongoDB` `passport``Compass`
- 마이페이지 : `MongoDB```mysql`
- 실시간채팅:`socket.io``rg.rok`





- Socket.io 의 ngrock 서버문제

[상황] ngrock의 문제

[문제] ngrock 이 최대 3명만가능해서 채팅기능이 3명이상 채팅이불가능했던 문제

[해결] ec2서버를열때 rgrock 을 사용하지않고 소켓을 자체적으로 서버를 열어서 인크루드방식으로 바꿈으로서 대량의인원도 문제없도록 해결함


### "NodeJs ,Socket.IO 를 통해 실시간 통신을 구현"

- 실시간 채팅 구현

[상황] 사용자들끼리 실시간 쇼핑관련 소통을하기위한 채팅

[문제] socket.io서버 를 코드에 추가할 떄 오류 발생

[해결] 서버의 순서도를 재배치 함으로써 문제 해결

- socket.id와 idx

[상황] 특정 사용자에게 메세지를 보내려면 해당 사용자의 socket.id를 알아야 함

[문제] socket.id는 브라우저를 킬 때마다 새롭게 생성됨, 고유한 key가 되지 못함

[해결] 해당 회원을 식별할 수 있는 idx 값을 활용:  idx값을 속성이름으로 추가, 그 속성의 값은 소켓의 ID값으로 할당함

- connected

[상황] 'ooo님이 채팅방에 입장/퇴장 하셨습니다' 기능을 추가해야함

[문제] 시간 부족으로 인해 미완성

[해결]  나중에 보안 할 예정

**NodeJs socket.io 깨달은 점 **

실시간 사용자들끼리의  채팅이기에 socket.id를 사용하는 1:1 채팅을 구현했음. socket.io 특성 상 사용자 지정에 대한 이해도 숙지  


데이터베이스 AWS EC2방식 체택 인스턴스 생성후 키값생성 후
서버 연결
<img width="400" alt="41" src="https://user-images.githubusercontent.com/89692626/137617959-d52e400d-d98f-4f6d-a035-948354a394ad.png">
<img width="400" alt="42" src="https://user-images.githubusercontent.com/89692626/137617960-c55d109a-dad6-4357-9c5f-87e25edac9dd.png">
<img width="400" alt="43" src="https://user-images.githubusercontent.com/89692626/137617961-14e02f5a-e202-41a7-9035-75bd4fd9eadd.png">
<img width="400" alt="44" src="https://user-images.githubusercontent.com/89692626/137617962-37868599-b3c1-4225-b012-1d4a30439032.png">


데이터베이스 AWS EC2방식 체택 인스턴스 생성후 키값생성 후
서버 연결

AWS의 RDS로 데이터베이스를 공유하여 협업하여 프로젝트를 진행하였고
AWS의 EC2를 활용하여 배포를 하였습니다

서버 연결하기 전 SCREEN을 활용 cmd창을 닫거나 컴퓨터를 종료하여도 서버가 계속 돌아가게 유지했습니다













////////////////////////////////////////////////////////////////////////////////////////////////////////

**여기서 부터 시작했습니다.!!**
---------------------------------

일단 서버단 코드구성은 아래 이미지와 같다.

<img width="400" alt="44" src="https://user-images.githubusercontent.com/89626182/143200197-1ee58e04-893b-4486-a10b-315c1ecc17c8.png">

아래의 기능에 공통적으로 설명이 들어가는것은 Controllers models routes server가 들어갈것이다.
Controllers

Controllers -> auth.js = login 부분

<img width="400" alt="44" src="https://user-images.githubusercontent.com/89626182/143200795-4107aa05-a85d-4e47-9e7c-efa0fe27729a.png">

models

models -> user.js

<img width="400" alt="44" src="https://user-images.githubusercontent.com/89626182/143201022-1e4bb97a-abc0-40f2-bcb3-95054eba831c.png">
                               
routes
     
routes -> authRouter.js부분

<img width="400" alt="44" src="https://user-images.githubusercontent.com/89626182/143201141-ad0ab66b-800f-4068-8efc-37e7e3a98dcd.png">

server

<img width="400" alt="44" src="https://user-images.githubusercontent.com/89626182/143201601-2f8f2f61-d13d-4bba-a37e-661dffbf32eb.png">

이렇게 서버 및 백단을 구성을 해놓고 client에서 각 라우터를 api형식으로 불어와준다.
일단 로그인과 로그아웃 회원정보변경등의 값을 불러오려면 client의 프론트단을 구별을 먼저해줘야한다.                     

/////////////////////////////////////////////////////////////////////////////////////////////////////////////

**Client front**

로그인 페이지 [Login]

<img width="400" alt="44" src="https://user-images.githubusercontent.com/89626182/143202567-a90e13e9-e44e-4d4c-bc2c-312c1199f281.png">

Client -> pages -> Login.jsx
                                                                                                                 
<img width="400" alt="44" src="https://user-images.githubusercontent.com/89626182/143202961-08cae5ad-431f-4cf0-92ad-aa2dda25fafe.png">
◈로그인 페이지 입니다. 텍스트 안에 값을 넣어주고 그값이 맞으면 홈페이지 화면으로 넘어가도록 하였습니다.


회원가입 페이지 [Register]

<img width="400" alt="44" src="https://user-images.githubusercontent.com/89626182/143237166-98f384d2-dd8b-4a95-a1eb-7e3d9ee7d208.png">

<img width="400" alt="44" src="https://user-images.githubusercontent.com/89626182/143237642-2ba87dc4-37c2-4d30-9101-bcc190d8f684.png">

◈회원가입 페이지 입니다. 텍스트 안에 값을 넣어주고 그값이 맞는지 유효성검사를 통해 조건문을 달아주었고 그값이 맞다면 다음 홈페이지 화면으로 넘어가도록 하였습니다.
                         
로그아웃 [Logout]

<img width="400" alt="44" src="https://user-images.githubusercontent.com/89626182/143204698-deb3c7eb-6e5c-41cc-97d2-51b8febd0b0a.png">
                                                                                                             Client -> component -> header -> Menu
                                                                                                                 
<img width="400" alt="44" src="https://user-images.githubusercontent.com/89626182/143204991-1a45ca39-834e-48ba-9d18-5f60751152f4.png">

◈로그아웃 기능 입니다. 메뉴에있는 로그아웃을 클릭을하면 onclick이벤트로인해 버튼을 눌렀을시 아이디값을 추적하고 그값을 공란 빈칸으로 만들어주고 그리고 페이지는 홈페이지인 로그인페이지로 넘어가도록 해두었다.


회원정보변경 [EditProfileEdit]
                                                                                                                
client -> component -> profile -> info[Editprofile]

<img width="400" alt="44" src="https://user-images.githubusercontent.com/89626182/143238277-fba1a632-e714-4823-bcea-ddf1cb22bc69.png">

<img width="400" alt="44" src="https://user-images.githubusercontent.com/89626182/143238414-352f34a1-7759-463a-aaad-83360e50368c.png">

<img width="400" alt="44" src="https://user-images.githubusercontent.com/89626182/143238564-904bcf57-5957-422c-aa15-a54209779aac.png">

<img width="400" alt="44" src="https://user-images.githubusercontent.com/89626182/143238632-a8ed4caf-27ab-4c11-8605-83a6ce915761.png">

<img width="400" alt="44" src="https://user-images.githubusercontent.com/89626182/143223226-96fff0fa-5061-49b4-8991-ef44a3101943.png">

◈프로필수정 기능 입니다. 메뉴에있는 프로필 클릭을하면 onclick이벤트로인해 버튼을 눌렀을시 아이디값을 추적하고 아이디값에 맞는 프로필의 페이지로 이동한다 이렇게 되면 info페이지로 넘어오게된다 거기에서 EditProfile 계정수정버튼을 누르면
아이지 계정의 값이 수정되는것을 알 수 있다.


알람기능

<img width="400" alt="44" src="https://user-images.githubusercontent.com/89626182/143239565-bc4abf2d-fd79-44cc-86b7-dec43ec10361.png">

<img width="400" alt="44" src="https://user-images.githubusercontent.com/89626182/143239761-876fa4b5-d2c5-4558-89fe-057b3843ab67.png">

Client -> components -> Menu -> nav-item dropdown

<img width="400" alt="44" src="https://user-images.githubusercontent.com/89626182/143240572-15609069-9e8a-4355-a6a1-1efe6d36567d.png">

◈알림 기능 입니다. 헤더메뉴에있는 알림 클릭을하면 onclick이벤트로인해 버튼을 눌렀을시 auth값을 추적하고 auth값에 맞는 프로필의 알림이 서치된다. 그러면 위의 내용과 같이 누가 나의 포스트의 댓글을 달았는지 누가 메세지를 하였는지 누가 나를 팔로우했는지에 대해서 알림을 확인할 수 있다.


포스트와 스테이터스 [posts, status]

포스트는 내가공유한글과 상대방이공유한그를 보는 부분이고

스테이터스는 내가 공유한부분을 보여주는 부분입니다.



Client -> Component -> Status.jsx 공유하는부분!

<img width="400" alt="44" src="https://user-images.githubusercontent.com/89626182/143227661-47537eeb-599b-4db0-9afd-acda9e2fc20e.png">

<img width="400" alt="44" src="https://user-images.githubusercontent.com/89626182/143256848-56c0a06f-2afc-4ae5-8386-875f6eadec4e.png">

공유한 부분을 보여주는 부분 [Post.jsx]

Client -> Component -> posts.jsx -> CardBody, CardFooter, CardHeader로 나눔

<img width="400" alt="44" src="https://user-images.githubusercontent.com/89626182/143227849-692ba1ac-9cc2-47d8-ae01-83937c5d926e.png">

<img width="400" alt="44" src="https://user-images.githubusercontent.com/89626182/143256740-04e2b21b-e79f-4564-a61f-ca9c69c3dac7.png">

◈ Status 기능 입니다. 헤더메뉴에있는 공유하는 input을 클릭하면 onclick이벤트로인해 버튼을 눌렀을시 auth값을 추적하고 status 텍스트에 맞는 이미지와 그리고 글을 값을 저장해주고 그값을 posts에 넘긴다. 그리고 posts.jsx에서 그값을 받고 postcard header,body,footer부분으로 부위를 나뉘어 표현해주었다.이로써 우리는 여러사람의 생각과 글을 볼수도 있었고 나의 글과 사진을 공유하기도 가능했다.

팔로우

client -> component -> RightSideBar[보여지는페이지]

RightSideBar에서 클릭하면 redux/action에서의 함수정의한것을 Follow버튼 컴포넌트를 실행시켜주어서 UserCard의 정보를 컴포넌트를 가져와서 그값을 저장해준다. 그렇게되면 서로 팔로우 했는지 안했는지 알 수 있다.

FollowBtn.jsx
<img width="400" alt="44" src="https://user-images.githubusercontent.com/89626182/143259066-442ec6c6-2e18-4971-ac65-e4d2cd799a0d.png">

UserCard.jsx
<img width="400" alt="44" src="https://user-images.githubusercontent.com/89626182/143259169-17cde38d-85fb-400a-94c6-0e038586be3f.png">

follow 전

<img width="400" alt="44" src="https://user-images.githubusercontent.com/89626182/143228000-48ae2891-351c-4113-847f-c1083e9dbc0a.png">
                                                                                                     follow 후

<img width="400" alt="44" src="https://user-images.githubusercontent.com/89626182/143228089-34e993a6-5969-4a85-a68f-efc277a879e6.png">
<img width="400" alt="44" src="https://user-images.githubusercontent.com/89626182/143228141-52b06736-029c-4e81-a4e6-e0a53b3cc6c2.png">


**여기까지가 제가 만진부분입니다.**
/////////////////////////////////////////////////////////////////////////////////////////////////////////















# 상세 설명
전체 코드 : [코드](https://github.com/dongnycklee/heavysmoker/tree/main/main)

    - 폴더 설명
    Client : 전반적인 모든 클라이언트 리액트 컴포넌트
      public : favicon.ico  , robots.txt 등 
      src : 컴포넌트,이미지스등 동적,정적인 모든 프론트영역
        auth : user db값이 들어가는 모든 PUT,GET,POST 요소들
        core : 최상위 컴포넌트 home , menu 컴포넌트
        css : front 영역
        images : 이미지 영역
        loading : 모든 랜더링 요소에 로딩바 불린값으로 트루 펄스로 넣음
        post : post css 요소
        user : user요소의 모든 컴포넌트의 
    
    

서버(app.js) : [코드]  
(https://github.com/dongnycklee/heavysmoker/blob/main/main/client/src/App.jsx)  
-> 각 페이지별 라우터 설정, passport 6002번, express 모듈로 넌적스를 http로 연결

### 페이지별 설명(views-html)   
## main :   
<img width="500" alt="KakaoTalk_20211017_155220801" src="https://user-images.githubusercontent.com/89692626/137615265-9ad9ddb3-b1d6-4db4-bdb6-f2149972f03c.png">
<img width="500" alt="KakaoTalk_20211017_155218347" src="https://user-images.githubusercontent.com/89692626/137615273-95e3f65a-1164-441f-ac89-3f3542a194f2.png">   

메인페이지입니다 풀스크린방식을 채용했으며 2페이지로 나눠서 제작했습니다   
1페이지같은경우는 무한슬라이드방식으로 이미지가 우에서 좌로 흘러가는형태로 제작했으며   
백그라운드 HOLIDAY 세일부분은 백그라운드이미지를 Z-INDEX를 줘서 해상도에따른 변화를 잡아줬습니다   

<img width="500" alt="KakaoTalk_20211017_155218347" src="https://user-images.githubusercontent.com/89692626/137615345-4e4a92e5-1c7b-4268-b977-5571b5c23497.png">

넌적스의 extends 형식으로 블록형태로 잡아둔후 작업한 페이지입니다  
저희가 리엑트를 배우기전 마지막으로쓸거같은 템플릿엔진이기에 최대한 많은부분을 활용해보고싶었습니다 페이지이동식은  
extends로 블록을잡아서 작업하였고  

<img width="500" alt="KakaoTalk_20211017_155218347" src="https://user-images.githubusercontent.com/89692626/137615352-e36c765b-5200-4258-8306-83da5bf6c4c5.png">   

include 방식입니다 해당html파일의 소스의 요소를 영역에 집어넣어 작업하였습니다  

<img width="172" alt="KakaoTalk_20211017_155727405" src="https://user-images.githubusercontent.com/89692626/137615422-0ed93abb-4870-485c-9070-72e82848ea4d.png">   
Socket.io 를 사용하여 접속자끼리의 실시간 채팅을만들었습니다 마찬가지로 include 방식을 사용하였습니다 

<img width="361" alt="3" src="https://user-images.githubusercontent.com/89692626/137615827-c3b77f28-3216-4a4b-9163-93e9c337a1a6.png">
<img width="500" alt="" src="https://user-images.githubusercontent.com/89692626/137615821-08abdf2e-32d1-421a-b38e-1b0539e8535d.png">

관리자와의 채팅이가능한 라이브러리입니다 "깃플" 이며 상담요청시 해당 상담사가 1:1매칭되어 채팅기록,db에서 사용자를 불러서  
어떤사용자가 시간표시와함께 어떠한 문의를 남겼는지 따로 저장이가능하며 어떠한상담사가 상담했는지도 표시가됩니다  
해당 라이브러리에는 상담요청이왔을때 카톡으로 알림이오는것도있어서 매우 편리하다 판단되어 사용해봤습니다  
   
## join :    
<img width="1133" alt="KakaoTalk_20211017_155846026" src="https://user-images.githubusercontent.com/89692626/137615487-14824b09-a5c0-4588-aeb9-65cf9e11a997.png">   
회원가입 페이지입니다 scss와 gif방식을사용하여 쇼케이스형식으로 왼쪽에는 이미지(GIF)와 음악이 재생이됩니다  
음악같은경우 크롬정책상 autoplay방식이 불가능한데  
오픈소스인 howl 을 사용하여 사용자가 헤더에있는 "회원가입" 버튼을 클릭했을시 사용자가 음악을 재생하는걸 동의하는걸로 간주하도록 제작하여  
회원가입페이지에 들어왔을때 자동으로 음악이 재생되도록 설정하였고 왼쪽하단엔 음악을 끄고 키는 기능을 넣었습니다  
<img width="500" alt="KakaoTalk_20211017_160205029" src="https://user-images.githubusercontent.com/89692626/137615680-b3adccde-f3f7-4179-bb99-24159d6414b0.png">회원가입하는 과정에서 유저의 정보를 담는다. 가입시 입력하는 email값이 unique로 중복이 안되게 하였고 비밀번호는 12자리의 해쉬값으로 저장하고 이미지는 변환된 url값으로 저장한다.
<img width="500" alt="KakaoTalk_20211017_160215296" src="https://user-images.githubusercontent.com/89692626/137615638-d9c36e30-2816-46b4-b8e1-084fc5a800f6.png">
회원가입하는 유저들의 profile img들을 담아 저장할 폴더를 만든다. multer를 통해 img를 profile img에 저장할 수 있게 하고 img 파일은 url 값으로 변환한다.
<img width="500" alt="KakaoTalk_20211017_160202804" src="https://user-images.githubusercontent.com/89692626/137615615-6990152d-b8ae-4fe8-b5f2-2176442aa8e1.png">    
login에는 passport에 넣어준 로그인 권한을 유지할 수 있게 하는 방법을 local과 kakao로 저장한다.
    
## login :  
<img width="500" alt="KakaoTalk_20211017_160529161" src="https://user-images.githubusercontent.com/89692626/137615748-7cd43437-9eaf-4d75-9412-ac5f205757fe.png">   
<img width="393" alt="1" src="https://user-images.githubusercontent.com/89692626/137615766-aa64fd48-cf60-4750-91ab-ad5a16be19e9.png">   
serializeUser와 deserailizeUser를 통해 로그인할 수 있게 해준다.
<img width="602" alt="2" src="https://user-images.githubusercontent.com/89692626/137615784-ec660a00-550b-460c-8a6c-e60582821ad0.png">  
이메일과 비밀번호값을 이용해 user의 정보와 맞는지 비교한다. 비밀번호가 일치하는지 알아보는 과정에서는 bcrypt 모듈을 사용했다.
  
## catlist(관리자 모드) :    
<img width="500" alt="11" src="https://user-images.githubusercontent.com/89692626/137615880-456ebc39-d110-4f02-9205-49ff64616984.png">  
<img width="500" alt="22" src="https://user-images.githubusercontent.com/89692626/137615883-06bfd619-cab6-4ac8-9962-fed0ca6b7b5e.png">  
<img width="500" alt="4" src="https://user-images.githubusercontent.com/89692626/137615937-b948f740-7fa6-4984-854e-3ceecaeda524.png">   
<img width="500" alt="5" src="https://user-images.githubusercontent.com/89692626/137615991-ee481480-c48d-4688-b3ab-402393c8d307.png">  
<img width="500" alt="6" src="https://user-images.githubusercontent.com/89692626/137615998-5f48cb9a-28ea-4ce7-a7df-04bae5e03e36.png">  
<img width="500" alt="7" src="https://user-images.githubusercontent.com/89692626/137616060-072cef9a-be20-4cd1-9a7c-986eb67b0af3.png">  
<img width="500" alt="8" src="https://user-images.githubusercontent.com/89692626/137616064-b5da190d-3dc7-4094-983e-b37d0deb1b16.png">  

관리자 모드에 post.js에 multer 모듈을 사용하여 이미지업로드를 uploads폴더에 자동 저장하며 url만 데이터베이스에 삽입

get에서 불러온 데이터베이스 값을 twits에 저장하여 하단부분에 렌더링

특정 id값만 불러와 전체상품 데이터베이스(Post)의 해당 특정 id값만 삭제

관리자만 관리자모드에접근하여사용할수있게했습니다
관리자모드에서 저장하게 되면 모든 물품이 하단에 위치하여생성 삭제가능하며
삭제할시 전체 물품 데이터베이스에서 데이터삭제


## nike :   
<img width="400" alt="9" src="https://user-images.githubusercontent.com/89692626/137616135-b6a95861-90cc-4617-8ee6-9fff2eaa6648.png">   
<img width="400" alt="10" src="https://user-images.githubusercontent.com/89692626/137616142-e4a68889-f17b-45a3-bd52-292862b73263.png">  
  
  
 넌적스 
  if문 및 for문을 사용하여 전체상품중 nike에 해당하는 컬럼값을가져와 
for문 사용하여 물품정렬

form 태그 사용 물품의 가격 이름 브랜드 등 
장바구니데이터베이스로 정보전달
사용하여
  
  
  
  
## alllist :   
<img width="400" alt="12" src="https://user-images.githubusercontent.com/89692626/137616175-df939fbd-b875-4f22-9397-02e1a6c74f71.png">  
<img width="400" alt="13" src="https://user-images.githubusercontent.com/89692626/137616199-51fcfcb7-9c24-41d2-9bbb-c7aa68749d06.png">  
<img width="400" alt="14" src="https://user-images.githubusercontent.com/89692626/137616235-8f035ea5-9872-436d-979f-72e7dbe269a0.png">  
  
  
  전체상품 페이지 라우터이며 관리자페이지에서 업로드한 전체물품을 확인할 수 있으며,get으로 전체물품을 불러온 후 정렬을 하였으며  전체물품을 확인가능하며 장바구니도 옆에구현하여 전체물품을 보면서 삭제할수 있게 기능을 구현하였음
  
## search :   
<img width="400" alt="15" src="https://user-images.githubusercontent.com/89692626/137616275-b2377a89-77a0-4d6a-bcf4-9c39fcc8980d.png">  

전체상품에서 입력한 태그값을 입력시 검색이 되게해놧으며 검색 후 물품 장바구니에 추가 가능
  
## cart :   
<img width="400" alt="16" src="https://user-images.githubusercontent.com/89692626/137616301-2a3ad31a-9388-41ca-8c30-43441c59d250.png">  
<img width="400" alt="23" src="https://user-images.githubusercontent.com/89692626/137616473-132a255d-463f-4262-8b9c-4590b29c0a63.png">  
<img width="400" alt="24" src="https://user-images.githubusercontent.com/89692626/137616469-5729844f-ff74-4501-8e6a-99b3eefc70a0.png">  
<img width="400" alt="25" src="https://user-images.githubusercontent.com/89692626/137616470-5e0e98b7-b71c-469f-97ca-057cbe332b87.png">  
<img width="400" alt="26" src="https://user-images.githubusercontent.com/89692626/137616471-affad1db-6fb4-4826-9051-7a5a6dce756d.png">  
<img width="400" alt="27" src="https://user-images.githubusercontent.com/89692626/137616472-50152136-9c18-4a2d-b19c-2ce2e6ba0c70.png">  
  
  
  브랜드상품페이지에서 상품을 post방식으로 보내게 되면 cart데이터베이스 담겨 userid에 해당하는 물품을 반환하여 랜더링 됨
  또한 삭제시 데이터베이에서 개별삭제되며
  전체삭제시 userid값에 해당하는 장바구니 데이터 전부삭제
  결제시 payments 데이터베이스에 주문정보가 입력되며
  결제 시 해당 유저id 값에 해당하는 장바구니 데이터베이스 데이터 전부 
## profile :  

<img width="400" alt="29" src="https://user-images.githubusercontent.com/89692626/137616555-458eeff7-4a3b-4def-acdc-52df795f3bea.png">  
<img width="400" alt="30" src="https://user-images.githubusercontent.com/89692626/137616556-720fc3a8-46c8-45d8-99c0-f79b3b6c3c3a.png">  
<img width="400" alt="31" src="https://user-images.githubusercontent.com/89692626/137616581-65a3d05e-04f8-4fb0-abb6-8c27912ad5f4.png">  

## board : board, boarddetail, boardwrite, boardedit  
<img width="400" alt="61" src="https://user-images.githubusercontent.com/89692626/137718441-cb96579f-a46a-4201-90aa-ed579aa468f2.png">  
게시판의 첫 페이지로 게시글 번호는 AUTO INCREMENT되는 게시글의 id로 나타나고 작성자에는 해당 게시글을 작성한 유저의 닉네임이 표시된다.  
<img width="400" alt="62" src="https://user-images.githubusercontent.com/89692626/137718449-2d768a39-72bf-4168-912a-4912b46db76f.png">  
글 작성 버튼을 누르면 넘어가는 페이지로 제목과 내용과 사진을 업로드할 수 있고 제출버튼을 누르면 게시글이 DB에 저장되어 board게시판에 보여진다.  
<img width="400" alt="63" src="https://user-images.githubusercontent.com/89692626/137718456-fb72d3c2-3d5a-43f0-b5aa-508ba4a759d7.png">  
방금 작성한 게시글이 닉네임과 함께 등록된 모습이다.  
<img width="400" alt="64" src="https://user-images.githubusercontent.com/89692626/137718457-02edc0c9-5f7f-4db9-a3ad-d6562be1a375.png">  
게시글의 제목을 누르면 해당 글을 더 자세히 볼 수 있는데 조회수는 해당 페이지로 렌더링되기 전에 게시글의 viewcount가 1 올라간다.   
<img width="368" alt="55" src="https://user-images.githubusercontent.com/89692626/137724837-d882c490-206e-4b66-b4b6-3b303fa00707.png">  
자기가 쓴 글에는 유저Nick과 해당 게시글의 Nick이 같기 때문에 수정과 삭제 버튼이 생긴다. 삭제를 하게 되면 DB에서 해당글이 사라진다.  
<img width="400" alt="65" src="https://user-images.githubusercontent.com/89692626/137718460-66a12fda-10f2-46c0-81b3-5f248780e6d4.png">  
수정을 하면 제목과 내용은 value를 가져와서 그대로 뜬 상태에서 글을 수정할 수 있게 되어있고 위와 같이 수정이 된다.  
<img width="400" alt="66" src="https://user-images.githubusercontent.com/89692626/137718461-a7c693cf-884f-454a-a2a1-a1f795c38a25.png">  
수정을 하면 위와 같이 수정된 내용이 게시글에 다시 표시가 된다.   
  

  












 







### DB구조 (ERD)
![KakaoTalk_Photo_2021-10-17-15-22-45](https://user-images.githubusercontent.com/88940298/137614327-8508ccda-6572-4a70-af91-d1f9063c4a60.jpeg)


### 프로젝트 기획안 PPT
[구글 프레젠테이션 파일](https://docs.google.com/presentation/d/1OXVmXrEVzpxmP-WIhmVdoelkYk0yFbY2/edit?usp=sharing&ouid=116791032032435384891&rtpof=true&sd=true)

# 시연 영상
[유튜브첫페이지 전반적인 틀,index 영상](https://www.youtube.com/watch?v=FMPbevKqxzo)    
[유튜브시연영상1](https://www.youtube.com/watch?v=tTzy22FIxdk)    
[유튜브시연영상2](https://www.youtube.com/watch?v=OSPGKf2z3MY)    

