# Powershell

- 查找有哪些成員或者方法可以使用: `Get-Member`
- 列出所有環境變數: `Get-ChildItem Env:`
- 查找變數: `Get-Variable` 後面可以接篩選條件
- 創建psd1指令: `New-ModuleManifest`
- 生成guid指令: `Microsoft.PowerShell.Utility\New-Guid`
- `Get-PSRepository`
- 更新powershell: `winget install --id Microsoft.Powershell --source winget`
    - 會安裝在: `%ProgramFiles%\PowerShell\7\pwsh.exe`
    - 查看powershell相關應用程式的位置: `gcm powershell, pwsh`
- [開啟powershell7](https://learn.microsoft.com/zh-tw/powershell/scripting/whats-new/differences-from-windows-powershell?view=powershell-7.3#powershell-executable-changes):
  pwsh
- [powershell github](https://github.com/PowerShell/PowerShell/)

## IDE

- [powershell_ise.exe](https://learn.microsoft.com/zh-tw/powershell/scripting/windows-powershell/ise/introducing-the-windows-powershell-ise?view=powershell-7.3):
  直接在powershell中輸入`powershell_ise.exe`即可開啟
    - 它的所在路徑: `%winDir%\%System32\WindowsPowerShell\v1.0\powershell_ise.exe`

## DEBUG

- `Wait-Debugger`在代碼中使用此指令，執行到此時會強制觸發debug程序
- Get-PSBreakPoint
- Set-PSBreakpoint -Script "calendar.psm1" -Line 137
- 設定好之後只要接下powershell有執行到該腳本，就會自動觸發debug模式，觸發debug的時候，前面會面`[DBG]:`的字眼
    - 在dbg的模式下，可以使用`?`來查看所有debug可以使用的命令

      ```yaml
      s, stepInto         Single step (step into functions, scripts, etc.)
      v, stepOver         Step to next statement (step over functions, scripts, etc.)
      o, stepOut          Step out of the current function, script, etc.

      c, continue         Continue operation
      q, quit             Stop operation and exit the debugger
      d, detach           Continue operation and detach the debugger.

      k, Get-PSCallStack Display call stack

      l, list             List source code for the current script.
      Use "list" to start from the current line, "list <m>"
      to start from line <m>, and "list <m> <n>" to list <n>
      lines starting from line <m>

       <enter>             Repeat last command if it was stepInto, stepOver or list

      ?, h                displays this help message.
      ```

        - `q`:放棄debug以及中斷不再繼續執行
        - 您可以善用`k`, 來得到目前您腳本的Location，就可以在使用[Set-PSBreakpoint](https://learn.microsoft.com/en-us/powershell/module/microsoft.powershell.utility/set-psbreakpoint?view=powershell-7.3)來增加中斷點
          > `Set-PSBreakpoint -Script "calendar.psm1" -Line 125` 設定好之後直接用`o`(stepOut)應該就可以很快地跑到您下中斷點的位置

      至於您對哪一個變數有興趣，直接在命令上打出該變數就可以查看


- 刪除所有中斷點 `Get-PSBreakpoint | Remove-PSBreakPoint`

## Command

- 查找command來至於哪一個模組: `Get-Command My-Command | select -exp Source`
- 查看該函數的實作內容: `get-command myCmd | format-list`當中的Definition可以看到
- `(get-command xxx).Definition`
- 查找command模組資訊: `(get-command xxx).Module`
    - 查找某個模組下所有可以使用的指令 `Get-Command -Module MyModule`
    - 查找模塊的路徑: `Get-Module -Name utils | format-List -Property Path, Name`
    - 取的可用的模組: `Get-Module`
- 找尋command `get-command -Name *id*`
- 找尋command開頭須符合 `get-command -Name id*`

如果想找command的定義，除了查看Definition以外，也可以去找該模塊的路徑，使用Get-Module有可能還是沒辦法查看到完整的路徑

但如果您是系統指令，例如`Get-command Get-StartApps` [Get-StartApps就位於](https://ccmcache.wordpress.com/2017/10/15/use-powershell-to-dynamically-manage-windows-10-start-menu-layout-xml-files/)

```yaml
Get-command Get-StartApps # 知道這個指令來自於StartLayout模塊
$env:PSModulePath.Split(";")
# C:\WINDOWS\system32\WindowsPowerShell\v1.0\Modules\StartLayout # 此路徑應該是PSModulePath的最後一個
# win+R: 輸入: shell:AppsFolder
```

### Alias

- `Get-Alias` 如果找不到, 請使用Get-Command查找一次該別名的函數，在使用`Get-Alias`就會出了. (雖然Get-Alias可能會看不到，但其實直接使用alias還是可以呼叫該command成功)
- `Get-Alias swa | Get-Member`
- 查找指令的alias(知道完整指令查找別名): `Get-Alias -Definition MyFullCmdName`

> 如果是PowerShell 5.1還是會受到ExecutionPolicy所影響，如果沒有設定，`Get-Help`
> 就可能會看不到該command的Alias `Set-ExecutionPolicy -Scope Process -ExecutionPolicy Bypass -F`

----

常用指令的別名

| Fullname       | Alias                                     | Description
|----------------|-------------------------------------------| ----
| ForEach-Object | foreach                                   |
| Write-Output   | write, **echo**                           |
| Where-Object   | where                                     |
| Remove-Item    | **ri**, rm, **rmdir**, **del**, erase, rd | 可以刪除物件、檔案、資料夾
| Get-ChildItem  | gci, **ls**, **dir**                      |
| Start-Process  | **start**                                 | 開啟目錄或者檔案

### Get-Help

除了查詢指定的命令以外，也可以查看參數的意思，例如:

```
Get-Help Select-Object -Parameter ExpandProperty
```

output:

| ExpandProperty | string                               |
|----------------|--------------------------------------|
| 必要?            | false                                |
| 位置?            | 已命名                                |
| 接受管線輸入?     | false                                |
| 參數集名稱        | SkipLastParameter, DefaultParameter  |
| 別名             | 無                                   |
| 動態?            | false                                |

## [UWP (Universal Windows Platform)](https://learn.microsoft.com/zh-tw/windows/apps/desktop/modernize/desktop-to-uwp-extend#show-a-modern-xaml-ui)

UWP(通用Windows平台)可以幫忙創建應用程式。
UWP是使用WinRT API來實現相關功能

WinRT API: 是使用c++語言所實現，它的底層技術使用windows API主要靠C語言實現。
WinRT是基於COM語言，這是一種與語言無關，只要滿足其接口就可以調用該API

- 列出所有UWP元件: `Get-AppxPackage -Name *xaml*` (不適用powershell7的版本)

## [Windows PowerShell 5.1 與 PowerShell 7.x 之間的差異](https://learn.microsoft.com/zh-tw/powershell/scripting/whats-new/differences-from-windows-powershell?view=powershell-7.3)

Windows PowerShell 5.1 建置在 .NET Framework v4.5 之上。

PowerShell 6.0 版成為以 .NET Core 2.0 為基礎的開放原始碼專案。

從.NET Framework移至 .NET Core 可讓 PowerShell 成為跨平臺解決方案使得PowerShell可在 Windows、macOS 和 Linux 上執行。


> 簡單來說`.NET Framework`只能在windows平台執行，而`.NET Core`是一個跨平台專案，可以在{Windows, macOS, Linux}上執行

## [powershell7支持那些系統](https://learn.microsoft.com/zh-tw/previous-versions/powershell/scripting/whats-new/what-s-new-in-powershell-70?view=powershell-7.1#where-can-i-install-powershell)

```
Windows 8.1 和 10
Windows Server 2012、2012 R2、2016 及 2019
macOS 10.13+
Red Hat Enterprise Linux (RHEL) / CentOS 7
Fedora 30+
Debian 9
Ubuntu LTS 16.04+
Alpine Linux 3.8+
```

其中[Powershell 7.3已確定不在能windows7執行](https://learn.microsoft.com/zh-tw/powershell/scripting/whats-new/what-s-new-in-powershell-73?view=powershell-7.3#breaking-changes-and-improvements)

## [powershell.editorconfig](https://github.com/PowerShell/PowerShell/blob/7fb867167e9702b292c643f6a4f4cc934acf4811/.editorconfig)

## cmdlet

Commandlet: 通常由一個`動詞+名詞`組成，例如`Get-Process`、`New-Item`等等

Cmdlet一般而言有三個條件

1. 函數名稱以動詞開頭，例如:{Get-, Set-, New-}等等: 算是慣例，沒有滿足，也不會怎樣。

   > 警告: 有些來自模組 'xxx' 的匯入命令名稱包含未核准的動詞，因此可能不易搜尋。如需核准動詞的清單，請輸入 [Get-Verb](https://learn.microsoft.com/en-us/powershell/module/microsoft.powershell.utility/get-verb?view=powershell-7.3)
   >
   > 找尋特定的動詞是否有在該列表 `Get-Verb | findstr XXX` 注意XXX有區分大小寫
   >
   > 如果你要找是哪一個函數不合法可以加上Verbose，例如 Import-Module xxx -Verbose 如果發現函數名稱不符合指定的動詞就會告訴您

   如果你的名稱不想要遵守指定的動詞規範，那麼函數名稱可以不要是用`-`來串接，例如`Create-Shortcut`改為`CreateShortcut`，這樣也不會跳出警告，但搜尋速度應該還是有些影響，建議盡量遵守規範！

3. 必須要有一個或多個參數:
   因為Cmdlet的設計是在powershell管道中提供小型命令，他們通常需要處理輸入和輸出，而輸入和輸出數據都是透過參數來傳遞，因此如果一個函數沒有參數，它就不能接受輸入或者向外輸出數據，就不符合Cmdlet的設計理念。
4. 必須要有回傳值

   在powershell中，任何沒有被附值給變量的語句或者表達式的結果，都會自動被視為返回值

   如果最後一行改成`echo $sum`，這種就不算返回值

   ```yaml
   function Add-Numbers ($a, $b) {
    $sum = $a + $b
    $sum   # 返回 $sum 的值
   }
   ```

Cmdlet函數可以透過管道`|`來串接

- 檢驗command是不是cmdlet的型別: `Get-Command myCmd | Select-Object CommandType` 如果是，回傳會是Cmdlet，否則可能返回其他類型，例如`Function`
- `Get-Command -Name C* -CommandType Cmdlet`: 列出所有C開頭且是Cmdlet型別的命令
    - 但是我試的結果都是Function，感覺不需要太糾結於是Function還是Cmdlet

## [ParameterSets](https://learn.microsoft.com/en-us/powershell/module/microsoft.powershell.core/about/about_parameter_sets?view=powershell-7.3)

假設我們參數可能有些可選項，然後有一些是要根據不同的流程，來決定是否要放那些參數，此時DefaultParameterSetName就可以幫上忙

它的好處是當模擬兩可的時候，自動提是可以自動幫您依據DefaultParameterSetName的內容帶出需要完成的參數

此外它有`$PSCmdlet.ParameterSetName`，可以知道當前是選擇哪一個，如果不使用，將要自己多一個參數來確定到底是執行哪一個

另外如果不使用ParameterSetName，每個流程的必填也沒辦法避免，變成要把所有的參數都打出來，例如你可能有OCR, Watch等流程，當您現在想走Watch流程，依此理應當只要填入Watch必要的參數即可，但是當您不使用ParameterSetName，那麼您沒辦法做到這件事情。


例如
```ps1
function TestParameterSet {
    [CmdletBinding(DefaultParameterSetName='OCR')] # 當我們只有輸入-Pram5的時候，它就會自動提示Param3的參數讓我們填，這就是DefaultParameterSetName的特色

    param (
        [Parameter(ParameterSetName='Watch', Mandatory=$true)]
        [string]$Param1,
        [Parameter(ParameterSetName='Watch', Mandatory=$false)]
        [string]$Param2,

        [Parameter(ParameterSetName='OCR', Mandatory=$true)]
        [string]$Param3,
        [Parameter(ParameterSetName='OCR', Mandatory=$false)]
        [string]$Param4,

        # 必填
        [Parameter(Mandatory=$true)]
        [string]$Param5,
        # 選填
        [Parameter(Mandatory=$false)]
        [string]$Param6
    )

    if ($PSCmdlet.ParameterSetName -eq 'Watch') { # ParameterSetName可以決定$PSCmdlet.ParameterSetName的內容，如果您不用ParameterSetName時，此時就必須再新增一個額外的變數來記錄您可能是要執行哪一個流程
            Write-Output "You've chosen Watch."
            Write-Output "Param1: $Param1"
            Write-Output "Param2: $Param2"
    } elseif ($PSCmdlet.ParameterSetName -eq 'OCR') {
            Write-Output "You've chosen OCR."
            Write-Output "Param3: $Param3"
            Write-Output "Param4: $Param4"
   }
}
```

列出名稱為`TestParameterSet`的使用方法

```yaml
PS> (Get-Command TestParameterSet).ParameterSets | Select-Object -Property @{n='ParameterSetName';e={$_.name}},@{n='Parameters';e={$_.ToString()}}

ParameterSetName Parameters
---------------- ----------
OCR              -Param3 <string> -Param5 <string> [-Param4 <string>] [-Param6 <string>] [<CommonParameters>]
Watch            -Param1 <string> -Param5 <string> [-Param2 <string>] [-Param6 <string>] [<CommonParameters>]

即
TestParameterSet -Param1 1 -Param5 5 # OCR
TestParameterSet -Param3 3 -Param5 5 # Watch
```

透過以上的指令，就可以列出這個命令的用法到底有哪些

## ~~[自定義開始佈局(磚塊牆)](https://learn.microsoft.com/en-us/windows/configuration/customize-and-export-start-layout)~~ (只能影響默認設定、不能修改當前的設定)

```yaml
Export-StartLayout -UseDesktopApplicationID -Path layout.xml # Path之後的名稱不重要，只是輸出到哪一個檔案而已
```

此指令會導出當前您開始選單中所看到的磚塊牆輸出到layout.xml之中，內容大概長成下面這樣

```xml

<LayoutModificationTemplate xmlns:defaultlayout="http://schemas.microsoft.com/Start/2014/FullDefaultLayout"
                            xmlns:start="http://schemas.microsoft.com/Start/2014/StartLayout" Version="1"
                            xmlns="http://schemas.microsoft.com/Start/2014/LayoutModification">
    <LayoutOptions StartTileGroupCellWidth="6"/>
    <DefaultLayoutOverride>
        <StartLayoutCollection>
            <defaultlayout:StartLayout GroupCellWidth="6">
                <start:Group Name="創作">
                    <start:Tile Size="4x2" Column="2" Row="0"
                                AppUserModelID="microsoft.windowscommunicationsapps_8wekyb3d8bbwe!Microsoft.WindowsLive.Mail"/>
                    <start:Tile Size="1x1" Column="5" Row="3"
                                AppUserModelID="Microsoft.Office.OneNote_8wekyb3d8bbwe!microsoft.onenoteim"/>
                    <start:Tile Size="1x1" Column="5" Row="2"
                                AppUserModelID="Microsoft.Office.Desktop_8wekyb3d8bbwe!PowerPoint"/>
                </start:Group>
                <start:Group Name="玩樂">
                    <start:Tile Size="2x2" Column="0" Row="0"
                                AppUserModelID="Microsoft.XboxApp_8wekyb3d8bbwe!Microsoft.XboxApp"/>
                    <start:Tile Size="1x1" Column="3" Row="1" AppUserModelID="Microsoft.WindowsMaps_8wekyb3d8bbwe!App"/>
                    <start:Tile Size="1x1" Column="2" Row="1"
                                AppUserModelID="Microsoft.WindowsCalculator_8wekyb3d8bbwe!App"/>
                    <start:Tile Size="1x1" Column="3" Row="0"
                                AppUserModelID="Microsoft.ZuneMusic_8wekyb3d8bbwe!Microsoft.ZuneMusic"/>
                    <start:Tile Size="1x1" Column="2" Row="0"
                                AppUserModelID="Microsoft.ZuneVideo_8wekyb3d8bbwe!Microsoft.ZuneVideo"/>
                </start:Group>
                <start:Group Name="探索">
                    <start:Tile Size="2x2" Column="4" Row="2" AppUserModelID="Microsoft.SkypeApp_kzf8qxf38zg5c!App"/>
                    <start:Tile Size="2x2" Column="2" Row="2"
                                AppUserModelID="Microsoft.MSPaint_8wekyb3d8bbwe!Microsoft.MSPaint"/>
                    <start:Tile Size="2x2" Column="0" Row="2" AppUserModelID="Microsoft.BingWeather_8wekyb3d8bbwe!App"/>
                    <start:DesktopApplicationTile Size="2x2" Column="4" Row="0" DesktopApplicationID="MSEdge"/>
                    <start:Tile Size="4x2" Column="0" Row="0"
                                AppUserModelID="Microsoft.WindowsStore_8wekyb3d8bbwe!App"/>
                </start:Group>
                <start:Group Name="Carson">
                    <!-- 注意如果不是appID，前面用的是start:DesktopApplicationTile -->
                    <start:DesktopApplicationTile Size="2x2" Column="0" Row="0"
                                                  DesktopApplicationLinkPath="%ProgramData%\Microsoft\Windows\Start Menu\Programs\dovego.lnk"/>
                </start:Group>
            </defaultlayout:StartLayout>
        </StartLayoutCollection>
    </DefaultLayoutOverride>
</LayoutModificationTemplate>
```

修改完xml之後可以透過

```yaml
# Import-StartLayout -LayoutPath my.xml -MountPath C:\ 建議用下面的命令
Import-StartLayout -LayoutPath my.xml -MountPath $env:SystemDrive\
```

----

而這裡面的項目，都必須新增捷徑在開始的目錄之中，要新增捷徑

就必須把它放在AppsFolder之中，要檢視AppsFolder的內容可以使用以下命令

```
win+R: 輸入: shell:AppsFolder
```

至於這個資料夾的內容怎麼新增，則是在以下目錄

```yaml
shell:AppsFolder
start "$env:APPDATA\Microsoft\Windows\Start Menu\Programs\" # 使用者
start "$env:ProgramData\Microsoft\Windows\Start Menu\Programs\" # LocalMachine
start "$env:ProgramData\Microsoft\Windows\Start Menu\Programs\<myFolder>\<my.lnk>" # 如果沒有東西，可能是您直接把目錄複製過去之類的，嘗試把檔案重新命名之後再使用命名回來再用Get-StartApps就會看到東西

  # 查找*dove*, *go-http*, *Example*的項目
Get-StartApps | where { $_.Name -match '(dove|go-http|Example)' }
```

放完之後在重新檢視AppsFolder就會看到它有在裡面了，此外`Get-StartApps`也會列出來，你就可以看到AppID

> 不過並非所有的項目都一定會在`Start Menu\Programs`中被找到，例如透過命令:
>
> > Add-AppxPackage -Path "C:\...\test\AppxManifest.xml" -Register
>
> 他的exe檔案位置可以取決於AppxManifest.xml，所以不一定只能放在`Start Menu\Programs`之中。
>
> 詳細的位置可以透過Get-AppxPackage找到，例如
>
> > (Get-AppxPackage -Name "MyPackage.Identity.Name").InstallLocation

```yaml
Name   AppID
----   -----
dovego C:\...\go\1.16\bin\dovego.exe

  # 如果你以此AppID去[Windows.UI.Notifications.ToastNotificationManager]::CreateToastNotifier("{{.AppID}}");
  # 那麼他會在:
HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Notifications\Settings
  # 建立出目錄，而由於AppID有用\分開，所以他會一層一層的建立目錄，最後得到的位置是:
\HKEY_CURRENT_USER\SOFTWARE\Microsoft\Windows\CurrentVersion\Notifications\Settings\C:\...\go\1.16\bin\dovego.exe
```

## LocalGroupMember

```yaml
Get-LocalGroup | ForEach-Object { Write-Host "Group Name: $($_.Name)"; Get-LocalGroupMember -Name "$($_.Name)"; } # $_ 是迴圈的變數, 如果要在字串中直接取得該變數的某個成員，不可以直接用$_.Name，要使用$($_.Name)
Get-LocalGroupMember -Name Administrators # 查看當前Administrators群組成員有哪些
Add-LocalGroupMember -Group "Administrators" -Member $env:UserName # 添加當前使用者成為Administrators成員
```

## env

列出所有env的項目

```
Get-ChildItem Env:
```

## 註冊AppID

這兩個都命令{`Get-AppxPackage`, `Add-AppxPackage`}都只在5.1支持

AppID要成功有以下幾種方法都可以辦到

1. 在`Start Menu\Programs`新增捷徑:

   ```yaml
   start "$env:APPDATA\Microsoft\Windows\Start Menu\Programs\" # User
   start "$env:ProgramData\Microsoft\Windows\Start Menu\Programs\" # Machine
   ```

   但是之後用`[Windows.UI.Notifications.ToastNotificationManager]::CreateToastNotifier("{{.AppID}}")`顯示訊息的時候，他會用捷徑的路徑寫到機碼去

   ```yaml
    # HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Notifications\Settings
    # 建立出目錄，而由於AppID有用\分開，所以他會一層一層的建立目錄，最後得到的位置是:
    \HKEY_CURRENT_USER\SOFTWARE\Microsoft\Windows\CurrentVersion\Notifications\Settings\C:\...\go\1.16\bin\dovego.exe
   ```

2. 使用Add-AppxPackage來註冊應用程式 (推薦用這個，最穩定)

   這項內容可能會涉及到:

    - [AppxManifest.xml](https://learn.microsoft.com/en-us/uwp/schemas/appxpackage/how-to-create-a-basic-package-manifest)
    - 憑證 (可選項)

   使用Add-AppxPackage的項目，可以不用在這兩個資料夾之中

    ```yaml
    "$env:APPDATA\Microsoft\Windows\Start Menu\Programs\" # User
    "$env:ProgramData\Microsoft\Windows\Start Menu\Programs\" # Machine
    ```

   大部份都是在:"$env:ProgramFiles\WindowsApps\"之中

   但是也不限定，像我安裝我寫的App用的位置就是其他地方

   ```yaml
    Get-AppxPackage | ForEach-Object { if ($_.Name -match "Example") { Write-Output "Object info:`n"; $_ | Format-List ; start $($_.InstallLocation); }}

    Name             ExampleApp
    Publisher        CN=Example
    Architecture     Neutral
    ResourceId
    Version          1.0.0.0
    PackageFullName  ExampleApp_1.0.0.0_neutral__s2ne61n4j7kre
    InstallLocation  C:\Users\...\src  # <-- 在其他的路徑
   ```

> 💡 不論是`建立捷徑`還是用`Add-AppxPackage`的方式，都可以在shell:AppsFolder的目錄中看到

### [Add-AppxPackage](https://learn.microsoft.com/en-us/powershell/module/appx/add-appxpackage?view=windowsserver2022-ps)

```yaml
  # Add-AppxPackage -Path $ManifestPath -Register -DisableDevelopmentMode # 所謂的DisableDevelopmentMode是指，如果您當前是開發模式，那麼它就會把這個app disable，也就是在開發模式下禁用這個項目
Add-AppxPackage -Path "C:\...\test\AppxManifest.xml" -Register # 後面一定要加上Register，不然會遇到參數錯誤
Get-AppxPackage -Name *Example* # 就可以查找到PackageFullName
Get-StartApps -Name  *Example* | Select-Object Name, AppID  # 可以查看到AppID
Remove-AppxPackage -Package "PackageFullName"
Remove-AppxPackage -Package "ExampleApp_1.0.0.0_neutral__s2ne61n4j7kre"

# \HKEY_USERS\S-1-5-21-3051027765-3782066248-1388807790-1001\SOFTWARE\Classes\Local Settings\Software\Microsoft\Windows\CurrentVersion\AppModel\Repository\Packages\ExampleApp_1.0.0.0_neutral__s2ne61n4j7kre
```

### 常用命令

```yaml
Get-StartApps
Get-AppxPackage # 這些的內容其實都保存在 %ProgramFiles%\WindowsApps\ 但是檔案總管不能直接瀏覽，可以用系統管理員dir $env:ProgramFiles\WindowsApps\
dir "$env:ProgramFiles\WindowsApps\" | foreach { if($_.Name -match 'Example') {$_} }
dir "$env:ProgramFiles\WindowsApps\" | foreach { if($_.Name -match '(AppUp|Notepad)') {$_} }
dir "$env:ProgramFiles\WindowsApps\" | foreach { if($_.Name -match '(s2ne61n4j7kre)') {$_} } # ExampleApp_s2ne61n4j7kre!ExampleApp
start "$env:ProgramFiles\WindowsApps\AppUp.IntelGraphicsExperience_1.100.4779.0_neutral_split.language-zh-hant_8j3eq9eme6ctt" # 在上面查到之後就可以開啟
echo "$env:ProgramFiles\WindowsApps\" # 會沒有權限瀏覽
echo "$env:APPDATA\Microsoft\Windows\Start Menu\Programs"
echo "$env:ProgramData\Microsoft\Windows\Start Menu\Programs"
  # Add-AppxPackage -Path "C:\...\out\MyApp.appx" # 這個一直試不成功
Add-AppxPackage -Path "C:\...src\AppxManifest.xml" -Register # 這個一定可以成功！ 後面一定要加上Register，不然會遇到參數錯誤

Get-AppxPackage | ForEach-Object { if ($_.Name -eq "Microsoft.WindowsCalculator") { Write-Output "Object info:`n"; $_ | Format-List ; Write-Output "InstallLocation:`n$($_.InstallLocation)"; }} # 小算盤
Get-AppxPackage | ForEach-Object { if ($_.Name -match "LinkedInforWindows") { Write-Output "Object info:`n"; $_ | Format-List ; Write-Output "InstallLocation:`n$($_.InstallLocation)"; }} # LinkedInforWindows
Get-AppxPackage | ForEach-Object { if ($_.Name -match "Example") { Write-Output "Object info:`n"; $_ | Format-List ; Write-Output "InstallLocation:`n$($_.InstallLocation)"; }}
Get-AppxPackage | ForEach-Object { if ($_.Name -match "Example") { Write-Output "Object info:`n"; $_ | Format-List ; start $($_.InstallLocation); }} # 開啟安裝的目錄
Get-AppxPackage | ForEach-Object { if ($_.Name -match "AppUp.IntelGraphicsExperience") { Write-Output "Object info:`n"; $_ | Format-List ; start $($_.InstallLocation); }} # 開啟安裝的目錄


  # 查看VisualElements
$appx = Get-AppxPackage -Name Microsoft.WindowsCalculator
$manifest = (Get-AppxPackageManifest $appx).Package
$manifest.Applications.Application.VisualElements

  # 查看AppID
Get-StartApps -Name "小算盤" | Select-Object Name, AppID
```

### 憑證相關

| 執行檔名稱       | 位置                                                                           | 描述 |
|-------------|----------------------------------------------------------------------------------| ---- |
| makeAppx    | echo "$env:ProgramFiles (x86)\Windows Kits\10\bin\10.0.17763.0\x64\makeappx.exe" | 產生出appx的檔案
| makeCert    | echo "$env:ProgramFiles (x86)\Windows Kits\10\bin\10.0.17763.0\x64\makecert.exe" | 可以被`New-SelfSignedCertificate`指令取代
| mmc.exe     |                                                                                  | 開啟主控台，可以匯入多種畫面，包含憑證(如果要檢視Local的內容就要靠它)
| certmgr.msc |                                                                                  | 檢視`User`的憑證

makeAppx

```yaml
cd C:\Program Files (x86)\Windows Kits\10\bin\10.0.17763.0\x64\makeappx.exe # 10.0.17763.0是您的版本號碼，如果沒有要去載windows SDK就會有了
makeappx pack /d . /p MyApp.appx
./makeappx pack /d "C:\...\src" /p "C:\...\out\MyApp.appx"
```

```yaml
# 列出所有憑證內容
Get-ChildItem cert:\CurrentUser

  # 會跑出以下資訊，就代表CurrentUser底下還有這些的目錄
Name TrustedPublisher
Name ClientAuthIssuer
Name Root # 受信任的跟憑證授權單位
Name UserDS
Name CA # 中繼憑證授權單位
Name REQUEST
Name AuthRoot
Name TrustedPeople
Name ADDRESSBOOK
Name Local NonRemovable Certificates
Name My
Name SmartCardRoot
Name Trust
Name Disallowed

  # 找尋名稱
Get-ChildItem -Path cert:\CurrentUser -Recurse | Where-Object {$_.Subject -eq "CN=MyApp Test Certificate"}
  # 如果你有需要刪除，找到他之後再使用Remove-Item去刪除
Get-ChildItem -Path cert:\CurrentUser -Recurse | Where-Object {$_.Subject -eq "CN=MyApp Test Certificate"} | foreach { Remove-Item -Path $_.PSPath }
Get-ChildItem -Path cert:\LocalMachine -Recurse | Where-Object {$_.Subject -eq "CN=MyApp Test Certificate"} | foreach { Remove-Item -Path $_.PSPath }
```

mmc.exe
certmgr.msc

```yaml

certmgr.msc # 這個只能顯示CurrentUser
mmc.exe # Local Computer. 開啟後是一個空白的介面，在檔案 > 嵌入式管理單元 > 選擇憑證 > Local Computer > 匯入

Get-ChildItem -Path cert:\
  # 輸出
  # Location   : CurrentUser
  # StoreNames : {TrustedPublisher, ClientAuthIssuer, Root, UserDS...}
  # Location   : LocalMachine
  # StoreNames : {TestSignRoot, ClientAuthIssuer, m, Root...}

C:\Program Files (x86)\Windows Kits\10\bin\10.0.17763.0\x64\makecert.exe # 不需要makecert.exe  可以靠New-SelfSignedCertificate來完成
cd C:\Program Files (x86)\Windows Kits\10\bin\10.0.17763.0\x64\
./makecert -r -pe -n "CN=MyApp Test Certificate" -b 01/01/2023 -e 01/01/2026 -ss my # CN為憑證的名稱 產生一個區間為(begin, end): 2023~2026年, 存放在: cert:\CurrentUser\My底下
  # 如果你建立錯誤，不是建立在my(my是一個關鍵字，表示個人)，那麼可以用certMgr.msc的搜尋，用名稱去搜，找到之後再把它刪除即可
./makecert -r -pe -n "CN=MyApp Test Certificate" -b 01/01/2023 -e 01/01/2026 -sky exchange -ss my  # 添加 "-sky exchange" 參數，以確保憑證包含可以用於程式碼簽署的私鑰
./makecert -r -pe -n "CN=MyApp Test Certificate" -b 01/01/2023 -e 01/01/2026 -sky exchange -ss my -a sha256 # 指定演算法為sha256, 不然預設的是sha1
./makecert -r -pe -n "CN=MyApp Test Certificate" -b 01/01/2023 -e 01/01/2026 -sky exchange -a sha256 -eku 1.3.6.1.5.5.7.3.3 -ss my # -eku 1.3.6.1.5.5.7.3.3表示要加上 Code Signing 的擴展屬性
$cert = New-SelfSignedCertificate -DnsName "CN=MyApp Test Certificate" -CertStoreLocation cert:\CurrentUser\My -KeySpec Signature -KeyUsage DigitalSignature -FriendlyName "MyApp Test Certificate" -NotBefore (Get-Date).Date -NotAfter (Get-Date).Date.AddDays(365) -HashAlgorithm SHA256 -KeyExportPolicy Exportable -TextExtension @("2.5.29.37={text}1.3.6.1.5.5.7.3.3")
Get-ChildItem cert:\CurrentUser\My
  # 上面的弄完之後，打開certMgr.msc => 個人 => 憑證 也可以查看到我們所簽屬的憑證
  # 要刪除憑證，可以用certMgr.msc的UI介面刪除即可

  # 完成之後還是會有問題，查看憑證會看到警告: 這個CA根憑證不受信任因為它不是位於受信任的根憑證授權單位存放區中
$cert = Get-ChildItem -Path cert:\CurrentUser\My -Recurse | Where-Object {$_.Subject -eq "CN=MyApp Test Certificate"}
Export-Certificate -Type CERT -Cert $cert -FilePath "C:\MyCert.cer" # 匯出憑證 (需要管理員權限)
  # 接著點擊右鍵 > 安裝憑證 > 目前使用者 > 將所有憑證放入以下的存放區 > 選擇: 受信任的根憑證授權單位
  # 或者是直接用certMgr.msc從個人.憑證資料夾中選擇該項目，把它移到 「受信任的根憑證授權單位.憑證目錄」 之中也可以

  # 不可以直接在一開始就安裝在
  # New-SelfSignedCertificate -DnsName "CN=MyApp Test Certificate" -CertStoreLocation cert:\CurrentUser\TrustedPublisher -KeySpec Signature -KeyUsage DigitalSignature -FriendlyName "MyApp Test Certificate" -NotBefore (Get-Date).Date -NotAfter (Get-Date).Date.AddDays(365) -HashAlgorithm SHA256 -KeyExportPolicy Exportable -TextExtension @("2.5.29.37={text}1.3.6.1.5.5.7.3.3")

# 可以安裝到My之後在寫入到Root，就不需要透過UI. 但還是會出現警告視窗(沒辦法避免)
$cert = New-SelfSignedCertificate -DnsName "CN=MyApp Test Certificate" -CertStoreLocation cert:\LocalMachine\My -KeySpec Signature -KeyUsage DigitalSignature -FriendlyName "MyApp Test Certificate" -NotBefore (Get-Date).Date -NotAfter (Get-Date).Date.AddDays(365) -HashAlgorithm SHA256 -KeyExportPolicy Exportable -TextExtension @("2.5.29.37={text}1.3.6.1.5.5.7.3.3")
Get-ChildItem -Path Cert:\LocalMachine\My -CodeSigningCert # 只顯示所有CodeSigningCert的項目
  # $store = New-Object System.Security.Cryptography.X509Certificates.X509Store -ArgumentList "TrustedPublisher", "LocalMachine"
$store = New-Object System.Security.Cryptography.X509Certificates.X509Store -ArgumentList "Root", "LocalMachine" # Root才是「受信任的**根**憑證授權單位」
$store.Open([System.Security.Cryptography.X509Certificates.OpenFlags]::ReadWrite)
$store.Add($cert) # 在將憑證新增至憑證存放區時，如果該憑證的發行者（Issuer）不在受信任的憑證授權單位存放區中，系統會發出安全警告。因此，即使透過程式碼進行新增，仍然會出現安全警告。 這個警告主要是要提醒使用者注意憑證的發行者是否可信，若確定憑證是可信的，可以選擇信任該憑證並新增至存放區中，才能正確地使用該憑證。
$store.Close() # 安裝在Root之後，除了Root，在CA: 中繼憑證授權單位 也會出現
  # Remove-Item -Path cert:\LocalMachine\My\$($cert.Thumbprint) # 移除也可以直接透過指紋來移除
Remove-Item -Path $cert.PSPath # 以上方法雖然可行，但前面捕的位置有點饒口，覺得用$cert.PSPath會比較清楚

Get-ChildItem -Path cert:\LocalMachine\My
  # 輸出
Thumbprint                                Subject
----------                                -------
37E7789F6FD45B573705CD9DB4D8D72C5AE5E8A7  CN=1F906F59-B093-4E7E-8564-5D8E5548A460 # 原本舊有的項目
852FDDC5739C2D7C55B01D0B5D16B9D9BFA67BF0  CN=MyApp Test Certificate # 我們所建立的憑證

  # 接著要確保我們的憑證可以被用來簽屬程式碼，不然會遇到錯誤：「Set-AuthenticodeSignature : 無法簽署程式碼。指定的憑證不適合程式碼簽署。」
$cert = Get-ChildItem -Path cert:\CurrentUser\My
$cert2 = New-Object System.Security.Cryptography.X509Certificates.X509Certificate2($cert)
  # $cert2.Extensions.Find("2.5.29.37") # 沒有Find的方法
$cert.Extensions # 總之這邊是確認. 一開始如果您有在New-SelfSignedCertificate有加上-eku 1.3.6.1.5.5.7.3.3就會是CodeSigningCert，表示此憑證可以用來代碼簽署
```

#### 如何刪除憑證

以下為刪除憑證的方法，有待檢驗:

```yaml
# 載入憑證
$store = New-Object System.Security.Cryptography.X509Certificates.X509Store -ArgumentList "My", CurrentUser
$store.Open([System.Security.Cryptography.X509Certificates.OpenFlags]::ReadWrite)

  # 找到要移除的憑證
$thumbprint = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
$cert = $store.Certificates.Find([System.Security.Cryptography.X509Certificates.X509FindType]::FindByThumbprint, $thumbprint, $false)

  # 如果找到要移除的憑證就移除
if ($cert.Count -gt 0) {
$store.Remove($cert[0])
}

  # 關閉憑證
    $store.Close()
```

其實不需要上面的指令那麼麻煩，透過以下的指令就能辦到

```yaml
Get-ChildItem -Path cert:\CurrentUser -Recurse | Where-Object {$_.Subject -eq "CN=MyApp Test Certificate"} | foreach {Remove-Item -Path $_.PSPath }
```

#### ~~簽署Appx~~ (沒有試成功過)

```yaml
# 簽署appx
# 注意請先用: Get-ChildItem cert:\CurrentUser\My 確定有幾個，如果只有一個可以直接用，如果有多個要改成(Get-ChildItem cert:\CurrentUser\My)[n] n為第幾個的下標值
Set-AuthenticodeSignature -FilePath "C:\...\out\MyApp.appx" -Certificate (Get-ChildItem cert:\CurrentUser\My)[1] -CodeSigningCert # 如果出現「找不到符合參數名稱 'CodeSigningCert' 的參數。」的錯誤，請確認有在New-SelfSignedCertificate有加上-eku 1.3.6.1.5.5.7.3.3就會是CodeSigningCert，表示此憑證可以用來代碼簽署

  # $cert = (Get-ChildItem cert:\CurrentUser\My)[1] # 不能用My, My的內容不被信任
Get-ChildItem -Path cert:\CurrentUser -Recurse | Where-Object {$_.Subject -eq "CN=MyApp Test Certificate"} | foreach { echo $_.PSPath } # 找出Root的PSPath
$cert = Get-ChildItem -Path "Microsoft.PowerShell.Security\Certificate::CurrentUser\Root\4ABEFB58180FBE1A82F0048956A5C828A214755F" # 後面就放PSPath的路徑名稱
Set-AuthenticodeSignature -FilePath "C:\...\out\MyApp.appx" -Certificate $cert -HashAlgorithm "SHA256" -TimestampServer "http://timestamp.digicert.com"

Set-AuthenticodeSignature -FilePath "C:\...\out\MyApp.appx" -Certificate (Get-ChildItem cert:\CurrentUser\My)[1] -HashAlgorithm "SHA256" -TimestampServer "http://timestamp.digicert.com"
Set-AuthenticodeSignature -FilePath "C:\...\out\MyApp.appx" -Certificate (Get-ChildItem cert:\CurrentUser\TrustedPublisher) -HashAlgorithm "SHA256" -TimestampServer "http://timestamp.digicert.com"

  # 確認有沒有被簽屬成功
Get-AuthenticodeSignature -FilePath "C:\...\out\MyApp.appx"

Add-AppxPackage -Path "C:\...\out\MyApp.appx" # 預設似乎是都裝在LocalMachine，所以如果失敗不仿試試看換到LocalMachine
```

## 新增一個項目到「新增或移除程式」中

```yaml
$RegPath = "HKLM:\SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall\MyTestApp123"
$AppName = "MyTestApp123"
$AppVersion = "1.0.0"
$Publisher = "My Company"
$UninstallString = "C:\Program Files\MyTestApp123\Uninstall.exe"
$InstallLocation = "C:\Program Files\MyTestApp123"

New-Item -Path $RegPath -Force | Out-Null
New-ItemProperty -Path $RegPath -Name "DisplayName" -Value $AppName -PropertyType "String" -Force | Out-Null
New-ItemProperty -Path $RegPath -Name "DisplayVersion" -Value $AppVersion -PropertyType "String" -Force | Out-Null
New-ItemProperty -Path $RegPath -Name "Publisher" -Value $Publisher -PropertyType "String" -Force | Out-Null
New-ItemProperty -Path $RegPath -Name "UninstallString" -Value $UninstallString -PropertyType "String" -Force | Out-Null
New-ItemProperty -Path $RegPath -Name "InstallLocation" -Value $InstallLocation -PropertyType "String" -Force | Out-Null
New-ItemProperty -Path $RegPath -Name "EstimatedSize" -Value 12345 -PropertyType "DWord" # 用的是byte, 所以12345相當於12.1MB左右
# New-ItemProperty -Path $RegPath -Name "Size" -Value "12" # 寫Size沒有用，不論有沒有加MB都沒用，正確的是寫EstimatedSize才可以在新增/移除的該項目中看到內容
# Set-ItemProperty -Path $RegPath -Name "Size" -Value "12 MB" # 如果要修改可以用Set-ItemProperty
```

## ShellNotifyIcon殘留

1. HKEY_CURRENT_USER\SOFTWARE\Classes\Local Settings\Software\Microsoft\Windows\CurrentVersion\TrayNotify
2. (備份整個TrayNotify資料夾，以防萬一)
3. 刪除IconStreams, PastIconsStream兩個機碼數值
4. 開啟工作管理員(taskmgr.exe)，刪除所有explorer.exe的項目
5. 再次執行explorer.exe

## Foreach參考

```
foreach ($list in @((1..10), (11..20)) ) { foreach ($j in $list) { echo $j } }
@((1..10), (11..20)) | foreach { $_.GetType(); $_ | foreach { echo $_ } }
$list = @((1..10), (11..20)); $list | foreach { $item = $_; $item.GetType()  }
$list = @((1..10), (11..20)); $list | foreach { $item = $_; $item | foreach { $_.GetType() }  }
@((1..10), (11..20)) | foreach { $_ | foreach { echo $_ } }
```

## Powershell5.1的注意事項

### 註解
如果您是用powershell5.1去開發，有可能會因為註解而影響到，因為他的編碼不是UTF8，您可以在註解的最後面加上「`;`」應該就可以執行了

### SupportsShouldProcess

有一些指令會缺少，例如Start-Process在5.1就沒有Confirm的選項，所以如果要兼容，可能程式碼要做判斷

## 參考資料

- [discord powershell社群](https://discord.com/channels/180528040881815552/)
- 黑暗執行緒
    - [Powershell 學習筆記](https://blog.darkthread.net/blog/powershell-learning-notes/)
    - [GET/POST參考](https://blog.darkthread.net/blog/test-webapi-without-tool/): `Invoke-WebRequest`
- [PowershellBook](https://books.goalkicker.com/PowerShellBook/)
- [Use PowerShell to Dynamically Manage Windows 10 Start Menu Layout XML Files](https://ccmcache.wordpress.com/2017/10/15/use-powershell-to-dynamically-manage-windows-10-start-menu-layout-xml-files/)
