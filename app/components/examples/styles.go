package examples

const styles = `
#todo #todo-main {
	border: 1px solid #dedede;
	border-radius: 4px;
	padding: 0px 0px 4px 0px;
}

#todo #todo-main:childes {
	color: red;
}

#todo ul {
	padding: 0 16px;
	list-style-type: none;
	margin-left: 0;
}

#todo ul li {
	display: flex;
	padding: 4px 8px;
	border-bottom: 1px solid #dedede;

	font-size: 18px;
}

#todo ul li div {
	width: 100%;
	display: flex;
}

#todo ul li input {
	width: 100%;
}

#todo ul li button {
	border: 0;
	padding: 0;
	background-color: inherit;
	cursor: pointer;
}
#todo ul li button#submit:hover, #todo button#submit:focus {
	color: #009966;
}
#todo ul li button#delete:hover, #todo button#delete:focus {
	color: #ff0033;
}

#todo ul li button#submit {
	margin: 0 12px 0 0;
}

#todo ul li button#delete {
	margin: 0 0 0 auto;
}

#todo nav {
	padding: 6px 16px;
	margin-bottom: 8px;
	border-bottom: 1px solid #dedede;
	background-color: #f1f1f1;
}

#todo nav button {
	margin-right: 6px;
	border: 0;
	padding: 0;
	color: #009966;
	background-color: inherit;
	cursor: pointer;
}
#todo nav button:focus, nav button:hover {
	color: #00CC99;
}
#todo nav button.active {
	text-decoration: underline;
}

#todo #todo-new {
	margin: 0 16px;
}

#todo .footer {
	margin-top: 18px;
	color: gray;
	font-size: 12px;
	text-align: center;
}
#todo .footer div {
	margin-bottom: 4px;
}
#todo .footer a {
	margin: 0 4px;
	color: inherit;
}
`
