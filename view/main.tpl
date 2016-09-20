<!DOCTYPE html>
<html>
    <head>
        <title>{{template "title" .}}</title>
        <script src="/js/jquery.min.js"></script>
        <script src="/js/app.js"></script>
        <link rel="stylesheet" href="/css/style.css" />
        <link rel="canonical" href="http://pokemon.qqqq.io/" />
        <link rel="shortcut icon" href="/img/favicon.png?" />
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <meta property="og:url" content="http://pokemon.qqqq.io/" />
        <meta property="og:type" content="website" />
        <meta property="og:image" content="http://pokemon.qqqq.io/img/pika.png" />
        <meta property="og:title" content="Pokemon Name Generator" />
        <meta property="og:description" content="The nation's premier Pokemon naming service" />
    <head>
    <body>
        <div class="container">
            <h1>{{template "title" .}}</h1>
            {{template "content" .}}
        </div>
    </body>
</html>