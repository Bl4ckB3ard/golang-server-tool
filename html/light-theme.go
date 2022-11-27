package html

var LightThemeTemplate string

func init() {
	LightThemeTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Poppins">
    <title>GO Server</title>
    <style>
        html {
            background-color: whitesmoke;
            padding: 0;
            font-family: Poppins;
            font-style: normal;
            font-weight: 300;
        }

        table {
            width: 100%;
            height: 100%;
            border-collapse: collapse;
        }

        /*top of table border*/
        tr.header-row {
            border-top: 1px solid gray;
            border-bottom: 2px solid black;
        }

        /*Make bottom row td left align and fix*/
        tr.last-row td {
            text-align: left;
            padding-top: 7px;
            padding-left: 0;
        }
        
        /* line up table header */
        tr.header-row th {
            text-align: left;
        }

        /*fix padding and add bottom border*/
        tr {
            padding: 8px;
            border-bottom: 1px solid black;
        }
        td {
            padding: 5px;
        }

        /*aligh View link and download link*/
        td.view, td.download {
            text-align: right;
        }

        /*stop color change after click link*/
        a:visited {
            color: blue;
        }

        /*change a tag hover color*/
        a:hover {
            color: purple;
        }

        /*bottom of table border*/
        tr.last-row {
            border-top: 2px solid black;
            border-bottom: none;
        }

		a.no-color:hover, a.no-color {
			color: inherit;
			text-decoration: none;
		}
    </style>
</head>
<body>
    <h1><a href="/" class="no-color">Driectory: {{.RootDirectory}}</a></h1>
    <br>
    <table>
        <tr class="header-row">
            <th></th>
            <th>Name</th>
            <th>Size</th>
            <th></th>
        </tr>

	{{ range $i, $val := .Items }}
	    <tr>
	    {{ if $val.IsDir }}
	        <td><img src="https://i1.wp.com/whateverbrightthings.com/wp-content/uploads/2016/10/1-5.png?fit=1400%2C1400" alt="folder" width="25" height="25"></td>
	        <td><a href="{{ $val.RelativePath }}">{{ $val.Name }}</a></td>
	    {{ else }}
	        <td><img src="https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fcdn.onlinewebfonts.com%2Fsvg%2Fimg_522255.png&f=1&nofb=1&ipt=5ffa15719fafc7071f4feb4769e8da85d0ba6571bdc2a27cda4875e04325b36f&ipo=images" alt="file" width="20" height="25"></td>
	        <td>{{ $val.Name }}</td>
	    {{ end }}
	        <td>{{ $val.Size }}</td>
	    {{ if $val.IsDir  }}
	        <td class="view"></td>
	        <td class="download"></td>
	    {{ else }}
	        <td class="view"><a href="{{ $val.ViewHref }}">View In Browser</a></td>
	        <td class="download"><a href="{{ $val.DownloadHref }}">Download</a></td>
	    {{ end }}
	    </tr>
	{{ end }}
	    <tr class="last-row">
	        <td>{{ .LenOfDirectory }}</td>
	    </tr>
	</table>
	</body>
	</html>`
}
