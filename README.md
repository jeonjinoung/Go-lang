
////////////////////////////////////////////////////////////////////////////////////////////////////////
**전진영**
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
