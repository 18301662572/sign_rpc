$(function(){
    //注册
    $("#register").click(function(){
        if($("#checkbox").is(':checked')){
            $.post("/register",{username:$("#username").val().trim(),
                password:$("#password").val().trim()
            });
        }
        else {
            alert("请勾选阅读协议！")
            return
        }
    });
})