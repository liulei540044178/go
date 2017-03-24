/**
 *
 * Created by Administrator on 2017/3/18 0018.
 */

//需要注意：    此处调试了很久
//             将插件放在 自定义js后面
(function () {
    // alert(1);
    $("#save").bind("click",function(e){
        //当点击保存之后，取消可点击事件
        $(this).attr("disabled","disabled")
        //获取editor编辑内容
        // alert(CKEDITOR.instances.editor1.getData())
        var content = CKEDITOR.instances.editor1.getData()
        //获取title的值
        // alert($("input[name='noteTitle']").val())
        var title = $("input[name='noteTitle']").val()
        // alert(title + "==== ")
        $.post("http://127.0.0.1:8080/note/addnote",{noteTitle:title,noteContent:content},
            function (data) {
            var obj = eval('(' + data + ')');
            console.log(typeof  obj.state);

            if(obj.state == "succeed"){
                console.log(obj.operation+"   true");
                //移除disabled属性，恢复可点击状态`
                $("#save").removeAttr("disabled");

            }else{
                obj.operation = "失败的操作 =.="
            }
            showAlert(obj.operation);
            // console.log(data);
        });

    });
    bindItems()
})()


function showAlert(msg) {
    var i_alert = $(".i_alert_")[0];
    $(i_alert.getElementsByTagName("span")[0]).text(msg);
    $(i_alert).show()
    // alert($(i_alert.getElementsByTagName("span")[0]).text(msg))
    //不需要双引号
    window.setTimeout(hideAlert,2000);
}

function hideAlert() {
    $(".i_alert_").hide();
}

function  bindItems() {
    $(".i_file_list_ .i_list_ol_ li a").each(function (index,ele) {
        $(ele).bind("click",function () {
           var url =  $(this)[0].href
           $.get(url,function (data) {
               var obj = eval('(' + data + ')');
               // obj = data.parseJSON()
               console.log(obj.title)
               $("input[name='noteTitle']").val(obj.title)
               CKEDITOR.instances.editor1.setData(obj.content)
           })
        })
    })
}