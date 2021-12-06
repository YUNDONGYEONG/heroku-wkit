(function($) {
    'use strict';
   $(function() {
       var total = 0
       var billpage = $('.cover-container'); // 2021-11-05, 영수증 인쇄시 키 값을 넣기 위함, 장재혁
    var cartItem = $('.table');
    if(window.location.href.includes("/")){ // 작동이 안되는 경우가 있어서 .includes()를 사용해보았습니다 이상하면 원래 방식으로 해주십시오.
        var addItem = function(item) {
            if (item) {
                total = total + item.quantity*item.price
                cartItem.append("<tr><th>"+item.sessioncode+"</th><td>"+item.contentname+"</td><td>"+item.quantity+"</td><td>"+item.price+"원</td><td>"+item.quantity*item.price+"원</td></tr>");
            }
        };
        $.get('/cart', function(items) {
            items.forEach(e => {
                addItem(e)
            });
            cartItem.append("<tr><th>합계 : </th><td></td><td></td><td></td><td>"+total+"원</td></tr>");
        });
    };

    var billkey
    var inputemail

    billpage.on('click', '[id="sendkey"]', function() { // 2021-11-05, 영수증 인쇄시 키 값을 넣기 위함, 장재혁
        inputemail = prompt("키를 받을 이메일 주소를 입력해주세요.");
            if(inputemail){
        $.get('/bill', function(key){
            // 메일 전송 함수 필요
                billkey = key
            var templateParams = {	
                //각 요소는 emailJS에서 설정한 템플릿과 동일한 명으로 작성!
                message: billkey,
                email : inputemail
             };
             emailjs.send("service_d20u618","template_f0xni6m",templateParams)
            alert("키를 메일로 전송하였습니다.")
            });
            }
        });
        
    billpage.on('click', '[id="print"]', function() {
        var inputkey = $('#inputkey').val()
        if (inputkey == billkey) {
            alert("키가 맞습니다.")
            var html = $('#print-layer').html();
            printArea(html);
        }
        else if(inputkey ==""){
            alert("키 값을 넣어주세요.")
        }
        else{
            alert("키가 틀렸습니다.")
        };
        billkey = ""     // 2021-11-05, 혹시 모르니 키와 이메일을 초기화, 장재혁
        inputkey = ""
        inputemail = ""
    });
        });
})(jQuery);