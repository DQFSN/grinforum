{{if account_normal_authority .Account}}
<div class="">
</div>
{{ end }}

<div class="body-sidebar">
  {{template "components/_sidebar-about.tpl" . }}
</div>

<div class="body-sidebar">
    {{template "components/_friend_link.tpl" . }}
</div>

{{template "tags/_taglist.tpl" . }}

<div class="body-sidebar">
    <div style="text-align: center;">
        <img src="/static/img/qrcode.jpg" style="width: 100%;"></img>
        <div>公众号（古灵阁）</div>
    </div>
</div>