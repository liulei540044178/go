<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
    <link rel="stylesheet" href="/static/css/common.css" >
    <link rel="stylesheet" href="/static/lib/bootstrap/bootstrap/css/bootstrap.css" >
    <link rel="stylesheet" href="/static/css/index.css" >
    <link rel="stylesheet" href="/static/lib/plugin/jquery.mCustomScrollbar.min.css" >


</head>
<body>

    <div class="i_msg_">
        <div class="i_t_logo_">我的博客</div>
        <div class="i_head_">
                <div class="i_head_box_">
                    <img src="/static/images/header.jpg" alt="">
                    <span class="i_user_"><i class="glyphicon glyphicon-user"></i><span>administrator</span></span>
                </div>

                <div class="i_menu_ mCustomScrollbar" data-mcs-theme="dark">
                   <div>
                        <ol>
                            <li>
                                <i class="glyphicon-ban-circle "></i><span>添加目录树</span><i class="glyphicon glyphicon-plus i_float_r_" style="    line-height: 40px;float: right;"></i>
                            </li>
                            <li class="i_menu_left_border_"><i class="glyphicon-ban-circle"></i><span>Dashboard</span></li>
                            <li class="i_click">
                                <i class="glyphicon-ban-circle "></i><span>Dashboard</span><i class="glyphicon glyphicon-chevron-down i_float_r_" style="    line-height: 40px;float: right;"></i>
                                <ol class="i_menu_child_">
                                    <li><a href="">asdasdasd</a></li>
                                    <li><a href="">asdasdasd</a></li>
                                    <li><a href="">asdasdasd</a></li>
                                    <li><a href="">asdasdasd</a></li>
                                    <li><a href="">asdasdasd</a></li>
                                </ol>
                            </li>
                            <li><i class="glyphicon-ban-circle"></i><span>Dashboard</span></li>
                            <li><i class="glyphicon-ban-circle"></i><span>Dashboard</span></li>
                            <li class="i_click">
                                <i class="glyphicon-ban-circle "></i><span>Dashboard</span><i class="glyphicon glyphicon-chevron-down i_float_r_" style="    line-height: 40px;float: right;"></i>
                                <ol class="i_menu_child_">
                                    <li><a href="">asdasdasd</a></li>
                                    <li><a href="">asdasdasd</a></li>
                                    <li><a href="">asdasdasd</a></li>
                                    <li><a href="">asdasdasd</a></li>
                                    <li><a href="">asdasdasd</a></li>
                                </ol>
                            </li>
                        </ol>
                    </div>
                </div>

        </div>

        <div class="i_footer_">
            <ol>
                <li><a href="">Terms</a></li>
                <li><a href="">Terms</a></li>
                <li><a href="">Terms</a></li>
                <li><a href="">Terms</a></li>
            </ol>
            <span>Nadhif  Collect from   - More Templates </span>
        </div>
    </div>


    <div class="i_content_">

        <div class="i_title_header" style="">
            <!--导航-->
            <nav class="navbar navbar-inverse i_nav_ .navbar-static-top">
                <div class="container-fluid">
                    <!-- Brand and toggle get grouped for better mobile display -->
                    <div class="navbar-header">
                        <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1" aria-expanded="false">
                            <span class="sr-only">Toggle navigation</span>
                            <span class="icon-bar"></span>
                            <span class="icon-bar"></span>
                            <span class="icon-bar"></span>
                        </button>
                        <a class="navbar-brand" href="#">Brand</a>
                    </div>

                    <!-- Collect the nav links, forms, and other content for toggling -->
                    <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
                        <ul class="nav navbar-nav">
                            <li class="active"><a href="#">Link <span class="sr-only">(current)</span></a></li>
                            <li><a href="#">Link</a></li>
                            <li class="dropdown">
                                <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Dropdown <span class="caret"></span></a>
                                <ul class="dropdown-menu">
                                    <li><a href="#">Action</a></li>
                                    <li><a href="#">Another action</a></li>
                                    <li><a href="#">Something else here</a></li>
                                    <li role="separator" class="divider"></li>
                                    <li><a href="#">Separated link</a></li>
                                    <li role="separator" class="divider"></li>
                                    <li><a href="#">One more separated link</a></li>
                                </ul>
                            </li>
                        </ul>
                        <form class="navbar-form navbar-left">
                            <div class="form-group">
                                <input type="text" class="form-control" placeholder="Search">
                            </div>

                        </form>
                        <ul class="nav navbar-nav navbar-right">
                            <li><a href="#">Link</a></li>
                            <li class="dropdown">
                                <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">用户 <span class="caret"></span></a>
                                <ul class="dropdown-menu">
                                    <li><a href="#">修改信息</a></li>
                                    <li><a href="#">退出 </a></li>
                                    <li><a href="#">Something else here</a></li>
                                    <li role="separator" class="divider"></li>
                                    <li><a href="#">Separated link</a></li>
                                </ul>
                            </li>
                        </ul>
                    </div><!-- /.navbar-collapse -->
                </div><!-- /.container-fluid -->
            </nav>

        </div>


        <div>
        <div class="i_file_list_ mCustomScrollbar" data-mcs-theme="dark">
                <ol class="i_list_ol_">
                    <li>
                        <a href="" class="i_fresh_">
                            <i class="glyphicon glyphicon-refresh"></i>
                        </a>
                    </li>
                    {{/*
                        当使用range时 在range范围内 使用不了上下文中的数据 .
                        需要重新赋一次值 使用$xx 输出
                    */}}
                    {{ $user :=  .FromName }}
                    {{range $index, $elem := .notes}}
                    <li>
                        <a href="/note/query/{{ $user }}/{{$elem.Id }}" onclick="return false">
                            <i class=""></i><span> {{ substr $elem.ArticelTitle 0 10}}</span>
                            <div>

                                {{ substr $elem.ArticelContent 0 80}}...
                            </div>
                        </a>
                    </li>
                     {{end}}


                    <li>
                    <a href="">
                                                <i class=""></i><span>开源项目</span>
                                                <div>
                                                    asd asd asd asd asd asd asd asd asd asd asd asd asd asd asd asd asd asd asd asd asd asd asd
                                                    asd asd asd asd asd asd asd asd asd asd asd asd as
                                                </div>
                                            </a>
                    </li>

                </ol>
        </div>









        <div class="i_edit_">

         <form action="http://127.0.0.1:8080/note/addnote" method="POST" id="NoteModel">
                <input class="btn" id="save" type="button" value="保存" />
                <input class="btn" type="submit" value="submit" />

                <input class="input" name="noteTitle" type="text"  >
                <textarea name="" id="editor1" style="width:100%;height:100%;">

                </textarea>

                <script>
                     window.onload = function() {

                            CKEDITOR.replace( 'editor1');


                        };
                </script>
        </form>
        </div>
        </div>
    </div>

   <div class="i_alert_" style="display:none">
            <i class="glyphicon glyphicon-ok-circle"></i>
            <span>恭喜插入成功</span>
   </div>


</body>
<script src="/static/lib/jquery/jquery.js" ></script>
<script src="/static/js/editor.js"></script>
<script src="/static/lib/bootstrap/bootstrap/js/bootstrap.js" ></script>
<script src="/static/js/index.js" ></script>
<script src="/static/lib/plugin/jquery.mCustomScrollbar.js"></script>
<script>
    var c = $("#cke_1_contents").css({"height":"500"})
</script>
 <script src="/static/lib/ckeditor/ckeditor.js"></script>
</html>