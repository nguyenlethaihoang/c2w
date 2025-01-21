# 🪟 c2w - Curl to Windows

A tool is used to convert curl from Linux OS to Windows OS and written by Golang language.

## Download

[Download execution file](https://github.com/nguyenlethaihoang/c2w/raw/refs/heads/main/c2w.exe)

## 🚀 Installation
To install **c2w**, simply clone the repository and follow the instructions below:
```
https://github.com/nguyenlethaihoang/c2w/
cd c2w
```
Running the below command will globally run the **c2w**.
```
go run main.go
```
## 💡 Usage
Firstly, paste the curl command into the file named **input.txt**. For example:
```
curl --request POST \
  --url http://promotion.abc.com/api/promotion/addCustomer \
  --header 'Content-Type: application/json' \
  --data '{
	"customerKey": "123",
	"customerName": "N/A",
	"customerPhone": "123",
	"registerDate": "2025-01-13 00:00:00",
	"sourceId": "32",
	"roleId": "2168",
	"locationId": 0
}'
```
Secondly, run this command:
```
go run main.go
```
Finally, you'll find the converted curl command for Windows in a file named **output.txt**.
## 📝 License
This project is licensed under the MIT License - see the [LICENSE](https://github.com/nguyenlethaihoang/c2w/blob/main/LICENSE) file for details.
