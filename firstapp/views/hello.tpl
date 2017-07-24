<!DOCTYPE html>
<html lang='en'>
<head>
    <meta charset='UTF-8'>
    <title> hello - {{.title}}</title>
</head>
<body>
    {{range $idx, $val := .images}}
    <h1>{{$val.Id}}={{$val.ImageName}}</h1>
    {{end}}
    <footer> this is footer </footer>
</body>
</html>
