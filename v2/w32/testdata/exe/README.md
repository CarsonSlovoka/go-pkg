| name | desc.|
| ---- | ---- |
write.exe | 這是從`%windir%/System32/write.exe`所取得的執行檔，因為這個檔案比較小，所以就選擇它
writeWithFont.exe | 這是write程式內多新增了[teamviewer15.otf](../fonts/teamviewer15.otf)


## MFC鑲嵌字型資源教學

1. Resource.h 新增定義

    ```
    #define IDF_XXX  666
    ```

2. 開啟.rc2的檔案，新增
    ```
    IDF_XXX          FONT                    "C:\\...\\xxx.ttf"
    ```
3. 再次建置即可完成

----

另一種方式是直接用UI新增

在資源中添加字型(要打`8`(RT_FONT))

之後右下角可以更改`ID`它的預設可能用IDR開頭，可以自行改成IDF

接著可以打開您的字型檔，直接複製二進位資料，貼在剛才透過8所新增出來的檔案，最後再把右下角的路徑名稱改成自己想要的即可。

不推薦這種方法，要熟悉UI操作很麻煩。
