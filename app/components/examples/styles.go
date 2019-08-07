package examples

const styles = `
#todo-main {
	border: 1px solid #dedede;
	border-radius: 4px;
	padding: 0px 0px 4px 0px;
}

#todo-main:childes {
	color: red;
}

#todo-main ul {
	padding: 0 16px;
	list-style-type: none;
	margin-left: 0;
}

#todo-main ul li {
	display: flex;
	padding: 4px 8px;
	border-bottom: 1px solid #dedede;

	font-size: 18px;
}

#todo-main ul li div {
	width: 100%;
	display: flex;
}

#todo-main ul li input {
	width: 100%;
}

#todo-main ul li button {
	border: 0;
	padding: 0;
	background-color: inherit;
	cursor: pointer;
}
#todo-main ul li button#submit:hover, #todo-main button#submit:focus {
	color: #009966;
}
#todo-main ul li button#delete:hover, #todo-main button#delete:focus {
	color: #ff0033;
}

#todo-main ul li button#submit {
	margin: 0 12px 0 0;
}

#todo-main ul li button#delete {
	margin: 0 0 0 auto;
}

#todo-main nav {
	padding: 6px 16px;
	margin-bottom: 8px;
	border-bottom: 1px solid #dedede;
}

#todo-main nav button {
	margin-right: 6px;
	border: 0;
	padding: 0;
	color: #009966;
	background-color: inherit;
	cursor: pointer;
}
#todo-main nav button:focus, nav button:hover {
	color: #00CC99;
}
#todo-main nav button.active {
	text-decoration: underline;
}

#todo-main #todo-new {
	margin: 0 16px;
}

.todo-input {
	background: #fff;
    border: .05rem solid #bcc3ce;
    border-radius: .1rem;
    color: #3b4351;
    display: inline-block;
    font-size: .8rem;
    height: 1.8rem;
    line-height: 1.2rem;
    outline: 0;
    padding: .25rem .4rem;
    text-decoration: none;
    vertical-align: middle;
    white-space: nowrap;
}
.todo-input:focus {
	border: .05rem solid #5755d9;
}

#todo-wrap footer {
	margin-top: 18px;
	color: gray;
	font-size: 12px;
	text-align: center;
}
#todo-wrap footer div {
	margin-bottom: 4px;
}
#todo-wrap footer a {
	margin: 0 4px;
	color: inherit;
}
`
