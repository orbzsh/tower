{{define "navbar"}}
		<a class="navbar-brand" href="/">GoChair</a>
		<div>
			<ul class="nav navbar-nav">
				
				<li {{if .IsIndex}}class="active"{{end}}><a href="/">首页</a></li>
				<li {{if .IsCategory}}class="active"{{end}}><a href="/category">分类</a></li>
				<li {{if .IsTopic}}class="active"{{end}}><a href="/topic">文章</a></li>
				<!--
				<li {{if .IsIndex}} class="active" {{end}}><a href="/">GoIndex</a></li>
				<li {{if .IsJobs}} class="active" {{end}}><a href="/jobs">GoJobs</a></li>
				<li {{if .IsHdns}} class="active" {{end}}><a href="/hdns">GoHttpDNS</a></li>
				<li {{if .IsIp}} class="active" {{end}}><a href="/qip">GoIP</a></li>
				<li {{if .IsTodo}} class="active" {{end}}><a href="/todo">GoTodo</a></li>
				-->
			</ul>
		</div>

		<!--
		<div class="pull-right">
			<ul class="nav navbar-nav">
				{{if .IsLogin}}
					<li><a href="/login?exit=true">退出</a></li>
				{{else}}
					<li><a href="/login">管理员登录</a></li>
				{{end}}
			</ul>
		</div>
		-->
{{end}}