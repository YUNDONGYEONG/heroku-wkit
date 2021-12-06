

(function($) {
    'use strict';

   $(function() {
        var file
       // var contentname = $('.contentname')        
       // var productname = $('.productname')        
       // var contentmain = $('.contentmain')        
       // var price = $('.price')         
       // var registered = $('.registered')
       
       $("#chooseFile").change(function(e) { // 파일 정보 불러오는 함수
        var input = e.target;
        var reader = new FileReader();
        reader.readAsDataURL(input.files[0]);
        reader.onload = function(){
            file = reader.result;
            };
            });
      

       $('.btn-upload').on("click", function(event) {
           event.preventDefault();
           
           var contentname = $('.contentname').val();
           var productname = $('.productname').val();
           var contentmain = $('.contentmain').val();
           var fileCheck = document.getElementById("chooseFile").value; // 파일이 입력되어있는지 확인하는 변수
           var price = $('.price').val();
           var registered = $('.registered').val();

           if(contentname == ""){
               alert("제목을 입력해주세요")
               return false
           }else if(productname == ""){
               alert("상품명을 입력해주세요")
               return false
           }else if(contentmain == ""){
               alert("내용을 입력해주세요")
               return false
            }else if(!fileCheck){
                alert("이미지를 입력해주세요")
                return false
           }else if(price == ""){
               alert("가격을 입력해주세요")
               return false
           }else if(registered == ""){
               alert("이름을 입력해주세요")
               return false
           }






            //    $.post("/upload", {contentname : contentname, productname : productname, contentmain : contentmain, img : img, price : price, registered : registered}, addItem)
            // 파일을 전송하기 위해 폼데이터 사용 후 ajax로 전송하는 것으로 변경하였습니다.
            
            const sendingData = new FormData();
            sendingData.append('contentname', contentname);
            sendingData.append('productname', productname); 
            sendingData.append('contentmain', contentmain);
            sendingData.append('file', file); 
            sendingData.append('price', price);
            sendingData.append('registered', registered); 
            $.ajax({
                url: '/upload',
                processData: false,  // 데이터 객체를 문자열로 바꿀지에 대한 값이다. true면 일반문자...
                contentType: false,  // 해당 타입을 true로 하면 일반 text로 구분되어 진다.
                data: sendingData,  //위에서 선언한 fromdata
                type: 'POST',
                success: function(result){
                    console.log(result);
                },
                error : function(e){
                console.log(e);
                }
                });
               alert("등록완료")
               window.location.href='/list.html'
               addItem
           
       });
   });
})(jQuery);