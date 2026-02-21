# School Bus Schedule API

本文件說明 `schoolbusscheduleapi.go` 中各個 API 的回傳 JSON 內容與格式。

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

### `Schedule` (班次資訊)
```json
{
  "stations": [
    {
      "departTime": {
        "hour": 0,
        "minute": 0
      },
      "name": "string",
      "type": 0
    }
  ],
  "isStudentOnly": false,
  "onHoliday": false,
  "vehicleType": "string",
  "daysOfWeek": ["Mon", "Tue"] // 根據 getDayFlagsDescription 回傳的字串陣列或字串
}
```

### `NextBusResult` (下一次班次查詢結果)
```json
{
  "index": 0,
  "schedule": { /* Schedule 物件 */ }
}
```

---

## API 列表

### 1. `LoadSavedDataApi`, `RefreshSchoolBusDataApi`
- **說明**：載入儲存的資料 / 重新抓取校車資料。
- **回傳內容**：無特定 JSON 資料（若成功則資料為空字串，發生錯誤則回傳錯誤訊息）。

### 2. `GetLastSchoolBusDataFetchTimeApi`
- **說明**：取得最後一次抓取校車資料的時間。
- **回傳內容**：
```json
"2006-01-02T15:04:05Z07:00" // JSON 序列化後的 ISO 8601 時間字串
```

### 3. `GetYcToHpScheduleApi`, `GetHpToYcScheduleApi`
- **說明**：取得 燕巢開往和平 / 和平開往燕巢 的完整班次表。
- **回傳內容**：`Schedule` 物件陣列。
```json
[
  { /* Schedule 物件 1 */ },
  { /* Schedule 物件 2 */ }
]
```

### 4. `GetYcToHpNextBusNowApi`, `GetHpToYcNextBusNowApi`
- **說明**：以當前時間查詢 燕巢開往和平 / 和平開往燕巢 的下一班車。
- **回傳內容**：包含該班次在當天班次表中的索引值與班次資料。
```json
{
  "index": 0,
  "schedule": { /* Schedule 物件 */ }
}
```

### 5. `GetYcToHpNextBusApi`, `GetHpToYcNextBusApi`
- **說明**：以指定時間查詢 燕巢開往和平 / 和平開往燕巢 的下一班車。
- **回傳內容**：與 `NextBusNowApi` 相同。
```json
{
  "index": 0,
  "schedule": { /* Schedule 物件 */ }
}
```

### 6. `GetYcToHpBusByIndexApi`, `GetHpToYcBusByIndexApi`
- **說明**：透過索引值查詢特定的校車班次。
- **回傳內容**：單一 `Schedule` 物件。
```json
{ /* Schedule 物件 */ }
```
