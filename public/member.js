

(function($) {
    'use strict';

   $(function() {
        var testItem = $('.table-striped'); // 리스트 html
       // var contentname = $('.contentname')        
       // var productname = $('.productname')        
       // var contentmain = $('.contentmain')        
       // var price = $('.price')         
       // var registered = $('.registered')
       
    //    $("#chooseFile").change(function(e) { // 등록 페이지에서 이미지 파일을 선택 했을 때 미리보기하는 함수 불필요하면 제거해주세요.
    //     var input = e.target;
    //     var reader = new FileReader();
    //     reader.readAsDataURL(input.files[0]);
    //     reader.onload = function(){
    //         file = reader.result;
    //         console.log(file)
    //         };
    //         });
      

    //    $('.btn-upload').on("click", function(event) {
    //        event.preventDefault();
           
    //        var contentname = $('.contentname').val();
    //        var productname = $('.productname').val();
    //        var contentmain = $('.contentmain').val();
    //        var file
    //        var fileCheck = document.getElementById("chooseFile").value;
    //        var price = $('.price').val();
    //        var registered = $('.registered').val();

    //        if(contentname == ""){
    //            alert("제목을 입력해주세요")
    //            return false
    //        }else if(productname == ""){
    //            alert("상품명을 입력해주세요")
    //            return false
    //        }else if(contentmain == ""){
    //            alert("내용을 입력해주세요")
    //            return false
    //         }else if(!fileCheck){
    //             alert("이미지를 입력해주세요")
    //             return false
    //        }else if(price == ""){
    //            alert("가격을 입력해주세요")
    //            return false
    //        }else if(registered == ""){
    //            alert("이름을 입력해주세요")
    //            return false
    //        }





    //        if (window.location.href == 'http://localhost:3000/upload.html') {
    //         //    $.post("/upload", {contentname : contentname, productname : productname, contentmain : contentmain, img : img, price : price, registered : registered}, addItem)
    //         // 파일을 전송하기 위해 폼데이터 사용 후 ajax로 전송하는 것으로 변경하였습니다.
    //         const sendingData = new FormData();
    //         sendingData.append('contentname', contentname);
    //         sendingData.append('productname', productname); 
    //         sendingData.append('contentmain', contentmain);
    //         sendingData.append('file',file); 
    //         sendingData.append('price', price);
    //         sendingData.append('registered', registered); 
    //         $.ajax({
    //             url: '/upload',
    //             processData: false,  // 데이터 객체를 문자열로 바꿀지에 대한 값이다. true면 일반문자...
    //             contentType: false,  // 해당 타입을 true로 하면 일반 text로 구분되어 진다.
    //             data: sendingData,  //위에서 선언한 fromdata
    //             type: 'POST',
    //             success: function(result){
    //                 console.log(result);
    //             },
    //             error : function(e){
    //             console.log(e);
    //             }
    //             });
    //            alert("등록완료")
    //            window.location.href='/list.html'
    //            addItem
    //        }
    //    });

       
       if(window.location.href.includes("/list.html")){ // 첫 로그인시 작동이 안되는 경우가 있어서 .includes()를 사용해보았습니다 이상하면 원래 방식으로 해주십시오.
           var addItem = function(item) {
               if (item) {
                   testItem.append("<tr><th scope='row' class='completed'" + " idx='" + item.idx + "'>"+item.sessioncode+"</th><td><a href=detail.html?idx="+item.idx+" class='contentname'>"+item.contentname+"</a></td><td>"+item.registered+"</td><td><img height='60' width='60' class='img-responsive' src='"+item.file+"'></img></td><td>"+item.createtime+"</td></tr>");
               } //else {
            //        testItem.append("<tr><th scope='row' class='completed'" + " idx='" + item.idx + "'>"+item.idx+"</th><td><a href=detail.html?idx="+item.idx+">"+item.contentname+"</a></td><td>"+item.registered+"</td><td>"+item.createtime+"</td></tr>");
            //    }
           };
           $.get('/list', function(items) {
               items.forEach(e => {
                   addItem(e)
               });
           });
       };


       function getParameterByName(name) {
            name = name.replace(/[\[]/, "\\[").replace(/[\]]/, "\\]"); 
            var regex = new RegExp("[\\?&]" + name + "=([^&#]*)"), results = regex.exec(location.search); 
            return results == null ? "" : decodeURIComponent(results[1].replace(/\+/g, " ")); 
        }

        var idx = getParameterByName('idx');




        if(window.location.href.includes('/detail.html')){
        var detailitem = function(item){
            if (item.idx == idx){
                $('input[name=contentname]').attr('value',item.contentname);
                $('input[name=productname]').attr('value',item.productname);
                $('textarea[name=contentmain]').val(item.contentmain);
                $('img[name=file]').attr('src', item.file);
                $('input[name=price]').attr('value',item.price);
                $('input[name=registered]').attr('value',item.registered);
            }
        };
    
        $.ajax({
            url: "detail/" + idx,
            type: "GET", 
            success: function(item) {
               detailitem(item)
            }
        })

        };

        $('#btn-delete').on("click", function(event) {
            event.preventDefault();
            //var id = $(this).closest("th").attr('id');
            var $self = $(this);
            $.ajax({
                    url: "detail/" + idx,
                    type: "DELETE", 
                    success: function(data) {
                        if(data.success) {
                            $self.parent().remove();
                        }
                    }
                })
            alert("삭제되었습니다.")
            window.location.href='/list.html'
        });

        $('#btn-addcart').on("click", function(event) { // 장바구니 담기 기능
            event.preventDefault();
            //var id = $(this).closest("th").attr('id');
            var quantity = $('#quantity').val(); // 수량 가져오기
            const cartData = new FormData();
            cartData.append('idx', idx);
            cartData.append('quantity', quantity);
            $.ajax({
                    url: "/add_cart",
                    processData: false,  // 데이터 객체를 문자열로 바꿀지에 대한 값이다. true면 일반문자...
                    contentType: false,  // 해당 타입을 true로 하면 일반 text로 구분되어 진다.
                    data: cartData,  //위에서 선언한 fromdata
                    type: "POST", 
                    success: function(result){
                        console.log(result);
                    },
                    error : function(e){
                    console.log(e);
                    }
                    });
            alert("장바구니에 추가되었습니다.")
        });


            






    // var idx = a
    //     alert(idx)

    // function getParameterByName(name) {
    //     name = name.replace(/[\[]/, "\\[").replace(/[\]]/, "\\]");
    //     var regex = new RegExp("[\\?&]" + name + "=([^&#]*)"),
    //         results = regex.exec(location.search);
    //     return results === null ? "" : decodeURIComponent(results[1].replace(/\+/g, " "));
    // }

    // var idx = getParameterByName('idx');

    //if(idx){
        // alert(idx);
        // alert("테스트");
        // const sqlite3 = require('sqlite3').verbose();

        // // open database in memory
        // let db = new sqlite3.Database(':memory:', (err) => {
        //   if (err) {
        //     return console.error(err.message);
        //   }
        //   console.log('Connected to the in-memory SQlite database.');
        // });
        
        // // close the database connection
        // db.close((err) => {
        //   if (err) {
        //     return console.error(err.message);
        //   }
        //   console.log('Close the database connection.');
        // });
        
        // const sqlite3 = require('sqlite3').verbose();

        // // let db = new sqlite3.Database('/Users/user/go/src/goWeb/Web100/test.db', sqlite3.OPEN_READWRITE, (err) => {
        // let db = new sqlite3.Database('../test.db', sqlite3.OPEN_READWRITE, (err) => {
        //     alert("DB커넥트 테스트")
        //     if (err) {
        //       alert(err.message);
        //     }
        //     console.log('Connected to the chinook database.');
        //   });

        // db.serialize(() => {
        // db.each(`SELECT * FROM members`, (err, row) => {
        //     if (err) {
        //     alert(err.message);
        //     }
        //     alert(row.idx + "\t" + row.contentmain);
        // });
        // });
        
        // db.close((err) => {
        // if (err) {
        //     alert(err.message);
        // }
        // alert('Close the database connection.');
        // });



        //var db = new sqlite3.Database('../test.db')
        //var db = sql.Open("sqlite3", '../test.db')

        // db.each("SELECT * FROM members", function (err, row) {
        //     alert(row.idx + ": " + row.contentmain);
        // });
        // db.close();
        

        // let db = new sqlite3.Database('../test.db',sqlite3.OPEN_READWRITE,(err) =>{
        //     if(err){
        //         console.error(err.message);
        //         alert(err);
        //     }
        //     alert("Connected to the database.");
        //     console.log('Connected to the database.');
        // });

        // db.serialize(()=>{
        //     db.each('SELECT * FROM members',(err,row) => {
        //         if (err){
        //             console.error(err.message);
        //         }
        //         console.log(row.idx + "\t" + row.contentname);
        //     });
        // });
        // db.close((err) => {
        //     if (err){
        //         console.error(err.message);
        //     }
        //     console.log('Close the database connection.')
        // })

        // const sqlite3 = require('sqlite3').verbose();
        // if(testing){
        //     alert("sqlite3테스트")
        // }

        // // open the database
        // let db = new sqlite3.Database('../test.db');
    
        // let sql = `SELECT * FROM members
        //         WHERE idx = `+idx;

        // alert(sql)
    
        // db.all(sql, [], (err, rows) => {
        // if (err) {
        //     throw err;
        // }
        // rows.forEach((row) => {
        //     console.log(row);
        // });
        // });
    
        // // close the database connection
        // db.close();

    //}




    //    var detailitem = function(item){
           
    //        $('input[name=contentname]').attr('value',item.contentname);
    //        $('input[name=productname]').attr('value',item.productname);
    //        $('textarea[name=contentmain]').val(item.contentmain);
    //        $('input[name=price]').attr('value',item.price);
    //        $('input[name=registered]').attr('value',item.registered);
    //    };

    //    $.get('/list', function(items) {
    //     items.forEach(e => {
    //         detailitem(e)
    //     });
    // });
       

    //    $.get('/detail?idx='+idx, function(items) {
    //        alert("디테일테스트")
    //        items.forEach(e => {
    //            detailitem(e)
    //        });
    //    });





       // memberItem.on('click', '.remove', function() {
       //     // url: todos/id method: DELETE
       //     var id = $(this).closest("li").attr('id');
       //     var $self = $(this);
       //     $.ajax({
       //             url: "members/" + id,
       //             type: "DELETE", 
       //             success: function(data) {
       //                 if(data.success) {
       //                     $self.parent().remove();
       //                 }
       //             }
       //         })
       //     //$(this).parent().remove();
       // });

   });
})(jQuery);