{{template "../manage/header.tpl"}}
<body>
<div style="padding:20px 20px 40px 80px;">
    <form id="form1" method="post">
        <table width="95%">
            <tr>
                <td width="10%">ID：</td>
                <td colspan="3">{{.cabinet.CabinetID}}</td>
            </tr>
            <tr>
                <td>门数：</td>
                <td>{{.cabinet.Doors}}</td>
            </tr>
            <tr>
                <td>使用：</td>
                <td>{{.cabinet.OnUse}}/{{.cabinet.Doors}}</td>
            </tr>
            <tr>
                <td>类型：</td>
                <td width="30%"><input id="type" type="text" disabled value="{{.cabinet.TypeName}}"/></td>
                <td><input type="button"
                           onclick="(function(){document.getElementById('type').removeAttribute('disabled')})()"
                           value="修改"></td>
            </tr>
            <tr>
                <td>位置：</td>
                <td><input id="address" type="text" disabled value="{{.cabinet.Address}}"/></td>
                <td><input type="button"
                           onclick="(function(){document.getElementById('address').removeAttribute('disabled')})()"
                           value="修改"></td>
            </tr>
            <tr>
                <td>编号：</td>
                <td><input id="number" type="text" disabled value="{{.cabinet.Number}}"/></td>
                <td><input type="button"
                           onclick="(function(){document.getElementById('number').removeAttribute('disabled')})()"
                           value="修改"></td>
            </tr>
            <tr>
                <td>备注：</td>
                <td><textarea id="desc" class="easyui-validatebox" disabled rows="2" cols="21"
                              value="{{.cabinet.Desc}}">{{.cabinet.Desc}}</textarea></td>
                <td><input type="button"
                           onclick="(function(){document.getElementById('desc').removeAttribute('disabled')})()"
                           value="修改"></td>
            </tr>
            <tr>
                <td></td>
            </tr>
            <tr>
                <td colspan="3">
                    <table width="100%">
                        <tr>
                            <th width="8%">门号</th>
                            <th width="8%">开关状态</th>
                            <th width="8%">占用</th>
                            <th width="20%">存物ID</th>
                            <th width="20%">存物时间</th>
                            <th width="8%">启用</th>
                            <th colspan="3">操作</th>
                        </tr>
                    {{range $index, $elem := .cabinet.Detail}}
                        <tr>
                            <td align="center">{{$elem.Door}}</td>
                            <td align="center">
                            {{if eq $elem.OpenState 1}}
                                关
                            {{else if eq $elem.OpenState 2}}
                                开
                            {{end}}</td>
                            <td align="center">
                            {{if eq $elem.Using 1}}
                                空闲
                            {{else if eq $elem.Using 2}}
                                占用
                            {{end}}</td>
                            <td align="center">
                            {{if eq $elem.UserID ""}}
                                --
                            {{else}}
                                {{$elem.UserID}}
                            {{end}}</td>
                            <td align="center">
                            {{if eq $elem.UserID ""}}
                                --
                            {{else}}
                                {{dateformat $elem.StoreTime "2006-01-02 15:04:05"}}
                            {{end}}</td>
                            <td align="center">
                                <label><input class="mui-switch" type="checkbox"{{if eq $elem.UseState 1}}checked{{end}}
                                              onclick="alert('change')"></label>
                            </td>
                            <td align="center">
                                <button>开门</button>
                                <button onclick="flush({{$elem.Id}})">清除</button>
                                <button>记录</button>
                            </td>
                        </tr>
                    {{end}}
                    </table>
                </td>
            </tr>
        </table>
    </form>
</div>
</div>
</body>
<script type="text/javascript">
    var URL = "/state";

    function flush(id) {
        alert(id)
        vac.ajax(URL + '/flush/' + id, null, 'POST', function (r) {
            if (!r.status) {
                vac.alert(r.info);
            } else {
                $("#datagrid").datagrid("reload");
            }
        })
    }
</script>
<style>
    .mui-switch {
        width: 52px;
        height: 31px;
        position: relative;
        border: 1px solid #dfdfdf;
        background-color: #fdfdfd;
        box-shadow: #dfdfdf 0 0 0 0 inset;
        border-radius: 20px;
        border-top-left-radius: 20px;
        border-top-right-radius: 20px;
        border-bottom-left-radius: 20px;
        border-bottom-right-radius: 20px;
        background-clip: content-box;
        display: inline-block;
        -webkit-appearance: none;
        user-select: none;
        outline: none;
    }

    .mui-switch:before {
        content: '';
        width: 29px;
        height: 29px;
        position: absolute;
        top: 0px;
        left: 0;
        border-radius: 20px;
        border-top-left-radius: 20px;
        border-top-right-radius: 20px;
        border-bottom-left-radius: 20px;
        border-bottom-right-radius: 20px;
        background-color: #fff;
        box-shadow: 0 1px 3px rgba(0, 0, 0, 0.4);
    }

    .mui-switch:checked {
        border-color: #64bd63;
        box-shadow: #64bd63 0 0 0 16px inset;
        background-color: #64bd63;
    }

    .mui-switch:checked:before {
        left: 21px;
    }
</style>
</html>