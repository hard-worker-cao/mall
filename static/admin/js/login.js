$(function(){
    app.init();
})
var app={
    init:function(){
        this.getCaptcha()
        this.captchaImgChange()
    },
    //获取验证码逻辑
    getCaptcha:function(){
        //
        $.get("/admin/captcha?t="+Math.random(),function(response){
            console.log(response)
            $("#captchaId").val(response.captchaId)
            $("#captchaImg").attr("src",response.captchaImage)
        })
    },
    //点击验证码更换
    captchaImgChange:function(){
        var that=this;
        $("#captchaImg").click(function(){
            that.getCaptcha()
        })
    }
}