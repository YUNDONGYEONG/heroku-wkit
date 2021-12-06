
// 필터 검색
// $(document).ready(function () {
//     $("#keyword").keyup(function () {
//         var selectOption = document.getElementById("search");
//         selectOption = selectOption.options[selectOption.selectedIndex].value;
//         //alert(selectOption)
//         var k = $(this).val();
//         if(k == "") {
//             $("#user-table > tbody > tr").show();
//         }
//         $("#user-table > tbody > tr").hide();
//         var temp = $("#user-table > tbody > tr > td:nth-of-type(.selectOption):contains('" + k + "')");

//         $(temp).parent().show();

//     })

// })


// 기본 상품명 검색
$(document).ready(function () {
    $("#keyword").keyup(function () {
        var selectOption = document.getElementById("search");
        selectOption =  Number(selectOption.options[selectOption.selectedIndex].value);
        //alert(selectOption)
        var k = $(this).val();
        $("#user-table > tbody > tr").hide();
        var temp = $("#user-table > tbody > tr > td:nth-child(5n+"+selectOption+"):contains('" + k + "')");

        $(temp).parent().show();
    })
})