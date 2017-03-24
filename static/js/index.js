/**
 * Created by Administrator on 2017/3/15 0015.
 */
(function(){
    initUI()
    initData()
    $(".i_menu_child_").css("display","none");
})();

function initData()
{
    window.flag = true;
}

//
function initUI()
{
    //var menu = $(".i_menu_child_");


    $(".i_click").each(function(ele,index){
        $(this).bind("click", function () {
            //this.menu = $(" .i_menu_child_");
            var item = $(this.getElementsByTagName("i")[1]);
            var menu = $(this.getElementsByTagName("ol")[0]);
            //alert(menu)
            $(menu).slideToggle()
            if(flag){
                flag = false;
                item.removeClass("glyphicon-chevron-down").addClass("glyphicon-chevron-up");
            }else{
                 flag = true;
                item.removeClass("glyphicon-chevron-up").addClass("glyphicon-chevron-down");
            }


            console.log( flag)
        })
    })
}