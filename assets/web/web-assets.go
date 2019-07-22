package web

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets7592fbad3344d795f68ec4e255bb49c840fcc75a = "<!doctype html>\r\n<html lang='ja'>\r\n<head>\r\n    <meta charset='utf-8'>\r\n    <title>Speech recognition</title>\r\n    <script>\r\n        const url = \"ws://\" + window.location.host + \"/ws\";\r\n        const ws = new WebSocket(url);\r\n\r\n        ws.onopen = function () {\r\n            ws.send('chrome');\r\n            run();\r\n            document.getElementById('status').innerHTML = 'Connected';\r\n        };\r\n\r\n        ws.onclose = function () {\r\n            document.getElementById('status').innerHTML = 'Closed';\r\n        };\r\n\r\n        ws.onerror = function (err) {\r\n            document.getElementById('status').innerHTML = 'Error';\r\n            console.error(err);\r\n        };\r\n\r\n        const run = function () {\r\n            let recognition = new webkitSpeechRecognition();\r\n            recognition.lang = 'ja-JP';\r\n            recognition.onresult = function (event) {\r\n                let result = event.results[event.resultIndex][0].transcript;\r\n                let tmp = document.getElementById('result').innerHTML;\r\n                document.getElementById('result').innerHTML = result + \"<br>\" + tmp;\r\n                ws.send(result);\r\n            };\r\n            recognition.onspeechend = function () {\r\n                recognition.stop();\r\n                run();\r\n            };\r\n            recognition.onerror = function (err) {\r\n                console.error(err);\r\n                recognition.stop();\r\n                run();\r\n            };\r\n            recognition.start();\r\n        };\r\n    </script>\r\n</head>\r\n<body>\r\n<div id=\"status\">Connecting...</div>\r\n<div id=\"result\"></div>\r\n</body>\r\n</html>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": {"web"}, "/web": {"index.html"}}, map[string]*assets.File{
	"/": {
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1563781435, 1563781435364515000),
		Data:     nil,
	}, "/web": {
		Path:     "/web",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1563781504, 1563781504797408300),
		Data:     nil,
	}, "/web/index.html": {
		Path:     "/web/index.html",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1563781504, 1563781504786417400),
		Data:     []byte(_Assets7592fbad3344d795f68ec4e255bb49c840fcc75a),
	}}, "")
