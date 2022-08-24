# sap-api-integrations-assignment-activity-reads  
sap-api-integrations-assignment-activity-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API 競合品目データを取得するマイクロサービスです。  
sap-api-integrations-assignment-activity-reads には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-assignment-activity-reads は、オンプレミス版である（＝クラウド版ではない）SAPC4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/serviceassignment/overview  

## 動作環境
sap-api-integrations-assignment-activity-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。   
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。   
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須） 

## クラウド環境での利用  
sap-api-integrations-assignment-activity-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-assignment-activity-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/serviceassignment/overview  
* APIサービス名(=baseURL): c4codataapi

## 本レポジトリ に 含まれる API名
sap-api-integrations-assignment-activity-reads には、次の API をコールするためのリソースが含まれています。  

* ServiceAssignmentCollection（アサインメント活動 - サービスアサインメント）  

## API への 値入力条件 の 初期値
sap-api-integrations-assignment-activity-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.ServiceAssignmentCollection.ID（ID）  


## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"ServiceAssignmentCollection" が指定されています。    
  
```
	"api_schema": "AssignmentActivity",
	"accepter": ["ServiceAssignmentCollection"],
	"assignment_activity_code": "3054",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "AssignmentActivity",
	"accepter": ["All"],
	"assignment_activity_code": "3054",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetAssignmentActivity(iD string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "ServiceAssignmentCollection":
			func() {
				c.ServiceAssignmentCollection(iD)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP アサインメント活動 の サービスアサインメントデータ が取得された結果の JSON の例です。  
以下の項目のうち、"ObjectID" ～ "EntityLastChangedOn" は、/SAP_API_Output_Formatter/type.go 内 の Type ServiceAssignmentCollection {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-assignment-activity-reads/SAP_API_Caller/caller.go#L53",
	"function": "sap-api-integrations-assignment-activity-reads/SAP_API_Caller.(*SAPAPICaller).ServiceAssignmentCollection",
	"level": "INFO",
	"message": [
		{
			"ObjectID": "00163E0C6CDB1ED6A7A12F76CFFD3994",
			"ETag": "2016-10-28T21:24:36+09:00",
			"ProcessingTypeCode": "AACT",
			"ProcessingTypeCodeText": "Assignment Activity",
			"CustomerTypeCode": "159",
			"CustomerTypeCodeText": "Account",
			"Customer": "Matthew Whitehart",
			"CustomerID": "000000000000000000000000000000000000000000000000000001004460",
			"ServiceTechnicianTypeCode": "167",
			"ServiceTechnicianTypeCodeText": "Employee",
			"ServiceTechnician": "Clemens K",
			"ServiceTechnicianID": "8000000313",
			"TypeCode": "2215",
			"TypeCodeText": "Assignment Activity",
			"Subject": "Repair",
			"LifeCycleStatusCode": "1",
			"LifeCycleStatusCodeText": "Open",
			"ID": "3054",
			"FullDayIndicator": false,
			"Status": "01",
			"StatusText": "Scheduled",
			"AssignmentUUID": "00163E0C-6CDB-1ED6-A7A1-2F76CFFD3994",
			"BusinessActivityUUID": "00163E0C-6CDB-1ED6-A7A1-2F76CFFD3994",
			"StartDateTimeZoneCode": "PST",
			"StartDateTimeZoneCodeText": "Pacific Time (Los Angeles)",
			"Priority": "3",
			"PriorityText": "Normal",
			"DivisionCode": "00",
			"DivisionCodeText": "Cross Division",
			"Released": "",
			"ReleasedText": "Not Released",
			"Fixed": "",
			"FixedText": "",
			"PrimaryContactID": "",
			"PrimaryContact": "",
			"PrimaryContactTypeCode": "",
			"PrimaryContactTypeCodeText": "",
			"CreationDateTime": "2016-10-28T21:24:36+09:00",
			"CreatedBy": "RESPLANNER",
			"LastChangedBy": "EASTSERVICE",
			"OrganizerID": "8000000313",
			"OrganizerName": "Clemens K",
			"OrganizerTypeCode": "167",
			"OrganizerTypeCodeText": "Employee",
			"EndDateTime": "2016-10-28T18:00:00+09:00",
			"EndDateTimeZoneCode": "PST",
			"EndDateTimeZoneCodeText": "Pacific Time (Los Angeles)",
			"StartDateTime": "2016-10-28T17:00:00+09:00",
			"AssignmentSchedulingSource": "",
			"AssignmentSchedulingSourceText": "Manual",
			"EntityLastChangedOn": "2016-10-28T21:24:36+09:00"
		}
	],
	"time": "2022-08-24T16:59:49+09:00"
}
```