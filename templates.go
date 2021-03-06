// Copyright 2016 Lars Wiegman. All rights reserved. Use of this source code is
// governed by a BSD-style license that can be found in the LICENSE file.

package multipass

import "html/template"

const (
	notfoundPage = iota
	loginPage
	tokenInvalidPage
	tokenSentPage
	continueOrSignoutPage
	errorPage
)

type page struct {
	PJAX         bool
	Page         int
	NextURL      string
	LoginPath    string
	SignoutPath  string
	CSRFField    template.HTML
	ErrorMessage string
}

func loadTemplates() *template.Template {
	tmpl := template.New("")
	template.Must(tmpl.New("head").Parse(`<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<meta name="description" content="Multipass">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
	<title>Multipass</title>
	<style>{{ template "style" }}</style>
</head>
<body>`))

	template.Must(tmpl.New("tail").Parse(`</body></html>`))

	template.Must(tmpl.New("loginform").Parse(`
<form action="{{ .LoginPath }}" method=POST class="login-form">
	<input type=hidden name=next value="{{ .NextURL }}" />
	<input type=email name=handle placeholder="Your handle ..." />
	<input type=submit value="Submit" class="btn btn-default">
	{{ .CSRFField }}
</form>`))

	template.Must(tmpl.New("signoutform").Parse(`
<form action="{{ .SignoutPath }}" method=POST class="login-form">
	<input type=hidden name=next value="{{ .NextURL }}" />
	<input type=submit value="Signout" class="btn btn-danger">
	{{ .CSRFField }}
</form>`))

	template.Must(tmpl.New("page").Parse(`
	{{ if ne .PJAX true }}{{ template "head"}}{{end}}
<div class="wrapper">
	<section>
		<article class="login-box">
			<h1>Multipass</h1>
		{{ if eq .Page 1 }}
			{{ template "loginform" . }}
			<p class="notice">This resource is protected. Submit your handle to gain access.</p>
		{{ else if eq .Page 2 }}
			<p class="notice">Your access token has expired or is invalid. Submit your handle to request a new one.</p>
			{{ template "loginform" . }}
			<p class="notice">This resource is protected.</p>
		{{ else if eq .Page 3 }}
			<p>A message with an access token was sent to your handle; Follow the login link in the message to gain access.</p>
		{{ else if eq .Page 4 }}
			<p class="notice">Continue to access the private resource or signout.</p>
			<a href="{{ .NextURL }}" class="btn btn-success">Continue</a>
			{{ template "signoutform" . }}
			<p class="notice">This resource is protected.</p>
		{{ else if eq .Page 5 }}
			<p>{{ .ErrorMessage }}</p>
		{{ else }}
			<h4>Not Found (404)</p>
			<p>Sorry, couldn't find what you're looking for.</p>
		{{ end }}
		</article>
	</section>
</div>
{{ if ne .PJAX true }}{{ template "tail" }}{{ end }}`))

	template.Must(tmpl.New("style").Parse(`
* {
	-webkit-box-sizing: border-box;
	-moz-box-sizing: border-box;
	box-sizing: border-box;
}
article, aside, details, figcaption, figure, footer, header, hgroup,
main, nav, section, summary {
	display: block;
}
body {
	margin: 0;
}
input { -webkit-appearance: none; }
.wrapper {
	display: -webkit-flex;
	-webkit-flex-direction: column;
}
.login-box {
	width: 480px;
	margin: 4rem auto;
	text-align: center;
	padding: 1rem 0;
	border: solid .2rem #0d8eba;
	border-radius: .8rem;
	padding: 1rem 2rem;
	background-color: #fff;
}
.login-box h1 {
	font-family: sans-serif;
	font-size: 3rem;
	color: #444;
}
.login-form {
}
.login-form input {
}
.login-form input[type=text], .login-form input[type=email] {
	font-size: 1.2rem;
	padding: 0 1rem;
	border-color: #ddd;
	color: #444;
	margin: .6rem 0;
	width: 100%;
	height: 2.8rem;
	border-radius: .4rem;
	border-style: solid;
	border-width: .2rem;
}
.login-form input[type=submit] {
}
.btn {
	font-family: sans-serif;
	font-size: 1.2rem;
	margin: 1rem 0;
	padding: .5rem 1rem;
	width: 100%;
	height: 2.8rem;
	border-radius: .4rem;
	border-style: solid;
	border-width: .16rem;
	text-transform: uppercase;
	color: #fff;
	height: 3rem;
	cursor: pointer;
	display: inline-block;
	line-height: 2rem;
}
a.btn {
	text-decoration: none;
	color: white;
}
.btn-default {
	background-color: #0d8eba;
	border-color: #0d8eba;
}
.btn-success {
	background-color: #0dba77;
	border-color: #0dba77;
}
.btn-danger {
	background-color: #ba2f0d;
	border-color: #ba2f0d;
}
.notice {
	font-style: italic;
	color: #666;
}
/* Narrow */
@media only screen and (max-width: 480px) {
	body {
		background-color: #0d8eba;
	}
	.login-box {
		width: 100%;
		border: none;
		border-radius: 0;
	}
}
/* Medium */
@media only screen and (min-width: 481px) and (max-width: 960px) {
}
/* Wide */
@media only screen and (min-width: 961px) {
}
`))
	return tmpl
}
