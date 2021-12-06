(function($) {
    'use strict';

   $(function() {
    var cartItem = $('.table-striped');
    if(window.location.href.includes("/cart.html")){ // 작동이 안되는 경우가 있어서 .includes()를 사용해보았습니다 이상하면 원래 방식으로 해주십시오.
        var addItem = function(item) {
            if (item) {
                cartItem.append("<tr  id='" + item.idx + "'><th scope='row'" + " idx='" + item.idx + "'>"+item.sessioncode+"</th><td><a href=detail.html?idx="+item.idx+">"+item.contentname+"</a></td><td>"+item.registered+"</td><td><img height='60' width='60' class='img-responsive' src='"+item.file+"'></img></td><td>"+item.addtime+"</td><td>"+item.quantity+"</td><td>"+item.quantity*item.price+"</td><td><button class='btn-remove float-right btn btn-primary font-weight-bold' id='btn-remove'>제거</button></td></tr>");
            }
        };
        $.get('/cart', function(items) {
            items.forEach(e => {
                addItem(e)
            });
        });
    };

    cartItem.on("click", '#btn-remove', function(event) {
        event.preventDefault();
        var idx = $(this).closest("tr").attr('id');
        var $self = $(this);
        $.ajax({
                url: "cart/" + idx,
                type: "DELETE", 
                success: function(data) {
                    if(data.success) {
                        $self.parent().parent().remove();
                    }
                }
            })
        alert("제거 되었습니다.")
        window.location.href='/list.html'
    });

    // 2021-11-05, 제거 눌렀을 때 팝업이 뜨는 경우가 있어서 주석처리 하였습니다., 장재혁
    // cartItem.on("click", '#btn-remove', function(){
    //     Popup($(".bill").html());
    // });

    // function Popup(data)
    // {
    // var mywindow = window.open('', 'my div', 'height=400,width=600');
    // mywindow.document.write('<html><head><title>my div</title>');
    // mywindow.document.write('</head><body >');
    // mywindow.document.write(data);
    // mywindow.document.write('</body></html>');
    // mywindow.document.close(); // IE >= 10에 필요
    // mywindow.focus(); // necessary for IE >= 10
    // mywindow.print();
    // mywindow.close();
    // return true;
    // }

});
})(jQuery);