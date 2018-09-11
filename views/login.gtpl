<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>login</title>
</head>
<body>
<form action="/login" method="post">
    <div>
        <label for="">用户名:</label>
        <input type="text" name="username">
    </div>

    <div>
        <label for="">密码:</label>
        <input type="password" name="password">
    </div>

    <div>
        <label for="">你最喜欢的运动:</label>
        <input type="checkbox" name="interest" value="football">足球
        <input type="checkbox" name="interest" value="basketball">篮球
        <input type="checkbox" name="interest" value="tennis">网球
    </div>

    <div>
        <label for="">年龄:</label>
        <input type="number" name="age">
    </div>

    <button type="submit">我要登录</button>
</form>
</body>
</html>