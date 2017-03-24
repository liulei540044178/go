<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
    <link rel="stylesheet" href="/static/css/common.css" >
    <link rel="stylesheet" href="/static/lib/bootstrap/bootstrap/css/bootstrap.css" >
    <link rel="stylesheet" href="/static/css/signup.css" >


</head>
<body>

            <div class="i_login_">
            <form>
                <div class="form-group">
                    <label for="exampleInputEmail1">USER NAME</label>
                    <input type="email" class="form-control" id="exampleInputEmail1" placeholder="Email">
                </div>
                <div class="form-group">
                    <label for="exampleInputPassword1">PASSWORD</label>
                    <input type="password" class="form-control" id="exampleInputPassword1" placeholder="Password">
                </div>
                <div class="form-group">
                    <label for="exampleInputPassword2">PASSWORD</label>
                    <input type="password" class="form-control" id="exampleInputPassword2" placeholder="REPassword">
                </div>
                <div class="form-group capture">
                    <a href="" id="exampleInputPassword3"><i class="glyphicon glyphicon-refresh"></i> </a>

                    <img src="./images/header.jpg" width="100" height="34"/>
                    <input type="text" class="form-control i_capture_input_"  placeholder="capture">
                </div>

                <div class="form-group">
                    <label for="check" class="check">请阅读以下条款</label>
                    <div class="checkbox">
                        <label>
                            <input type="checkbox" id="check"> 我已经同意 《个人博客条例》
                        </label>
                    </div>
                    <p class="help-block">Example block-level help text here.</p>
                </div>

                <button type="submit" class="btn btn-default btn_width">Submit</button>
            </form>
        </div>

</body>
<script  src="/static/lib/jquery/jquery.js" ></script>
<script  src="/static/lib/bootstrap/bootstrap/js/bootstrap.js" ></script>
<script src="/static/js/login.js" ></script>

</html>