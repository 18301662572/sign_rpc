$(function () {
    //判断签到按钮是否可用
    if($("#state").val()=="true"){
        $("#sign").show()
    }else{
        $("#sign").hide()
    }
    //签到
    $("#sign").click(function () {
        $.get("/sign",{uid:$("#uid").val()});
    });
})