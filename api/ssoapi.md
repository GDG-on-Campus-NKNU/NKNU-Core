# SSO API

本文件說明 `ssoapi.go` 中各個 API 的回傳 JSON 內容與格式。

所有 API 的最終輸出皆會經過 `utils.FormatBase64Output` 處理，將實際結果包裝成固定的 JSON 結構後，再整體進行 **Base64 編碼** 回傳給客戶端。
將 API 的結果進行 Base64 解碼後，將得到以下統一的 JSON 結構：

```json
{
  "data": <各自API定義的回傳結構或是"">,
  "err": "錯誤訊息(字串)，若無錯誤則為空字串"
}
```

以下章節定義的以及 API 列表中的「回傳內容」，皆是指上述共同結構中 `data` 欄位裡面的內容（除非該 API 單純不回傳資料，則為 `""` 空字串）。

## 資料結構定義 (Models)

### `Session` (連線階段資訊)
```json
{
  "aspNETSessionId": "string",
  "viewState": "string"
}
```

### `HistoryScore` (歷年成績)
成積查詢將回傳一個陣列，陣列中包含多個學期的歷年成績物件：
```json
[
  {
    "title": "111學年度 第1學期",
    "courses": [
      {
        "name": "課程名稱",
        "credit": "學分",
        "category": "必/選修",
        "requirement": "全/半",
        "score": "分數"
      }
    ],
    "averageScore": "平均分數",
    "conductScore": "操行成績",
    "credits": "修習學分",
    "earnedCredits": "實得學分",
    "semesterRanking": "名次/班級人數",
    "classSize": "班級人數"
  }
]
```

### `MailServiceAccount` (郵件帳號資訊)
取得 Google Workspace 及 Office 365 學生信箱帳號與預設密碼：
```json
{
  "google": {
    "account": "學號@mail.nknu.edu.tw",
    "password": "密碼"
  },
  "o365": {
    "account": "學號@o365.nknu.edu.tw",
    "password": "密碼"
  }
}
```

---

## API 列表

### 1. `GetSessionInfoApi`
- **說明**：取得單一登入系統 (SSO) 的 Session 資訊 (AspNETSessionId 與 ViewState)。
- **回傳內容**：
```json
{
  "aspNETSessionId": "string",
  "viewState": "string"
}
```

### 2. `LoginApi`
- **說明**：進行 SSO 登入。
- **回傳內容**：無特定 JSON 資料（若成功則資料為空字串，發生錯誤則回傳錯誤訊息）。

### 3. `GetHistoryScoreApi`
- **說明**：取得學生的歷年成績資料。
- **回傳內容**：歷年成績（陣列格式），包含每個學期的成績單與各科成績。
```json
[
  {
    "title": "string",
    "courses": [ /* Course 物件陣列 */ ],
    "averageScore": "string",
    "conductScore": "string",
    "credits": "string",
    "earnedCredits": "string",
    "semesterRanking": "string",
    "classSize": "string"
  }
]
```

### 4. `GetMailServiceAccountApi`
- **說明**：取得學校提供的外部郵件服務 (Google & O365) 帳號及預設密碼。
- **回傳內容**：包含 Google 與 O365 帳密資訊的物件。
```json
{
  "google": {
    "account": "string",
    "password": "string"
  },
  "o365": {
    "account": "string",
    "password": "string"
  }
}
```
