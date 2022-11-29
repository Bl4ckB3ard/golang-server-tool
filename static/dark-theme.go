package static

var DarkThemeTemplate string

func init() {
	DarkThemeTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Poppins">
    <link href='https://fonts.googleapis.com/css?family=Allerta Stencil' rel='stylesheet'>
    <title>GO Server</title>
    <style>
        html {
            background-color: #181A1B;
            padding: 0;
            font-family: Poppins;
            font-style: normal;
            font-weight: 300;
            color: #E8E6E3;
        }
        
        table {
            width: 100%;
            height: 100%;
            border-collapse: collapse;
        }
                
        /*top of table border*/
        tr.header-row {
            border-top: 1px solid #545B5E;
            border-bottom: 2px solid #8C8273;
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
            border-bottom: 1px solid #545B5E;
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
            color: rgb(56, 56, 238);
        }
        
        /*change a tag hover color*/
        a:hover {
            color: purple;
        }
        
        /*bottom of table border*/
        tr.last-row {
            border-top: 2px solid #8C8273;
            border-bottom: none;
        }
        
        a.no-color:hover, a.no-color {
        color: inherit;
        text-decoration: none;
        }
        
        h4.bottom-num {
            padding: 0;
            margin: 0;
        }
        
        img.file {
            content: url("/78805a221a988e79ef3f42d7c5bfd41859b514174bffe4ae402b3d63aad79fe0/file_icon_dark.png");
            width: 25px; 
            height: 25px;
        }
        
        img.folder {
            content: url("/78805a221a988e79ef3f42d7c5bfd41859b514174bffe4ae402b3d63aad79fe0/folder_icon_dark.png");
            width: 25px;
            height: 25px;
        }
        
        img.sort-icon {
            content: url("/78805a221a988e79ef3f42d7c5bfd41859b514174bffe4ae402b3d63aad79fe0/sort_arrow_icon_dark.png");
            width: 10px;
            height: 10px;
        }
        
        button.sort-button {
            background-color: transparent;
            padding-left: 2.5px;
            border: none;
        }
        
        button.sort-button:hover {
            cursor: pointer;
        }
        
        img.logo-img {
            content: url("/78805a221a988e79ef3f42d7c5bfd41859b514174bffe4ae402b3d63aad79fe0/logo_dark.png");
            width: 200px;
            height: 200px;
            padding: 0;
            margin: 0;
        }
        
        h1.title {
            font-family: 'Allerta Stencil'; 
            padding: 0;
            margin: 0;
        
        }
        
        div.logo {
            margin-left: 10px;
            border-radius: 15px;
        }
        
        span.U {
            color: aqua;
            padding: 0;
            margin: 0;
        }
        
        span.server {
            color: #FFA72B;
        }
    </style>
</head>
<body>
    <div class="logo">
        <h1 class="title">
            T
            <span class="U">U</span>
            C
            <span class="server">SERVER</span>
        </h1>
        <img class="logo-img"  alt="logo">
    </div>
    <h2><a href="/" class="no-color">Driectory: C:\Users\Owner\</a></h2>
    <br>
    <table>
	{{ range $i, $val := .Items }}
	    <tr>
	    {{ if $val.IsDir }}
	        <td><img class="folder" alt="folder"></td>
	        <td><a href="{{ $val.RelativePath }}">{{ $val.Name }}</a></td>
	    {{ else }}
            <td><img class="file" alt="file"></td>
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
